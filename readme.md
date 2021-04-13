# 机器参数配置

# 客户端

```bash
# 最大打开文件数量
ulimit -n 200000

# 端口范围
echo "1024 65535" > /proc/sys/net/ipv4/ip_local_port_range

# 
```

# 服务端

```bash
# 最大打开文件数量
ulimit -n 1000000

# 半连接队列
echo 10000 >/proc/sys/net/ipv4/tcp_max_syn_backlog

# 全连接队列
echo 10000 >/proc/sys/net/core/somaxconn

# 最大连接追踪
sysctl -w net.nf_conntrack_max=1000000
```


# 执行测试

```bash
./client --remote-ip 127.0.0.1 --concurrent-num 60000  --local-ip 0.0.0.0
```

# 编译出Linux包

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
```# baiwan-bingfa
