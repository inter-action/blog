[nodejs stream 入门介绍](https://github.com/substack/stream-handbook)  
这个链接对nodejs stream做了一些简单的介绍, 包括nodejs的stream的类型有哪几种。这篇文章只简单介绍 transform stream. 这篇文章的代码是取自下面这两个链接。如果你有兴趣也可以直接读原文，跳过本文的介绍。
[My First Look At Streams In Node.js](http://www.bennadel.com/blog/2662-my-first-look-at-streams-in-node-js.htm)  
[Node.js Transform Streams vs. Through2 Streams](http://www.bennadel.com/blog/2663-node-js-transform-streams-vs-through2-streams.htm)  


## 首先 什么是 transform stream
trasform stream 是指从 input stream 输入后它输出另一个 stream, (当然如何输出这个stream就由我们来实现了)。本质有点像jdk实现的inputstream 各种 wrapper。

好 先说下这段代码是要做什么的， 这段代码的本质就是让流的一端输入字符串，然后根据自己定义的 transform stream 来根据正则表达式过滤出我们在意的字符串然后打印输出到 console 中。


### 原生实现  


```js
// Include module references.
var stream = require("stream");
var util = require("util");
var chalk = require("chalk");

// ---------------------------------------------------------- //
// ---------------------------------------------------------- //

// I am a Transform stream (writable/readable) that takes input and finds matches to the
// given regular expression. As each match is found, I push each match onto the output
// stream individually.
function RegExStream(pattern) {
    // If this wasnt' invoked with "new", return the newable instance.
    if (!(this instanceof RegExStream)) {
        return (new RegExStream(pattern));
    }

    // Call super-constructor to set up proper options. We want to set objectMode here
    // since each call to read() should result in a single-match, never a partial match
    // of the given regular expression pattern.
    //调用父类初始化函数
    stream.Transform.call( this, { objectMode: true } );

    // Make sure the pattern is an actual instance of the RegExp object and not just a
    // string. This way, we can treat it uniformly later on.
    // 构造正则表达式
    if (!(pattern instanceof RegExp)) {
        pattern = new RegExp(pattern, "g");
    }

    // Since the patter is passed-in by reference, we need to create a clone of it
    // locally. We're doing to be changing the RegExp properties and we need to make
    // sure we're not breaking encapsulation by letting the calling scope alter it.
    this._pattern = this._clonePattern(pattern);

    // I hold the unprocessed portion of the input stream.
    // 内部的缓存 处理断截的字符串
    this._inputBuffer = "";

}

// Extend the Transform class.
// --
// NOTE: This only extends the class methods - not the internal properties. As such we
// have to make sure to call the Transform constructor (above).
// 继承 node stream
util.inherits(RegExStream, stream.Transform);


// I clone the given regular expression instance, ensuring a unique refernce that is
// also set to include the "g" (global) flag.
RegExStream.prototype._clonePattern = function(pattern) {

    // Split the pattern into the pattern and the flags.
    var parts = pattern.toString().slice(1).split("/");
    var regex = parts[0];
    var flags = (parts[1] || "g");

    // Make sure the pattern uses the global flag so our exec() will run as expected.
    if (flags.indexOf("g") === -1) {
        flags += "g";
    }

    return (new RegExp(regex, flags));
};


// I finalize the internal state when the write stream has finished writing. This gives
// us one more opportunity to transform data and push values onto the output stream.
// 覆盖 node stream 的 flush 逻辑
RegExStream.prototype._flush = function(flushCompleted) {

    logInput("@flush - buffer:", this._inputBuffer);
    var match = null;

    // Loop over any remaining matches in the internal buffer.
    while ((match = this._pattern.exec(this._inputBuffer)) !== null) {
        logInput("Push( _flush ):", match[0]);
        this.push(match[0]);
    }

    // Clean up the internal buffer (for memory management).
    this._inputBuffer = "";

    // Signal the end of the output stream.
    this.push(null);

    // Signal that the input has been fully processed.
    flushCompleted();
};


// I transform the given input chunk into zero or more output chunks.
// 自定义逻辑的核心代码
RegExStream.prototype._transform = function(chunk, encoding, getNextChunk) {

    logInput(">>> Chunk:", chunk.toString("utf8"));

    // Add the chunk to the internal buffer. Since we might be matching values across
    // multiple chunks, we need to build up the buffer with each unused chunk.
    this._inputBuffer += chunk.toString("utf8");

    // Since we don't want to keep building a large internal buffer, we want to pair-
    // down the content that we no longer need. As such, we're going to keep track of the
    // the position of the last relevant index so that we can drop any portion of the
    // content that will not be needed in the next chunk-processing.
    var nextOffset = null;

    var match = null;

    // Loop over the matches on the buffered input.
    while ((match = this._pattern.exec(this._inputBuffer)) !== null) {

        // If the current match is within the bounds (exclusive) of the input buffer,
        // then we know we haven't matched a partial input. As such, we can safely push
        // the match into the output.
        if (this._pattern.lastIndex < this._inputBuffer.length) {

            logInput("Push:", match[0]);
            this.push(match[0]);
            // The next relevant offset will be after this match.
            nextOffset = this._pattern.lastIndex;

            // If the current match butts up against the end of the input buffer, we are in
            // danger of an invalid match - a match that will actually span across two (or
            // more) successive _write() actions. As such, we can't use it until the next
            // write (or finish) event.
        } else { // 如果匹配的文字在最后，这里就是可能有截断的情况，缓存到下次处理

            logInput("Need to defer '" + match[0] + "' since its at end of the chunk.");
            // The next relevant offset will be BEFORE this match (since we haven't
            // transformed it yet).
            nextOffset = match.index;
        }

    }

    // If we have successfully consumed a portion of the input, we need to reduce the
    // current input buffer to be only the unused portion.
    if (nextOffset !== null) {
        this._inputBuffer = this._inputBuffer.slice(nextOffset);
        // If no match was found at all, then we can reset the internal buffer entirely. We
        // know we won't need to be matching across chunks.
    } else {
        this._inputBuffer = "";
    }
    // Reset the regular expression so that it can pick up at the start of the internal
    // buffer when the next chunk is ready to be processed.
    this._pattern.lastIndex = 0;

    // Tell the source that we've fully processed this chunk.
    getNextChunk();

};


// ---------------------------------------------------------- //
// ---------------------------------------------------------- //


// Create our regex pattern matching stream.
var regexStream = new RegExStream(/\w+/i);

// Read matches from the stream.
regexStream.on( "readable", function() { 
        var content = null;

        // Since the RegExStream operates on "object mode", we know that we'll get a
        // single match with each .read() call.
        while (content = this.read()) {
            logOutput("Pattern match: " + content.toString("utf8"));
        }

    }
);

// Write input to the stream. I am writing the input in very small chunks so that we
// can't rely on the fact that the entire content will be available on the first (or
// any single) transform function.
"How funky is your chicken? How loose is your goose?".match(/.{1,3}/gi)
    .forEach( function(chunk) {
            regexStream.write(chunk, "utf8");
        }
    );

// Close the write-portion of the stream to make sure the last write() gets flushed.
regexStream.end();


// ---------------------------------------------------------- //
// ---------------------------------------------------------- //


// I log the given input values with a distinct color.
function logInput() {

    var chalkedArguments = Array.prototype.slice.call(arguments).map( function(value) {
            return (chalk.magenta(value));
        }
    );
    console.log.apply(console, chalkedArguments);
}


// I log the given output values with a distinct color.
function logOutput() {

    var chalkedArguments = Array.prototype.slice.call(arguments).map( function(value) {
            return (chalk.bgMagenta.white(value));
        }
    );
    console.log.apply(console, chalkedArguments);

}
```

### through2 实现  


```js
// Include module references.
var through2 = require("through2");
var chalk = require("chalk");


// ---------------------------------------------------------- //
// ---------------------------------------------------------- //


// I am a Transform stream (writable/readable) that takes input and finds matches to the
// given regular expression. As each match is found, I push each match onto the output
// stream individually.
function RegExStream(patternIn) {

    // Make sure the pattern is an actual instance of the RegExp object and not just a
    // string. This way, we can treat it uniformly later on.
    if (!(patternIn instanceof RegExp)) {

        patternIn = new RegExp(patternIn, "g");

    }

    // Since the patter is passed-in by reference, we need to create a clone of it
    // locally. We're doing to be changing the RegExp properties and we need to make
    // sure we're not breaking encapsulation by letting the calling scope alter it.
    var pattern = clonePattern(patternIn);

    // I hold the unprocessed portion of the input stream.
    var inputBuffer = "";

    // Return the Transform stream wrapper. We're using the "obj" convenience method
    // since we want out read stream to operate in object mode - this way, each read()
    // invocation will constitute a single pattern match.
    return (through2.obj(transform, flush));

    // ---
    // PRIVATE METHODS.
    // ---

    // I clone the given regular expression instance, ensuring a unique reference that
    // is also set to include the "g" (global) flag.
    function clonePattern(pattern) {
        // Split the pattern into the pattern and the flags.
        var parts = pattern.toString().slice(1).split("/");
        var regex = parts[0];
        var flags = (parts[1] || "g");

        // Make sure the pattern uses the global flag so our exec() will run as expected.
        if (flags.indexOf("g") === -1) {
            flags += "g";
        }

        return (new RegExp(regex, flags));
    }


    // Since we are no longer using Prototype methods, we're creating functions that
    // are bound to instances of the RegExStream. As such, we should probably clean
    // up variable references to help the garbage collector, especially since we're
    // also passing those methods "out of scope."
    // --
    // CAUTION: I'm not entirely sure this is necessary in NodeJS. In JavaScript on the
    // browser, the worst-case is that the user refreshes their page. But, when running
    // JavaScript on the server, we might need to be more vigilant about this stuff.
    function destroy() {
        patternIn = pattern = inputBuffer = clonePattern = destroy = flush = transform = null;
    }


    // I finalize the internal state when the write stream has finished writing. This gives
    // us one more opportunity to transform data and push values onto the output stream.
    function flush(flushCompleted) {

        logInput("@flush - buffer:", inputBuffer);
        var match = null;

        // Loop over any remaining matches in the internal buffer.
        while ((match = pattern.exec(inputBuffer)) !== null) {
            logInput("Push( _flush ):", match[0]);
            this.push(match[0]);
        }
        // Clean up the internal buffer (for memory management).
        inputBuffer = "";
        // Signal the end of the output stream.
        this.push(null);
        // Signal that the input has been fully processed.
        flushCompleted();
        // Tear down the variables to help garbage collection.
        destroy();
    }


    // I transform the given input chunk into zero or more output chunks.
    function transform(chunk, encoding, getNextChunk) {
        logInput(">>> Chunk:", chunk.toString("utf8"));

        // Add the chunk to the internal buffer. Since we might be matching values across
        // multiple chunks, we need to build up the buffer with each unused chunk.
        inputBuffer += chunk.toString("utf8");

        // Since we don't want to keep building a large internal buffer, we want to pair-
        // down the content that we no longer need. As such, we're going to keep track of
        // the the position of the last relevant index so that we can drop any portion of
        // the content that will not be needed in the next chunk-processing.
        var nextOffset = null;
        var match = null;

        // Loop over the matches on the buffered input.
        while ((match = pattern.exec(inputBuffer)) !== null) {
            // If the current match is within the bounds (exclusive) of the input
            // buffer, then we know we haven't matched a partial input. As such, we can
            // safely push the match into the output.
            if (pattern.lastIndex < inputBuffer.length) {
                logInput("Push:", match[0]);

                this.push(match[0]);

                // The next relevant offset will be after this match.
                nextOffset = pattern.lastIndex;

                // If the current match butts up against the end of the input buffer, we are
                // in danger of an invalid match - a match that will actually span across two
                // (or more) successive _write() actions. As such, we can't use it until the
                // next write (or finish) event.
            } else {

                logInput("Need to defer '" + match[0] + "' since its at end of the chunk.");

                // The next relevant offset will be BEFORE this match (since we haven't
                // transformed it yet).
                nextOffset = match.index;

            }
        }

        // If we have successfully consumed a portion of the input, we need to reduce
        // the current input buffer to be only the unused portion.
        if (nextOffset !== null) {
            inputBuffer = inputBuffer.slice(nextOffset);
            // If no match was found at all, then we can reset the internal buffer entirely.
            // We know we won't need to be matching across chunks.
        } else {
            inputBuffer = "";
        }
        // Reset the regular expression so that it can pick up at the start of the
        // internal buffer when the next chunk is ready to be processed.
        pattern.lastIndex = 0;
        // Tell the source that we've fully processed this chunk.
        getNextChunk();

    }

}

// ---------------------------------------------------------- //
// ---------------------------------------------------------- //

// Create our regex pattern matching stream.
var regexStream = new RegExStream(/\w+/i);

// Read matches from the stream.
regexStream.on( "readable", function() {
        var content = null;
        // Since the RegExStream operates on "object mode", we know that we'll get a
        // single match with each .read() call.
        while (content = this.read()) {
            logOutput("Pattern match: " + content.toString("utf8"));
        }
    }
);

// Write input to the stream. I am writing the input in very small chunks so that we
// can't rely on the fact that the entire content will be available on the first (or
// any single) transform function.
"How funky is your chicken? How loose is your goose?".match(/.{1,3}/gi)
    .forEach(
        function(chunk) {

            regexStream.write(chunk, "utf8");

        }
    );

// Close the write-portion of the stream to make sure the last write() gets flushed.
regexStream.end();

// ---------------------------------------------------------- //
// ---------------------------------------------------------- //

// I log the given input values with a distinct color.
function logInput() {
    var chalkedArguments = Array.prototype.slice.call(arguments).map(
        function(value) {
            return (chalk.cyan(value));
        }
    );
    console.log.apply(console, chalkedArguments);
}


// I log the given output values with a distinct color.
function logOutput() {
    var chalkedArguments = Array.prototype.slice.call(arguments).map(
        function(value) {
            return (chalk.bgCyan.white(value));
        }
    );
    console.log.apply(console, chalkedArguments);
}
```


原文中对 throught2 的本质做了精确的解释
> To be clear, Through2 streams aren't really a special kind of stream - they are Node.js Transform streams. The Through2 module just encapsulates the construction of the Transform stream so that you don't have to worry about prototypal inheritance and defining prototype methods. Instead, you just pass in your transform() and flush() functions and Through2 will construct the stream object and wire up your methods.


through2 仅仅是对 transform stream 做了封装, 所以你不必关心原生实现的细节，仅需要你做的事是 定义 transform 和 flush 函数