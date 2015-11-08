说起这个gulp的串行和并行的知识点，我也是被坑才了解的。这个源自公司项目一直暴露的bug引起的。这个bug是这样的，每次打测试包到服务器上的时候都需要弄两次，页面才不会出样式上的错误。详细调查了下得知是由于 gulp-rev 插件生成的 MD5 校验后的文件的 元信息(metadata) 丢失造成的。由于 php 端会根据 metadata 的文件生成的 合并文件去加载各个样式和 js 文件。如果有某一个 metadata 文件丢失， 这意味着最终构建的 metadata 合并文件的信息不完整。但是为什么会丢失，好几天一直找不到头绪，一直在乱撞。终于在今天找到了问题的根源。

这就是由于 gulp 的 task 的执行的并行引起的。

gulp 的所有的 task 支持并行执行，只要满足[如下条件](https://github.com/gulpjs/gulp/blob/master/docs/API.md#async-task-support):  

    1. 处理回调
    2. 返回流
    3. 返回 Promise

而要串行执行的话, 则需要满足如下条件：

>In task "one" you add a hint to tell it when the task is done. Either take in a callback and call it when you're done or return a promise or stream that the engine should wait to resolve or end respectively.

>In task "two" you add a hint telling the engine that it depends on completion of the first task.  


###
   
```js
var gulp = require('gulp');

// takes in a callback so the engine knows when it'll be done
// 接受回调
gulp.task('one', function(cb) {
    // do stuff -- async or otherwise
    cb(err); // if err is not null and not undefined, the run will stop, and note that it failed
});

// identifies a dependent task must be complete before this one begins
// 添加 task 的依赖关系
gulp.task('two', ['one'], function() {
    // task 'one' is done now
});

gulp.task('default', ['one', 'two']);
```


而实际的使用，我们一般会用到 run-sequence 来控制 gulp 的任务串行执行。

通过 


```js
// 任务从左往右顺序执行, 数组中的任务表示 数组里边的任务可以并行执行
runSequence( 'clean', [parallel_1, parallel_2], 'inserthash')
```


然而用这个插件来串行执行任务的时候，[文档](https://www.npmjs.com/package/run-sequence)要求有几点需要满足:

函数必须返回一个stream 或者 处理 callback


```js
// configure build-clean, build-scripts, build-styles, build-html as you 
// wish, but make sure they either return a stream or handle the callback 

// Example: 
 
gulp.task('build-clean', function() {
    return gulp.src(BUILD_DIRECTORY).pipe(clean());
//  ^^^^^^ 
//   This is the key here, to make sure tasks run asynchronously! 
});
```


这时候再来看公司项目原来项目写的方式:

```js
gulp.task('inserthash', function(done) {
    gulp.src(jsonTemp + '*.json')
        .pipe(genDspFile())
        .pipe(gulp.dest('./www/ad-management/config/')).on('end', done);
});
```

修改后


```js
gulp.task('inserthash', function() {
    return gulp.src(jsonTemp + '*.json')
        .pipe(genDspFile())
        .pipe(gulp.dest('./www/ad-management/config/'));
});
```

虽然以前的方式处理回调， 但是这种 on('end') 这种方式貌似是有问题的。最终将所有任务都修改成了返回 stream 的形式。问题就解决了。rev.mainifest 生成的文件就没再丢失了。可见写 gulp task 的时候, 总是返回流是个很好的习惯。


```js
.pipe(rev.manifest({
                path: filename + '.json'
            }))

```