# Docker 在前端的应用

## Docker
2015年可能是react爆炸的一年, 同样也是容器技术Docker火热的一年。Docker本质是一个轻量级的容器。你可以
理解一个轻量版的VM，它构建在虚拟化技术之上，内部是一个Linux系统。当然不只这些，它真正让服务器部署
和开发的环境无缝衔接起来，使得易用的批量自动部署成为肯能。如今Docker基本统治着云端部署。

Docker 组件核心分为

* Docker client:  
    Docker client 一般是命令行, 用于和 Docker Engine 交互  

* Docker Engine:  
    最核心的模块，它管理启动每个Docker实例  
    核心知识点
    * volumes
    * network

* Docker Compose   
    可以将各个Docker整合起来，比如一个简单的web应用可以拆分HaProxy/Nginx, 中间件(express or spring mvc),
    redis,数据库, 这些可以分拆成4个docker示例组成一个Docker Compose。然后你便可以统一管理这个compose组合
    的产物。

* Docker Swarm...etc  
    没研究过 :(


## 如何应用到前端
这受启发于今天我看的一个[视频](https://www.youtube.com/watch?v=zcSbOl8DYXM). 本质上来说, 我也是为了巩固
自己学习的知识才想办法应用的，但这并不影响事情的核心。我手头上的项目，需要启动nginx，处理一些服务端模板(SSI module)
，和接口的代理(proxy_pass)。这个项目有alpha和beta环境, 项目交接的时候，总是需要copy一份nginx配置文件给下一个人。
当然项目中放份拷贝也是可以的。但是，作为一个懒码农来说, 能自动化的都自动化当然是最好的选择。好处是下一个人启动这个项目
不必再装nginx了, 当然docker你是要装滴 :) 

so, 这篇文章就讲如何部署beta, alpha, 开发和伪生产不同环境的前段项目的一个解决方案的敲门砖。


## Bunch of .... codes !


### 前置条件:

#### 安装 dnsmasq
dnsmasq可以让你的本机成为一个dns服务器, 这样可以免于你每次新增项目都在host文件增加条目了。
我们要做的呢就是吧所有以 .dev 结尾的域名都回路到本机去


[Mac 下安装 dnsmasq 来配置开发环境](https://www.goodspb.net/mac-%E4%B8%8B%E5%AE%89%E8%A3%85-dnsmasq-%E6%9D%A5%E9%85%8D%E7%BD%AE%E5%BC%80%E5%8F%91%E7%8E%AF%E5%A2%83/)
[The Perfect Web Development Environment for Your New Mac](https://mallinson.ca/osx-web-development/)

`>brew install dnsmasq`
`>cp /usr/local/opt/dnsmasq/dnsmasq.conf.example /usr/local/etc/dnsmasq.conf`

open `/usr/local/etc/dnsmasq.conf` in your editor, search for `address=`, and append following content to below

    # this tells dnsmasq to route all dns request ends with `.dev` to localhost
    address=/.dev/127.0.0.1

create dns resolver, this make localhost a dns server

    sudo mkdir /etc/resolver
    sudo bash -c 'echo "nameserver 127.0.0.1" > /etc/resolver/dev'

start dnsmasq

    > cd /usr/local/Cellar/dnsmasq/2.76/sbin
    > ./dnsmasq

now test it:

    >dig makedup.dev @127.0.0.1

    # should response like, if not that means you didnt set it up properly

    ;; QUESTION SECTION:
    ;test.dev.     			IN     	A

    ;; ANSWER SECTION:
    test.dev.      		0      	IN     	A      	127.0.0.1


#### 安装 Docker
下载Docker for mac 注意不是 Docker toolbox !

[installation on mac](https://docs.docker.com/engine/installation/mac/)



### 正文开始了... ！ 


create docker-compose.yml file

    > touch docker-compose.yml

docker-compose.yml

    version: '2'
    services:
      nginx-dev-alpha: #创建的一个名为nginx-dev-alpha服务
        # official image doesnt include SSI module :(, bugged me for a while to find out this fact
        # 这里如果你不用SSI module，建议你用docker hub中官方的module替代
        image: dperson/nginx
        volumes: #创建共享文件夹, 宿主<->docker实例内部的共享， 以:号分割, 宿主更改会docker实例内部也会更改, vice versa
         - ./docker-build/nginx.alpha.conf:/etc/nginx/conf.d/mysite.conf #nginx.alpha.conf 是这个项目的nginx配置, 会在下面粘贴出来
         - ./app:/htmls # 
        ports:
         - "80:80" #宿主和docker实例的接口映射
        command: /bin/bash -c "nginx -g 'daemon off;'"


/docker-build/nginx.alpha.conf

    server {
        listen       80;
        server_name  open.alpha.dev;

        location / {
            root /htmls;
            index  index.html index.htm;

            ssi on;
            ssi_silent_errors on;
            ssi_types text/shtml;
        }

        location /anubis/v1 {
            proxy_pass http://maked-up.com:8080; # ！替换你自己的中间件ip地址
        }

    }

 
这里我本就创建了一个dev的alpha环境的nginx实例, 以此类推, 你可以创建更多的类似的nginx Docker 实例(docker compose service)

启动(在 docker-compose.yml 目录下):

    > docker-compose up

关闭:

    > docker-compose stop # or ctrl/cmd+c


打开浏览器 http://open.alpha.dev,应该就能看到你创建的实例结果。


### 后语
有的前段项目需要前段编译, 并不能简单用nginx做代理, 这个理论上也可以解决。
首先你需要创建一个 Dockerfile, 添加到docker compose file 的service中, 这里有几个重点

* 创建 shrinkwrap.json. npm shrinkwrap。这步是充分利用docker的cache机制。
* 创建 docker 实例中的以 node_module 目录的 named volumn。这步是为了不同步宿主机器和docker实例的node binary.
比如mac和linux就是不同的内核，binary会彼此不兼容。
* 在Dockerfile写入对应的开发命令， 比如 npm run dev。

这部分, 有机会我会将对应代码贴出来。

最后, 上面的代码可以写的更好, 将 nginx_host, nginx_port通过环境变量的形式写到docker-compose.yml文件中,
更加明显和易于维护。

核心的代码

docker-compose.yml：

    environment:
     - NGINX_HOST=open.alpha.dev
     - NGINX_PORT=80
     - SSI=y
    # about envsubst, view http://stackoverflow.com/questions/14155596/how-to-substitute-shell-variables-in-complex-text-files
    command: /bin/bash -c "envsubst < /etc/nginx/conf.d/mysite.template > /etc/nginx/conf.d/default.conf && nginx -g 'daemon off;'"


/docker-build/nginx.alpha.conf


    server {
        #listen       ${NGINX_PORT}; # 这里的变量会经docker-compose.yml对应的environment替换掉, 核心是 envsubst 工具
        #server_name  ${NGINX_HOST};

        ....





















































