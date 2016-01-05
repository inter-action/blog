

this post list all my understandings on scala. I'll try to write it all in English. 
Please forgive me for certain syntax error and some chinese mixins in this post.

# null elimination
scala has every powerful type system. this means it can provide things much more powerful if you can master it.
Java also is a strong type language. But its type system is quite limited. if you familiar with its generics.
What make things worse is in java null can be assign to any types, this means null itself is every type, 
yet isnt non of these types. eg:

    List<String> a = getNullFromDataBase()//
    a.length //whoops, null pointer reference at runtime. wake up and fix it!

But in scala, Null itself is a specific Type. It cant be assign to other type. This means you dont always need to check on a variable 
nullable or not. eg:

    def getListLength(as: List[Any]): Int = as.length // save your lots of trouble from checking nullable here. because variable as cant be null.

Other than this, scala provide you tools like Option, Either to handle semantic null or error things that traditional null
or Exception provided. You can also wrap nullable third party interface (typically java libs) with Option or Either. 



## todo

    FP pattern in scala (monad)
    recursive data structure
