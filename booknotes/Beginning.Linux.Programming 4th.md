# Beginning.Linux.Programming 4th


# Denotes


under links context:

    #   //in links context, # denotes segment of the page
    !   //in links context, ! denotes an important link
    ?   // not fully understand yet, todo


# mac step up

    install xcode
    >xcode-select --install
# gcc options

    -L: library location
    -l: library
    -o: output


## chapter 1: Getting Started

> Applications are usually kept in directories reserved for them. Applications supplied by the system for general use, including program development, are found in /usr/bin. Applications added by system administrators for a specific host computer or local network are often found in /usr/local/bin or /opt.
 
* header files: 
    * what is header files:
    >For programming in C and other languages, you need header files to provide definitions of constants and declarations for system and library function calls. 

    * where it located:
    >For C, these are almost always located in /usr/include and subdirectories thereof. You can normally find header files that depend on the particular incarnation of Linux that you are running in /usr/include/sys and /usr/include/linux.


* Library Files:
    * two types
        * static libraries:  .a for traditional, static libraries
            * static libraries 一般是你有改对应的lib文件, 每个工程都需要一份对应的拷贝。会被编译到二进制文件中
        * shared libraries: .so for shared libraries (see the following)
            * shared libraries 是多个binary 程序公用一个 lib， 这个lib不会编译到二进制文件中，而是在运行时动态调用。
            * ldd <program_name>
                >You can see which shared libraries are required by a program by running the utility ldd

    * notes:
        * You must use the –l option to indicate which libraries other than the standard C runtime library are required.

        * In many ways, shared libraries are similar to dynamic-link libraries used under Windows. The .so libraries correspond to .DLL files and are required at run time, and the .a libraries are similar to .LIB files included in the program executable.



## chapter 2: Shell Programming


file descriptor:
    0 - input, 1 - output, 2 - error

    $ kill -1 1234 >killouterr.txt 2>&1
    redirect standard output to the file killouterr.txt, and then direct standard error to the same place as the standard output.

    
    $ kill -1 1234 >/dev/null 2>&1
    discard all std & std error msg

glob:

    * : any
    ? : any single char
    [^set] : invert, not
    {finger,toe}: options

Shell Syntax:

references:
* https://www.tutorialspoint.com/unix/index.htm


* variables
    * By default, all variables are consid- ered and stored as strings, even when they are assigned numeric values. 

* Quoting
    double quote("): with substitution
    single quote('): no substitution

* Conditions
    * ! test: Page 33
    * page 36, a variable problem
    * Statement Blocks, page 45

* Function:
    * You must always define a function before you can invoke it
    * function args: 
    >When a function is invoked, the positional parameters to the script, $*, $@, $#, $1, $2, and so on, are replaced by the parameters to the function. That’s how you read the parameters passed to the function. When the function finishes, they are restored to their previous values.






## todos

    bash set









    

