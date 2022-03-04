#### 1、 Create image

```shell
docker build -t httpserver:v0.0.1 .
```

#### 2、Create container

```shell
/bin/bash start.sh
```

####  3、Result

```shell
[root@gitlab ~]# docker ps -a |grep Http
2e929ff0ef1c        httpserver:v0.0.1               "./HttpServer"      9 minutes ago       Up 9 minutes                 0.0.0.0:8080->8080/tcp                               HttpServer
[root@gitlab ~]# curl 'http://192.168.100.202:8080'
<h1> This is a test page</h1>[root@gitlab ~]#
[root@gitlab ~]# docker logs HttpServer
os version11:
os version11:
[root@gitlab ~]# curl 'http://192.168.100.202:8080/healthz'
server is working[root@gitlab ~]#
[root@gitlab ~]# cat /data/logs/HttpServer.log
2022/03/04 22:56:37.662280 E:/Go_code/src/cncamp/week03/HttpServer.go:44: Success! clientIp: 192.168.100.202
2022/03/04 23:07:13.297937 E:/Go_code/src/cncamp/week03/HttpServer.go:44: Success! clientIp: 192.168.100.202
[root@gitlab ~]#
[root@gitlab ~]# ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host
       valid_lft forever preferred_lft forever
2: ens33: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000
    link/ether 00:0c:29:5d:20:51 brd ff:ff:ff:ff:ff:ff
    inet 192.168.100.202/24 brd 192.168.100.255 scope global noprefixroute ens33
       valid_lft forever preferred_lft forever
    inet6 fe80::ee30:f326:75f5:cb0c/64 scope link noprefixroute
       valid_lft forever preferred_lft forever
3: docker0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default
    link/ether 02:42:a4:33:6b:ef brd ff:ff:ff:ff:ff:ff
    inet 172.17.0.1/16 brd 172.17.255.255 scope global docker0
       valid_lft forever preferred_lft forever
    inet6 fe80::42:a4ff:fe33:6bef/64 scope link
       valid_lft forever preferred_lft forever
11: vethad94dc6@if10: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue master docker0 state UP group default
    link/ether 2a:f7:0e:6a:89:9e brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet6 fe80::28f7:eff:fe6a:899e/64 scope link
       valid_lft forever preferred_lft forever
21: vethde387f6@if20: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue master docker0 state UP group default
    link/ether 56:be:8e:33:ba:a3 brd ff:ff:ff:ff:ff:ff link-netnsid 1
    inet6 fe80::54be:8eff:fe33:baa3/64 scope link
       valid_lft forever preferred_lft forever
[root@gitlab ~]# docker inspect -f {{.State.Pid}} HttpServer
13135
[root@gitlab ~]# nsenter -n -t13135
[root@gitlab ~]# ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
20: eth0@if21: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default
    link/ether 02:42:ac:11:00:03 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.3/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever


```

