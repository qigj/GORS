# GORS
远程执行本地的脚本

## 编译
```
go build -o GORS main.go
```
## 使用

- `-f`是必须添加的
- `-p`和`-k`二选一，同时添加使用密码

```
root@c18e40ed9187:/data/test-go/GORS# ./GORS -h
Usage of ./GORS:
  -f string
        * 本地脚本文件的路径.
  -k string
        * SSH私钥文件的路径.
  -p string
        * 远程服务器的密码.
  -s string
        远程服务器的IP地址或主机名加端口. (default "127.0.0.1:22")
  -u string
        远程服务器的用户名. (default "root")
```