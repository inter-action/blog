# 大纲
* quick demo
* service worker
  用途，特点，生命周期&工作机制

* 工具: sw-cache & Lighthouse

* PWA


# 主要讲cache

# quick demo



# Service Worker & PWA

## 什么是service worker:
service worker 类似一个 daemon 线程，在浏览器UI线程后面执行着任务。它有着自己的生命周期。
浏览器tab关闭之后，它仍会在后面运行着。



## 用途
离线文件是PWA的核心功能之一。

* 执行background tasks, 接口轮询，发送通知
* 缓存文件
* 拦截网络请求


## 特点:
* 不能更新UI
* 在完全独立的上下文中执行


## 生命周期 & 工作机制

前提条件:
只能在在localhost和配置了https的网站下安装

生命周期:
https://developers.google.com/web/fundamentals/getting-started/primers/service-workers
在service worker中，所有缓存的文件成功后，service worker 才会被激活。

作用域:
在根路径下的，能拦截所有的网络请求，在对应的自路径下，只能拦截自路径下的网络请求。
/example/

## demo:

    //简单的cache文件demo

    var CACHE_NAME = 'my-site-cache-v1';
    var urlsToCache = [
      '/',
      '/styles/main.css',
      '/script/main.js'
    ];

    self.addEventListener('install', function(event) {
      // Perform install steps
      event.waitUntil(
        caches.open(CACHE_NAME)
          .then(function(cache) {
            console.log('Opened cache');
            return cache.addAll(urlsToCache);
          })
      );
    });


## 其他:

* http cache 和 service worker的cache有什么区别：

相比于 http cache, service worker + cache api 
可以让你决定该以什么样的策略cache文件。

https://developers.google.com/web/fundamentals/instant-and-offline/offline-cookbook/#on-background-sync



## 调试 

chrome://inspect/#service-workers


# Tools
## sw-cache 
Google 开发的用于简便生成缓存网站资源的service worker文件。


    npm install --save-dev sw-precache
    sw-precache --config=path/to/sw-precache-config.js --verbose


### 生成的代码文件解释:




### cache hit 的机制

https://open.alpha.elenet.me/tip?status=registerSuccess&email=l1643162@mvrht.com
/dist/tip.html

simply wont match.



### 现在我想更新文件，怎么办？


重新生成一份service-worker.js。
这部分一定要注意注册service-worker.js的文件，一定不要被service worker缓存掉.
否则会造成死锁，导致你没有有效的方式更新文件或者service worker，用户手动清除浏览器数据，或者更新浏览器。

最好的方式就是把注册代码放置到 index.html 中，然后禁用其 http cache。


### 好的实践
* 不缓存html文件, 除非是angularjs那样的模板
* 小心使用 navigateFallback 选项



## Lighthouse




#PWA - The cutting edge of web.

离线资源，数据同步, 消息推送，接近原生app的效果。

## app shell model:
* [App Shell Model](https://developers.google.com/web/fundamentals/architecture/app-shell)


## app install banners
https://developers.google.com/web/fundamentals/engage-and-retain/app-install-banners/


### 


# Links

必看:
* [Service Worker](https://developers.google.com/web/fundamentals/getting-started/primers/service-workers)
* [App Shell Model](https://developers.google.com/web/fundamentals/architecture/app-shell)
* [sw-precache](https://github.com/GoogleChrome/sw-precache)

选看:
* [service worker explained](https://github.com/w3c/ServiceWorker/blob/master/explainer.md)
* [! offline cookbook](https://developers.google.com/web/fundamentals/instant-and-offline/offline-cookbook/#on-background-sync)
* [https://developers.google.com/web/fundamentals/instant-and-offline/offline-cookbook/](https://developers.google.com/web/fundamentals/instant-and-offline/offline-cookbook/)
* [https://github.com/GoogleChrome/sw-toolbox](https://github.com/GoogleChrome/sw-toolbox)
* [https://developers.google.com/web/showcase/2016/iowa2016](https://developers.google.com/web/showcase/2016/iowa2016)
* [下一代 Web 应用模型 —— Progressive Web App](https://mp.weixin.qq.com/s?__biz=MzAwNTAzMjcxNg==&mid=2651424849&idx=1&sn=4339fc1e71169159b81d9bf925bf68ca&chksm=80dff632b7a87f2463d5100bec33ffb6a1c62398366458b92ee643c435aa2fa714682031cd01&mpshare=1&scene=1&srcid=0209O24qPzjAkApLzs3pzHRN&key=f813d0a51cab06334f1dadd83dd128e7fd13da3658299f279d93792315c75b1795f374d626d64de53ff7a45e7cbd2082df069e7d730588ec3513d7c032e5d770df6868428b0e7ddf0fb071076821fa50&ascene=0&uin=MjIxMzg2NDgwMA%3D%3D&devicetype=iMac+MacBookPro12%2C1+OSX+OSX+10.12.2+build(16C67)&version=12010210&nettype=WIFI&fontScale=100&pass_ticket=H1rvbxc812xcT71iPbCSO%2FY7fDq0w%2Flzi8nayRUonEqPJJXDeuOjRE2FES0yShEf)
* [安全问题](https://sakurity.com/blog/2016/12/10/serviceworker_botnet.html)

    ========
    -title: service worker & PWA
    -tags: service worker, PWA
    -last_edit: 2017-02-07 14:19:22
    -created: 2017-02-07 14:19:22
    ========