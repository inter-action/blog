
## chapter 9: Type System

* 注意下 156 页面的 size_t 的定义，unsigned long, 用于 array 的 index. 具体大小跟底层平台有关, 基本上就是 pointer size
* sizeof 的使用方式也是, 轻易不要用在数组上，在函数中作为传入参数的数组计算会有问题
* p159, const 和 * 的组合也需要注意下

* string:
    * The type of string literals is char*. Modifying them, however, while being syntactically possible (e.g., "hello"[1] = 32), yields an undefined result. It is one of the cases of undefined behavior in C.

* 9.4-Polymorphism in C， 需要重点看下

## Links

* [online c playground](https://www.tutorialspoint.com/compile_c_online.php)
* [What is the difference between NULL, '\0' and 0](https://stackoverflow.com/questions/1296843/what-is-the-difference-between-null-0-and-0)



## chapter 10: Code Structure
