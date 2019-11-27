# openccgo [![Build Status](https://github.com/MartialBE/openccgo/workflows/Go/badge.svg)](https://github.com/MartialBE/openccgo/releases)  [![Go Report Card](https://goreportcard.com/badge/github.com/MartialBE/openccgo)](https://goreportcard.com/report/github.com/MartialBE/openccgo)


用于简繁体转换， opencc 是一个很好用的简繁体转化工具， 不过由于分词比较弱，可能导致转化时，出现某些词转换失败的情况。
所以 该工具在opencc基础上 增加了 分词的功能， 通过 先分词， 然后在进行 转化，这样得到的 转换 错误率将会降低。


## 安装

```
go get -u github.com/MartialBE/openccgo

make

```

## 使用

```
openccgo -i [待转换文件地址] -o [输出文件地址] -m [s2t,t2s,s2tw,tw2s,s2hk,hk2s,s2twp,tw2sp,t2tw,t2hk]
```