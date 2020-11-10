# Apress.Low-Level.Programming


## chapter 3: Remote connectivity: Safely accessing networked machines

```
pstree -p #ps in tree mode
```

-> init process and systemd
```
$ file /sbin/init
/sbin/init: symbolic link to /lib/systemd/systemd
```


## chapter 11: System monitoring: Working with log files

-> journalctl vs Syslogd
Syslogd 以前是负责管控log的, journalctl 是新的工具来管控log


```
journalctl -n 20
```

-> about syslogd
> With syslogd, the way messages are distributed is determined by the contents of the 50-default.conf file that lives in the /etc/rsyslog.d/ directory. 

```
# syslogd 文件目录
/var/log/syslog
```

`/etc/logrotate.conf` 文件可以控制日志如何被生成.


## chapter 12: Sharing data over a private network

## chapter 13: Troubleshooting system performance issues

