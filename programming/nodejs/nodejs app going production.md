    
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


# Links:

[Real World Lessons on the Pain Points of Node.js Applications](https://www.youtube.com/watch?v=eZF1t7Cv_7o&list=LLuhDiGxQ78K0p35lrkwZLjA&index=2)
[nodejs habits](https://blog.heroku.com/node-habits-2017)

