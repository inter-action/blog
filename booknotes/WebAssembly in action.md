
# WebAssebmly In Action

说下这本书的总览吧, 因为在ipad上看完的, 所以也没啥心情做笔记了.加上这本书其实也没有必要.
这本书主要是给用C/C++, Emscripten 工具链的读者来看的. 虽然包括了 WebAssebmly 通用的
知识. 

就其工具链的部分, 我是几乎没有在意的. 估计看了也会忘. 也没啥太大意义. 核心的是 Emscripten
工具会内置一些方法, 帮你生成 Bootstrap JS 文件. 这个文件核心就是帮助你去管理内存, 将 C 的
内存管理函数映射到 WebAssembly 上. 因为 WebAssembly 的内存是 Linear 的. 所以肯定会有更高级
的内存抽象技术建立在其之上. 使其用起来方便.

关于 WebAssembly 核心的部分, 这本书讲的比较有意义的就是 
* WebAssembly 的 s-expression. 
* 还有 binary 的layout 的规范, 包括 global, data, type section 等等. 
* WebAssembly 的 stack machine 的模型, 包括指令集的介绍.

对于我个人而言, 工具链不是我在意的部分, 所以本书对我来说有价值的就是 WebAssembly 核心的部分.
因为我想做的是用 rust 来写 WebAssembly 我理解其生态和工具都比这个 emscripten 用着舒服.
