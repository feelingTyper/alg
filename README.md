# 一些实用算法

## ACAutoMachine
```Go
patterns := []string{
    "我的",
    "tpu",
}
content := "我的最爱是tpu"
results := ac.NewAcMachine().AddPatterns(patterns...).Build().SimpleQuery(content)
```
