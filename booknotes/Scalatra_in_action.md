


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


two way:
* setting scalatra env key:
  org.scalatra.environment key either as a system property or using the web.xml file,

    a system property can be set through a command-line parameter: -Dorg.scalatra.environment=production.

    web.xml:
      <context-param>
        <param-name>org.scalatra.environment</param-name>
        <param-value>production</param-value>
      </context-param>

* use typesafe config lib, which is more robust.

  
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
    


#todos

http://www.foundweekends.org/conscript/how.html
scala TrieMap
