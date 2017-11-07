## ?
这里, 记录下我开发用的工具, 方便以后查阅。


## wireshark 抓包
wireshark, 也许我该用charles, 但是考虑到需要购买授权。我就还是算了。charles能提供的功能我想wireshark都能
提供。唯一的缺陷就是wireshark貌似有些复杂, (edit： 并不能很方便的抓取https流量). 

wireshark 抓 iphone 的流量:
* https://dingtwo.github.io/2016/03/17/mac%E4%B8%8B%E4%BD%BF%E7%94%A8WireShark%E6%8A%93%E5%8C%85/
* https://ask.wireshark.org/questions/17559/packet-capturing-application-for-the-iphone


* Connect your iOS device to your Mac via USB.
* Get the UDID for the connected device from iTunes or organiser.
  or open xcode, window->devices, grab uuid there 
* Open terminal in your Mac
* type the following commands in the terminal:

  ```
  // First get the current list of interfaces.
  $ ifconfig -l 

  // Then run the tool with the UDID of the device.
  // This adds a new virtual network interface rvi0.
  $ rvictl -s <udid>

  // Get the list of interfaces again, and you can see the new virtual network interface,
  // rvi0, added by the previous command.
  $ ifconfig -l 

  ```

启动wireshark， 监听rvi0这个网卡的请求就可以了。


wireshark 抓 chrome https 请求:

这个也是可以的, 忘了自己去搜！然后再补在这里。


## mitmproxy 抓取ios https请求
mitmproxy 其实不止可以抓ios的请求的, 还可以抓正常的网络请求。这个到时候我会研究下。
这里说下如何用mitmproxy抓ios的https请求。这个主要来自这个链接 [intercepting-ios-traffic](http://jasdev.me/intercepting-ios-traffic)

1. `pip install mitmproxy`
2. 启动mitmproxy, `> mitmproxy`
3. 配置你iphone wifi的proxy, 你电脑的 ip + 8080 默认的端口号。注意, 你不能启用其他占用了8080端口号的
nginx或者webpack inline server。
 
4. 安装证书, 在iphone中输入这个地址(safari): mitm.it, 这步必须先启用mitmproxy。安装ios的证书就可以了。
然后你就可以在命令行中看拦截的网络请求了。

 






















