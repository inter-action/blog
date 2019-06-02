
# OReilly.C.in.a.Nutshell.2nd.Edition.2015.12

## annoation
```
:- , finished todo
```

## 
* `ar`, 工具是用来创建c library 的
* c library:
    * `.a`: for archive, `-l -L`参数, 在link阶段
    * `.so`: shared objects
         * Shared libraries are special object files that can be linked to a program at runtime.
         * LIBRARY_PATH
* Freestanding Programs: 
```
// -v 很好用能看到各种include path的search方式, 和目标
// 输出 assembly language 的数据, 很有用.
gcc -S circle.c
// 
gcc -v -o circle -Wall,-as=circle.sym,-L circle.c
// 在 `/usr/local/lib` 下搜索 `libncurses.a` 文件
gcc -o circle -L/usr/local/lib -lncurses circle.c

// 这个参数将编译的中间产物留下来.  .i, .s, and .o for preprocessor output
// , assembly language, and object files, respectively.
-save-temps
// compile all .c files to .o files
gcc -c circle.c circulararea.c
// link all objects
gcc -o circle circle.o circulararea.o -lncurses

// create a shared library
gcc -c circulararea.c
gcc -shared -o libcirculararea.so circulararea.o

// use shared library
gcc -c circle.c
gcc -o circle circle.o libcirculararea.so -lncurses
```


## chapter 01: Language Basics

* `Wide characters`, in which the same bit width is used for every character in a character set. `wchar_t` 类型表示

* `Multibyte characters`, in which a given character can be represented by one or several bytes, and the character value of a given byte sequence can depend on its context in a string or stream

* Multibyte char 不适合程序处理

* `execution character`: 是指程序运行时候的 character set 设置吧. 

* `Digraphs and Trigraphs`, 没看懂, 估计也没啥大用处



## chapter 02: Types
* `size_t`, 
    * > size_t is the unsigned integer type of the result of the sizeof operator
    * https://stackoverflow.com/questions/1119370/where-do-i-find-the-definition-of-size-t
* `intptr_t` and `uintptr_t`, used for convert pointer to integer type
* `type void`, 可以是类型,返回值, 无类型指针


The Alignment of Objects in Memory: 
* `_Alignas(int|type)`:  indicate to compiler that that type requires a alignment of n bytes
* `_Alignof(type)`:  return a number of that type 

虽然还是没有很看懂这一块. 估计还是要查下

todo: 
* :- memory alignment.

## chapter 03: Literals
string: 
* `u8` prefix, can be used to create a utf8 string literal.
* string 的处理还是要先理解下 chapter 01的 `wide character` 和 `Multibyte character`的概念

todo: 
* `character constant`: 还有 `wide-character` 没怎么看懂, p43


## chapter 04: Type Conversion
* `typedef double (func_t)(double);`, function type.
* `strlen(msg)` vs `sizeof(msg)`, `strlen`返回的是string的长度, `sizeof`返回的是该指针申请的内存的长度

incomplete types:
* https://stackoverflow.com/questions/3917712/what-is-the-definition-of-incomplete-type-and-object-type-in-c
* incomplete types, 在我理解就是rust所说的unsize type.

todo:
* :- a pointer to a complete or incomplete object type, p60


## chapter 05: Expressions and Operators

* `Lvalue`, 书中给的解释就是operator左侧, 主要用于对象寻址的操作, 比如指针`*ptr`, 数组`arr[1]`.而`Rvalue`在operator右侧, 主要是常量, 数值.

* `Sequential evaluation`, `x, y`, Evaluates first x, then y; the result of the expression is the value of y
* `Compound literals`
    * syntax: `(type name ){ list of initializers }`
    * 我的理解就是用来动态的初始化对象的, 而省去了创建变量的过程.
* `sizeof`,
    * >sizof the operand has a structure type, the result is the total size that the object occupies in memory, including any gaps that may occur due to the alignment of the structure members. 
    *  Variable-length arrays are the only type that got evaluated at runtime.

* 三目运算符有些规则, 在书里我就不记录了, 也记不住.遇到再说


## chapter 07: functions

