# github.com/hold7techs/go-shim

通常在开发过程中会用到一些垫片函数，比如针对一个数值数组去重、分片处理，该库就提供了在这些场景下的基础函数。


## Feature
1. 在开发过程中常用的一些垫片函数，比如数值、日期、时间、金钱等 
2. 在开发过程中第三方包的扩展，比如数据库、Cron
3. Debug相关，比如GoVal、格式化对象输出等 
4. 函数命名尽量遵循Linux KISS原则，比如`shim.Uniq()`
5. 数值部分采用了范型，使之更加通用

## 示例

### Number

```
# Uniq
shim.Uniq([]uint{1,2,2} => []uint{1,2}
shim.Uniq([]uint64{1,2,2} => []uint64{1,2}

# Sharding Numbers
shim.ShardingNumbers([]int{1, 2, 3}, 2) => [][]int{{1, 2}, {3}}
```

### Debug

```
type Fav struct {
    Name string
}
type user struct {
    Id   int
    Name string
    Fav  []*Fav
}

t.Logf(ToJsonString(&user{1, "user1", []*Fav{{"sport"}}}, false))
// output
{"Id":1,"Name":"user1","Fav":[{"Name":"sport"}]}

t.Logf("%+v", &user{1, "user1", []*Fav{{"sport"}}})
// outoput
got: &{Id:1 Name:user1 Fav:[0x1400010b460]}
```


