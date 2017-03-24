


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
