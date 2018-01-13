# Webassembly
<img src="https://webassembly.github.io/spec/_static/webassembly.png" width="200">

这篇分享的目的, 只是简单给大家介绍下webassembly, 所以不会在细节上有过多的深入。以后如果谁有兴趣，可以自行查阅文档，或者阅读本篇下面的链接。

# overview

execution model:
* basic model

  <img src="https://2r4s9p1yi1fa2jd7j43zph8r-wpengine.netdna-ssl.com/files/2017/02/05-02-diagram_past01.png" width="600">

* JIT model
  <img src="https://2r4s9p1yi1fa2jd7j43zph8r-wpengine.netdna-ssl.com/files/2017/02/05-01-diagram_now01.png" width="600">

  * JIT optimzation, base line mark.
    * why some optimization may backfire.

  <img src="https://2r4s9p1yi1fa2jd7j43zph8r-wpengine.netdna-ssl.com/files/2017/02/02-07-jit11.png" width="600">


* webassembly:
  <img src="https://2r4s9p1yi1fa2jd7j43zph8r-wpengine.netdna-ssl.com/files/2017/02/05-03-diagram_future01.png" width="600">

## introduction to webassebly

  RAM, Register 都是作为一个存储工具. 

  <img src="https://2r4s9p1yi1fa2jd7j43zph8r-wpengine.netdna-ssl.com/files/2017/02/03-04-computer_architecture17.png" width="600">


a hello world example of assembly code:

```as
section     .text
global      _start                              ;must be declared for linker (ld)

_start:                                         ;tell linker entry point

    mov     edx,len                             ;message length
    mov    ecx,msg                             ;message to write
    mov     ebx,1                               ;file descriptor (stdout)
    mov     eax,4                               ;system call number (sys_write)
    int     0x80                                ;call kernel

    mov     eax,1                               ;system call number (sys_exit)
    int     0x80                                ;call kernel

section     .data

msg     db  'Hello, world!',0xa                 ;our dear string
len     equ $ - msg                             ;length of our dear string
```


传统的方式是通过compiler转换成对应的机器code, assembly 代码.但这种方式并不利于维护, 因为每一种语言都要实现下到不同target
的汇编实现
  <img src="https://2r4s9p1yi1fa2jd7j43zph8r-wpengine.netdna-ssl.com/files/2017/02/03-05-langs05.png" width="600">

IR, intemidiate Representation, LLVM

  <img src="https://2r4s9p1yi1fa2jd7j43zph8r-wpengine.netdna-ssl.com/files/2017/02/03-07-langs09.png" width="600">

webassembly:

<img src="https://2r4s9p1yi1fa2jd7j43zph8r-wpengine.netdna-ssl.com/files/2017/02/04-02-langs08.png" width="600">

<img src="https://2r4s9p1yi1fa2jd7j43zph8r-wpengine.netdna-ssl.com/files/2017/02/04-03-toolchain07.png" width="600">

<img src="https://mdn.mozillademos.org/files/14647/emscripten-diagram.png" width="600">
a add example of webassebmly: 
  * [Understanding WebAssembly text format](https://mdn.mozillademos.org/files/14647/emscripten-diagram.png)

  *http://mbebenita.github.io/WasmExplorer/

  ```c
  int add42(int num) {
    return num + 42;
  }


  ```


```webassembly
sub rsp, 8                            ; 0x000000 48 83 ec 08
mov ecx, edi                          ; 0x000004 8b cf
mov eax, ecx                          ; 0x000006 8b c1
add eax, 0x2a                         ; 0x000008 83 c0 2a
nop                                   ; 0x00000b 66 90
add rsp, 8                            ; 0x00000d 48 83 c4 08
ret                                   ; 0x000011 c3
```

.wasm file:

* a compiled wasm file like this:
```
00 61 73 6D 0D 00 00 00 01 86 80 80 80 00 01 60
01 7F 01 7F 03 82 80 80 80 00 01 00 04 84 80 80
80 00 01 70 00 00 05 83 80 80 80 00 01 00 01 06
81 80 80 80 00 00 07 96 80 80 80 00 02 06 6D 65
6D 6F 72 79 02 00 09 5F 5A 35 61 64 64 34 32 69
00 00 0A 8D 80 80 80 00 01 87 80 80 80 00 00 20
00 41 2A 6A 0B
```

<img src="https://2r4s9p1yi1fa2jd7j43zph8r-wpengine.netdna-ssl.com/files/2017/02/04-06-hex_binary_asm01.png" width="600">

<img src="https://2r4s9p1yi1fa2jd7j43zph8r-wpengine.netdna-ssl.com/files/2017/02/04-07-hex_binary_asm02.png" width="600">

loading wasm module into js:

```js
function fetchAndInstantiate(url, importObject) {
  return fetch(url).then(response =>
    response.arrayBuffer()
  ).then(bytes =>
    WebAssembly.instantiate(bytes, importObject)
  ).then(results =>
    results.instance
    // results.instance.export.doSomething()
  );
}
```

* webassembly limitation:
  * no actual file io api, need simulation
  * can only pass integer between js & webassembly module.
    * using `Table` to pass more complex data structor.
  * no multithread support yet.
  * no network/dom capability

* 什么是webassembly:
webassembly is a conceptional machine, it's stack machine, it's a compiled target, it's not a language.

webassembly 的前景: (目前核心浏览器都支持webassembly)

* 嵌入语言的vm, 可以在前端写其他语言, PHP, Go, Java, C#,
  * https://github.com/SteveSanderson/Blazor
  * 这意味着以后的代码同构的另一种方式(代码复用, etc)
* 更强大的web性能, 当前web还是不能很好处理图片压缩, Encryption, Custom Media Decoder(自定义媒体的类型).
* 游戏引擎
* web平台对于Native Mobile具有更强的竞争能力,
* React 等核心库如果能够重写, 则意味着更小的前端代码和更快的性能。

WebAssembly可以说是近些年，前端技术出现的比较大的变革了。


## [Emscripten: An LLVM-to-JavaScript Compiler](https://github.com/kripken/emscripten)

* [installation](http://webassembly.org/getting-started/developers-guide/)

* standalone and glued js + wasm output.





## Advanced:
* [Sections of the module](https://hacks.mozilla.org/2017/02/creating-and-working-with-webassembly-modules/)
* cache

* two types of translation IR.
>There’s another tool called Emscripten which is a bit easier to use at the moment. It has its own back-end that can produce WebAssembly by compiling to another target (called asm.js) and then converting that to WebAssembly. It uses LLVM under the hood, though, so you can switch between the two back-ends from Emscripten.

* webassembly don't specify registers
>Even though WebAssembly is specified in terms of a stack machine, that’s not how it works on the physical machine. When the browser translates WebAssembly to the machine code for the machine the browser is running on, it will use registers. Since the WebAssembly code doesn’t specify registers, it gives the browser more flexibility to use the best register allocation for that machine.



links:
* https://hacks.mozilla.org/category/code-cartoons/a-cartoon-intro-to-webassembly/
* https://developer.mozilla.org/en-US/docs/WebAssembly
* https://github.com/mdn/webassembly-examples
* https://github.com/Hanks10100/wasm-examples
* https://github.com/inter-action/simple-webassembly-demo
* rust
  * https://medium.com/@ianjsikes/get-started-with-rust-webassembly-and-webpack-58d28e219635
  * https://www.hellorust.com/news/native-wasm-target.html