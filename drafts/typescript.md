Option Type:
what is type

Programming Paradiam

  Object:
    java

  Procedure
    c

  Functional
    Haskell

  Paradiam:

    swift, scala, python, ....


Type Complex Explosion

  type constraint

    extends, super, can be view as, interface. trait, pattern matching.


Immutable Objects

Javascript: good & bad parts
  [jslint]()

  Why Javascript/Flexiable is Bad,
  Why Type is good:

    count on programmer's discipline to guarante code quality
      include good doc, reasonable unit test, organize code in a reasonable way.
      monkey patch things in an explicit way.

    hard to refactor
      no ide check it for you.
      you may include way too much internal state in a method.

    not IDE friendly:
      you get basically no ide prompt when u typing
      u cant peek a definition of certain method
      u cant view doc inside ide.
      u cant easily jump around

    type as document:
      xceptor's handle object:
        (request, response)=> Boolean | Promise

    do crazy stuff & get compiler check it for you.


  Why Javascript is good:
    no compile time
    with popular platform, de facto "write once run everywhere"
    simple


    ...



Type System:

  type inference
    val x = 3


  Type Ops:
    Or, Union, Algreba Types


  Function VS Method.
    function does not attached objects


expression vs statement
  expression always return value

  val x = if (x > 3) 3 else 4
