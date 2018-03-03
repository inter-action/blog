

## chapter 1: Security


summary:
this chapter is about security set up, which consist of 4 parts
* right way to set up ssh
    * enable ssh login by rsa keys
    * disable root remote login
    * ease the login process with ssh conf file
    * assign user with right privilege
* setting up firewall with `iptables`
* config fail2ban that bans too many illegal ssh login attempt
* make server automatically install security patchs



notes:
* iptalbes:
    * [Linux Firewall Tutorial: IPTables Tables, Chains, Rules Fundamentals](https://www.thegeekstuff.com/2011/01/iptables-fundamentals/)
    * iptables -have-> multiple tables -have-> multiple rules
    * >By default, iptables does not save firewall rules after a reboot,

* iptable modules:
    * 提供基本的filter之上的扩展功能
    * http://ipset.netfilter.org/iptables-extensions.man.html
* ufw - firewall/iptables alternative:
  * https://www.digitalocean.com/community/tutorials/how-to-setup-a-firewall-with-ufw-on-an-ubuntu-and-debian-cloud-server  

links:
* [iptables](https://wiki.archlinux.org/index.php/iptables)
* [tee](http://codingstandards.iteye.com/blog/833695)



## chapter 2: Package Manager - Apt

## chapter 3: Permissions and User Management
summary:
all about creating users & how to properly set up permissions.

notes:

* create a `www-data` user:

    ```shell
    # Create directory as root
    sudo mkdir /var/www

    # Change to owner/group www-data:
    sudo chown www-data:www-data /var/www
    sudo chmod ug=rwx /var/www
    sudo chmod o=rx /var/www

    sudo adduser deployer
    sudo usermod -a -G www-data deployer
    sudo su - deployer
    # be sure to set this (umask) in .bashrc or .bash_profile
    umask 002

    sudo chgrp www-data /var/www # change group to www-data
    sudo chmod g+s /var/www # Set group id bit of directory /var/www, see below
    ```

* `chomo g+s` means:
    * 就是说任何用户在该文件夹下创建的文件优先继承这个文件夹下的group的权限设定, 其他user or others权限不受影响
    * https://unix.stackexchange.com/questions/182212/chmod-gs-command


## chapter 4: Webservers

summary:
* how web server works in general way
* how to properly config nginx
    * 
* other php-related thing



## chapter 5: SSL Certificates

summary:
* how to create self-signed SSL Certificates & how to install them on nginx



## chapter 6: Multi-Server Environments

summary:
* challeges when face with mulit-server environments
    * assets, log, session, load blancing,  ssl termination, etc... how these should be managed

* load blancing with nginx & HaProxy

## chapter 7: Web Cache
summary:
* give a general notes how cache are categorized 
    * object cache 
        * redis or memcache
    * http cache

* give a general introduction how http cache works & categorized
    * validation cache
    * expiration cache

* give a brief intro how to implement http cache through
    * nginx
    * varnish



notes:

* http cache:
    * validation cache:
        * E-Tag:
            * The response contains an ETag. ETags are generated based on the contents of a file. If a file changes,
            the ETag returned by the origin server will be changed.

    * expiration cache:
        * Modern browsers will prefer expiration caching. Validation caching is typically used if no expiration information is present. 

* cache architecture
    * Load Balancer -> Cache Server -> Origin Server (backend server)

* 


## chapter 8: Logs

summary:
* what's the main responsiblities that a log system should meet
* introduction to Logrotate & Rsyslog


notes:
* Logrotate:
    * help rotate your logs

* Rsyslog
    * help centerlize your logs


## chapter 9: File Management, Deployment & Configuration Management

summary:
* manually sync files with rsync tool
* auto deployment with github hook & bash scripts
* Ansible for auto deployment

links:
* [Ansible中文权威指南](https://ansible-tran.readthedocs.io/en/latest/)

## chapter 10: SSH

summary:
* how to properly config ssh
* what ssh tunneling is



notes:
* ssh tunneling:
    * 这个东西是一个很好的特性, 可以通过ssh 管道将本地端口和远程端口打通. 可以访问到正常情况防火墙屏蔽的端口, 比如mysql数据库3306端口
    * Local Port Forwarding: 将本机端口和远程服务的端口map上
    * Remote Port Forwarding: 在远程server端将端口和目标端口映射上, 然后可以让远程的client端通过映射过的ip+端口, 访问到目标端口的服务


## chapter 11: Monitoring Processes
summary:

* System Level
    * System V Init (这个书中就直接跳过了)
    * Upstart (老的版本linux会用到这个)
    * Systemd (新版本的linux会用到这个工具)

* Software Level

    * supervisor, (this is cool)
    * forever (written in js)
    
    


notes:
* 需要注意的是不论使用upstart还是systemd, 两种方式都不建议在最终的执行的 shell 及脚本以daemon方式执行
    否则无法实现监控, 也不能正常启动

* 优先级
    * Systemd -> Upstart -> SysV 




links:
[Getting started with systemd](https://coreos.com/os/docs/latest/getting-started-with-systemd.html)



## chapter 12: Development and Servers






# todo:
Logrotate
Rsyslog
ansible














