


## notes
在读这本书之前, 我就已经看过了 google web performance 网站上列的资源. 这本书里边的内容也有些重叠.
所以本书中有些内容没有必要看的很细, 速读就好了.


## chapter 2: Using assessment tools

summary: 
这章讲了chrome调试工具的一些介绍.


notes:
* benchmarking tool: console

    ```
    console.timeStamp("Modal open."); # create a marker in chrome dev tool
    # benchmarking using time function    
    console.time("jQuery"); jQuery("#schedule"); console.timeEnd("jQuery");
    ```

## chapter 3: Optimizing CSS

summary:
这章讲了css一些best practices, 大部分都已经知道, css 优化工具还是蛮有意思的


## chapter 4: Understanding critical CSS

notes:
* Below-the-fold asynchronously loaded styles

    ```html
    <link rel="preload" href="css/styles.min.css" as="style" onload="this.rel='stylesheet'">
    ```


## chapter 5: Making images responsive
summary:
这章主要讲了图片适配的问题, 还有图片格式的区别, loss compression & lossless compression 

notes:
* 图片适配一直是一个比较麻烦的问题, 
    * css 图片适配基本上是通过 media query 覆盖 image src 实现的, 但是我隐约还记得好像 background 属性直接
        就支持 media query
    * 还有就是 html img 标签的图片适配, 这种方式比较麻烦
        * 首先就是 `max-width: 100` 这个关键属性, 默认是使用图片的自然尺寸, 最大不超过容器宽度

        * img 标签 srcset 属性, 这个属性可以指定不同图片的宽度, 让浏览器自己去选择加载, 这个属性有个特点就是最好
            保证图片的 aspect ratio 相同, 否则会出现 undefined behaviour, 浏览器会根据 viewport 宽度自己加载
            如果加载大的, viewport 缩小, 图片也不会重新下载, 如果是从小的 viewport 拉伸到大的, 浏览器则会去加载
            大的图片, 保证质量

        * picture 标签, 这个标签可以说是 html img 适配终极利器了, 这个东西可以 根据 dpi, width, media query去定制
            要下载那张图片, 还有根据 size 属性, 指定显示的图片的大小, 以100 view port的比例来算 

            ```html
            <picture>
                <!-- 384w, w表示图片的像素大小 -->
                <source media="(min-width: 704px)"
                    srcset="img/amp-medium.jpg 384w, img/amp-large.jpg 512w"
                    sizes="33.3vw">
                <!-- 1x dpi, 2x dpi 加载不同的图片 -->
                <source srcset="img/amp-cropped-small.jpg 1x, img/amp-cropped-medium.jpg 2x" sizes="75vw">
                <!-- fallback -->
                <img src="img/amp-small.jpg"> amp-cropped-medium.jpg is 
            </picture> 
            ```


## chapter 6: Going further with images

summary:
这章主要讲了图片的优化方法,包括 sprites, compress, webp. 最后还有图片 lazy loading 的实现方式(
无非就是计算图片是否在 viewport + buffer 区域中, 然后改变图片loading的attribute的值).


## chapter 7: Faster fonts
skipped


## chapter 8: Keeping JavaScript lean and fast

skipped


## chapter 9: Boosting performance with service workers

skipped, 内容我都通过其他方式了解到了, 没有啥新颖的了.


## chapter 10: Fine-tuning asset delivery

notes:
* integrity attibute:
    <script src="https://code.jquery.com/jquery-2.2.3.min.js" integrity="sha256-a23g1Nt4dtEYOj7bR+vTu7+T8VP13humZFBJNIYoEJo=" crossorgin="anonymous">
    </script>


## chapter 11: Looking to the future with HTTP/2

summary:
这章内容主要是涉及http2带来的改变, 既然我对HTTP2已经有所了解, 我只看了我感兴趣的内容.就是http2向下兼容的方法.
无非就是通过检测http协议版本, 然后选择加载 concated asset or not.

## chapter 12: automating optimization with gulp

skipped

