
## Dynamic Mixins, Enrich, Implicit Pimping

https://coderwall.com/p/k_1jzw/scala-s-pimp-my-library-pattern-example


    class BlingString(string: String) {
        def bling = "*" + string + "*"
    }

    implicit def blingYoString(string: String) = new BlingString(string)


    scala> "Let's get blinged out!".bling
    res0: java.lang.String = *Let's get blinged out!*

    // is equally

    implicit class BlingString(string: String){
        def bling = "*" + string + "*"
    }


## extends class with Constructor call

    //bm: extends class with Constructor call
    /** A LinearLayout that is preset to be vertical */
    class VerticalLinearLayout(ctx: Context) extends LinearLayout(ctx) {
      setOrientation(LinearLayout.VERTICAL)
    }


## ! Cake Pattern
简单来说这里定义了 两个 trait, 一个 trait 显式的制定了另一个 trait 的依赖. 
这个模式的好处：对代码更加的灵活的拆分. 据说是scala的DI的一种解决方式

http://www.cakesolutions.net/teamblogs/2011/12/19/cake-pattern-in-depth



## Type Class
Type Class 的主要目的是为了对现有对象进行行为上的扩展。当然用传统的 implicit conversion 也可以扩展。
但是implicit有3个根本的问题：
  1.搜索方式的规则复杂
  2.只是根据 type 去区分, 很容易两个同样的 implicit 冲突
  3.没有定义 implicit conversion 的 target 对象的行为, 别人方便没法扩展

这就是 Type class 处理的问题

http://www.cakesolutions.net/teamblogs/demystifying-implicits-and-typeclasses-in-scala
http://danielwestheide.com/blog/2013/02/06/the-neophytes-guide-to-scala-part-12-type-classes.html


