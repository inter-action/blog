## 开发日常的工具记录
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


# mitmproxy
## mitmproxy 抓取ios https请求
mitmproxy 其实不止可以抓ios的请求的, 还可以抓正常的网络请求。这个到时候我会研究下。
这里说下如何用mitmproxy抓ios的https请求。这个主要来自这个链接 [intercepting-ios-traffic](http://jasdev.me/intercepting-ios-traffic)

1. `pip install mitmproxy`
2. 启动mitmproxy, `> mitmproxy`
3. 配置你iphone wifi的proxy, 你电脑的 ip + 8080 默认的端口号。注意, 你不能启用其他占用了8080端口号的
nginx或者webpack inline server。
 
4. 安装证书, 在iphone中输入这个地址(safari): mitm.it, 这步必须先启用mitmproxy。安装ios的证书就可以了。
然后你就可以在命令行中看拦截的网络请求了。

## mitmproxy 给应用加上https, 通过reverse proxy 模式

```shell
#对应的需要修改host machine的 hosts 文件或者通过自定义的dns服务解析这个配置的域名.
sudo mitmdump -p 443 --mode reverse:http://minimart-crm.faas.alpha.elenet.me:9001/ --ssl-insecure

# 安装 mitmproxy 根证书, 
# 这个例子是在android的emulator装证书的例子, 通过文件上传. 当然也可以通过 mitm.it 来下载安装, 需要先启动mitmproxy代理
cd ~/.mitmproxy 
adb push mitmproxy-ca-cert.pem storage/self/primary


# 下面这两行是为了真机调试, 不是真机不用这个
mitmproxy --listen-port 8888 --ssl-insecure

# 配置host
127.0.0.1   minimart-crm.faas.alpha.elenet.me
```

## mitmproxy 代理mac机器的网络请求, 通过 socks mode
通过socks mode来代理电脑的网络请求, 一是非常强大, 第二配置简单. 当然也可以用浏览器插件来显式制定proxy server.

核心的链接: [Tracing All Network Machine Traffic Using MITMProxy for Mac OSX](https://blogs.msdn.microsoft.com/aaddevsup/2018/04/11/tracing-all-network-machine-traffic-using-mitmproxy-for-mac-osx/)

```shell
mitmproxy --mode socks5 --showhost
```

打开mac网络设置, 点击到你链接的Internet网络上, 点击Advanced-> Proxies -> toogle SOCKS Proxy -> 在 sockets proxy server 下输入 `127.0.0.1` 端口 `8080`


# homebrew
## homebrew 命令
```
brew info dnsmasq
brew list dnsmasq
```

## 将本机配置为dns服务器, 用dnsmasq

brew install dnsmasq

```
# 编辑文件 /usr/local/etc, 添加一行
# 表示所有以faas.alpha.elenet.me结尾的域名都解析到127.0.0.1上
address=/faas.alpha.elenet.me/127.0.0.1

#重启dnsmasq
sudo launchctl stop homebrew.mxcl.dnsmasq
sudo launchctl start homebrew.mxcl.dnsmasq

#test dnsmasq
dig minimart-crm.faas.alpha.elenet.me @127.0.0.1


# 绑定根域名解析
sudo mkdir -p /etc/resolver
# 注意这里的文件名一定要和你解析的域名根部相同, 可以为 `me` or `elenet.me` or `faas.alpha.elenet.me`
cd /etc/resolver && vim faas.alpha.elenet.me 


# 用 ping 来测试下
```

链接:
* [Using Dnsmasq For Local Development On MacOS](https://www.michaelpporter.com/2017/11/using-dnsmasq-for-local-development-on-macos/)


# android emulator 调试应用:

### 下载配置emulator

* install jdk 8, 高版本会报错:
    * https://www.oracle.com/technetwork/java/javase/downloads/jdk8-downloads-2133151.html

* 下载command line tools
    创建 $HOME/workspace/tool/android_sdk 文件夹, 并将解压完的 tools 文件夹移动到该下面

* 设置PATH:
    ```raw
    // ~/.zshrc
    export ANDROID_SDK_ROOT=$HOME/workspace/tools/android_sdk
    export PATH=$ANDROID_SDK_ROOT/tools:$ANDROID_SDK_ROOT/platform-tools:$PATH
    ```
*  sdkmanager下载镜像
    * [sdkmanager](https://developer.android.com/studio/command-line/sdkmanager)
    * `sdkmanager "system-images;android-25;google_apis;x86" "extras;intel;Hardware_Accelerated_Execution_Manager" "emulator" "platforms;android-27" "build-tools;28.0.3" "platform-tools"`

* avdmanager 创建虚拟机
    * [avdmanager](https://developer.android.com/studio/command-line/avdmanager)
    * `avdmanager create avd -n test -k "system-images;android-25;google_apis;x86" -d "pixel"`

    * 在虚拟机中关联键盘输入
        * vim `vim ~/.android/avd/test.avd/config.ini`, 其中 test.avd 的 test 是你创建avd的时候所指定的名称
            * add new line `hw.keyboard=yes`
        * https://stackoverflow.com/questions/11235370/android-emulator-doesnt-take-keyboard-input-sdk-tools-rev-20


* emulator 启动虚拟机
    * cd 到 $ANDROID_SDK_ROOT/emulator 文件夹下面, 执行 `./emulator @test  -writable-system -dns-server 127.0.0.1`



### 配置网络, 将 emulator 中的网络请求映射到本机的资源
虽然emulator和host machine共享同一个网络, 但是emulator并不会遵守host machine 配置的`/etc/hosts` 文件, 所以为了让android的网络请求映射到host machine 上, 有两种方式: 
* a) 修改 adnroid 的 `/system/etc/hosts` 文件 
* b) 启动的时候, 带上 dns-server 参数, 
* c) 记得关闭掉emulator中的4G/蜂窝网络的访问权限

需要看的核心内容就是这个链接, 去理解emulator中的网络配置 [Set up Android Emulator networking](https://developer.android.com/studio/run/emulator-networking)

#### a)
```shell
# 将hosts文件拉去下来

adb root 
adb remount
adb pull /system/etc/hosts

# 编辑拉去下来的hosts文件, 添加一行 `10.0.2.2	 minimart-crm.faas.alpha.elenet.me` 
# 10.0.2.2 是一个特殊的地址, 指向了emulator的宿主地址

adb push hosts /system/etc/hosts
adb reboot

```

#### b) 
这种方式需要通过dnsmasq将本机配置成dnsserver, 然后启动 emulator 的时候带上 dns-server 参数

### cmd tools

* adb 
  * 常用命令
    ```shell
    adb pull /system/etc/hosts
    adb push hosts /system/etc/hosts
    ```

  * links:
    * [Url mappings in hosts file in emulator is being ignored](https://stackoverflow.com/questions/33869775/url-mappings-in-hosts-file-in-emulator-is-being-ignored)
    * [Use modified hosts file on Android Emulator](https://medium.com/code-procedure-and-rants/use-modified-hosts-file-on-android-emulator-4f29f5d12ac1)
    [How to connect to my http://localhost web server from Android Emulator in Eclipse](https://stackoverflow.com/questions/5806220/how-to-connect-to-my-http-localhost-web-server-from-android-emulator-in-eclips)


### chrome 调试 emulator:

chrome dev tools -> click more options on the top right corner -> more tools ->  remote devices
