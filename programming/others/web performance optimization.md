
# Web Performance Optimization （网站应用性能优化）
性能优化一直是一个比较高级的主题，因为涉及到的知识方方面面，从网络(dns寻址，http协议, DNS, CDN, compression)，浏览器机制(dom, cssom, render tree, repaint, reflow, resource loading, service worker, caching), 项目工程化(code splitting, image optimization, tree shaking, css & js minification), 最后还需要对自己的项目有着深刻的理解，知道哪些点可以优化。总之，这是展示前端工程最顶级的技术之一，也是能体现一个高级前端开发人员技能的点。

核心的链接:
* [Google developer doc on Web Performance](https://developers.google.com/web/fundamentals/performance/rail)


* don't over tweak performance, your dev time matters too!
* Know your tools! If you can't measure it, you can't improve it.

## road map

* Performance
  * load perf (all about loading time)
    * network
      * compression, gzip etc
      * http, http2 protocal
      * dns, cdn
    * browser cache mechanism
      * http protocal cache
        * cache-control header
        * private & public header
          * private: only browser can cache this resource
          * public: any intermidate route including cdn can cache this resource
        * ETag
        * `"no-cache" and "no-store"`

      * Cache API
        * `cache.add`, 
    * Critical Path: 
      * `<link rel="preload">`, `<link rel="preconnect">`, and `<link rel="prefetch">`
        * preload: resource need to be loaded quickly as possible
          * 注意这种方式不能预加载ajax请求的资源, 不然chrome中会请求2次同样请求
        * preconnect: hint to browser a request to another domain may occur, let it resolve dns as quickly as possible
        * prefectch: load this resource when others is done

        * Below-the-fold asynchronously loaded styles
          <link rel="preload" href="css/styles.min.css" as="style" onload="this.rel='stylesheet'">
      * script tag attribute `async, defer`
        * async
        * defer: put scripts right before body closing tag
      * js script would require cssom to be constructed, because browser don't know in advance whether js would manipulate cssdom, which in return blocks onDomLoaded event
      * code splitting with webpack
        * webpack-bundle-analyzer
          * 
        * import() syntax, webpack would know this syntax, and create an async module automatically
          * https://github.com/thejameskyle/react-loadable
            * support component level & route level code splitting

        * react-router, code splitting
        * webpack common trunk plugin
          * https://medium.com/@adamrackis/vendor-and-code-splitting-in-webpack-2-6376358f1923


    * offline with `service worker & cache api`
      * https://developers.google.com/web/fundamentals/instant-and-offline/offline-cookbook/
    
    * js optimization:
      * js minification:
         ```js   
         new UglifyJSPlugin({
          parallel: true,
          compress: {
            warnings: false,
            screw_ie8: true
          },
        }),
        ```
        
    * image optimization:
      * compress image:
        * https://github.com/Klathmon/imagemin-webpack-plugin, this is recommanded

          ```js
          // Make sure that the plugin is after any plugins that add images
          new ImageminPlugin({
            jpegtran: {
              progressive: true,
            },
            pngquant: {
              quality: '70-85'
            }
          })
          ```

        * https://github.com/lovell/sharp
        * https://tinypng.com/

      * choose the right image format, vector (svg) or rasterized (png, jpeg)
      * responsive load with `srcset`, `<picture>`
      * lazy load image:
        * LQIP technique:
          * https://github.com/aFarkas/lazysizes
          * https://github.com/loktar00/react-lazy-load
    * PWA
      * case study:
        * https://medium.com/@addyosmani/a-tinder-progressive-web-app-performance-case-study-78919d98ece0
        
      * PRPL Pattern:
        * Push critical resources for the initial URL route.
        * Render initial route.
        * Pre-cache remaining routes.
        * Lazy-load and create remaining routes on demand.
      * APP Shell Pattern
      * service worker:
        * [workbox google doc](https://developers.google.com/web/tools/workbox/get-started/webpack)
          * [runtime caching](https://developers.google.com/web/tools/workbox/reference-docs/latest/module-workbox-runtime-caching.StaleWhileRevalidate)

  * rendering performance:
    * css:
      * avoid complex selectors, use class name instead. 
        * `.box:nth-last-child(-n+1) .title`
      * split css for different page
      * inline css for initial load (by using http2, this approach should'nt be considered)
      * 

    * batch read first then do writes.

      ```js
      // wrong 
      function resizeAllParagraphsToMatchBlockWidth() {
        // Puts the browser into a read-write-read-write cycle.
        for (var i = 0; i < paragraphs.length; i++) {
          paragraphs[i].style.width = box.offsetWidth + 'px';
        }
      }

      // right
      var width = box.offsetWidth;

      function resizeAllParagraphsToMatchBlockWidth() {
        for (var i = 0; i < paragraphs.length; i++) {
          // Now write.
          paragraphs[i].style.width = width + 'px';
        }
      }
      ```

    * `will-change: transform;`
