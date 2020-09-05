#!/bin/bash
set -e 
set -o pipefail

docker_china_mirror()
{
    sudo mkdir -p /etc/docker
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
    sudo systemctl daemon-reload
    sudo systemctl restart docker
}

check_docker()
{
    if ! type docker >/dev/null 2>&1; then
        echo 'install docker'
        curl -fsSL https://get.docker.com | bash -s docker --mirror Aliyun
        docker_china_mirror
        echo 'installed docker successfully'
    else
    	echo 'docker has been installed'
    fi
}

install_mysql()
{
    echo 'install mysql'
    sudo docker pull mysql:5.6
    sudo docker run --name docker-mysql -p 33061:3306 -v ~/mysql-data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=11 -d mysql:5.6
    echo 'installed mysql successfully'
    echo 'mysql config: '
    echo 'user:root password:11 port:33060'
}


check_docker
install_mysql
