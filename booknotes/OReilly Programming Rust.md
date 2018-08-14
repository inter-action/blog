
# OReilly Programming Rust

好难呀 :(, !!!

## chapter 2:

demos:

notes:

* about trait:
    > Although we never use the names Write or FromStr elsewhere in the program, a trait must be in scope in order to use its meth‐ ods. 


## chapter 3: Basic Types


demos:

notes:

* Trait Object:
    > Trait object: reference to any value that implements a given set of methods (namely some Trait)
    > 

* usize, isize:
    >The usize and isize types are analogous to size_t and ptrdiff_t in C and C++.
    >Rust requires array indices to be usize values.

* byte literals:
    > `b'abc'`, 表示 abc 的字符串实际的字节流
    > `b'\xAB'`, 表示以hex的方式创造的HH字节流，即 `1010 1011`

* examples:

    ```
    0b101101u8
    ```

* String , str, &str
    * 默认的情况是 `"this is &str"` 字符串默认是 string literal 即 `&str`
    * str 代表着最原始的内存 slice, 即 [u8] 数组, 这段信息应该是保存在 compiled binary 中的 .constants 字段中的
    * &str, 和 String 都是对最原始 string [u8] 的引用, 只不过 String 貌似是在 heap 中分配的对象, 因为是 Owned type, 肯定要 Drop
        的.
        * > For example, String implements From<&str>, which copies the string slice into a new heap-allocated buffer for the String.  @page299

    * &str 可以指向 .constants 中的 string literal 也可以指向 String (在 heap中分配的 [u8])


todos:
* add notes here


## chapter 4: Ownership

demos:

notes:

todos:


## chapter 5: References


## chapter 6: Expressions

demos:

notes:

todos:


## chapter 7: Error Handling


demos:
note:
todos:
* Panic is per thread. One thread can be panicking while other threads are going on about their normal business. 


## chapter 8: Crates and Modules



## chapter 9: Structs


* Interior Mutability: 
    * 内部变更的 trait 有两种 Cell<T> & RefCell<T>, 这两种都允许你用 & 的 ref 去做 mutation.
        * Cell<T>: 
            * T 必须是 Copyable的, 因为 get 方法需要返回这个 T
            * 用法是 

                ```rust
                cell.get()
                cell.set(value)
                ```
        * RefCell<T>:
            * 这个 T 是没有限制的, 用法是:

                ```rust
                ref_cell.borrow()
                ref_cell.borrow_mut()
                ```
            * 需要注意的是, 虽然这两者都是在 compile time 规避掉 rust 的borrow rule 的check, 但RefCell会在运行时去enforce这个check, 所以如果不通过的话, 会出现一个runtime的panic

* unsized struct:
    * A struct type’s last field (but only its last) may be unsized, and such a struct is itself unsized.


## chpater 10: Enums and Patterns

* rust enum type can hold
    * a number
    * struct like tuple
    * struct like data


## chapter 11: Traits and Generics

notes:

* >There is one unusual rule about trait methods: the trait itself must be in scope. Otherwise, all its methods are hidden.
    * > The reason Clone and Iterator methods work without any special imports is that they’re always in scope by default
* trait object:
    * > A variable’s size has to be known at compile time
    * > A reference to a trait type, like writer, is called a trait object.
    * made by `&` or `Box` or any other pointer types
    * 
    ```rust
    struct Salad<V: Vegetable> { veggies: Vec<V> } // V: 只允许 Vegetable type
    struct Salad {
        veggies: Vec<Box<Vegetable>> // veggies 经过 Box 包装后允许任何实现了 Vegetable type 的type
    }
    ```
    * >The second advantage of generics is that not every trait can support trait objects. Traits support several features, such as static methods, that work only with generics: they rule out trait objects entirely. We’ll point out these features as we come to them.
    * trait object vs trait
        * type with trait will be statically compiled to different code with different actual type
        * trait object will be compile with a fat pointer(one point to actual instance, and one with ref to instance function implementation vtable), so it's a dynamic dispatch mechanism. due to pointer's inherent behavior, it has to create data on the heap. a pointer has
        to refs to something.
    * vs `existential types`
        * https://blog.rust-lang.org/2018/05/10/Rust-1.26.html



* trait
    * Rust lets you implement any trait on any type, as long as either the trait or the type is introduced in the current crate.

    * implement trait for an abstract type
    ```rust
    use std::io::{self, Write};
    /// Trait for values to which you can send HTML.
    trait WriteHtml {
        fn write_html(&mut self, &HtmlDocument) -> io::Result<()>;
    }
            /// You can write HTML to any std::io writer.
    impl<W: Write> WriteHtml for W {
        fn write_html(&mut self, html: &HtmlDocument) -> io::Result<()> {
        ... }
    }

    ```



* Fully Qualied Method Calls, 这一节需要注意下, 这一节讲述了如何解决 trait 方法调用的不确定性


todos:
* p252, 需要注意下


## chapter 12: Operator Overloading

summary:

这章主要讲了如何 overload operator in rust.

notes:

* page 273 的  PartialEq 的解释终于让我明白了这个东西是啥意思 !
* ParticalOrder <-> Order 还有 PartialEq <-> Eq 的关系是一致的


## chapter 13: Utility Traits
summary:
chapter 13将的这些 trait 还是等以后慢慢品吧, 感觉现在短期看不出来是干啥的


notes:

* If a type implements Drop, it cannot implement the Copy trait. If a type is Copy.
* Sized, 这个 trait 很重要, 书里边也讲的非常详细. 有可能需要找个时间再重读下这个位置
    * A sized type is one whose values all have the same size in memory. 
    * unsized type: Write, str, [T], Display etc...
    * 



* AsRef 和 DeRef 的区别, 
    * 这个地方我想了好久, 其实 AsRef 就是 DeRef 的反向操作
    * 这两个trait需要好好看下, 还是蛮重要的.

    

todos:



## chapter 14: Closures
summary:
这章讲了 rust 如何实现的 closure, 以及 closure 不同的类型和区别, function type & closure 在函数类型上
兼容的方式的处理. 还有Rust和其他语言由GC衍变上的Closure实现的核心区别. 还有 Boxed closure

notes:


todos:


## chapter 15: Iterators
summary:

notes:
* iterator 的 shared 和 mutable type获得方式, 跟 root path 中的 shared or mutable type有关 
    ```rust
    (&favorites).into_iter() 
    (&mut vector).into_iter() 
    ```

* rust 的 iterable type 里边拿到 iterable type 有两种方式, `iter` `iter_mut` 和 `into_iter`.
    其中只有 `into_iter` 是 trait 中定义的, 所以你可以用 ta 来写些more generic code .

* `by_ref` 这个实现非常有意思, 正常的 adapter 会 take ownership of the target, `by_ref` 把 iterator
    的 ownership 转换成一个 `&mut` reference, 然后这个 `&mut` reference 给 adapter 去 take ownership.

    
* > Of course, a for loop uses IntoIterator::into_iter to convert its operand into an iterator. 
    But the standard library provides a blanket implementation of IntoIterator for every type that implements Iterator,
    



todos:
* how drain method is implemented?
    ```rust
    use std::iter::FromIterator;
    let mut outer = "Earth".to_string();
    let inner = String::from_iter(outer.drain(1..4)); assert_eq!(outer, "Eh");
    assert_eq!(inner, "art");
    ```
*




## chapter 15: Collections
summary:


notes:



todos:











## others

todos:

    pending:
        virtual methods, inline methods
        

