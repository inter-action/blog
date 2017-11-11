# NOTICE & SUMMERY

总体上这本书，是3本书的合集，应该。书的时间也比较久，出的，导致有些库(package)里边介绍的，都已经过时了，仓库都很难找到了。书中
提供的价值和参考并没有我期望的那么大（可能好多内容我已经熟悉了）, 就当对前年golang的复习了。也就算ok了整体上的内容。 3.5-4星吧。而且书中有些细小的错误，实例代码也有。

本书的代码我也有Fork，加了些标注， 地址: [https://github.com/CrazyFork/Go-Building-Web-Applications](https://github.com/CrazyFork/Go-Building-Web-Applications)


# golang
### golang basics

* go get:
    * go get is way to install package dep. all your dep will be installed into dir specified in GOPATH variable
    * go get always gets code on master branch, 

* var vs :=
    * `var` is basically same with `:=` . used to declared mutable variable

* interface{}, can be the type of any other complex type

* coersion in go
    * type(srcType) or srcType.(type)

* mysql query should call `Close` after done .


# Book notes

## module 1
### chapter 1: Introducing and Setting Up Go


### chapter 2: Serving and Routing

### chapter 3: Connecting to Data




### chapter 4: Using Templates
* this chapter demos:
    *how use template, unescape html content, define func that be used in template.

* by default go template escapes any html content, if you dont want this behaviour, you can reference page 38 which contains detailed instructions
    that tell you how to do it


### chapter 5: Frontend Integration with RESTful APIs

* demos:
    * create api endpoint
    * secure api request with https



### chapter 6: Sessions and Cookies

* demos:
    * 如何用mysql db 来实现Session storage.
        * 
    * flash message 实现的机制: （书中这一部分介绍的有点混乱）
        * 第一次请求中在 controller 的 cookie 中设置值, 加上合适的过期时间 60s
        * 然后在第二次请求中清除cookie
        * 其实这部分的实现可以先加两个middleware, 第一个拦截对应 flash_ 开头的cookie,并清除掉, 第二个middleware可以在此基础上设置 flash 的值


links:
* [HTTPS and Go](https://www.kaihag.com/https-and-go/)



### chapter 7: Microservices and Communication

demos:
* how to use go to build a rabbitmq server/publisher
* how to use python to build a rabbitmq client/consumer



* micro service
    * pros:
        * more robust, one failure service wouldn't take down the entire system
    * cons:
        * decouple nature make it not easily to spot a failure service

notes:
* message quene vs direct RESTFul API Call
    * message quene guarantee a payload's recieve. It doesn't require recipient to be available all the time.



### chapter 8: Logging and Testing


demos:
* logging tool
    * create multiple file logger
* unit test
    * simple unit test with go
    * httptest module


todo:
* login module
* go  vet tool?

### chapter 9: Security

demos:
* 这章就是介绍了如何处理安全问题, 几种常见的安全漏洞攻击的方式，还有推荐了 http://github.com/unrolled/secure 库来处理安全问题


Cross-site request forgery:
就是第三方网站把你的网站的表单地址嵌入到了他们网站里边，然后欺骗用户去点击, 解决方式是在 form 中加一个 token, 然后在表单
处理的时候验证这个token

 
notes:
*  There's nothing preventing you
from using another port for HTTPS; the bene t here is that the browser directs https:// requests to port 443 by default, just as it directs HTTP requests to ports 80 and sometimes fallback to port 8080


todo: 
HTTP Strict Transport Security (HSTS)



### chapter 10: Caching, Proxies and Improved Performance

demos:
* reverse proxy with nginx
* implement file based cache system using go
* memcache
* http2, which go net/http module supports out of box


todo:
godoc.org/github.com/ bradfitz/gomemcache/memcache


## module 2

### chapter 1: Chat Application with Web Sockets


demos:
* tracing
    * Tracing is a practice by which we log or print key steps in the  ow of a program to make what is going on under the covers visible.
* write a chatroom utilizing
    * websocket
    * go concurrency with select & channel
* write a package 


notes:
* defer func() { r.leave <- client }(), page 133
    * create a function and invoke before the scope is about done

* sync.Once 

* channel:
    * You can think of channels as an in-memory thread-safe message queue where senders pass data and receivers read data in a non-blocking, thread-safe way.


### chapter 2: Adding Authentication

demos:
* how to create a http handler middleware using go to do the authentication work
* how to use OAuth2 library.



todo:
* http://github.com/stretchr/objx


### chapter 3: Three Ways to Implement Profile Pictures

demos:
* TDD
    * mock api using testify package
* create profile image using Gavatar
* uploading image file
* demo how to leverage interface to create a flexiable package design



### chapter 4: Command-line Tools to Find Domain Names

demos:
* how to build cli tool
* how to use std io to pipe data
* how to use golang cmd to pipe subcommand together



### chapter 5: Building Distributed Systems and Working with Flexible Data

demos:
* how to capture signal in order to gracefully shutdown program
    ```golang
    termChan := make(chan os.Signal, 1)
    signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
    ```

* how to use nsq to subscribe live info feed
* how to use mgo to manipulate mongodb
* using mutx & channel 
* 

nsq: (A realtime distributed messaging platform)
* github - https://github.com/nsqio/nsq

* `nsqlookupd` start nsq deamon
* `nsqd --lookupd-tcp-address=127.0.0.1:4160` 


notes:
* Because the deferred statements are run in LIFO (last in, first out) order, the first function we defer will be the last function to be executed, which is why the first thing we do in the main function is to defer the exiting code. p260
* p263, use `timer.AfterFunc & timer.Reset` to create a `setInterval` effect

todo:
p265
https://github.com/joeshaw/envdecode
nsqd


### chapter 6: Exposing Data and Functionality through a RESTful Data Web Service API


demos:
* create a helper package that aids sharing data through out entire request
    * p271, 实现方式可以好好借鉴下
* use decorator pattern to extract common logic, decorator http handlers
    ```golang
    func withCORS(fn http.HandlerFunc) http.HandlerFunc {
        return func(w http.ResponseWriter, r *http.Request) {
            w.Header().Set("Access-Control-Allow-Origin", "*")
            w.Header().Set("Access-Control-Expose-Headers", "Location")
            fn(w, r)
        } 
    }
    ```

* how to gracefully shutdown server using [graceful](https://github.com/tylerb/graceful) package @p281
* how to create simple RESTFul api and with CORS.



notes:
* sync.RWLock:
    * RWLock, RLock 会和 WLock 排斥，也就是说 RLock 上锁就不能有 WLock, 而 WLock 上锁就不能有RLock
    * RLock 之间是没有锁的， 而WLock之间是有锁的

* p277 的对于http请求的encode，decode & response 的封装需要注意下

todo:

http.ServeMux



### chapter 7: Random Recommendations Web Service

demos:
* implement enumerators in Go, using Iota
* waitGroup with go routine
* 

todos:

* p317, The url.Values type is actually a map[string][]string type, which is why we use make rather than new.
* go random.seed()


### chapter 8: Filesystem Backup

demos:
* how to create a cli tool project with go



notes:
* common package structure
    ```plain
    -backup # root folder for other programmer to use your package.
        -cmds # cli command tools folder
            -<cli_cmd_name>
    ```



## module 3 - Mastering Concurrency in Go
### chapter 1 - An Introduction to Concurrency in Go

demos:
* sync.WaitGroup usage
* 


notes:
* Gosched: 将当前线程让出，让其他go routine 执行完，之后的代码会自动执行，@p361
* By default, channels are unbuffered, which means they will accept anything sent on them if there is a channel ready to receive
    * channel need to be closed after done.
    * The inverse is also true; a deadlock can result from a channel continuing without sending anything, leaving its receiving channel hanging indefinitely.



todos:
p359
p373-A nil or uninitialized channel will always get blocked




### chapter 2 - Understanding the Concurrency Model
感觉这章读的没什么用，应该是所有的概念都熟悉的缘故


### chapter 3 - Developing a Concurrent Strategy

demos:
* what is race condition
* how to avoid race condtion by using go provided tools
    * mutex
    * channel
    * go race checking tool






notes:
* The race detector is guaranteed to not produce false positives, so you can take the results as strong evidence that there is a potential problem in your code
    * `go run -race race-test.go`
* load test:
    > We can then run a concurrent load tester against the action. There are a number of such testers available, including Apache's ab and Siege. For our purposes, we'll use JMeter, primarily because it permits us to test against multiple URLs concurrently.

*  a note on string type:
    > The string type is the sole immutable type in Go; this is noteworthy if you end up assigning and reassigning values to a string. Assuming that memory is yielded after a string is converted to a copy, this is not a problem. However, in Go (and a couple of other languages), it's entirely possible to keep the original value in memory.

    > If you ever encounter a place where a string is logical, but you want or could benefit from a mutable type, consider a byte slice instead.



todos:
* bufferred channel vs unbuffered channel


### chapter 4: Data Integrity in an Application

demos:
* go with c, how to call each other in their domain
* Getting even lower – assembly in Go, 
    * this section is skipped due to unnecessary
* demos different concurrency pattern
    * Distributed shared memory
    * First-in-first-out – PRAM (Pipelined RAM (PRAM) )
    * master-slave model
    * producer-consumer
    * leader-follower

* introduce memcached
* introduce Circuit


notes:

todos:


### chapter 5: Locks, Blocks, and Better Channels

demos:
* Pprof – yet another awesome tool, 这一节之前好像就介绍了 channel 的几种用法, 然后就没搞什么了
    * channel type of struct, interface , function & channel
    * use channel to create a block(serialized) execution
    * 
* Pprof tool, how to use it.



notes:
* QuoteMeta: escape regular expression meta characters.


todos:
p474 到底说的是啥

### chapter 6: C10K – A Non-blocking Web Server in Go

demos:
* this chapter starts off by introducing the c10k problem, that is serving 10k concurrent connection on a single server
    then the author compared a bare file server performance between apache web server and a go implemented simple file server.
    after that, the author then 
    * added dynamic template parsing to the server, then benchmarked it. 
    * precompiled the template, then benchmarek it
    * added mysql db, then benchedmarked it
    * suggested futhur optimization by limited Read & Write timeout



notes:

todos:
apache ab tool


### chapter 7: Performance and Scalability


demos:
* more example usage of Pprof tool
* Distributed go, some topologies (with some level implementation of course), including:
    * star
    * mesh
    * The Publish and Subscribe model
        * Serialized data
        * Remote code execution (not recommend, various security reason )
    * 

* Message Passing Interface:
    * MPI was borne from early 1990s academia as a standard for distributed communication.
    * it is still a protocol, so it's largely language agnostic.
    * For the most part, MPI is used by the scientic community; it is a highly concurrent and analogous method for building large-scale distributed systems.

* some profiling tool, most of them are outdated

* memory , stack info using `runtime.MemProfileRecord` 

notes:

todos:
`pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)`
Google App Engine





### chapter 8: Concurrent Application Architecture

demos:

notes:
* about mongodb: 
    > Mongo has also gotten a bit of a bad rap as it pertains to fault tolerance and data loss, but this is equally true of many NoSQL solutions. 
* give brief introduction on various nosql, the outstanding ones are:
    * redis
    * cassandra
    * couchbase

* create a example project that versions file changes:
    * a core file that monitor all file changes, then broadcast changes to interested clients
    * a example client that can make connection to server's message, then do the backup job
    * a cli tool that revision file(cover file from a old backup)




todos:


### chapter 9: Logging and Testing Concurrency in Go



demos:
* different logging methods
    * to console
    * to file
    * to a network
    * concurrent logging each goroutine with a dedicated logger
    * how utilize log4go package
    * how to capture panic error with recover

notes:

todos:




# dictionary
* crany: ill tempered
* mitigation: the action of reducing severity, seriousness, painfulness



### chapter 10: Logging and Testing Concurrency in Go


demos:
* a various use cases of channel
    * timeout, interval
    * typeless channel
    * notify channel close with `tomb` package

notes:

* go channel:
    * closed channel & nil channel
        * closed channel is channel that has been closed
        * a nil channel is channel that has never been initialized, 
            * select a nil channel will cause a permenent block.
            * send would sliently fail
    * can channel be closed at sender end? 
        * if so, can messege be guarantee to reach recieving end?
            * ANS: it seems so, but after that, sending to a closed channel causes panic

todos:
go channel: `case _,ok := <-conChan:`
    * difference between a buffered channel with a unbuffered channel






# libs
* [cors](https://github.com/fasterness/cors) - cors
* [graceful](https://github.com/tylerb/graceful) - Graceful is a Go package enabling graceful shutdown of an http.Handler server.
* [is](https://github.com/cheekybits/is) - A mini testing helper for Go
* [tomb](https://github.com/go-tomb/tomb/tree/v2) - https://github.com/go-tomb/tomb/tree/v2
* db
    * sql 
        * [mysql](https://github.com/go-sql-driver/mysql) - mysql db driver

    * nosql
        * [redis](https://github.com/go-redis/redis) - Type-safe Redis client for Golang
        * [mgo](https://github.com/go-mgo/mgo) - The MongoDB driver for Go
        

* distributed system
    * [nsq](https://github.com/nsqio/nsq) - A realtime distributed messaging platform 
    * [gomemcache](https://github.com/bradfitz/gomemcache) - Go Memcached client library
    * [Circuit](https://github.com/gocircuit/circuit) - Dynamic cloud orchestration
    * [consul](https://github.com/hashicorp/consul) - Service Discovery and Configuration Made Easy

    * serialization
        * [protobuf](https://github.com/golang/protobuf) - Go support for Google's protocol buffers
    * rpc
        * [gRpc](https://github.com/grpc/grpc-go) - The Go language implementation of gRPC. HTTP/2 based RPC



# Todos:


book src code - https://github.com/PacktPublishing/Go-Building-Web-Applications


gofmt with max column number ? break line seems not supported by go


https://marketplace.visualstudio.com/items?itemName=lukehoban.Go


https://github.com/go-sql-driver/mysql/


debug in go - https://github.com/derekparker/delve
    - https://github.com/derekparker/delve/blob/master/Documentation/installation/osx/install.md
    - brew install go-delve/delve/delve

mysql/root/j*p

https://github.com/avelino/awesome-go
