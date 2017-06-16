
## scaffold scala project with bash
create a bash file:

    #!/bin/sh
    mkdir -p src/{main,test}/{java,resources,scala}
    mkdir lib project target

    touch project/plugins.sbt
    echo '
    ' > project/plugins.sbt

    # create an initial build.sbt file
    echo 'name := "Scala.js Tutorial"

    version := "1.0"

    scalaVersion := "2.11.5" // or any other Scala version >= 2.10.2' > build.sbt

    touch .gitingore

    echo '.idea/
    target/
    project/**/target
    ' > .gitingore


make bash executable:

[how-do-i-run-a-shell-script-without-using-sh-or-bash-commands](http://stackoverflow.com/questions/8779951/how-do-i-run-a-shell-script-without-using-sh-or-bash-commands)

    chmod +x <bash_file_name>

## sbt 使用国内镜像

touch ~/.sbt/repositories

    [repositories]
    local
    oschina:http://maven.oschina.net/content/groups/public/
    oschina-ivy:http://maven.oschina.net/content/groups/public/, [organization]/[module]/(scala_[scalaVersion]/)(sbt_[sbtVersion]/)[revision]/[type]s/[artifact](-[classifier]).[ext]
    typesafe: http://repo.typesafe.com/typesafe/ivy-releases/, [organization]/[module]/(scala_[scalaVersion]/)(sbt_[sbtVersion]/)[revision]/[type]s/[artifact](-[classifier]).[ext], bootOnly
    sonatype-oss-releases
    maven-central
    sonatype-oss-snapshots

[加速 SBT 下载依赖库的速度](https://segmentfault.com/a/1190000002474507)
[Scala sbt 添加国内镜像](http://blog.csdn.net/mmical/article/details/41925823)
