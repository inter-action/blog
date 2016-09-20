
# Out Line:
  two sections:
  * general language introduction
  * typescript


## Overview language & Programming Paradiam
Categories:
dynamic, static, strong type, weak type

#### Programming Paradiam
  Object:
    java

  Procedure
    c

  Functional
    Haskell

  multi Paradiam:
    swift, scala, python, ....

#### Language birdview

* primitive types
* control statements, if while break for
* error handling
* coding structual mechanism, module, namespace
* others.... inheritance,



## Javascript: good & bad parts
  [jslint]()

### Why Javascript/Flexiable is Bad,
more info, see #Links - 1, 3

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
  *  with popular platform, de facto "write once run everywhere"
  *  simple
  * ...


#  why learning another language:

    * help you think differently

    * battery included features

        * Immutable Objects
        * builted-in full fledged collections

    * different language excel at different application domain

        python: data mining & machine learning & system administration & devops
        scala: big data & java alternative
        java, golang: backend
        c: embeded systems
        js: front end.
        swift & object c: ios related

    * abstract your code in more elegant ways
    * reads lots of books


## Types:


### Why Type is good:

* type as document:  
  xceptor's handle object:
      (request, response)=> Boolean | Promise

* fit for small & large project  
  do crazy stuff & get compiler check it for you.

* type means safer

### Type headache:

* type system can be simple, also can be extreme complex
* tedious & rigid
* have compile time
* you may get bunch of problem related to types realm
* learning curve

in spite of all these...,  
  what type really means is more constraint on programmer to write better structured code with decent quality

### the notorious null pointer exception:

reason two-fold: bad language design & lazy programmer

what is it, how to eliminate it.

go approach:

    let err, result = dosomething()

java approach:

    @NotNull annotation

scala & rust & swift

      Option type # much powerful.

typescript:


### The modern language: Option Type & deconstruct & Pattern matching


deconstruct
  var {name, age} = {name: 3, age: 4}

Pattern Matching:

    class Person(name)
    val p = new Person("sam")
    p match {
      case Person(name) => println(name)
      case _ => not found // if not matches throws runtime exception
    }

commbined example:

    val map = Map.empty[String, Int]()
    val age1 = map.get() //javascript indexOf
    val age2 = map.get()

    (age1, age2) match {
      case ( Some(age1), Some(age2) ) => println(age1, age2)
      case _ => println("we dont care")
    }

    // or
    age1.flatMap( age1 => age2.map(age2=> (age1, age2)) ).getOrElse((-1, -1))


### Type System:

*  type inference
        val x = 3

*  Type Ops:
        Or, Union, Algreba Types


*  Type alias:
        type Success = Int

*  Generics:

        def get[T](idx: Int): T


*  High Order Types:
        def w(h: List[Map[String, Int]]): Int

        covariant vs contravariant
          class Fruit
          class Orange
          class Apple

          List[Apple].add(new Orange)


*  Function VS Method.  
    function does not attached instances


*  OOP programming:

  *  Single Parent inheritance / trait based inheritance
  *  interface vs trait based contract, structual types
  *  type abstraction


  Type morphing: ???


* Type Complex Explosion

    type constraint

      extends, super, can be view as, interface. trait, pattern matching.



### expression vs statement
  expression always return value

  val x = if (x > 3) 3 else 4

  statement:
    declare what i gonna do

  Some(3).getOrElse(4)

## TypeScript:

  > load scalajs

### why typescript:

* thriving community
* with most compatibility with Javascript
* not so complex
* better fullstack dev, combinded with nodejs & angular

###  Language itself:

    - https://www.typescriptlang.org/docs/handbook/advanced-types.html

    todo:

    namespace vs modules
      namespace: com.github.interaction
      modules: single file or bunch of files as a single unit

###  Application:

    load types
    write type for existing js
    write pure types
      Option Type:

## VSCode:

  existing project add typing hint

  https://www.youtube.com/watch?v=e3tPWAq74v4

  cmd + p + ?


Functional Programming

  key elements:
    immutable Objects
    recursive data structure
    lazy evaluation
    tail call elimination - http://stackoverflow.com/questions/1240539/what-is-tail-recursion-elimination



## Links:

1. [Hands-on Scala.js - About Javascript](http://www.lihaoyi.com/hands-on-scala-js/)
2. [A high level look at Angular 2](http://www.developerhandbook.com/angular/high-level-look-angular-2/)
3. [The advantages of static typing, simply stated](https://pchiusano.github.io/2016-09-15/static-vs-dynamic.html)
