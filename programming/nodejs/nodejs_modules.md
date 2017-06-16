
## Stream

4 种类型的stream:
* Readable
* Writable
* Duplex
* Transform


Object Mode:
Stream Api 只接受两种类型 string or buffer. 如果将Object Mode打开，Stream可以接受除null以外的js对象。

Buffer:
Readable&Writable Stream 内部都有buffer，可以通过api获得。
highwaterMark: 这个限制了Readable, Writable内部的buffer可以容纳的限制。当Object Mode启用的时候，这个数值代表写入的对象书。
  * reach high water mark 的影响
    * writable.write(chunk) return false.
      * 继续写入数据的影响: 首先nodejs会buffer此次写入，后续写入:
        * While calling write() on a stream that is not draining is allowed, Node.js will buffer all written chunks until maximum memory usage occurs, at which point it will abort unconditionally. 
    * readable._read() 不会从底层数据流中读取数据

Duplex， Transform 
  * 这duplex流内部有两个buffer。
  * both are readable & writable

核心的class:
* stream.Writable

  * events:
    * drain: 如果 stream.write(chunk) returns false，当这个流可以再次被write的时候，drain会被触发。
    * unpipe: The 'unpipe' event is emitted when the stream.unpipe() method is called on a Readable stream, removing this Writable from its set of destinations.

  * methods:
    * cork: The writable.cork() method forces all written data to be buffered in memory. 
    * write: 的方法的callback的error有可能不正确, 文档有写, 优先用 event 的 error

  * example:

    ```js
    // respect backend pressure on write.
    function write (data, cb) {
      if (!stream.write(data)) {// data written in here would be buffered.
        stream.once('drain', cb) // prefered way to do next write
      } else {
        process.nextTick(cb) // prefered way to do next write
      }
    }

    // Wait for cb to be called before doing any other write.
    write('hello', () => {
      console.log('write completed, do more writes now')
    })
    ```

* stream.Readable: 
  * Two Modes:
    两种模式如何转换，文档有写
    * flow: 这种模式，数据会自动从底层读出，然后用EventEmitter发送出去
    * paused: 这种模式下，数据必须explicitly读出

  * methods:
    * .pipe(destination[, options]): option 中的end参数，如果设置成false，并且readable stream在读取的时候发生了错误, writable stream 不会自动关闭。
      > The process.stderr and process.stdout Writable streams are never closed until the Node.js process exits, regardless of the specified options.
    * read(size): 
      * normal mode:
        * return size bytes of Buffer. null if no data to read
        * return all the bytes inside internal buffer if this stream has ended.(eg. underlying file descriptor closed)
      * object mode:
        * return a single object
      
    * unshift:
      *  Unlike stream.push(chunk), stream.unshift(chunk) will not end the reading process by resetting the internal reading state of the stream.
    
    
* stream.Duplex: readable & writable
* stream.Transform
  * 实现 _transform & _flush 方法
  * transform stream 数据下游的数据如果没有被消费掉的时候需要代码处理掉
  * 错误处理是又 _transform 的 callback 来做

* stream.PassThrough

错误处理:

readable & writable: ???


