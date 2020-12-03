# TIC 接口调用示例

## 最简单的流程

1. 创建资源栈，得到资源栈ID和版本ID
2. 使用上一步的资源栈ID和版本ID进行调用`PlanStack`，`ApplyStack`，`DestroyStack`分别进行预览，执行和销毁操作，其中销毁会删除资源栈创建的资源

可以参考[main.go](main.go)代码示例。

## 在已有的资源栈或版本上操作

1. 已有资源栈ID，可以给调用`CreateStackVersion`增加一个版本，或者调用`UpdateStackVersion`更新版本内容，然后进行plan，apply，destroy操作资源栈

## 查询接口

- `DescribeStacks`用于查询一个或多个资源栈信息
- `DescribeStackVersions`用于查询一个或多个版本的详细信息，可以通过名字(Name)和状态(Status)过滤
- `DescribeStackEvents`用于查看一个或多个事件详细信息，可以通过版本ID(VersionId)，资源栈ID（StackId)，事件类型(Type)，事件状态(Status)过滤
- `DescribeStackEvent`用于获取单个事件详情，尤其是可以得到事件的详细控制台输出文本

## 注意事项

执行事件的过程可能比较耗时，需要等待事件状态结束，并且处于plan，apply状态下的版本无法再执行其他事件，需要等待状态变成success或fail。