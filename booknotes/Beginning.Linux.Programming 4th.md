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


## chapter 11: Processes and Signals

what is process, what is it made of:
>Additionally, a process has its own stack space, used for local variables in functions and for controlling function calls and returns. It also has its own environment space, containing environment variables that may be established solely for this process to use, as you saw with putenv and getenv in Chapter 4. A process must also maintain its own program counter, a record of where it has gotten to in its execution, which is the execution thread. In the next chapter you will see that when you use threads, processes can have more than one thread of execution.

how process work:

>In general, each process is started by another process known as its parent process. A process so started is known as a child process. When Linux starts, it runs a single program, the prime ancestor and process number 1, init. This is, if you like, the operating system process manager and the grandparent of all processes. Other system processes you’ll meet soon are started by init or by other processes started by init.


### Process Scheduling:

* The Linux kernel uses a process scheduler to decide which process will receive the next time slice.

the "nice" value:
>In a multitasking system such as Linux where several programs are likely to be competing for the same resource, programs that perform short bursts of work and pause for input are considered better behaved than those that hog the processor by continually calculating some value or continually query- ing the system to see if new input is available. Well-behaved programs are termed nice programs, and in a sense this “niceness” can be measured. The operating system determines the priority of a process based on a “nice” value, which defaults to 0, and on the behavior of the program. Programs that run for long periods without pausing generally get lower priorities. Programs that pause while, for exam- ple, waiting for input, get rewarded. This helps keep a program that interacts with the user respon- sive; while it is waiting for some input from the user, the system increases its priority, so that when
it’s ready to resume, it has a high priority. You can set the process nice value using nice and adjust it using renice. The nice command increases the nice value of a process by 10, giving it a lower prior- ity. You can view the nice values of active processes using the –l or –f (for long output) option to ps. The value you are interested in is shown in the NI (nice) column.


### Starting New Processes:

`system ` call: is inefficent, it requires specific shell to load another process.

`exec ` call: efficent, it start execute new process & quit current one.

`wait` call: this api wait current child process fork to be done in order to do some child process stats check.


### Zombie Processes

how is it created:
>Using fork to create processes can be very useful, but you must keep track of child processes. When a child process terminates, an association with its parent survives until the parent in turn either terminates normally or calls wait. The child process entry in the process table is therefore not freed up immediately. Although no longer active, the child process is still in the system because its exit code needs to be stored in case the parent subsequently calls wait. It becomes what is known as defunct, or a zombie process.

 

### Signals

how singal affect processes:
* If a process receives one of these signals without first arranging to catch it, the process will be termi- nated immediately. Usually, a core dump file is created. This file, called core and placed in the current directory, is an image of the process that can be useful in debugging.

* The default action for the signals in the following table is abnormal termination of the process with all the consequences of _exit (which is like exit but performs no cleanup before returning to the kernel).

* a process can also be suspend be some specific signals like(SIGSTOP, SIGTSTP...)



>You must program your signals carefully, because there are a number of “race conditions” that can occur in programs that use them.

Race condition:
* Signal can be catched or restored. 
* When a process is on handling a singal there's a chance that a new signal could arrive
    * the default mechanism is the handling function wouldn't be trigger twice, because the current signal has been added to the signal mask set which filter out signals that matched the ones already been catched.


the signal mask set:
>Ordinarily, when a signal handling function is being executed, the signal received is added to the process signal mask for the duration of the handling function. This prevents a subsequent occurrence of the same signal, causing the signal handling function to run again. If the function is not re-entrant, hav- ing it called by another occurrence of a signal before it finishes handling the first may cause problems. If, however, the SA_NODEFER flag is set, the signal mask is not altered when it receives this signal.



## chapter 12: POSIX Threads
skiped.

## chapter 13:  Inter-Process Communication: Pipes

desc:

* popen, pipe