* > A function cannot return a function or an array. However, you can define a function that returns a pointer to a function or a pointer to an array.
* Functions and Storage Class Specifiers: `extern` 这是default, `static` 作为private关键字, 只在当前source file中有意义, `inline` :todo.
* c 中的function分为prototype和implementation两种, 书中在p121中详细介绍了区别, 尤其是variable array的声明. 在prototype declaration中, the name of function arguments and variable representing length of array are optional.
* c 的函数必须先声明才能被调用
* `Non-Returning Functions`, >The function specifier _Noreturn is new in C11.
* `variadic functions`, c处理这种函数还真的是原始.需要用macro.


## chapter 08: Arrays


## chapter 20: Using make to Build C Programs

* `@`: target
* `^`: first prerequisite
* `<`: all prerequisites

pattern rule:

```
// % 表示wildcard, 显示指定规则: match任何.o结尾的文件, 去替换.c
    circulararea.o circle.o: %.o: %.c
            $(CC) $(CFLAGS) -o $@ -c $<
// implicit style, match all .o file with .c file
 %.o: %.c
            $(CC) $(CFLAGS) -o $@ -c $<
```

* `make -p`, display default rules, variables
* `Double-Colon Rules`: 可以having alternatives的rules
* `$(name:ending=new_ending)`, 
```
OBJ = circle.o circulararea.o
SYMTABS = $(OBJ:.o=.sym)
```

The Automatic Variables: 是指 `$` 开头的默认变量


Attributes:
* `.PHONY`: >Any targets that are prerequisites of .PHONY are always treated as out of date.

Functions:
* function 分两种, 一种是built-in, 一种是user defined, user defined需要用`call` 关键字触发

gcc with makefile
* `gcc -M hello.c`: 会生成make file rule, 把所有hello.c的依赖动态生成makefile规则.

todo: 

* review default rules
* make file 的phony 始终没怎么看懂, 感觉和make file的依赖build有关, as it only build when targets are older than its prerequisites.
* `Generating Header Dependencies`, 一节没怎么看懂

Links:
* [makefile reference](https://www.gnu.org/software/make/manual/html_node/Quick-Reference.html)

--------------------------------------
# Others

## static 
* when used with function, it means private
* A static variable inside a function keeps its value between invocations.
    * 这个我记得是compile完, 会在某个section区域声明一个holder去维持这个变量, 在整个进程中.
* Static global variables are not visible outside of the C file they are defined in.

links:
* https://stackoverflow.com/questions/572547/what-does-static-mean-in-c

## scope in c

links:
* http://aelinik.free.fr/c/ch14.htm

## memory alignment:
* https://en.wikipedia.org/wiki/Data_structure_alignment
* 非align的, 会拖慢计算机性能, 也有脏读的bug, 很少见但也很难排查
* C and C++ do not allow the compiler to reorder structure members to save space
* 在多线程模式中, malloc 函数生成的block最好能align到cache block, 以满足性能上的要求.
* 文中最后一节没有特别看懂.
* c compiler will automaticlly generate padding for structs & union.
* disable automatically alignment will result a incompatible ABI.

links:
* https://stackoverflow.com/questions/381244/purpose-of-memory-alignment


## pointers
* used as general type for functions:
    ```c
    int floatcmp( const void* p1, const void* p2 ) {
        float x = *(float *)p1, y = *(float *)p2;
        return(x<y)?-1:((x==y)?0:1); 
    }
    ```
* `NULL pointer`: A null pointer constant is an integer constant with the value 0

* const pointer: readonly pointer `const struct Node *pNode `

    ```c
    struct Article *pArticle = &sw, 
        const *pcArticle = &sw;
    ++(pArticle->number);
    // illegal, const pointer.
    pcArticle->price += 50;
    ```
* 
* function pointer in a array
    ```c
    #include <stdio.h>
    void func0() { puts("This is the function func0(). "); }
    void func1() { puts("This is the function func1(). "); }
    /* ... */
    void (*funcTable[2])(void) = { func0, func1 }; // Array of two pointers
                                                    // to functions
    // returning void. for ( int i = 0; i < 2; ++i ) // Use the loop counter as the array
    funcTable[i](); // index.
    ```

