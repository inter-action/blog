# Title: General introduction to Programming language ?


# Outlines:
  * why learn new language
  * general language introduction
  * typescript


#  why learning another language:

>to be a better progammer

* things are connected  
  Rx = functional combinators + nodejs stream
  typescript opens new door to a strong type language, angularjs2.
* help you think differently
  OOP, functional, Procedure
* battery included features
    * Immutable Objects
    * builted-in full fledged collections
    * elegant ways to handling errors
* different language excel at different application domain
    python: data mining & machine learning & system administration & devops
    scala: big data & java alternative
    java, golang: backend
    c: embeded systems
    js: front end.
    nodejs: IoT & backend
    swift & object c: ios related
* abstract your code in more elegant ways
* reads lots of books
* compile to JS

## Overview language & Programming Paradiam
![alt tag](/others/typescript/assets/language-spectrum.png)

#### Programming Paradiam
> almost every other language can do some multi paradiam programming, but some language just feel natrual & easy
 
* OOP:
  java, c++
* Procedure
    c
* Functional
    Haskell
* multi Paradiam:
    swift, scala, python, ....

#### Language birdview

* primitive types
* control statements, if while break for
* error handling
* coding structual mechanism, module, namespace
* others.... inheritance,


## Javascript: good & bad parts
  [jslint]()  got removed, cant find it anymore :(

### Why Javascript/Flexiable is Bad.
more info, see #Links - 1, 3
> most these problem are born with the language interited natrue which is dynamic language 

* too dynamic.

*  count on programmer's discipline to guarantee code quality  
    * include good doc, reasonable unit test, organize code in a reasonable way.
    * monkey patch things in an explicit way.
    * etc...

* hard to refactor  
  no ide check it for you.  
  you may include way too much internal state in a method.

* not IDE/tools friendly:
  * you get basically no ide prompt when u typing
  * u cant peek a definition of certain method
  * u cant view doc inside ide.
  * u cant easily jump around

* Why Javascript is good:
  *  no compile time
  *  the web programming language, with popular platform, de facto "write once run everywhere"
  *  simple
  * ...
  

## Types:


### Why Type is good:

* type as document:
  https://github.com/CrazyFork/xceptor
    you have no idea what `res, resp` really is & what you can do with it.
  
  xceptor's handler object:
    /Users/interaction/workspace/work/projects/eleme-dep/xceptor
      line: 28

      (request, response)=> Boolean | Promise


* fit for small & large project  
  do crazy stuff & get compiler check it for you.

* type means safer, the reason most backend language are static types

core difference between static type vs dynamic:
> write first, debug later


### when to apply a strong type language
see link #4

### Type headache:

* type system can be simple, also can be extreme complex
* tedious & rigid
* have compile time
* you may get bunch of problem related to types realm
* all come down to learning curve...

in spite of all these...,  
  what type really means is more constraint on programmer to write better structured code with decent quality

### the notorious null pointer exception:

reason: bad language design

what is it, how to eliminate it:

* go approach:

  ```
  let err, result = dosomething()
  ```    

* java approach:

  ```
  @NotNull annotation
  ```    

* scala & rust & swift

    ```
    Option type # much powerful.
    ```

* typescript, swift:

    hero?.fullName

since typescript 2.0, typescript add a strict null complier option.    


### The modern language: Option Type,  deconstruct , Pattern matching, for comprehension


* expression vs statement

  expression always return value
  ```
  val x = if (x > 3) 3 else 4
  ```
  statement:
    declare what i gonna do

* Option Type:

   Option[T], Some[T], None.

   /Users/interaction/workspace/typescript/typescript-samples

* deconstruct

  ```
  var {name, age} = {name: 3, age: 4}
  ```
* Pattern Matching:

  ```scala
  class Person(name)
  val p = new Person("sam")
  p match {
    case Person(name) => println(name)
    case _ => not found // if not matches throws runtime exception
  }
  ```

* commbined example:

  ``` scala
  val map = Map.empty[String, Int]()
  val age1 = map.get() //javascript indexOf
  val age2 = map.get()

  (age1, age2) match {
    case ( Some(age1), Some(age2) ) => println(age1, age2)
    case _ => println("we dont care")
  }

  // or
  age1.flatMap( age1 => age2.map(age2=> (age1, age2)) ).getOrElse((-1, -1))
  ```

* for comprehension

  scala:
  ```
  for (i <- List(1, 2, 3)) yield i * 2 // List(2, 4, 6)
  ```
  python:
  ```
  [x*2 for x in [1, 2, 3]]
  ```


### OOP

what is OOP?
//todo:

*  Single Parent inheritance vs trait based inheritance
*  interface vs trait based contract, structual types


### Type System:

*  type inference
    ```scala
    val x = 3
    ```

*  Type Ops:
        Or, Union, Algreba Types

*  Type alias:
    ```scala
    type Success = Int
    ```
* Type Coercing

  ```
  interface HeroService{
    getHeroes: Heros[]
  }

  <HeroService> {getHeroes: () => expectedHeroes }
  ```

*  Generics:

    List
    ```
    def get[T](idx: Int): T
    ```

*  High Order Types:

    ```scala
    def size[F, T](h: F[T]): Int // F can be List, Option, or any other container type
    ```

* covariant vs contravariant

    ```scala
    class Fruit
    class Orange
    class Apple

    List[Apple].add(new Orange)
    ```


*  Function VS Method.  
    function does not attached instances





  Type morphing: ???


* Type Complex Explosion

    type constraint

      extends, super, can be view as, interface. trait, pattern matching.

* Error handling
  try catch

  value, err = dosomethingInGo()

  scala:
    Option[T], Either[Error, Value]




## TypeScript:

### why typescript:

https://www.scala-js.org/

* thriving community
* not so complex
* better fullstack dev, combinded with nodejs & angular
* es6 superset & with most capability to interop with js


###  Language itself:

https://www.typescriptlang.org/docs/handbook/advanced-types.html


namespace vs modules:
* namespace: com.github.interaction, a prefix mechanism to avoid confliction
* modules: single file or bunch of files as a single unit

### tools
[ts-node](https://www.npmjs.com/package/ts-node)



###  Application demos:
//todo:
load types
write type for existing js
write pure types
  Option Type:

## VSCode:

[introduction of vs code](https://www.youtube.com/watch?v=e3tPWAq74v4)

existing project add typing hint //todo: 

cmd + p + ?


## Looking ahead / Future Sharing :

* Functional Programming:
for anyone interested in functional programming,   
*the red book* [Functional programming in scala](http://www.salttiger.com/functional-programming-in-scala/)


key elements:
* immutable Objects
* recursive data structure
* lazy evaluation
* tail call elimination - http://stackoverflow.com/questions/1240539/what-is-tail-recursion-elimination
* combinators
    applicator & monad

* Angularjs2 or RxJS:





## Links:

1. [Hands-on Scala.js - About Javascript](http://www.lihaoyi.com/hands-on-scala-js/)
2. [A high level look at Angular 2](http://www.developerhandbook.com/angular/high-level-look-angular-2/)
3. [The advantages of static typing, simply stated](https://pchiusano.github.io/2016-09-15/static-vs-dynamic.html)
4. [flow vs typescript](http://djcordhose.github.io/flow-vs-typescript/flow-typescript-2.html#/)
    comparison, why types, null types, type variantion
