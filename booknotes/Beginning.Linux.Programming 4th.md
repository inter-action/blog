# Beginning.Linux.Programming 4th


# Denotes


under links context:

    #   //in links context, # denotes segment of the page
    !   //in links context, ! denotes an important link
    ?   // not fully understand yet, todo


# mac step up

    install xcode
    >xcode-select --install

## common notes:
gcc options

    -L: library location
    -l: library
    -o: output

??:
    * These /dev files are used to access hardware in a specific way using low-level system calls.
    * Linux provides a special file system, procfs, that is usually made available as the directory /proc. It contains many special files that allow higher-level access to driver and kernel information. Applications can read and write these files to get information and set parameters as long as they are running with the correct access permissions.
    
c lang:
    * [define](http://www.cprogramming.com/tutorial/cpreprocessor.html)


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
    backquote(\`): , you should use  `$(command)` instead
    ```plain
        All new scripts should use the $(...) form, which was introduced to avoid some rather complex rules covering the use of the characters $, `, and \ inside the backquoted command.
    ```

* Conditions
    * ! test: Page 33
    * page 36, a variable problem
    * Statement Blocks, page 45

* Function:
    * You must always define a function before you can invoke it
    * function args: 
    >When a function is invoked, the positional parameters to the script, $*, $@, $#, $1, $2, and so on, are replaced by the parameters to the function. That’s how you read the parameters passed to the function. When the function finishes, they are restored to their previous values.


* Commands:
    * The : Command, 这货啥也不干, 相当于 true 
    >The colon command is a null command. It’s occasionally useful to simplify the logic of conditions, being an alias for true. Since it’s built-in, : runs faster than true

    * The dot (.) command executes the command in the current shell: `. ./shell_script`
    > In shell scripts, the dot command works a little like the #include directive in C or C++. Though it doesn’t literally include the script, it does execute the command in the current context, so you can use it to incorporate variable and function definitions into a script.

    * export:
    > The export command makes the variable named as its parameter available in subshells. By default, variables created in a shell are not available in further (sub)shells invoked from that shell. The export command creates an environment variable from its parameter that can be seen by other scripts and pro- grams invoked from the current program. More technically, the exported variables form the environ- ment variables in any child processes derived from the shell.

    * unset variables:
    >Writing foo= would have a very similar, but not identical, effect to unset in the preceding program. Writing foo= has the effect of setting foo to null, but foo still exists. Usingunset foohastheeffectofremovingthevariablefoofromtheenvironment.

    * eval command:

    * find command: 
        * `$ find . -newer while2 -type f -exec ls -l {} \;` 
        * ` find . \( -name “_*“ -or -newer while2 \) -type f -print`


    * expr & $((...)): the both of two do arithmetic calc, but the later form is supposed to replace the former.

* Parameter Expansion:

    ${param:-default}, ${#param}, ${param%word}, ${param%%word}, ${param#word}, ${param##word}

>${foo:=bar}, however, would set the variable to $foo. This string operator checks that foo exists and isn’t null. If it isn’t null, then it returns its value, but otherwise it sets foo to bar and returns that instead.

* Debugging Scripts: page74

* Page 75, section: Going Graphical — The dialog Utility, skiped


## chapter 3: Working with Files

* Libraries function vs System calls
    * System call come with an overhead (without buffering), you in most case you should not make system call directly, instead you should use library func, let them pass control to System call instead.

* umask, block file permissions, same with file permissions the sequence is `rwx`, it take preceduene to file permissions.


* The Standard I/O Library:
    >The standard I/O library (stdio) and its header file, stdio.h, provide a versatile interface to low-level I/O system calls. The library, now part of ANSI standard C, whereas the system calls you met earlier are not, provides many sophisticated functions for formatting output and scanning input. It also takes care of the buffering requirements for devices.


## chapter 4: The Linux Environment

* /etc/syslog.conf
* syslog facility: https://en.wikipedia.org/wiki/Syslog#Facility
* CPU time consumed by a program is separated into user time (the time that the program itself has con- sumed executing its own instructions) and system time (the CPU time consumed by the operating system on the program’s behalf; that is, the time spent in system calls performing input and output or other sys- tem functions).

* priority value: the less value, the higher priority
    > The default priority is 0. Positive priorities are used for background tasks that run when no other higher priority task is ready to run. Negative priorities cause a program to run more frequently, taking a larger share of the available CPU time. The range of valid priorities is -20 to +20. This is often confusing because the higher the numerical value, the lower the execution precedence.

* soft limit & hard limit:
    > Typically, the soft limit is an advisory limit that shouldn’t be exceeded; doing so may cause library functions to return errors. The hard limit, if exceeded, may cause the system to attempt to terminate the program by sending a signal to it. Examples would be the signal SIGXCPU on exceeding the CPU time limit and the signal SIGSEGV on exceeding a data size limit. 





# links
* [Linux system call references](http://syscalls.kernelgrok.com/)
* [POSIX OS header files references](https://en.wikipedia.org/wiki/Unistd.h)
* [c lange tutorialspoint](https://www.tutorialspoint.com/cprogramming/index.htm)
* [cpu GHZ means](http://smallbusiness.chron.com/ghz-mean-computer-processor-66857.html)

## todos

    bash set
    umask, how to use it.
    grep O_RDONLY inside usr/include dir








    

