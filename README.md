# scip
扫描 IP 地址的开启端口

## 安装

```
go get -u github.com/zangdale/scip@latest
```

## 使用

```
scip -h
Usage of scip:
  -port uint
        port ...
  -proxy
        use proxy ...
```

```
scip
scip 192.168.10.10
scip -proxy 192.168.10.10
scip -port 8080
```
