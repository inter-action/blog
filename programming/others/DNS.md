
# DNS 入门/笔记/教程


过程:
  首先要到对应网站注册域名

* Name Server: 用于解析DNS的服务器。有三种途径
  * 由你域名注册商提供，你可以在他们提供的面板里进行配置
  * 你自己的网站提供。
  * 外包给第三方服务商



* DNS Zone:
  DNS Zone 是多种DNS Record的集合。
  * A DNS Zone is like a container of all the DNS records for a specific domain and only that domain. For example: pressable.com, www.pressable.com, blog.pressable.com, and mail.pressable.com are four DNS Records inside a single DNS Zone for pressable.com.


* DNS Record: https://pressable.com/blog/2014/12/11/understanding-dns-record-types/
  * 什么是 DNS record: 一条DNS配置
    
  
  * 一条Record需要有三个信息
    * Record Name
      * 比如 blog.pressable.com, blog 就是一个Record Name
      * Record Name 的类型:
        * Blank Name: 就是domain前面啥也没有的那种
          >To refer to a previous example, pressable.com and www.pressable.com are two different DNS records with separate values for their name. The www.pressable.com DNS record uses “www” as its record name and pressable.com uses nothing/blank for its record name.
        * @ Symbol：重定向到 Blank Name
        * `* Symbol (Wildcard)`: 如果其他都匹配不到就走这个规则

    * Record Value/Data

    * TTL(Time to Live): 缓存过期时间，秒
      > TTL is the numerical value, in seconds, of how long a DNS record will be cached before it needs to be refreshed. Whenever a nameserver is queried for a DNS record, it will check to see if it has delivered that same DNS record within the time period specified by the TTL and if so, will deliver the cached version of that DNS record. Once that period of time specified by the TTL passes, the nameserver will query the zone for record data and cache it once more for the specified period of time.



  * A DNS Record type: https://pressable.com/blog/2014/12/23/dns-record-types-explained/
  
    * SOA: https://support.dnsimple.com/articles/soa-record/
      * 记录父级的DNS server
      
      > An SOA record is a Start of Authority. Every domain must have a Start of Authority record at the cutover point where the domain is delegated from its parent domain. For example if the domain mycompany.com is delegated to DNSimple name servers, we must include an SOA record for the name mycompany.com in our authoritative DNS records. We add this record automatically for every domain that is added to DNSimple and we show this record to you as a System Record in your domain’s Manage page.



    * A: IPV4的
      > A Records are the most basic type of DNS record and are used to point a domain or subdomain to an IP address.

    * AAAA: IPV6的 

    * MX: 
      * MX 用来route邮件的，比 A Type 和 CNAME 多了一个priority value，值越小优先级越高
      > Mail Exchanger (MX) records are used to help route email according the domain owners preference. The MX record itself specifies which server(s) to attempt to use to deliver mail to when this type of request is made to the domain. They differ from A Records and CNAMEs in the way that they also require a “priority” value as a part of their entry. The priority number is used to indicate which of the servers listed as MX records it should attempt to use first.
      * In MX records, the value/data information indicates what mail servers email should be routed to.

    * CNAME:
      * 别名，比如下面的例子就是将主机 blog.myapp.com 重定向到另一个 host 地址。这个value，必须是一个域名，不能是ip

        host           |         value          | TTL
        blog.myapp.com | another-blog.myapp.com | 3600

    * Text：
      * value 是一段字符串，用于存储一些元信息。应该是随便存。

    * SPF 
      * SPF records use the value/data field so specify what servers are allowed to legitimately use your domain name for the sending of emails.