[API for Stream Implementers](https://nodejs.org/api/stream.html#stream_api_for_stream_implementers): 
这一节演示了如何创建Stream。

Writable:
* 主要是继承Writable， 然后实现_write & _writev方法

Readable:
* 继承Readable， 实现 _read(size) 方法
  * size: 参数是optional的, 可以ignore掉
  * 有内容的时候需要调用 this.push(buffer, encoding) 方法， 这个方法会emit data这个event
  * _read 方法里边不能throw error, 要 emit error.
  * 当 readable.push() 返回 false 的时候，就不能调用 this.push 方法了， 直到 _read 再次被调用




## Buffer

Instances of the Buffer class are similar to arrays of integers but correspond to fixed-sized, raw memory allocations outside the V8 heap. The size of the Buffer is established when it is created and cannot be resized.

it's global


## Net

classes:
* net.Server:
* net.Socket:
  * 注意回收socket




## Http

classes:
* http.Agent:
  用于保存http连接池, 同一个请求(target domain, port , localAdress 为一个unique请求)保持着一个socket和对应的这个socket的请求队列。链接池的行为跟server端的行为也有关。使用完Agent之后应该注意关闭掉这个Agent

* http.ClientRequest:

  ```javascript
  // 如果需要公用agent，则在header中制定
  //使用方式
  let req = http.request(headers)
  req.once('response', ...)
  req.once('timeout', ...)
  req.once('error', ...)
  req.write(body)
  req.end()
  ```


## Query String: 
用于解析和构建url的path部分的请求参数，和url module中的parse方法类似




## DNS

node:
  dns module的方法有两种类型，一种是 lookup() 这个， 一种是其他。
  lookup 这个函数会走本机系统的dns机制，去解析对应的ip地址。这意味着 /etc/hosts 里边的配置会生效。
  而其他函数，则会通过网络请求去解析对应的结果。


## Native Modules:

* https://nodejs.org/api/addons.html

Write a native addon需要处理的点:

* V8: v8 c++ api 用于和 v8 引擎交互，创建js对象，执行等等，非io之类的东西
* libuv: IO。包括文件，网络等等
* Internal Node.js libraries：主要的item是node::ObjectWrap，用于 Wrapping C++ objects, 将c++对象暴露给js， 用于 new <ClassName> 这种方式的调用。ClassName的实现需要继承 node::ObjectWrap
* Node.js includes a number of other statically linked libraries including OpenSSL.（nodejs自己带的一些linked libraries）


所有的 native 库的header文件:

  /Users/interaction/.node-gyp/7.1.0/include/node/v8.h:


node-gyp:
  * install https://github.com/nodejs/node-gyp#installation
  * npm install -g node-gyp
  * node-gyp configure
  * node-gyp build


  ```js
  //bind.gyp
  {
    "targets": [
      {
        "target_name": "addon",         //编译完的二进制文件名称
        "sources": [ "hello.cc" ]
      }
    ]
  }

  ```



nan: 工程
  * https://github.com/nodejs/nan
  * 这是在v8 api上的做的一层封装，用于减小v8 api 改动受到的影响


v8:

### [v8 C++ api 关键概念解析](https://github.com/v8/v8/wiki/Embedder's%20Guide)
* Handles and Garbage Collection
  * 概述:
    * handles 这个机制是涉及Garbage Collection如何运作的。注意这一部分是C++代码和V8交互的问题，没有和V8的交互的部分的垃圾回收还是要自己做的
  * LocalHandle: 
    >Local handles are held on a stack and are deleted when the appropriate destructor is called. 

    * Local handles have the class `Local<SomeType>`.
  * PersistentHandle:
    > Persistent handles provide a reference to a heap-allocated JavaScript Object,

    * 两种类型:
      * A `UniquePersistent<SomeType>` handle relies on C++ constructors and destructors to manage the lifetime of the underlying object. 
      * A `Persistent<SomeType>` can be constructed with its constructor, but must be explicitly cleared with Persistent::Reset.     
  * 其他的handle

  * Scope: 用于方便管理各个handle的生命周期，Scope是handle的容器。如果scope失效了，里边的handle会自动调用自毁代码。

  * pitfall
    > you cannot return a local handle directly from a function that declares a handle scope.
    具体这一部分怎么弄看这个文档链接，有示例

  


  
* Context
  > In V8, a context is an execution environment that allows separate, unrelated, JavaScript applications to run in a single instance of V8. You must explicitly specify the context in which you want any JavaScript code to be run.

  * Context 第一次创建成本高，但后面由于v8的缓存机制会很便宜
  * Context 是重入可以覆盖的：
    >When you have created a context you can enter and exit it any number of times. While you are in context A you can also enter a different context, B, which means that you replace A as the current context with B. When you exit B then A is restored as the current context. This

* Templates
  >A template is a blueprint for JavaScript functions and objects in a context. You can use a template to wrap C++ functions and data structures within JavaScript objects so that they can be manipulated by JavaScript scripts.

  * 两种Template类型：
    * Function: 
      >A function template is the blueprint for a single function. You create a JavaScript instance of the template by calling the template's GetFunction method from within the context in which you wish to instantiate the JavaScript function. 
    * Object
      >Each function template has an associated object template. This is used to configure objects created with this function as their constructor.

    


* Accessors: get 或者 set 时候会被调用的c++函数
  > An accessor is a C++ callback that calculates and returns a value when an object property is accessed by a JavaScript script. Accessors are configured through an object template, using the `SetAccessor` method. 


* Interceptors:
  用于拦截 `jsobj.<proper_name>` 这种方式的请求的

  有两种: 
    * `named property interceptors` - called when accessing properties with string names. An example of this, in a browser environment, is `document.theFormName.elementName`.
    * `indexed property interceptors` - called when accessing indexed properties. An example of this, in a browser environment, is `document.forms.elements[0]`.


* Security Model

  * same origin policy:
    * protocal, domain, port 3者都相同才能认为相等

* Inheritance:
  * 创建inheritance的方式就是, 创建一个 function template 延后调用 PrototypeTemplate 方法

    ```c++

    Local<FunctionTemplate> biketemplate = FunctionTemplate::New(isolate);
    biketemplate->PrototypeTemplate().Set(
        String::NewFromUtf8(isolate, "wheels"),
        FunctionTemplate::New(isolate, MyWheelsMethodCallback)->GetFunction();
    )
    ```






* notes:

  * two types of export
    * eg.1, export as a sub variable
      ```c++
      void Init(Local<Object> exports) {
        NODE_SET_METHOD(exports, "add", Add);
        //const addon = require('./build/Release/addon');
        // addon.add(xx)
      }

      NODE_MODULE(addon, Init)
      ```

    * eg.2, export as global variable
      ```c++
      void Init(Local<Object> exports, Local<Object> module) {
        NODE_SET_METHOD(module, "exports", RunCallback);
        // const addon = require('./build/Release/addon');
        // addon(xx)
      }

      NODE_MODULE(addon, Init)
      ```

  * examples:
    ```c++
    // 创建返回一个js object， {msg: arg[0]}
    void CreateObject(const FunctionCallbackInfo<Value>& args) {
      Isolate* isolate = args.GetIsolate();// 创建一个隔离域

      Local<Object> obj = Object::New(isolate);//创建一个js object
      obj->Set(String::NewFromUtf8(isolate, "msg"), args[0]->ToString());// 对这个object设置值

      args.GetReturnValue().Set(obj);//返回
    }

    ```
