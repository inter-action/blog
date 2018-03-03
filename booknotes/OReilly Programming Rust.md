
# OReilly Programming Rust

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









## others

todos:

    pending:
        virtual methods, inline methods
        

