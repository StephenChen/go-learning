### 函数类型

| 类型   | 格式               | 作用              |
|------|------------------|-----------------|
| 测试函数 | 函数名前缀为 Test      | 测试程序的一些逻辑性为是否正确 |
| 基准函数 | 函数名前缀为 Benchmark | 测试函数的性能         |
| 示例函数 | 函数名前缀为 Example   | 为文档提供示例         |

#### 单元测试函数

#### 表格驱动测试

```shell
# gotest 
go get -u github.com/cweill/gotests/...
gotests -all -w split.go
```

#### 测试覆盖率

```shell
go test -cover
go test -cover -coverprofile=c.out
go tool cover -html=c.out
```

### testify/assert

```shell
go get github.com/stretchr/testify
```

### httptest

### gock

