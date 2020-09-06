---
title: Gin流量处理流程
author: GreenHatHG
date: 2020-09-06 19:30:28
---

# Gin demo

![gin.jpg](https://i.loli.net/2020/09/06/fa9H45F3sRexAUj.jpg)

上面最终调用的是`http.ListenAndServe(address, engine)`，这个函数是`net/http`包的函数，所以我们可以先了解http包下对请求的处理流程

# net/http

## socket

为了更方便地开发网络应用程序，人们推出了一种应用程序访问通信协议的操作系统调用接字(`Socket`)，使得程序员可以很方便地访问 `TCP/IP`，从而开发各种网络应用程序

所谓套接字(`Socket`)，就是**对网络中不同主机上的应用进程之间进行双向通信的端点的抽象**，一般由IP地址和端口号决定

要通过互联网进行通信，至少需要一对套接字( `socket` )，其中一个运行于客户端，我们称之为 `Client Socket `，另一个运行于服务器端，我们称之为 `Server Socket`

其基本流程主要包括

- 服务器监听

- 客户端请求

- 连接确认



![socket.png](https://i.loli.net/2020/09/06/QyvJHYSVmhqD3iP.png)

`net/http`则是实现了上述的过程







---

参考：

[Http包源码分析 · Golang 学习笔记](https://www.huweihuang.com/golang-notes/web/golang-http-execution-flow.html)

[gin源码阅读之一 -- net/http的大概流程 | HHF技术博客](https://www.haohongfan.com/post/2019-02-17-gin-01/#top)

[Go框架解析:gin - TIGERB](http://tigerb.cn/2019/07/06/go-gin/)