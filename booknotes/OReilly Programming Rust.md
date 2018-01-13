
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



