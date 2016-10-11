
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
Type Class 的主要目的是为了对现有Type进行行为上的扩展, 通过将现有Type动态转换成另一个Type Class
实现来达到ad-hoc type polymorphism（type的多态）。当然用传统的 implicit conversion 也可以扩展。

但是implicit有3个根本的问题：
  1.搜索方式的规则复杂
  2.只是根据 type 去区分, 很容易两个同样的 implicit 冲突
  3.没有定义 implicit conversion 的 target 对象的行为, 别人方便没法扩展

这就是 Type class 处理的问题

创建type class的规则就是
* 定义好对应type的trait行为
* 创建不同的type的implicit object实现继承这个 trait.
* import 对应的implicit object.


http://www.cakesolutions.net/teamblogs/demystifying-implicits-and-typeclasses-in-scala
http://danielwestheide.com/blog/2013/02/06/the-neophytes-guide-to-scala-part-12-type-classes.html


! http://stackoverflow.com/questions/5408861/what-are-type-classes-in-scala-useful-for
! http://eed3si9n.com/learning-scalaz/a+Yes-No+typeclass.html

type class 定义了装饰在不同type上的行为, 达到动态添加行为的目的, dynamically lift some types。

```scala
trait Addable[T] {
  def zero: T
  def append(a: T, b: T): T
}

implicit object IntIsAddable extends Addable[Int] {
  def zero = 0
  def append(a: Int, b: Int) = a + b
}

def sum[T](xs: List[T])(implicit addable: Addable[T]) =
  xs.FoldLeft(addable.zero)(addable.append)

```



## Lifting

    trait Lift[F] {
      def apply(value: F):(Int, Int)
    }

    case class Point(x: Int, y: Int)
    object Point{
      implicit def liftable = new Lift[Point] {
        override def apply(value: Point) = (value.x, value.y)
      }
    }

    def testLift[T](v: T)(implicit conv: Lift[T]) = {
      val result: (Int, Int) = conv(v)
      println("after lift", result)
    }

    testLift(Point(1, 2))



## Procedure Recording
    

1. Simple recording: https://github.com/47deg/macroid/blob/master/macroid-core/src/main/scala/macroid/ToastDsl.scala

    // Loaf is an operation mutate Toast
    case class Loaf(f: Toast ⇒ Unit) {
      def apply(t: Toast) = f(t)
    }

    val loaf = Loaf(_.setDuration(Toast.LENGTH_LONG))
    loaf(new Toast(getActivity()))

2. More complex one, with flatMap
    
    














