# Learn Vimscript the Hard Way


## chapter 6:

local leader
> Vim has a second "leader" key called "local leader". This is meant to be a prefix for mappings that only take effect for certain types of files, like Python files or HTML files.



## chapter 11: 

`:nnoremap <buffer> <leader>x dd`, `<buffer>` 表示在当前buffer定义, 在当前buffer生效

> The convention of using `<localleader>` for local mappings will prevent your plugin from overwriting someone else's `<leader>` mapping that they've painstakingly burned into their fingers over time.

local setting will have a higher priority

two ways of specify local
* `<localleader>`
* `<buffer>`



## chapter 12: autocmd

```
# syntax
# :autocmd <format> <filter> <command>
:autocmd BufNewFile *.txt :write

:help autocmd-events
```

## chapter 14: Autocommand Groups

```vim
augroup filetype_html
    autocmd!                        "clear cmd in same group
    autocmd FileType html nnoremap <buffer> <localleader>f Vatzf
augroup END
```


## Operator-Pending Mappings

`onoremap` map motion,  

```
# print foo(bar)
:onoremap in( :<c-u>normal! f(vi(<cr>   # <c-u> clear any text to head of line in command mode 
```


## More Operator-Pending Mappings

abount normal command
> The problem is that normal! doesn't recognize "special characters" like `<cr>`. There are a number of ways around this, but the easiest to use and read is execute.


execute command:

```
:onoremap ah :<c-u>execute "normal! ?^==\\+$\r:nohlsearch\rg_vk0"<cr>
```


## Variables

:help registers 

> Strings that start with a number are coerced to that number, otherwise they're coerced to 0.

`@ ` 开头的都是registers

## Comparisons

Read :help expr4 to see all the available comparison operators.

## Functions

> The second shows us that if a Vimscript function doesn't return a value, it 
implicitly returns 0.


## Strings
escape string 的方式
* 默认用 `\`, 比如 `\\`, `\<esc>` 表示按下 `esc` 键, `:execute "normal! mqA;\<esc>`q"`


> The one exception is that two single quotes in a row will produce one single quote.

`:help expr-quote`, see how to escape string

## functions

`:help functions`


## regular expression

> Vim has four different "modes" of parsing regular expressions! The default mode requires a backslash before the + character to make it mean "1 or more of the preceding character" instead of "a literal plus sign".


```
# 这句是有问题的, 因为 / 搜索会将后边都判断为 regex
:execute 'normal! gg/for .\+ in .\+:\<cr>'

# 解法, 拆开
:execute "normal! gg" . '/for .\+ in .\+:' . "\<cr>"
```

>We've split the pattern out from the rest of the command into its own literal string again, and this time we started the pattern with \v. This tells Vim to use its "very magic" regex parsing mode, which is pretty much the same as you're used to in any other programming language.



## Case Study: Grep Operator, Part Three

这一节的function定义还是非常有意思的.












