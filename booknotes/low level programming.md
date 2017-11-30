# Apress.Low-Level.Programming

* 这本书可以算是我读过的比较好的书的一种了，内容可以说是很全面也很有深度。
* codes - git@github.com:CrazyFork/low-level-programming.git

## Compilation Pipeline

demos:
* explaination of a internal representation of compiled binaries.


## chapter 9: Type System

* 注意下 156 页面的 size_t 的定义，unsigned long, 用于 array 的 index. 具体大小跟底层平台有关, 基本上就是 pointer size
* sizeof 的使用方式也是, 轻易不要用在数组上，在函数中作为传入参数的数组计算会有问题
* p159, const 和 * 的组合也需要注意下

* string:
    * The type of string literals is char*. Modifying them, however, while being syntactically possible (e.g., "hello"[1] = 32), yields an undefined result. It is one of the cases of undefined behavior in C.

* 9.4-Polymorphism in C， 需要重点看下

## chapter 10: Code Structure
demos:

* Preprocessor:
    * preprocessor is like rust macro, it replace the preprocessor symbol define with its body and param included.
    * it can be defined in two ways
        * in c code:

            ```c
            #define FLAG
            #define MY_CONST 42
            ```
        * gcc cmd line with `gcc -D<name>=<value>` or `gcc -D<name>`

    * Include Guard:
        * include guard is a technique to solve the issue that a duplicated symbol's declaration inclusion results a compile error.
        ```
        #ifndef _FILE_H_  # pattern is _<filename with dot replaced with _>_H_
        #define _FILE_H_

        void a(void);     # the actual defination, so if the file been included more then once, _FILE_H_ should already been defined, hence the code here would not be included again.

        #endif
        ```
    * preprocessor is eval, 

notes:
* c 语言中函数声明必须放在前面的原因
>and due to the single-pass translation, the compiler can’t look ahead and try to find the definition.

* forward declaration, help to resolve circular reference.

    ```
    struct b; /* forward declaration */
    struct a {
        int value;
        struct b* next;
    };
    /* no need to forward declare struct a because it is already defined */
    struct b {
        struct a* other;
    };
    ```
* two types of libraries:
    * static lib, .o or .a file: this kind lib would be included in the final compilied binary
    * shared/dynamic lib, .dll(windows) or .so(linux): wouldn't be included in the binary, just ffi call using some kind of protocal.

* c header file:
    * 这么多年，老夫总于读懂 header file了。c 的编译是 single-pass, 上边说过了，即必须所有用到的东西都需要在使用之前声明,
        所以 header file 是为了不同文件能够公用一份声明(declaration), 用于代码的可维护

    * #include 指令 " 和 < 的区别
        * < search 的directory包括
            * /usr/local/include
            * <libdir>/gcc/target/version/include
                * Here <libdir> stands for the directory that holds libraries (a GCC setting) and is usually /usr/lib or /usr/local/lib by default.
            * /usr/target/include
            * /usr/include
            * 任何 -I 参数包括的其他directories
        * " search除了包括 < 的directory之外还包括当前目录



* static:
    * static local variable:

        ```c
        int demo (void)         
        {
            static int a = 42;      // 这种变量，只能由enclosing函数可见, 会保存上一次demo函数调用后的值
            printf("%d\n", a++);
        }
        ```


## chapter 11: Memory

demos:
* Pointer Arithmetic
    * NULL pointer
    * void* pointer
    * pointer diff, result should always be size_t, to match target machine pointer size.
    * Function Pointer

        ```c
        // normal usage        
        double doubler (int a) { return a * 2.5; }
        ...
        double (*fptr)( int );
        double a;
        fptr = &doubler;
        a = fptr(10); /* a = 25.0 */

        // with type def
        double doubler (int a) { return a * 2.5; }
        typedef double (megapointer_type)( int );
        ...
        double a;
        megapointer_type*  variable  =  &doubler;
        a  =  variable(10);  /*  a  =  25.0  */

        ```
* Memory Allocation:
    * Automatic memory allocation, in stack
    * Static memory allocation, .bss, .rodata, .data sections
    * Dynamic memory allocation
        * malloc, calloc, free, realoc

* pointer with array

* string in c
    * it end with a byte of zero, namely `\0`
    * can be allocated in .rodata section (local variable declaration), dump(malloc system call), .data section(global variable declaration) 
    * it's immutable

* platform independent types, since C99:
    * 解决不同target的类型type不统一的问题, 感觉蛮复杂的如果真正用起来，不过如果只是target one specific os, this issue should be significant easier.

* Data streams(files):
    * `fread, fread, fwrite, fprintf, fscanf, fopen, fclose, fflush`. is user space function that working with `FILE` struct that hide all the os difference that wrapped around `open` system call
    * Again, descriptors are integers, FILE instances are not, and FILE always allocated on heap.
    * EOF constant means end of file.
    * streams can be buffererd to avoid expansive system call(that involving context switching)

notes:
* linux man tool:
    * man tool can be used to reference c language functions, or linux command
    * the manual is divided in different sections, you can take a glance at https://linux.die.net/man/, eg, `LS(1)` is display at left top corner of termial when `man ls` is typed
    * man <section number> <command>, to limit your search




## chpater 12: Syntax, Semantics, and Pragmatics

demos:

notes:


## others

notes:
* 学习 c 的思路, 先找一本简单的入门书籍, 然后是<深入理解操作系统>(Operating system in programmer's view), 然后就是这本书了。最后可以考虑备一本高阶的c语言书作为参考。

* gcc
    * examples:
        ```bash
        gcc -E main.c  # view preprocess result
        ```


## Links

* [online c playground](https://www.tutorialspoint.com/compile_c_online.php)
* [What is the difference between NULL, '\0' and 0](https://stackoverflow.com/questions/1296843/what-is-the-difference-between-null-0-and-0)
* [gcc options reference](https://gcc.gnu.org/onlinedocs/gcc/Warning-Options.html#index-Wno-pedantic-ms-format)
* [linux tools ref, including nm, objdump](http://linuxtools-rst.readthedocs.io/zh_CN/latest/tool/nm.html)
* [! code section](https://en.wikipedia.org/wiki/Data_segment)
# todos

* elf file format, objdump tool
* [linux tools ref, including nm, objdump](http://linuxtools-rst.readthedocs.io/zh_CN/latest/tool/nm.html)
    * nm, objdump
