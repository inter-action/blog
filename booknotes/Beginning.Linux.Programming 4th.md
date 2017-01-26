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




## chapter 5: Terminals

* LF vs CR:
    >The character the program actually sees isn’t an ASCII carriage return, CR (decimal 13, hex 0D), but a line feed, LF (decimal 10, hex 0A). This is because, internally, Linux (like UNIX) always uses a line feed to end lines of text; that is, UNIX uses a line feed alone to mean a newline, where other systems, such as MS-DOS, use a carriage return and a line feed together as a pair. 

* `/dev/tty` 
    > /dev/tty, which is always the current terminal, or login session. Because Linux treats everything as a file, you can use normal file operations to read and write to /dev/tty.

* The termios Structure:
    * [Two Styles of Input: Canonical or Not](http://www.gnu.org/software/libc/manual/html_node/Canonical-or-Not.html#Canonical-or-Not)

    * 控制terminal的关键struct:
    >termios is the standard interface specified by POSIX and is similar to the System V interface termio. The terminal interface is controlled by setting values in a structure of type termios and using a small set of function calls. Both are defined in the header file termios.h.

    * terminal控制的模式（mode）:
        * input: 控制字符如何读取到程序
        * output: 控制输出如何输出到driver
        * control: These modes control the hardware characteristics of the terminal
        * Local Modes: These modes control various characteristics of the terminal. 
        * Special Control Characters: Special control characters are a collection of characters, like Ctrl+C, acted upon in particular ways when the user types them. 

    * `stty -a`: view current terminal settings


* The TIME and MIN Values:
    >The values of TIME and MIN are used only in non-canonical mode and act together to control the reading of input. Together, they control what happens when a program attempts to read a file descriptor associ- ated with a terminal.

* terminfo package:
    >The terminfo package contains a database of capabilities and escape sequences for a large number of terminals and provides a uniform programming interface for using them


* [GUN terminal mode manual](http://www.gnu.org/software/libc/manual/html_node/Terminal-Modes.html)



## chapter 6: Managing Text-Based Screens with curses

skip, right now. not interested

## chapter 7: Data Management

* swap space:
    >Initially, the kernel was simply able to use free physical memory to satisfy the application’s request for memory, but once physical memory was full, it started using what’s called swap space. On Linux, this is a separate disk area allocated when the system was installed. If you’re familiar with Windows, the Linux swap space acts a little like the hidden Windows swap file. However, unlike Windows, there are no local heap, global heap, or discardable memory segments to worry about in code — the Linux kernel does all the management for you.

    > Behind the scenes, Linux is managing the blocks of memory the programmer is using as a set of physical “pages,” usually 4K bytes each, in memory. However, if a page of memory is not being used, then the Linux mem- ory manager will be able to move it from physical memory to swap space (termed paging), where it has little impact on the use of resources. If the program tries to access data inside the memory page that has be moved to swap space, then Linux will very briefly suspend the program, move the memory page back from swap space into physical memory again, and then allow the program to continue, just as though the data had been in memory all along.


* when space run out:
    > Eventually, when the application exhausts both the physical memory and the swap space, or when the maximum stack size is exceeded, the kernel finally refuses the request for further memory and may pre- emptively terminate the program.


* 2 ways of file locking:
    * The simplest method is a technique to create lock files in an atomic way, so that nothing else can happen while the lock is being created.
    * The second method is more advanced; it enables programs to lock parts of a file for exclusive access.


## chapter 8: Mysql

    Post-Install Configuration:
    删除root以外的账户
    create user & grant related privilege

links:

* [! tutorialspoint mysql](https://www.tutorialspoint.com/mysql/)    
* [mysql configration](https://dev.mysql.com/doc/refman/5.5/en/server-configuration.html)


## chapter 9: Development Tools

Makefile option:

* `-` tells make to ignore any errors. For example, if you wanted to make a directory but wished to ignore any errors, perhaps because the directory might already exist, you just precede mkdir with a minus sign. You will see - in use a bit later in this chapter.

* `@` tells make not to print the command to standard output before executing it. This character is handy if you want to use echo to display some instructions.



Makefile macros:

* $?: List of prerequisites (files the target depends on) changed more recently than the current target
* `$@`: Name of the current target
* `$<`: Name of the current prerequisite
* `$*`: Name of the current prerequisite, without any suffix

Makefile build-in rules:
* You can ask make to print its built-in rules with the -p option. Th


# links
* [ ! The GNU C Library](http://www.gnu.org/software/libc/manual/html_node/index.html)
* [Linux system call references](http://syscalls.kernelgrok.com/)
* [POSIX OS header files references](https://en.wikipedia.org/wiki/Unistd.h)
* [c lange tutorialspoint](https://www.tutorialspoint.com/cprogramming/index.htm)
* [cpu GHZ means](http://smallbusiness.chron.com/ghz-mean-computer-processor-66857.html)

## todos

    bash set
    umask, how to use it.
    grep O_RDONLY inside usr/include dir
    clang: size_t 类型









    