* [Shell Parameter Expansion] (https://www.gnu.org/software/bash/manual/html_node/Shell-Parameter-Expansion.html)

    > `cat popen*.c | wc -l` In Linux (as in all UNIX-like systems), all parameter expansion is done by the shell, so invoking the shell to parse the command string before the program is invoked allows any shell expansion, such as determining what files *.c actually refers to, to be done before the program starts. This is often quite useful, and it allows complex shell commands to be started with popen. Other process creation functions, such as execl, can be much more complex to invoke, because the calling process has to perform its own shell expansion

    书中这段说的意思是 *.c 要被shell完全解析之后传给底层的program处理


* pipe & dup api 的联合使用:
    in order to be able to communicate from pipe with child process created by using `execlp` api, the book demo
    this technique to handle such situation.

    dup: create a new file descriptor pointed to the same underlying file pointed by the target file dscriptor argument.

    >The trick is knowing that the standard input file descriptor is always 0 and that dup always returns a new file descriptor using the lowest available number. By first closing file descriptor 0 and then calling dup, the new file descriptor will have the number 0. Because the new descriptor is a duplicate of an existing one, standard input will have been changed to access the file or pipe whose file descriptor you passed to dup. You will have created two file descriptors that refer to the same file or pipe, and one of them will be the standard input.


* FIFO(named pipe) atomic write among many processes:
    >Well, if you ensure that all your write requests are to a blocking FIFO and are less than PIPE_BUF bytes in size, the system will ensure that data never gets interleaved.

    >Linux arranges the scheduling of the two processes so that they both run when they can and are blocked when they can’t. Thus, the writer is blocked when the pipe is full, and the reader is blocked when the pipe is empty.


  links:
  *[linux mode 0777 vs 777](http://unix.stackexchange.com/questions/103413/is-there-any-difference-between-mode-value-0777-and-777)  


## chapter 14:  Semaphores, Shared Memory, and Message Queues

Semaphores用的太少了，书中仅仅是当做锁用的，如果是当做锁用的又有什么优势呢。
进程锁，这算是优势吧。
书中这一块没怎么看懂，以后用到了再说。

ipcs linux command: 用于查看系统ipc的状态的
ipc（inter process communication?）:
* Semaphores
* shared memory
* message Queues


## chapter 15: Sockets

* local sockets: 
    > Local sockets are given a filename in the Linux file system, often to be found in /tmp or /usr/tmp. 

* 


### socket attributes:

>Sockets are characterized by three attributes: domain, type, and protocol. They also have an address used as their name. The formats of the addresses vary depending on the domain, also known as the protocol family. Each protocol family can use one or more address families to define the address format.

* domain:
    >Domains specify the network medium that the socket communication will use.
    * AF_INET:
        * AF_INET: internet
        * AF_INET6: ipv6 internet
    * AF_UNIX:  which can be used by sockets based on a single computer that perhaps isn’t networked. 


* types:
    * internet domain:
        * streams(tcp, SOCK_STREAM)
            >Stream sockets (in some ways similar to standard input/output streams) provide a connection that is a sequenced and reliable two-way byte stream. Thus, data sent is guaranteed not to be lost, duplicated, or reordered without an indication that an error has occurred.

        * datagrams(UDP, SOCK_DGRAM)
            >In contrast, a datagram socket, specified by the type SOCK_DGRAM, doesn’t establish and maintain a con- nection. There is also a limit on the size of a datagram that can be sent. It’s transmitted as a single network message that may get lost, duplicated, or arrive out of sequence — ahead of datagrams sent after it.

* Protocols:
    >Where the underlying transport mechanism allows for more than one protocol to provide the requested socket type, you can select a specific protocol for a socket.
    



如何创建一个 server socket:
* create a unnamed socket
* bind socket to address
* listen sockets connection (allocate a connection queue)
* accept connection
* read from sockets

if num of pending request exceeds the maximum number specified in the `listen` api, the exceeded request will simply be reject.
`accept` would block if no request is send to server.



### Network Information:

这里有些获得host info的api都是通过查找 `/etc/host` 文件和DNS信息获得的。具体可以看链接 - https://linux.die.net/man/5/hosts


### The Internet Daemon (xinetd/inetd):

这货表示linux的web 服务可以在用户请求的时候才启动，不用一直运行着。这就需要 internet daemon来监听多个端口，然后又它来启动服务。


### Select 机制
linux Select 机制是可以同时监听 `read_fds, write_fds, error_fds` 3中fd集合。任何一个状态可以操作则会停止block，然后走到下边的代码进行处理。下面的代码通常需要检测哪个fd被设置了，然后操作对应的fd。如果设置了timeout，对应timeout时间内没有触发也会走到下边，如果没有timeout则会一直block掉。

```c
// @return -1, 表示出错
// @param nfds, The nfds argument specifies the number of file descriptors to be tested, and descriptors from 0 to nfds-1 are considered.  
int select(int nfds, fd_set *readfds, fd_set *writefds, fd_set *errorfds, struct timeval *timeout);
```


### notes:
* port number:
    > Usually, port numbers less than 1024 are reserved for system services and may only be served by processes with superuser privileges.

* file socket:  The file system socket has the disadvantage that, unless the author uses an absolute pathname, it’s created in the server program’s current directory. 
    


## chapter 18: Standards for Linux
这章讲了些关于linux标准化的一些事情。C语言，gcc，各种standard的组织和意思。自启动service，File System hierarchy的定义等等。

* C 语言的标准化进程
* GCC的选项和标准，参数解释等等
* Interfaces and the Linux Standards Base:
    * 总的来说 Linux 的标准化是由Linux Standards Base这个组织定义的。
        * 程序接口方面（LSB Standard Libraries）这个组织定义了两种
            * Linux 自己的接口标准，实现
            * Unix 的接口标准，这一部分由于历史原因比较复杂。Unix标准化的历史问题
                * 简单来说 Unix 的第一个流行的标准是POSIX(IEEE 1003),但这个标准太保守，有很多功能不满足
                * 然后 1994年 X/OPEN 这个公司做了一个 X/OPEN CAE, or Common Applications Environment 的标准，是上边POSIX标准的超集
                * 然后 X/OPEN 这个组织和OSF合并了，成立了 The Open Group, http://www.opengroup.org/.


        * LSB System Initialization:
            * 标准定义了启动状态各个阶段的数值
            * 定义了 `/etc/init.d` 作为服务的自启动目录，对应的服务的命名规范由下面的组织规定，并定义了启动服务需要提供 `start, stop, restart ...` 语义参数做的事情
                * The Linux Assigned Names And Numbers Authority (LANANA), which you can find at http://www.lanana.org/


        * The Filesystem Hierarchy Standard:
            * 文件系统组织结构
            * The last of the standards we are going to look at in this chapter is the Filesystem Hierarchy Standard (FHS), which you can find at http://www.pathname.com/fhs/.
            * 定义了文件系统如何组织，构成的。



# links
* [unix os tutorial](http://www.tutorialspoint.com/unix/index.htm)
* [! using libraries](https://rufflewind.com/2017-02-25/using-libraries)
* [! c_standard_library](https://www.tutorialspoint.com/c_standard_library/stdio_h.htm)
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









    

