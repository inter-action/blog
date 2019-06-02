
# attack network protocal
这本书也算是质量上乘的一本书了, 这本书的组织结构非常适合一个刚刚看完cs的小白, 涵盖了网络,操作系统,网络分析软件, 加密算法的各个方面.




## CHAPTER 1: 

## chapter 03:
第三章讲解了各种 binary encoding / text encoding 的方式. 这让我对这方面有了豁然开朗的感觉. 以前我从来没有从这种维度去看待计算机encoding的方式.

* 常见的binary encoding 比如 signed number, float number. text encoding 的代表有xml, json, base64.

todos:
* abstract syntax notation 1 (ASN.1) 需要再看下, 没看懂.

## chapter 04: advanced application traffic capture

* NAT 分成两种
    * SNAT, DNAT, SNAT 改变ip 中的source地址, SNAT 在把packet send back 的过程中判断destination & port numer 来映射回去的链接.
    * DNAT改版IP packet 中的destination地址. NAT 是在router中启用的, 

用DNAT 去做proxy 的network forward还是...诡异的没法很好理解呢, 虽然明白是怎么说的.

而且在启用的时候我看还需要用iptable进行路由更改

* DCHP spoofing: 伪装成DCHP server去发送假的DCHP response请求.
* 讲了ARP在Ethenet上的应用, 原理. 还有可以通过操作系统的`arp`命令去查看本机arp的状况


## chapter 05: analyze from the wire

* python analyze binary dump with unpack lib
* wireshark scripting using lua, creating a Dissector.

todo: 
* 最后的xor一节没有看懂


## chapter 06: application reverse engineering
memory alignment:....

register(寄存器):
* `EIP`: 表示下一次指令的位置(memory address of next instruction to be run)
* `ESP`: 控制stack push pop, 还有 subroutine 的 current memory location of the base of stack. `(啥意思?)  应该是stack的根基`


reverse static analysis:
* `IDA pro` dissemble binary file 
* `ISA`( instruction set architecture. )
* `debug sumbols`: 
    * 用来储存binary 的源码映射关系的文件. windows里边这个文件叫做`PDB(program database)`
    * in mac it's called `dSYM, debugging symbol package`
    * in linux, 所有的debug信息都存放在最终的ELF文件的debug section当中
* 加密算法里边会写magic number嵌入到binary中
* `Static Reverse Engineering`, `Dynamic Reverse Engineering`:
   * Static : 使用工具去查看ISA代码, 看看代码是如何执行的
   * Dynamic: 是用debugger工具去动态审查程序的运行, 去检查各个寄存器的值, 主要是 `ESP` `EIP` 还有 `General Purpose Registor` 的值去查看代码.

todo:
* 看下深入理解计算机操作系统里的这些信息

links:
ELF - http://refspecs.linuxbase.org/elf/elf.pdf


## chapter 07: network protocol security
* 密码学的两个应用: 加密和签名
* >we usally refer encrption  algorithms as cipher or code


#### 加密算法的分类:

* substiution cipher: 通过reference的1v1映射, 将plaintext和加密的char做mapping
* XOR encription: 主要利用两个字节异或的操作, 并利用异或的对称性来加密解密, 字节A和key B进行异或操作得到字节C, 并再次用C和B异或得到plaintext. 这种加密方式的缺陷就是如果要保证这种加密的安全就必须保证key的值完全random. 而且每个key值以为都需要一个random的key就会造成, key和加密的data是一个长度, 而且不能复用key, 因为一个被破解, 所有的都会被破解
* `加密中完全随机的重要性`: 加密的随机性保证了加密算法的强度, 计算机没法做到随机, 因为同样的输入, 计算机总会给出同样的输出. 真正随机就是读取自然中的信息, 这种方式的缺陷在于读取自然信息的速度太慢. `PRNGs( pseudorandom number generator)` 的算法是用 seed value 去生成一连串数字, 理论上来说如果不知道原始的seed value, 加密都是无法被破解的.

Symmetic Key Encription:
* Block Cipher: `AES DES`都输入这类, 这类算法的原理就是对data的size进行分割成n个bytes的block, 然后对每个block应用加密算法. 需要一个key.
   * block cipher 的 operation mode:
        * ECB(Electronic Code Block): 就是每个block去应用一次encription, 这种mode的问题就是data比距是n*block size, 没有提供校验机制, 意味着破坏者可以随意的shift 一个block size 到cipher text 中, 去破坏解密, 然后呢校验者还没法感知.
        * CBC (Cipher Block Chainning):  这种工作方式就是将前一个block和当前block进行XOR操作, 然后在用key去加密. 初始的时候没有前一个block, 所以需要一个随机生成的`IV(intialization vector)`
* Stream Cipher: 不知道是个啥? 没咋看懂


#### Asymmetic Key Cryptograhpy:
* p &amp; q 两个数, 都是质数
* `e`: public exponent, 65537 一般是
* n: modulus, 是 p&amp;q的product
* d: privte exponent , 用于解密

RSA 加密非常耗费计算能力, 一般方式是通过AES加密, 然后通过RSA把key传给解密方.

* `Diffie-Hellman key exchange`: 是一种交换shared key的一种方式, 涉及到两方在不reveal彼此private key的前提下, 交换key. 具体还是要看194页的描述, 回头看看下吧.
* `cryptographic hashing algorithms`: 用于保证data integrity.
* `MAC(Message Authentication Codes)`: 是一种采用对称加密做签名的方式, 
* `HMAC(Hashed Message Authentication Codes)`: 

#### X.509 Certificates:
* it 是现代浏览器证书校验的底层机制, 理解它是理解浏览器证书验证的基础.
* 讲了`X.509 Certificates`的工作机制, 证书作用用途分为 `Code Signing Certificate` 和`Web Server Certificate`
* 证书是有层级的, 父级可以sign子级的证书
* `basic constraint` flag: 是用来指示证书的类型, 可不可以act as a CA
* `key usage`:  标识这个certificate的用途,  是 `Code Signing` 还是 `Web Server`
* `Certificate revocation list`: CA用来解除证书的有效性的机制
* `Root Certificate`: root certificate 不能被其他 certificate sign.


#### TLS(Transport Layer Security):

* 可以说TLS是对本章讲解的所有topic都有cover, 所以回去后要好好看下todo里面的内容.
* 先协商symmetic croptograhpic的key
* 然后server 端把证书和any itimidiate certificates send 给client.
* 客户端从上到下验证server的certificate的合法性.



todo: 
* Diffie-Hellman key exchange.
* HMAC
* MAC




## 

