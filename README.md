# go-shim

基于DRY原则（Don't Repeat Yourself），Go垫片库主要提供一些技术(`x`)和业务(`shim`)的垫片库和方法，用于常见的软件开发过程。

> DRY原则（Don't Repeat Yourself）是软件开发中的一项重要原则，它强调避免代码重复。该原则的核心思想是，在一个系统中，每个功能或信息应该只有唯一的、明确的表达方式。

## Feature

1. 在开发过程中常用的一些垫片函数，比如数值、日期、时间、金钱等
2. 在开发过程中第三方包的扩展，比如数据库、Cron
3. Debug相关，比如GoVal、格式化对象输出等
4. 函数命名尽量遵循Linux KISS原则，比如`shim.UniqElems()`
5. 数值部分采用了范型，使之更加通用

## 目录说明

```
$ tree -L 2 -d     
.
├── shim // 业务垫片
│   ├── goval
│   ├── grayacc
│   └── log
└── x  // 技术垫片
    ├── crond
    ├── mysqlx
    ├── openaix
    └── redisx
```

## 示例

### 函数示例

`go-shim`主要for业务场景简单封装，例如 从字符串、数字切片中

- 元素是否存在: `InElems[T comparable](elem T, elems []T) bool`
- 元素去重: `UniqElems[T comparable](elems []T) []T`
- 元素分片: `ShardingElems[T comparable](elems []T, batchSize int) (batches [][]T)`
- 元素分页: `PagingElems[T interface{}](elems []T, page int, size int) []T`
- 拼接成SQL字符串: `JoinElems[T comparable](elems []T, sep string) string `


