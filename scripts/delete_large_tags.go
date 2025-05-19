package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"sync"
)

// 用于传递删除结果的结构体
type deleteResult struct {
	TagName string
	Err     error
	Output  string // 记录执行 git push 的输出，方便调试
}

func main() {
	// 检查参数
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <tag_pattern>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Example: %s '^old-.*'\n", os.Args[0]) // 删除所有以 'old-' 开头的标签
		os.Exit(1)
	}

	tagPattern := os.Args[1]
	remoteName := "origin" // 默认远程名称，如果你的远程不是 origin，请修改

	// 检查是否在 Git 仓库目录下
	if _, err := os.Stat(".git"); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Error: Must be run inside a Git repository clone.\n")
		os.Exit(1)
	}

	// 编译正则表达式
	regex, err := regexp.Compile(tagPattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error compiling regex pattern '%s': %v\n", tagPattern, err)
		os.Exit(1)
	}

	fmt.Printf("Fetching tags from remote '%s'...\n", remoteName)

	// 1. 获取远程仓库所有标签列表
	// git ls-remote --tags origin
	cmdLsRemote := exec.CommandContext(context.Background(), "git", "ls-remote", "--tags", remoteName)
	stdout, err := cmdLsRemote.StdoutPipe()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating stdout pipe for git ls-remote: %v\n", err)
		os.Exit(1)
	}
	if err := cmdLsRemote.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting git ls-remote: %v\n", err)
		os.Exit(1)
	}

	var allRemoteTags []string
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 2 {
			ref := parts[1]
			// 提取标签名称，确保是 refs/tags/ 开头
			if strings.HasPrefix(ref, "refs/tags/") {
				tagName := strings.TrimPrefix(ref, "refs/tags/")
				allRemoteTags = append(allRemoteTags, tagName)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading git ls-remote output: %v\n", err)
		// 尝试等待命令结束，获取可能的错误信息
		cmdLsRemote.Wait()
		os.Exit(1)
	}

	// 等待 git ls-remote 命令完成
	if err := cmdLsRemote.Wait(); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing git ls-remote: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Found %d remote tags.\n", len(allRemoteTags))
	fmt.Printf("Filtering tags matching pattern: '%s'\n", tagPattern)

	// 2. 筛选要删除的标签
	var tagsToDelete []string
	for _, tagName := range allRemoteTags {
		if regex.MatchString(tagName) {
			tagsToDelete = append(tagsToDelete, tagName)
		}
	}

	if len(tagsToDelete) == 0 {
		fmt.Println("No tags found matching the pattern. Nothing to delete.")
		os.Exit(0)
	}

	fmt.Printf("Found %d tags to delete:\n", len(tagsToDelete))
	// 为了避免打印 200k 个标签，只打印前 100 个或一部分
	displayLimit := 100
	if len(tagsToDelete) > displayLimit {
		for i := 0; i < displayLimit; i++ {
			fmt.Println(tagsToDelete[i])
		}
		fmt.Printf("... and %d more.\n", len(tagsToDelete)-displayLimit)
	} else {
		for _, tag := range tagsToDelete {
			fmt.Println(tag)
		}
	}

	// 3. 确认删除
	fmt.Print("\nAre you sure you want to delete the listed tags? (yes/no): ")
	reader := bufio.NewReader(os.Stdin)
	confirmation, _ := reader.ReadString('\n')
	confirmation = strings.ToLower(strings.TrimSpace(confirmation))

	if confirmation != "yes" {
		fmt.Println("Deletion cancelled.")
		os.Exit(0)
	}

	fmt.Println("Starting deletion process...")

	// 4. 并发删除标签
	// 使用 Goroutines 池来并发执行 git push --delete
	const numWorkers = 20 // 可以根据你的网络和机器性能调整并发数
	tagsChan := make(chan string, len(tagsToDelete))
	resultsChan := make(chan deleteResult, len(tagsToDelete)) // Buffered channel for results

	var wg sync.WaitGroup

	// 启动 worker Goroutines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go deleteWorker(tagsChan, resultsChan, &wg, remoteName)
	}

	// 将要删除的标签发送到 channel
	for _, tag := range tagsToDelete {
		tagsChan <- tag
	}
	close(tagsChan) // 关闭 channel，信号 worker 没有更多任务了

	// 等待所有 worker 完成
	wg.Wait()
	close(resultsChan) // 关闭结果 channel

	// 5. 收集和报告结果
	deletedCount := 0
	var failedTags []deleteResult
	for res := range resultsChan {
		if res.Err == nil {
			deletedCount++
			// 可以选择在这里打印成功的标签，但对于大量标签可能刷屏
			// fmt.Printf("Successfully deleted %s\n", res.TagName)
		} else {
			failedTags = append(failedTags, res)
			fmt.Fprintf(os.Stderr, "Error deleting %s: %v\nOutput: %s\n", res.TagName, res.Err, res.Output)
		}
	}

	fmt.Println("\nDeletion process finished.")
	fmt.Printf("Deleted %d tags.\n", deletedCount)

	if len(failedTags) > 0 {
		fmt.Fprintf(os.Stderr, "Failed to delete %d tags.\n", len(failedTags))
		fmt.Fprintf(os.Stderr, "Failed tags:\n")
		// 只打印一部分失败的标签，避免输出过多
		displayLimit = 100
		for i, res := range failedTags {
			if i >= displayLimit {
				fmt.Fprintf(os.Stderr, "... and %d more.\n", len(failedTags)-displayLimit)
				break
			}
			fmt.Fprintf(os.Stderr, "- %s\n", res.TagName)
		}
	}
}

// deleteWorker 是一个 Goroutine，从 tagsChan 读取标签并执行删除命令
func deleteWorker(tagsChan <-chan string, resultsChan chan<- deleteResult, wg *sync.WaitGroup, remoteName string) {
	defer wg.Done()

	for tagName := range tagsChan {
		// git push origin --delete <tag_name>
		cmdDelete := exec.CommandContext(context.Background(), "git", "push", remoteName, "--delete", tagName)
		output, err := cmdDelete.CombinedOutput() // 捕获 stdout 和 stderr

		// 将结果发送回主 Goroutine
		resultsChan <- deleteResult{
			TagName: tagName,
			Err:     err,
			Output:  string(output),
		}

		// 可选：在 worker 之间添加短暂的延迟
		// time.Sleep(50 * time.Millisecond)
	}
}
