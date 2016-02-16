## Operators
http://nerd-is.in/2013-08/scala-learning-operators/

>结合性决定了当有一系列同一优先级的操作符时，它们的求值顺序。在Scala中，绝大部分的操作符都是左结合（从左到右）的，
除了：1. 以冒号:结尾的操作符；2. 赋值操作符。而且，在右结合时，二元操作符的参数顺序也发生了变化：二元操作符是其第二个参数的方法。



## Function vs Method

Function is instance extends FunctionX, while method is just a type, not a value or
anything so it cant be assign to a variable.

convert a method to a function instance:

    m _

define a method:

    def <function_name>(s: String): Int = s.toInt + i

define a function:

    val strToInt = (s: String) => s.toInt

this links explain it very well
[http://stackoverflow.com/questions/2529184/difference-between-method-and-function-in-scala](http://stackoverflow.com/questions/2529184/difference-between-method-and-function-in-scala)


## Complex Type System

We all know scala has a complex system. Most of the time we simply refer to its Object Type graph and
Covariant, Contra Variant ... But how to efficently use its type constructor is another thing you may need to
consider.

Let's look at the code from book `Functional Programming In Scala`. The `Id[A]` type contructor is a tricky one.
In order to implement the map method with traverse method, he contructs the `Id` type, and use it created
an `Monad[Id]`  (Since all Monads are Applicative).


    trait Traverse[F[_]] extends Functor[F] with Foldable[F] { self =>
      // 注意这个地方和书中的定义不同, implicit 和 M[_]:Applicative 是一个意思
      // 但是用法的时候, 看 @map 的定义, 后面显式传递了 Applicative 的限定 (Monad 是 Applicative 的子类型)
      // 所以用 M[_]:Applicative 这种写法限定 M 的时候，编译器会动态加上 implicit G: Applicative[G] 这样的限定
      // def traverse[G[_],A,B](fa: F[A])(f: A => G[B])(implicit G: Applicative[G]): G[F[B]] =
      def traverse[M[_]:Applicative,A,B](fa: F[A])(f: A => M[B]): M[F[B]] =
        sequence(map(fa)(f))
      def sequence[M[_]:Applicative,A](fma: F[M[A]]): M[F[A]] =
        traverse(fma)(ma => ma)

      type Id[A] = A
      // this idMonad is a hard one to understand. Tricky
      val idMonad = new Monad[Id] {
        // def unit[A](a: => Id[A])
        def unit[A](a: => A) = a // a is a Type Id[A]
        // def flatMap[A, B](a: Id[A])(f: A => Id[B]): Id[B]
        override def flatMap[A,B](a: A)(f: A => B): B = f(a)
      }

      // traverse[Id, A, B](fa)(f)(idMonad) actually returns Id[F[B]], so here it makes sense
      // f: A=> B is actually f: A=> Id[B]
      def map[A,B](fa: F[A])(f: A => B): F[B] =
        traverse[Id, A, B](fa)(f)(idMonad)

    }

    trait Monad[F[_]] extends Applicative[F] {
      ....
    }


Another use case is Type Lambda, which dyamically create a new type.

Syntax of Type Lambda:

    ({type f[x] = Validation[E,x]})#f

Look the use case below:
Applicative only take a Type that accept one Type Param. But `Validation` Type takes two.
So here he use Type Lambda dynamically create new wrapper type `f`. Defination and usage in one line.

    object Applicative{
      def validationApplicative[E]: Applicative[({type f[x] = Validation[E,x]})#f] =
        new Applicative[({type f[x] = Validation[E,x]})#f] {
          def unit[A](a: => A) = Success(a)

          override def map2[A,B,C](fa: Validation[E,A], fb: Validation[E,B])(f: (A, B) => C) =
            (fa, fb) match {
              case (Success(a), Success(b)) => Success(f(a, b))
              case (Failure(h1, t1), Failure(h2, t2)) =>
                Failure(h1, t1 ++ Vector(h2) ++ t2)
              case (e@Failure(_, _), _) => e
              case (_, e@Failure(_, _)) => e
            }
        }

    }

    trait Applicative[F[_]] extends Functor[F] {
      ...
    }


    sealed trait Validation[+E, +A]

    case class Failure[E](head: E, tail: Vector[E]) extends Validation[E, Nothing]

    case class Success[A](a: A) extends Validation[Nothing, A]

Evaluation of Type constraint

    // in this case, type U constraint is  U>:T &&  U <% Ordered[U]
    def addValue[U >: T <% Ordered[U]](x: U): Tree[U]
























