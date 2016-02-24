


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