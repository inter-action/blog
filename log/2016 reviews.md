我会在这里分享一些个人的事情，尽可能的不涉及个人的情感和隐私


### 2016.08.05 - 00:03

8月, 对于时间的感觉已经从流动变成静止了。只是日历上的数字增加。
anyway, 最近的状况就是, 原本定的限制自己的·繁杂兴趣·变成了空话。除了 *scalatra in action* 最近又看了
*grokking algorightms*, *linux in a nutshell*, rustlang, typescript, koajs.
比较坑的就是读 scalatra 这本书了, scala 语言学习了近2年, 一直没有时间和动力去好好应用下，突然看到这本书
我就想这我看完了是不是能做个web项目呢，就几天兴致勃勃的看完了。但是当我装上了giter8试图去跑个demo的时候，
这sbt从maven仓库拉lib的速度简直是让我蛋疼的无以复加, 加上这几年前端的历练让我对IntelliJ这IDE越来越蛋疼。
但主要还是lib下载太慢，加上配置oss的maven镜像不起作用。我决定放弃Java做web这条路。我决定投身koajs来弄。
这几年nodejs社区的火热，让nodejs的库质量非常高，加上nodejs这种异步特性, 不做计算密集的应用还是非常有优势的，
而且koa的demo也很多，不像scalatra我搜都搜不到。所以我决心投身nodejs web开发了。但是js本身呢，有些问题，
弱类型的语言，导致很多问题不能在编译期间发现，还好typescript这门语言这几年的火热，让我看到了node环境中应用
强类型语言的曙光。我自己也比较过scalajs, 我觉得scalajs哪方面都应该比typescript要吊，无奈社区没有typescript
强大(导致好多type defination没有typescript全)，加上typescript是javascript超集，这两个特性让我决定还
是好好学习应用typescript来做nodejs应用。

Java的优势我觉得在于生态的强大, 加上多年在服务端多年的应用，让它应用场景很多, 所以Java我还是会持续关注的，
只不过目前的重心会倾斜些。

还有最近对rustlang的兴趣, 我一直希望掌握一个接近native的语言, c, c++这么复杂的我就直接放弃了，golang 学习初期感觉
很好但是槽点也不少, 蛋疼的大小写约定, 代码怎么看都丑, 没有统一的包管理, 除了提供些 goroutine 这些貌似看不到目前
现代高级语言的特征... 所以最近就看上rust了, 还在学, 不知道这条路到后期回报会怎么样。 不过目前感觉良好，至少感觉没有
scala那么复杂，看着就头疼的样子.. :(

android 开发最近是停滞了, 应用了一段 macroid 库编写 android 应用（搭界面）, 槽点也很多, 首先感觉好多类型不知怎么就
check不到了 。而且前期需要写大量的Tweak(虽然 47seg 提供了对应的Tweak扩展，但我不喜欢用)。 主要还是编译, 我看我同事
用java写的应用最新google出的instant run 不知有多爽。 反正打算看看 android 上应用scala的别的路吧。最后还是要夸奖
下 macroid 库本身写的确实吊。代码有时间我还要再研习下。

不得不说今年是我读源码最多的一年，看了好多别人写的库的源码, scala, react-native, golang, nodejs的都有些。
而且好多点都连在了一起，让我这个想成为全栈的选手有增添了些信心。初步有想法做一个简单的web应用, 把我学的这些都串一串。
我觉得目前这个应用可以涉及的点可以有:

* koajs 作为项目的基础base, 源码由typescript编写
* reactjs 作为前端的UI，界面完全和后端分离
* 加入服务端reactjs渲染
* 用户验证&创建流程(jwt)
* mysql，基础测试框架整合, logging, 系统检测 (DevOps...)
* nodejs 和 java 或者其他语言的交互, 通过http协议或者google的protobuf协议, 或者初期就用linux的process socket也行
  好处就是我可以用java写些lucene的搜索。
* 加入 redis .
* 整合 docker, 用docker 整合mysql, redis, nginx ...

自己工作的事情，就密到evernote里吧，槽点永远比G点多。

自己生活方面, 由于15年到现在一直在看技术书籍，没有时间和心情看些杂七杂八的生活类书籍，我觉得这样也不是很好的状态。
人嘛，还是要把生活活明白些才算成功，未来我会汲取些其他非技术方面的知识。


### 2016.07.01
不知不觉已经来到7月份了, 时间真是过得好快。近2个月的快节奏工作，疲惫感到现在仍没有消除。
最近断断续续地持续着安卓开发的学习, macroid框架从刚开始的混乱无序，到渐渐可以使用。不得不说, android开发
远没有web开发的规则那么一致顺手。除了一些quirk情况外, 频繁的代码编译安装就让我不可耐烦。我不知道目前google 
新出的instant run 现在情况如何。不论怎么说, 我觉得react-native 开发的方式将是我非常喜欢的。我想等我
尝试做一个原生的app之后, 打算切到react-natvie上, 加以 gradle+scala compiler plugin来开发android 应用。

web端技术，从我3月份入职以来到现在的收获，就是熟练了些 reactjs 。和到最近的RxJS。尤其是RxJS，在我看过
functional programming in scala 一书之后, 觉得里边的api倍感亲切。又有点被连在了一起的感觉。RxJS虽然
目前只是入门，有时间或者机会的话，我会继续深入的了解下。不管是在web端还是node端的应用。TypeScript是我当前
想要了解的兴趣点之一, 对于Js这种基本都是字符串的语言。我在开发的时候如果没有lint工具会蛋疼很多，但在eslint
等工具的帮助下，依然有一些对象的属性引用出错, 方法名写错的尴尬。对于TypeScript这种目前看似门槛很低的语言，我觉得
只要花些简单的时间, 不论实在web前端还是在node应用后端，都会有着不错的应用。

后端技术,我到现在也没有成功的尝试过什么,也许我会找机会和时间点处理下这方面的问题。

要学的东西好多, 然而我会到哪里停下...或者找到真正的方向，抑或是没有方向...

### 2016.04.21 - 20:17
貌似只有晚上下班之前可以静下来写些个人的东西, 这几年技术方面看的杂乱不精。什么都想看，分散了太多精力去
深入的做些事情。所以今年的计划尽可能的限制下自己的习惯。

首先, 由于去年的 *functional programming in scala* 这本著名 **red book** 勾起了我对 scala 语言
的兴趣。但这本书太难啃，囫囵吞枣的啃到差最后一章。所以今年要把最后一章读完。然后看情况在今年或者明年再重读
一遍(yeah it will be frustrating)

一直在看 scala 没有真正的好好应用过, 所以几年会用scala写个android app(初步打算做一个rss阅读器), 
争取能部署到应用商店上。

其次, 今年会抽时间把 *OReilly.Beautiful.JavaScript* 这本书读完, 今年不会再看其他技术方向的书籍
（至少不会深入的看）然后会试着读读 *Pearls of Functional Algorithm Design* (不知道能不能看下去
，书中的标注太多了 不见得能搞定)。主要的想法就是专注下算法，不再追求些乱七杂八的书籍了。

明年? 明年会来的很快






