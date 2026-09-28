# go-learning

Go 后端学习与实战记录。目标：以 Go 后端开发为主线（Java 为辅），积累可展示的工程项目，为 2028 年 3 月首次投递做准备。

## 仓库结构

```
.
├── basics/                 # 语言基础
│   ├── variables/          # 变量声明、常量
│   ├── io-output/          # fmt 输出、格式化
│   ├── io-input/           # 标准输入
│   ├── types/              # 基本数据类型、类型转换
│   ├── array/              # 数组
│   ├── slice/              # 切片
│   └── map/                # map
├── control-flow/           # 流程控制
│   ├── if/                 # if 条件
│   ├── switch/             # switch 分支
│   └── for/                # for 循环
├── functions/              # 函数
│   ├── basic/              # 函数定义、可变参数
│   ├── closure/            # 闭包
│   ├── args/               # 值传递与引用传递
│   └── init/               # init 函数
├── structs/                # 结构体
│   ├── basic/              # 结构体定义、方法
│   ├── tag/                # 结构体标签
│   └── custom-type/        # 自定义类型、错误码
├── interfaces/             # 接口
│   └── basic/              # 接口定义与实现
├── cmd/hello/              # 可执行程序入口示例
├── internal/greet/         # 内部包（含单元测试）
└── pkg/version/            # 可复用包
```

## 设计说明

### 为什么每个知识点独立成目录

每个示例都是独立的 `package main`，包含自己的 `main()` 函数。Go 不允许同一目录下存在多个 `main()`，因此**每个知识点一个目录**，可单独运行：

```bash
go run ./basics/variables
go run ./functions/closure
```

### 目录分层约定

| 目录 | 用途 | 说明 |
|---|---|---|
| `basics/` `functions/` 等 | 学习示例 | 每个子目录一个独立可运行程序 |
| `cmd/` | 程序入口 | 每个子目录编译一个可执行文件，业务逻辑不写这里 |
| `internal/` | 内部包 | Go 强制约束：只能被本模块引用，外部无法导入 |
| `pkg/` | 可复用包 | 可对外暴露的公共库 |

## 运行

```bash
# 运行全部测试
go test ./...

# 查看覆盖率
go test -cover ./...

# 编译全部包
go build ./...

# 静态检查
go vet ./...

# 代码格式化
gofmt -w .
```

## 学习进度

- [x] 变量与常量
- [x] 输入输出
- [x] 基本数据类型
- [x] 数组、切片、map
- [x] 流程控制（if / switch / for）
- [x] 函数、闭包、值传递与引用传递
- [x] init 函数
- [x] 结构体、标签、自定义类型
- [x] 接口
- [ ] 并发（goroutine / channel）
- [ ] 错误处理进阶
- [ ] 网络编程（net/http）
- [ ] 数据库操作

## 提交约定

按「学一个点 → 写一段代码 → 提交一次」的节奏推进，保证提交历史真实反映学习轨迹。

提交信息格式：

```
feat: 新增 xxx 知识点示例
fix: 修复 xxx
docs: 更新文档
```
