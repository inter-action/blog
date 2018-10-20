# [核心网络协议袖珍手册].(Packet.Guide.to.Core.Network.Protocols).Bruce.Hartpence.文字版.pdf

* 这本书可以算是我读过的比较好的书的一种了，内容可以说是很全面也很有深度。
* codes - git@github.com:CrazyFork/low-level-programming.git

# Denotes


under links context:

    #   //in links context, # denotes segment of the page
    !   //in links context, ! denotes an important link
    ?   // not fully understand yet, todo

## chapter 2: Ethernet

这章讲了Ethernet的构成, 作为网络模型的第二层 (Data Link), 还有对这个协议比较重要的MAC 地址的讲解. 还有Physical Layer的讲解, 说实在的讲的是RJ45的网络插口的构成, 还有对应cable线的组成, 以及对传输速率的影响. 还有每根线的作用, 还有如何组建Ethernet的合理方式, 局限等. 我是没有特别看懂这部分


MAC:
* > MAC addresses have no significance beyond a computer’s own network, so the MAC addresses of machines beyond the local network are unknown
* > Generally, MAC addresses are divided into two parts: a three-byte vendor code and a three-byte host ID
* There are three different types of MAC addresses on an Ethernet network: unicast, broadcast, and multicast.


headers:
* Control Field (Type), what payload type are (ARP or IP)

special address:
ff:ff:ff:ff:ff:ff, broadcast address to the current netork ofc.

## chapter 3: Internet Protocol
这章主要讲IP协议部分. 包括协议的构成.

IP header:
* Identification: this field determines wether two fragments are from same data unit.
* Fragment Offset: used with ID field, to determine how to assemble this fragement.
* Time to live: the maxium hops allow for total router, everytime the packet got transmitted through a router, this field is substract by 1.



Addressing:
ip 地址氛围 prefix 和 suffix 两部分, 前面一部分为 network id, 后面为该 network id 下的该host的唯一地址. prefix 和 suffix 的分隔是通过 mask (子网掩码) 于该ip地址取&得到的.


Reserved IP addresses:
* 129.21.255.255, suffix id 都以为1, 这种情况是 Broadcast packet to a particular network, 就是外部network对内部的广播
* All ones, 255.255.255.255, Broadcast packet to the current network




## chapter 4: Address Resolution Protocol
这章主要讲, ARP, 一种local network用来找到对应主机的mac地址的. ARP 总共有请求和返回两种消息, 请求用了MAC地址broadcast(广播), 谁拥有这个ip地址就用unicast(单播)回复.

注意IPV6 没有包括ARP协议, 用了不同方式的实现.

* ARP poisoning: 是一种攻击的方式, 攻击者伪装成回复ARP协议对象, 让所有的请求都打到攻击者机器那里, 然后攻击者再将请求转发的正确的接受者那里. 是一种 Man in the middle 形式的attack



## chapter 5: Network Equipment

这章主要讲了network设备的内容 switch, hub, router, AP 等:

我的理解是

hub 操作在 physical 那层, 会broadcast 0 and 1 to all the ports.
switch 操作在 link 层, 会根据 SAT (将MAC address to port的table) 来转发 LAN 请求. 在传递的过程中会修改 MAC 的 SRC ADDR .
AP 也是类似于switch, 但是具体协议有些不同, 而且是wireless
router 是操作在 network 层, 即 IP 协议层, 会修改IP的地址和头.



## chapter 6: Internet Control Message Protocol

这章主要讲了ICMP协议. 这个协议的目的在于网络中节点的沟通和调试. 这个协议包含的主要type类型为:
* echo / reply: debug. ping 程序会用到
* redirect: router 发现自己不能处理这个请求之后, 会发送给 initiatied host 一个 ICMP redrect, 告诉host你去访问这个router. 然后host会将这个信息加在自己的dynamic routing table 中.
* Destination Unreachable: 网络不通. 中间有节点不能正确的到达目的地. 会将原始的请求通过payload带给initiatied host.


还讲了 TTL 的作用, traceroute 程序用ICMP TTL header来判断网络正常的边界. 通过发送多次请求, 并 increment TTL field by 1. 来判断网络的问题.
中间涉及的NAT协议没有特别懂. traceroute 也需要再看看.

## chapter 7: Subnetting and Other Masking Acrobatics

这章主要讲了supernetting, subnetting 的方式

* supernetting 是通过自网络去计算 gateway 的 ip 和 mask 的
    * 计算network id, 找到相同的bits(两个子网络),和不同的bits, 不同的bits全部置位0, 与相同的bits组成最终的network id
    * 计算mask, 相同的bits置位1, 不同的bits置位0, 组成最终的mask
* subnetting 是一个相反的过程 ()


links:
* [Supernetting Explained Very Well](https://www.youtube.com/watch?v=N938u4_Jfc4)
* [Computer Networks Lecture 9 -- Supernetting or aggregation](https://www.youtube.com/watch?v=MnqP_TVwkbs)


# Links

* [Multicast address](https://en.wikipedia.org/wiki/Multicast_address)

# todos
* https://www.youtube.com/watch?v=MnqP_TVwkbs 这个关于computer network的视频需要看下