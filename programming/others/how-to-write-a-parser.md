# How to write a parser | 如何构建一个parser

Parser这个东西，以前认为很玄，感觉自己并不能搞懂，就没有看，直到最近看了一些资料，感觉没有那么难，只是过程相对繁琐。
首先，我看到的是[how to be a compiler - by,  kosamari](https://medium.com/@kosamari/how-to-be-a-compiler-make-a-compiler-with-javascript-4a8a13d473b4)
但是这篇文章说实话我并没有怎么看懂。
后来看到了[ReactNL 2016 James Kyle - How to Build a Compiler](https://www.youtube.com/watch?v=ZYFOWesCm_0)
这个视频, 再看了这个视频的代码[the super tiny compiler](https://github.com/CrazyFork/the-super-tiny-compiler) 顿时让我觉得有了些方向。基本的过程就是 tokenize-lexier-code generation, 这几步。
* tokenize 将文本提取成有意义的token标示, 
* lexier - 将生成的tokens parse成有效的 ast.
* 根据 ast 做转换或者生成代码

后来看了我们公司一位牛人写的 [thrift-parser](https://github.com/CrazyFork/thrift-parser), 基本上就是根据string的position的offset，和state去进行对应的parse, 这个库的特点就是移除了tokenize的过程, 直接parse生成ast。

之后又看了[thrust](https://github.com/CrazyFork/tokio-thrift/tree/reading_thrust)的代码, 这个库的代码由于语言的关系，看起来要舒服多了, 这个库会根据string的offset去生成对应的token，在具体的parse的过程的时候，是根据token的状态去解析合法与否。这个parser是作者完全手写的。后来的这个库的版本借助了rust的parser combinator的库 [nom](https://github.com/Geal/nom) 重写了parser部分应该。还有thrust这个库由于历史原因, 我阅读的这个版本已经不是新版了，好多thrift格式感觉和官网也不搭，不过并不影响这个库的质量。

后来又看了[fpscala](https://github.com/CrazyFork/fpinscala/tree/master/answers/src/main/scala/fpinscala/parsing) 中的parsing部分， 可以说3个仓库中，写的最复杂最抽象的就是这个了, 这个仓库的写法就是抽象Parser，然后根据情况组装对应的parser，combinator大致有or, flatMap, map等等，代码很难懂，很抽象。核心的代码就是

  ```scala
  // 可以看到 Parser 就是输入 ParseState 然后产生 Result 的过程, Result := Success | Failure
  // ParseState := {current_offset, string}
  type Parser[+A] = ParseState => Result[A]
  // run 的过程就是接受 Parser， string 最后生成结果
  def run[A](p: Parser[A])(s: String): Either[ParseError,A]
  ```

最后我要说下, 基本上各个语言都有生成compiler的工具库, 比如yacc, rust的nom等等。没有必要从头手写一个，当然如果你喜欢挑战的话。
  