    
    load balancing:
      nginx + docker
    clustering
      docker + nginx
    health check point

    yarn replace npm:
    profile end point, dump data for v8 to analyze
    monitoring
    processUncaughtException:
    explicit define your errors 
    dealing with problems
      clustering + load balancing cause socket-io to fail:
        solution - sticky session

    how to check outdated project dependencies
      any better way to do it?

    SEO:
      schema.org - https://schema.org/docs/gs.html

    db:
      mysql driver- https://github.com/mysqljs/mysql


    maintenance:
      module.exports = function maintenance(req, res, next) {
          if (config.get('maintenance').enabled) {
              return next(new errors.MaintenanceError({message: i18n.t('errors.general.maintenance')}));
          }

          next();
      };


    `trust proxy`?, express
        blogApp.enable('trust proxy');
        https://expressjs.com/en/guide/behind-proxies.html
        启用之后，注意这里的默认值的意思


    https://blog.risingstack.com/node-js-interview-questions-and-answers-2017/

## security:

rate-limit:
* conf by nginx: 
* config by nodejs:
  * pros: easy to control
  * cons: by default, it only control request each node process, if you work around this by plugin redis, and store info there.

csrf:


## profile tool
* memory-usage
* heapdump
* heap-profile
* low level: llnode, mdb
* <Debugging node.js in prod at netflix>
* v8-profiler:
  * https://github.com/node-inspector/v8-profiler

##server config:
* google config sample - https://github.com/h5bp/server-configs
* cache what, api & other sensitive request shouldn't be cached, `blogApp.use(routes.apiBaseUri, cacheControl('private'));`

nginx:
http2:
* [](https://certsimple.com/blog/nginx-http2-load-balancing-config)
    

## metrics
  https://github.com/RuntimeTools/appmetrics


# Ref:
* [mysql data types](https://www.tutorialspoint.com/mysql/mysql-data-types.htm)
* [sql ref](https://www.tutorialspoint.com/sql/)
* [nodejs at scale series](https://blog.risingstack.com/node-js-at-scale-understanding-node-js-event-loop/)
* [https: let encrypt](https://my.oschina.net/u/2328699/blog/829503)

security:
* [](https://blog.risingstack.com/node-hero-node-js-security-tutorial/)


performance:
* [Keeping Node.js Fast: Tools, Techniques, And Tips For Making High-Performance Node.js Servers](https://medium.com/@smashingmag/keeping-node-js-fast-tools-techniques-and-tips-for-making-high-performance-node-js-servers-8cfcb55e3d7)
  讲了 AutoCannon 和 Clinic 两种工具识别线上问题的方式方法
  --trace-warnings, tag 的使用
  还有就是系统过载 nodejs 处理方式 503 service not available
  
  * Active Handle in Clinic
    >Active Handles are unaffected by the Event Loop delay. An active handle is an object that represents 
    either I/O (such as a socket or file handle) or a timer (such as a setInterval). We instructed AutoCannon 
    to open 100 connections (-c100). Active handles stay a consistent count of 103. The other three are 
    handles for STDOUT, STDERR, and the handle for the server itself.
  

# Links:

[Real World Lessons on the Pain Points of Node.js Applications](https://www.youtube.com/watch?v=eZF1t7Cv_7o&list=LLuhDiGxQ78K0p35lrkwZLjA&index=2)
[nodejs habits](https://blog.heroku.com/node-habits-2017)

