---
title: MySQL脚本解释
author: GreenHatHG
date: 2020-09-05 12:02:28
---

# 简介

## 功能

1. 自动安装Docker并配置国内源
2. 使用Docker启动MySQL

## MySQL配置

- 版本：5.6
- 数据存放位置：`~/mysql-data`
- 端口：33061
- 账户：root
- 密码：11

## 已实验Linux发行版

http://mirrors.aliyun.com/ubuntu-releases/18.04.5/ubuntu-18.04.5-live-server-amd64.iso

按道理Linux通用

# 出现错误退出脚本

```shell
set -e 在"set -e"之后出现的代码，一旦出现了返回值非零，整个脚本就会立即退出
```

```shell
set -o pipefail 表示在管道连接的命令序列中，只要有任何一个命令返回非0值，则整个管道返回非0值，即使最后一个命令返回0
```

# 检查Docker是否安装

```shell
if ! type docker >/dev/null 2>&1; then
	echo 'ok'
else
	echo 'not found'
fi
```

## type命令

type命令用于显示要查找的命令的信息

![image.png](https://i.loli.net/2020/09/05/oLND7t53yEdwKzl.png)

## > /dev/null

`>`起到将屏幕上显示出输出的结果重定向到某个地方

`/dev/null`是一个特殊的文件，写入到它的内容都会被丢弃；如果尝试从该文件读取内容，那么什么也读取不到

## 2>&1

`>&` 是将流重定向到另一个文件描述符的语法

0为标准输入`stdin`，1为标准输出`stdout`和2为标准错误`stderr`

`2>`将`stderr`重定向到一个（未指定的）文件，追加`&1`将`stderr`重定向到`stdout`

# 安装docker

```shell
curl -fsSL https://get.docker.com | bash -s docker --mirror Aliyun
```

## curl

curl是一个利用URL规则在命令行下工作的文件传输工具，简单来说这里用到的功能就是请求URL

参数：

`-f`：连接失败时不显示http错误，通常情况下，当HTTP服务器请求失败后，它会返回一个HTML内容，说明其内容（通常还说明了为什么等），将阻止curl输出该错误并返回22错误码。

`-s`：静默模式。不显示进度或错误消息

`-S`：与`-s`一起使用时，如果curl失败，它将使curl显示一条错误消息

`-L`：如果服务器报告的内容已移至其他位置，此选项将使curl重新执行以访问新的地址

## | 管道

将前一个命令的标准输出通过管道连接到命令2的标准输入， 管道中的每个命令都作为子进程执行

## bash -s

如果存在`-s`选项，或者在选项处理后没有剩余参数，那么bash将从标准输入中读取命令

# 添加内容到文件结尾

```shell
sudo tee <<EOF  /etc/docker/daemon.json > /dev/null
    {
        "registry-mirrors": [
            "https://1nj0zren.mirror.aliyuncs.com",
            "https://docker.mirrors.ustc.edu.cn",
            "http://f1361db2.m.daocloud.io",
            "https://registry.docker-cn.com"
        ]
    }
EOF
```

## tee

会从标准输入设备读取数据，将其内容输出到标准输出设备，同时保存成文件

## << EOF

`EOF`是`END Of File`的缩写，表示自定义终止符

既然自定义，那么`EOF`就不是固定的，可以随意设置别名，在`Linux`按`ctrl-d`就代表`EOF`

其用法如下:

```shell
<<EOF     //开始
....
EOF      //结束
```

还可以自定义，比如自定义：

```shell
<<BBB     //开始
....
BBB       //结束
```

`<<`：从标准输入中读入，直到遇到分界符停止