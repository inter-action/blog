# installation

brew instsall wget

then install conscript
then g8

g8 scalatra/scalatra-sbt 


## Chapter 6 - Handling files

this chapter shows u how to handle file uploading and downloading with scalatra
also how to set file upload limits & handling related errors.


Serving static resources
这种方式来 serve 静态文件的比较陈旧, 一般都用nginx做reverse proxy, 用 nginx 来serve静态文件



## Chapter 8 - Testing


difficulty with testing Servlet API.

Integration testing with Specs2
    
  demos how to write simple spec2 test with scalatra

Unit testing with Specs2
  


## Chapter 9 - Configuration, build, and deployment

keys:
    
  split config base on different enviorment

staging env
  In the staging environment, the application behaves similar to the production 
  environment, but it runs in a non-public area.


two way of spliting config base:
* setting scalatra env key:
  org.scalatra.environment key either as a system property or using the web.xml file,

    a system property can be set through a command-line parameter: -Dorg.scalatra.environment=production.

    web.xml:
      <context-param>
        <param-name>org.scalatra.environment</param-name>
        <param-value>production</param-value>
      </context-param>

* use typesafe config lib, which is more robust.

  
ways of deploy app:
* pack as .war file
* pack as standalone .zip file
* pack as docker image.





the scalatra-sbt plugin:

The scalatra-sbt plugin adds a servlet container that can be used during development. 
It also supports building deployable web archive (WAR) packages of an application


xsbt-web-plugin:

xsbt-web-plugin is an extension to sbt that integrates a servlet-based web application into an sbt build. 
The scalatra-sbt plugin depends on it.

9.2.3 - Using sbt-web to simplify working with web assets
9.2.4 - Precompiling Scalate templates with the scalate-generator plugin

these sections are skiped due to ... i'm pro front-end engineer!

9.4 - Deploying as a standalone distribution

standalone distribution:
  
  instead of package a .war file and deploy it into servlet container, standalone distribution package
  along with servlet container into a single file.

  benifits of doing this enable you:

    * more consistent deployment process
    
9.5 - Running Scalatra as a Docker container

The sbt-docker plugin (https://github.com/marcuslonnberg/sbt-docker) integrates
the building of a Docker image into an sbt build.



## Chapter 10 - Working with a database

DBIO[T]
Each interaction with the database is encapsulated in a DBIO[T] action, 
where T is the return type of that action. 

Database
a database instance

TableQuery[E]
represent a table query on a specific table, use this to create query update or delete
table






## Chapter 11 - Authentication
this chapter demos 3 types authentication strategies you can apply to you app:
* basic auth
* username/password
* remember me 



Scentry Framework:

  "org.scalatra" %% "scalatra-auth" % ScalatraVersion,



ScentryStrategy:
  //extend this strategy to provide basic auth support
  authenticate//handle authentication
  unauthenticated//handle scenario when user is not authenticated



## Chapter 12 - Asynchronous programming
use AsyncResult whenever you can, try not to Use Future with scalatra on Action Route


ExecutionContext.global & ForkJoinPool
  If in doubt, use ExecutionContext.global. It uses a ForkJoinPool, which helps to minimize 
  context switches and starts up the thread pool with a size equal to the number of processors

Tomcat Servlet Threads limit:
  Servlet containers maintain a thread pool for dealing with incoming requests. 
  By default, Apache Tomcat has a pool of 200 threads. 


Servlet Thread Pool vs Scala Thread pool:
  For instance, the servlet container makes the variable request available to you inside 
  your Scalatra actions. The request is a reference to something that, by definition, 
  lives inside the servlet container’s thread pool. This raises a conundrum: 
  the request is in the servlet container, but everything inside the Future executes in
  a totally different thread pool. What happens if you attempt to access the request
  from inside the Future? The answer is simple and potentially unexpected: it will be null,
   because ExecutionContext.global doesn’t know anything about it.


Akka vs Scala Future:

Like Futures, Akka Actors run in their own thread pool, which is detached from your 
Scalatra application’s serv- let thread pool. Unlike Futures, they can run on either
a single machine or across a cluster of machines. The Akka library does all the thread
management and schedul- ing and takes care of inter-Actor communication. 
On the other hand, setting up a dis- tributed Akka ActorSystem is a lot more 
complex than just firing off a Future—each approach has its place.


Ways of doing Async Jobs
* Scala Future
* Akka
* Spark

## Chapter 13 - Creating a RESTful JSON API with Swagger

feel the demo:
  you can feel it @
    http://petstore.swagger.io/#/pet

swagger-codegen:

  https://github.com/swagger-api/swagger-codegen

  swagger-codegen contains a template-driven engine to generate client code in
  different languages by parsing your Swagger Resource Declaration



swagger with scalatra
  
  You might be thinking that all of these benefits sound wonderful, 
  but if you need to build a JSON file by hand for each of your APIs, 
  and keep it up to date manually whenever you make any changes, 
  you’re not very far ahead. Luckily, this isn’t the case


HMAC

amazon authenication:
  http://docs.aws.amazon.com/AmazonS3/latest/dev/RESTAuthentication.html

HMAC add timestamp to prevent replay attack.








# scalatra traits

FutureSupport - handle automatically processing Future object. @ page 162
  also notice you need to return a  AsyncResult from Route/Service Endpoint from scalatra @page162

SessionSupport - 
FlashMapSupport - @page 175, will turn on SessionSupport.
FutureSupport - @page 200, scala Future Support 

MethodOverride -
DatabaseSessionSupport -

ScalatraServlet - 
ScalateSupport - 生成html 用的？
JacksonJsonSupport - 支持JSON





#todos

http://www.foundweekends.org/conscript/how.html
scala TrieMap
relation database with docker
chapter 10 , page 160, <>
h2 database 
Await.result
ORM framework Squeryl vs Slick
goback to 11.3, revisit the whole thing.

different thread pool
  CachedThreadPool, FixedThreadPool, and WorkStealingThreadPool

scala future on akka doc:
  http://doc.akka.io/docs/akka/2.4.8/scala/futures.html
