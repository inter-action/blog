## ?
这里, 记录下我开发用的工具, 解释为什么。方便以后查阅。


## wireshark 抓包
wireshark, 也许我该用charles, 但是考虑到需要购买授权。我就还是算了。charles能提供的功能我想wireshark都能
提供。唯一的缺陷就是wireshark貌似有些复杂, 但是意味着更加强大。这个程序员的鄙视链是不, 复杂才是好东西, 
装的一手好B。

wireshark 抓本机网络没啥意思就不写了。

能抱怨吗？后端都是爷，你得明确定位问题，后端才能抬出他的手去改！

wireshark 抓 iphone 的流量:
* https://dingtwo.github.io/2016/03/17/mac%E4%B8%8B%E4%BD%BF%E7%94%A8WireShark%E6%8A%93%E5%8C%85/
* https://ask.wireshark.org/questions/17559/packet-capturing-application-for-the-iphone


      * Connect your iOS device to your Mac via USB.
      * Get the UDID for the connected device from iTunes or organiser.
        or open xcode, window->devices, grab uuid there 
      * Open terminal in your Mac
      * type the following commands in the terminal:
        // First get the current list of interfaces.
        $ ifconfig -l 

        // Then run the tool with the UDID of the device.
        // This adds a new virtual network interface rvi0.
        $ rvictl -s <udid>

        // Get the list of interfaces again, and you can see the new virtual network interface,
        // rvi0, added by the previous command.
        $ ifconfig -l 

      启动wireshark， 监听rvi0这个网卡的请求就可以了。


wireshark 抓 chrome https 请求:

这个也是可以的, 忘了自己去搜！然后再补在这里。























