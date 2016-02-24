
本书主要使用 php 和 python 语言讲解演示如何构建 restful 的 webservice

在 router 中不要使用自增id, 自增id会暴露你的数据量, 也方便了爬虫的爬取, 可以使用 UUID 替代

## todos:

    POST vs PUT: FIGHT! @page 17




## chapter 4: Status Codes, Errors and Messages
如何正确的返回 status code, Errors, Message

status code:
不要用太多的 http response code

Message:

    {
        data:{},
        pagination:{}
    }


errors response: JSON-API 中有对于这个地方的规范 @ page 44
    {errors:[{<error>},...]}

    //error
    {
        code:String
        title: String
        details: String
        href: String
    }

// error structure defined by JSON-API
* "id" - A unique identifier for this particular occurrence of the problem.
* "href" - A URI that MAY yield further details about this particular occurrence of the problem.
* "status" - The HTTP status code applicable to this problem, expressed as a string value.
* "code" - An application-specific error code, expressed as a string value.
* "title" - A short, human-readable summary of the problem. It SHOULD NOT change from occurrence to occurrence of the problem, except for purposes of localization.
* "detail" - A human-readable explanation specific to this occurrence of the problem.
* "links" - Associated resources, which can be dereferenced from the request document.
* "path" - The relative path to the relevant attribute within the associated resource(s). Only appropriate for problems that apply to a single resource or type of resource.





## chapter 5: Endpoint Testing

讲了如何测试你的 web 接口, 并建议用 test-driven 的开发方式去做
也讲了 什么是 feature 什么是 scenario 


## chapter 6: Outputting Data

如何将 database 的数据输出出去

principle:

不能将 orm 层的数据直接输出到前端, controller 层有 orm query 都不被允许, 坏处主要有: @page 57

* Security: 敏感数据容易被直接输出
* Display: 数据需要格式化, 比如 mysql 中的 boolean 是 0, 1 你不希望客户端拿到的 boolean 值也是 0 or 1
* Stability: orm 底层的 field 肯定会改动, 但是你的接口使用者则不会期望他拿到的数据格式经常变动




## chapter 7:  Data Relationships

如何返回数据的relationship

strategies:

* Subresources: /users/x/contacts
* Foreign Key Arrays:

        users:{
            name:,
            "_links":{
                "contacts":[1, 2, 3]//id seqs
            }
            
        }

* Compound Documents(side loading)

        {
            users:{
                "_links":{
                    contacts:["1", "2"]
                }
                
            },

            "linked":{
                contacts:[{id:1}, {id: 2}]
            }
        }

* Embedded Documents:

        {
            users:{ contacts:[{id: "1", data:"this is a comment"}, ...]}
        }


## chapter 8: Debugging

如何调试接口

chrome extension: Postman

Debug Panel:
    Anything to do with a slow page return, silent fails, unexpected results, etc., needs more information, 
    and to do that you probably need another extension.

    chrome extention for this:
        Chrome Logger5 - Chrome Logger only for Python, PHP, Ruby, Node, .NET, CF and Go


network debugging:
    
* charles
* Wireshark is also handy for Linux/OS X users, and Fiddler is fun for Windows users.



## chapter 9: Authentication

Restful webservice authentication 主流的方式就两种:

* Oauth 1 or 2
* JSON Web Token (this is prefered, easy to implement also secure)


## chapter 10: Pagination

pagination example:
    
        {
            data:[],
            pagination:{
                cursors: {
                    after: 12,
                    next_url: "/places?cursor=12&number=12"
                }
            }
        }

challenges:

* Counting lots of Data is Hard: 分页计算总数的时候 会影响性能 如何有效地返回数据将是一个挑战
* 

pageination ways:

* embeded pagination field in json
* Pagination with the Link Header (作者不建议使用这种方式, 我也不喜欢)


## chapter 11: Documentation

介绍了如何构建 API 文档

[blueprint syntax](http://apiary.io/blueprint)

## chapter 12:  HATEOAS


>HATEOAS is a tricky subject to explain, but it is actually rather simple. It stands for Hypermedia as the Engine of Application State

sounds a little like a cereal for API developers.
However you want to try and say it, it basically means two things for your API:

* 1.Content negotiation
* 2.Hypermedia controls


## chapter 13:  API Versioning

这章主要讲了如何进行 API Versioning. 讨论了几种常见的方式和各自的优缺点, 最后总结道没有最好的方式, 只有最适合自己的方式。
还有不同的方式, 可能会对 http cache 产生影响

* versioning by url (不建议)

        https://api.example.com/v1/places

* Approach #2: Hostname (不建议)

        https://api-v1.example.com/places

* Approach #3: Body and Query Params

        POST /places?version=1.0 HTTP/1.1
        Host: api.example.com
        
        header1,header2
        value1,value2

* Approach #4: Custom Request Header (没看懂)

        GET /places HTTP/1.1
        Host: api.example.com
        BadApiVersion: 1.0
        
        
        HTTP/1.1 200 OK
        BadAPIVersion: 1.1
        Vary: BadAPIVersion

* Approach #5: Content Negotiation
    
        application/vnd.github[.version].param[+json]

* Approach #6: Content Negotiation for Resources
    
        Accept: application/vnd.github.user+json; version=4.0


* Approach #7: Feature Flagging (没看懂)




























