
# todo

* https://en.wikipedia.org/wiki/Shift-reduce_parser
  >A table-driven parser has all of its knowledge about the grammar encoded into unchanging data called parser tables. 

* https://en.wikipedia.org/wiki/LR_parser
  > Most LR parsers are table driven. The parser's program code is a simple generic loop that is the same for all grammars and languages. The knowledge of the grammar and its syntactic implications are encoded into unchanging data tables called parse tables. Entries in a table show whether to shift or reduce (and by which grammar rule), for every legal combination of parser state and lookahead symbol. The parse tables also tell how to compute the next state, given just a current state and a next symbol
    * 说白了就是 parser 的loop是通用的, 真正不同的是不用语法生成的 table

  * LR parser & table 的作用方式
    >The LR parser begins with a nearly empty parse stack containing just the start state 0, and with the lookahead holding the input stream's first scanned symbol. The parser then repeats the following loop step until done, or stuck on a syntax error:

    >The topmost state on the parse stack is some state s, and the current lookahead is some terminal symbol t. Look up the next parser action from row s and column t of the Lookahead Action table. That action is either Shift, Reduce, Done, or Error:

  * LL(k)/LR(k), 都表示需要 lookahead 的 k 个token, 更大的k能解析更复杂的语法



# youtube notes
* G=(V, T, P, S), @https://www.youtube.com/watch?v=WccZQSERfCM&index=2&list=PLEbnTDJUr_IcPtUXFy2b1sGRPsLFMghhS
  * G: grammars
  * V: variable, V left hand side, 
  * T: Terminal
  * S: symbol


  