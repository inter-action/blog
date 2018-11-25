# Practical.Vim.2nd.Edition.2015.10

moving around:
* `yap`: yank a whole paragraph. you can test text range with `vap`

delete & copy & insert

* `s` = `xi<Esc>, or cl`


repeat:
* `;`: 用于F, f之后
* `.`: any change 
* `n, N`: search
* `@@`: repeat last macro
* `g&`: repeat last subsitude


seperate window
* `q\` `q?`: open search window
* 

others:
* `<C-r><C-w>` 会把光标下面的word粘贴到prompt里边


## CHAPTER 2: Normal Mode
* `<C-a>, <C-x>`, 当光标在数字上的时候, 这两个命令分别对数字进行加减, 如果不在数字上, 它会寻找第一个数字
* 

## CHAPTER 6: Manage Multiple Files

```
// group *.rb 
:args *.rb
// open first file in the group
:first 
// repeat macro a for all files in this group
:argdo normal @a 

// write all files 
:wall
```

#chapter 07: Manage Multiple Files

* `buffer`: editing an in-memory representation of a file, which is called a buffer in Vim’s terminology.

#chapter 08: Navigate Inside Files with Motions


## CHAPTER 10: Copy and Paste

* unnamed register, `""`, 默认的register.
* black hole register, `"_`
* The Yank Register ("0), 任何`y`开头的yank会到 unnamed register & yank register
* The System Clipboard ("+) 
* The Expression Register ("=): 
    * 这个register不仅可以输出算数表达式, 还能输出vim变量, insert mode 可以用`<C-r>=`activate.

tips:
* copy & delete & paste 是我写代码比较常规的操作, 通常我都是paste到前面用P,然后delete掉后面.书中还有一种操作方式是copy, 然后用v去选择, 然后去paste达到替换的目的. 被替换的文字会put到unnamed register下面


## CHAPTER 11: Macros
* `<cout>@<register>`: playback with a count
* `:normal @a`: 用v选中后用 normal mode 去执行repeat.
* macro 和 yank 用的是同一个register, 这意味着macro的内容可以被打印出来被彼此覆盖. 修改之后可以再被yank到register上去执行修改过的macro
* `:let @a=substitute(@a, '\~', 'vU', 'g')`, change register content with a function


## CHAPTER 12: Matching Patterns and Literals

vim 的正则还是有点怪, 是posix标准, 不是perl
* `/\Va.k.a.`, `\V`后面的pattern除了`\`其他都按literal走. 不按special meaning
* `\v`, 后面的pattern按照perl的类似规则走. 但还是有细微区别
* `\c` vs `\C`, search 大小写开关
* `/\v<(\w+)\_s+\1>`, `\1` reference first group match.
* `/\v%(And|D)rew Neil`, `%`前面表示不capture 后面的group
* `<>`: word boundary 类似 perl 正则里的 `\b`
* `\zs` & `\ze`, 用于高亮match中的sub match. p195有介绍

## CHAPTER 13: Search

* `gn`, 作为一个motion指向当前match的character, like `hjkl`


## CHAPTER 14: Substitution
* `:[range]s[ubstitute]/{pattern}/{string}/[flags]`
* `:%s//\=@0/g`, `@0` reference 0 register
* > Here’s the good news: leaving the search field blank tells Vim to use the cur- rent pattern.
* flags
    * `c`: 替换时候给予提示
    * 
* eg:

    ```
    // search Pragmatic in "Pragmatic Vim"
    /Pragmatic\ze Vim
    // replace it with Practical
    :%s//Practical/g
    ```

## CHAPTER 15: Global Commands
* `:global :g` & `:vglobal :v`, 用于执行全局命令, `:g` 会搜索所有match pattern 的行, 然后执行cmd命令, `:v`, 是invserse, 会执行所有unmatch的行数.
* `sort`: vim 自带的sort cmd
* `  :g/{pattern}/[range][cmd]`


# plugins

* https://github.com/tpope/vim-abolish
