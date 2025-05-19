# Go Script for Bulk GitHub Tag Deletion

这是一个使用 Go 语言编写的命令行工具，用于高效地从 GitHub 仓库中批量删除标签。它特别适用于需要删除大量（例如，数千个甚至数十万个）标签的场景，因为传统的 GitHub API 逐个删除方法效率低下且容易触发速率限制。

本工具利用 Git 协议的 `git push --delete` 命令来进行标签删除，这比通过 GitHub REST API 进行大量删除通常更高效。

## 功能特性

* 使用 `git ls-remote --tags` 高效获取所有远程标签列表，避免 REST API 的分页限制。
* 支持通过正则表达式模式精确匹配要删除的标签。
* 使用 `git push --delete` 命令通过 Git 协议进行删除。
* 支持并发删除，提高处理速度。
* 包含重要的用户确认步骤，防止误删。
* 报告删除成功和失败的结果。

## 先决条件

在运行此脚本之前，你需要准备以下环境：

1.  **Go 环境：** 需要安装 Go 语言环境（推荐最新版本）。
2.  **Git：** 需要在你的系统上安装 Git。
3.  **本地仓库克隆：** 你需要将目标 GitHub 仓库克隆到本地。**脚本必须在克隆的仓库目录下运行。**
4.  **Git 推送权限：** 你的本地 Git 环境需要配置好推送到远程仓库的权限。这通常意味着：
    * **使用 SSH Keys (推荐)：** 将你的 SSH 公钥添加到你的 GitHub 账号中，并确保本地 SSH Agent 正在运行。
    * **使用 HTTPS 和 Personal Access Token (PAT)：** 配置 Git Credential Helper 来缓存一个具有足够权限（至少 `repo` 或 `write:packages` 范围）的 GitHub PAT。**脚本本身不直接使用 PAT，它依赖你的 Git 配置来处理认证。**

## 构建脚本

1.  将脚本代码保存到你的本地仓库目录下，例如命名为 `delete_large_tags.go`。
2.  在终端中进入到你的本地仓库目录。
3.  运行 Go 构建命令：
    ```bash
    go build delete_large_tags.go
    ```
    这会在当前目录下生成一个名为 `delete_large_tags` (或者在 Windows 上是 `delete_large_tags.exe`) 的可执行文件。

## 使用方法

在终端中，**切换到你的本地仓库目录**，然后运行生成的可执行文件，并提供一个正则表达式作为参数来指定要删除哪些标签。

```bash
./delete_large_tags <tag_pattern>
```

- `tag_pattern`：一个必需参数，表示用于匹配要删除标签名称的正则表达式。

**重要提示：** 在开始删除之前，脚本会列出匹配到的标签数量（以及部分名称），并要求你输入 yes 来确认删除。请务必仔细检查！

## 示例

1. 删除所有以 old- 开头的标签：
```bash
./delete_large_tags '^old-.*'
```
2. 删除所有版本号小于 2.0 的标签 (需要更复杂的模式或二次过滤)：
（例如，删除 v1.0, v1.1, v1.9 但保留 v2.0, v2.1 等）
```bash
./delete_large_tags '^v1\.\d+$'
# 或者删除所有以 v1. 开头的标签
# ./delete_large_tags '^v1\.'
```
3. 删除所有标签（极度危险！请务必确认！）：
```bash
./delete_large_tags '.*'
```
