

## nmap

https://nmap.org/man/zh/

用于网段, 服务器端口扫描

* install:

  ```
  //mac
  brew install nmap
  ```

* synatx:

  ```
  nmap [ <扫描类型> ...] [ <选项> ] { <扫描目标说明> }
  ```

* usage:

  ```
  // 扫描特定主机的端口
  nmap <host | ip_address>

  // 扫描 对应网段 范围，注意加sudo，要不然没有mac信息
  nmap -sn 192.168.1.0/24
  ```

* 状态解释:

  > 状态可能是 open(开放的)，filtered(被过滤的)， closed(关闭的)，或者unfiltered(未被过滤的)。 Open(开放的)意味着目标机器上的应用程序正在该端口监听连接/报文。 filtered(被过滤的) 意味着防火墙，过滤器或者其它网络障碍阻止了该端口被访问，Nmap无法得知 它是 open(开放的) 还是 closed(关闭的)。 closed(关闭的) 端口没有应用程序在它上面监听，但是他们随时可能开放。 当端口对Nmap的探测做出响应，但是Nmap无法确定它们是关闭还是开放时，这些端口就被认为是 unfiltered(未被过滤的) 如果Nmap报告状态组合 open|filtered 和 closed|filtered时，那说明Nmap无法确定该端口处于两个状态中的哪一个状态。


* options:
  
    https://nmap.org/man/zh/man-briefoptions.html

* links

  * https://nmap.org/man/zh/man-target-specification.html
  * 