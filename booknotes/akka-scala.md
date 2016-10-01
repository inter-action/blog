# Akka Core:
! http://doc.akka.io/docs/akka/2.4.10/scala.html

## General (重要):

! http://doc.akka.io/docs/akka/2.4.10/general/index.html

这节讲了核心的actor的概念和机制, 重启，失败，各个情况对actor和mail box的影响。



### Actor System:

what is actor:
>Actors are objects which encapsulate state and behavior, they communicate exclusively by exchanging messages which are placed into the recipient’s mailbox.
>An actor is a container for State, Behavior, a Mailbox, Child Actors and a Supervisor Strategy.

actor not automactically destroyed itself:
>One noteworthy aspect is that actors have an explicit lifecycle, they are not automatically destroyed when no longer referenced; after having created one, it is your responsibility to make sure that it will eventually be terminated as well

> An actor system will during its creation start at least three actors,
* /user: The Guardian Actor
* /system: The System Guardian 
* /: The Root Guardian

Actor Core components:

* Actor Reference:?
* State
* Behavior:?
* Mailbox:
    >The piece which connects sender and receiver is the actor’s mailbox: each actor has exactly one mailbox to which all senders enqueue their messages. Enqueuing happens in the time-order of send operations, which means that messages sent from different actors may not have a defined order at runtime due to the apparent randomness of distributing actors across threads. Sending multiple messages to the same target from the same actor, on the other hand, will enqueue them in the same order.

* Child Actors:
    >Each actor is potentially a supervisor: if it creates children for delegating sub-tasks, it will automatically supervise them. The list of children is maintained within the actor’s context and the actor has access to it.
* paths
    * logical
    * physical

    
Notes:

* --

    !Supervisor Strategy:
    >Fault handling is then done transparently by Akka, applying one of the strategies described in Supervision and Monitoring for each incoming failure. As this strategy is fundamental to how an actor system is structured, it cannot be changed once an actor has been created.
    When an Actor Terminates:

    >draining all remaining messages from its mailbox into the system’s “dead letter mailbox” which will forward them to the EventStream as DeadLetters.
    what is actor system:

    >The actor system as a collaborating ensemble of actors is the natural unit for managing shared facilities like scheduling services, configuration, logging, etc. Several actor systems with different configuration may co-exist within the same JVM without problems, there is no global shared state within Akka itself

2.4 Supervision and Monitoring:

* --

    Hierarchical Structure:

    >One actor, which is to oversee a certain function in the program might want to split up its task into smaller, more manageable pieces. For this purpose it starts child actors which it supervises.
    >The recursive structure then allows to handle failure at the right level.

    supervisor:
    一个actor可以创建子actor，父actor是子actor的supervisor

    当supervisor 停掉 any sub actor, any sub actors of this sub actor will also be stoped, same with started

    Delayed restarts with the BackoffSupervisor pattern:
        可以创建一个supervisor, 如果sub actor failed to start, it will be sceduled to restart after a increasing time interval

    2.4.5 One-For-One Strategy vs. All-For-One Strategy:?






important sections:
* 2.2.3 Actor Best Practices
* 2.2.4 Blocking Needs Careful Management
* 2.4.2 The Top-Level Supervisors
* 2.5.1 What is an Actor Reference?


## akka packaging:
http://doc.akka.io/docs/akka/2.4.10/intro/deployment-scenarios.html

##General Notes:

> [Do not use -optimize Scala compiler flag](http://doc.akka.io/docs/akka/2.4.10/intro/getting-started.html)

# Akka-Http:
http://doc.akka.io/docs/akka/2.4.10/scala/http/index.html
