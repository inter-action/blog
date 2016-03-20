


## chapter 1: Fine-Tuning Your Development Environment

$ android update sdk --no-ui # update sdk

#### The adb (Android Debug Bridge) Tool
------------------------------

$ adb logcat MyActivity:* *:S # filter log

    adb devices #List all connected Android devices and Emulators.
    adb push <local> <remote> #Copy a file from your computer to a device (usually on the SD card). 
    adb pull <remote> <local> #Copy a file from the device to your computer.

    adb shell # Executing Commands on an Android Device

        am startservice –a <intent action>  #start a Service using an Intent action

        # The Package Manager lets you interact with the installed applications (packages), 
        # allowing you to list, install, uninstall, and inspect the features and permissions on the device.
        pm 
            list packages


        # Stress-Testing an Application’s UI with Monkey

        monkey –p <package name> <event count> #executes the Monkey tool on an application with the specific <package name> and injects as many random events as specified by <event count>.





    emulator 
        -list-avds
        @<device_name>



connect device via wifi

    $ adb devices
    List of devices attached
    0070015947d30e4b
    $ adb tcpip 5555
    $ adb connect 192.168.1.104
    $ adb devices
    List of devices attached
    192.168.1.104:5555     device



$ANDROID_HOME/extras

>The Android SDK comes with a number of ready-to-use library projects that can be found under extras/ 
google in the SDK folder. More specifically, you can find library projects for the Play Services, 
APK extensions, In-app Billing and Licensing features. 


## chapter 2: Efficient Java Code for Android


#### Dalvik
-----------
JIT compilation, also known as dynamic translation, takes the byte-code and translates it into native code prior to execution 

Another difference between the Java SE VM and the Dalvik VM is that the latter is 
optimized for running in multiple instances on the same machine.


#### AsyncTask
-----------
>The only problem with AsyncTask is that you can use each instance of this class only 
once, which means that you have to call new MyAsyncTask() every time you want to perform this operation. 

>it is not suitable for operations that you perform frequently because you would 
quickly gather up objects that need to be garbage collected and eventually cause your 
application’s UI to stutter.

>In addition, you cannot schedule the time for the execution of this operation or perform the operation 
at a certain interval. The AsyncTask class is suitable for things like file downloads or similar
situations that will happen relatively infrequently or by user interaction.


#### Handler(more powerful than AsyncTask)
-----------
This class allows you to schedule operations with exact precision, and you can reuse it as
many times as you want. The thread that executes the operations loops until you explicitly terminate it; 

two types of handle creation

    HandlerThread handlerThread = new HandlerThread(“BackgroundThread”); //off main thread
    mMainHandler = new Handler(getMainLooper(), this);  //main thread



## Chapter 3: Components, Manifests, and Resources

The three core concepts of any Android application are the components, the manifest, and the resources.



components type:

* Activity           # with ui
* Services           # background task, doesnt constraint by Activity
* BroadcastReceivers # listen for system events
* ContentProviders   # store application data
* Application        #


Android Data storage:

* file
* ContentProvider
* sqlite
* SharedPreferences


#### Application:
------------------------------

>You can consider the Application component as a top-level component that’s created before Activities,
Services, and BroadcastReceivers.

>You can always retrieve a reference to the Application component through the method Context.getApplication().
Because all Android apps will have one and only one instance of this component, you can use it to share 
variables and communicate across the other components within your app. 


#### The Manifest Element
------------------------------


>The android:sharedUserId and android:sharedUserLabel are the Linux user ID and name that 
your application will run on. By default, these are assigned by the system, but if you set them 
to the same value for all the applications you publish using the certificate, you can access the
same data and even share the process of these applications. If you’re building a suite of 
applications or if you have a free and a paid version, this can be very helpful.


Google Play Filter: pass



The Application Element:
    be sure set label and description, android:backupAgent 


Components:

set android:enabled=”false”, then set it back when user finish certain configuration.


Intents:
    
* explicit:

    Explicit Intents contain information about the package and the name of the component, 

* implicit:

    Implicit Intent resolution depends on three factors: the action of the Intent, the data URI and type, and the category.             


    All components in Android are accessed using Intents.



Resource vs Assets

All resources belong to a certain type (such as layout, drawable, and so on), whereas assets are
simply generic files stored inside your application.

always put a default resource, or you app may crash

>However, because the assets directory supports subfolders, in some situations, you may 
still want to use assets instead of resources.


Further Resources:


## Chapter 4: Android User Experience and Interface Design


When developing Android applications you have a large set of ready-made UI elements, called widgets,
in the Android SDK that you can use.


Dimensions and Sizes:

One dp is roughly equal to one physical pixel on a 160 dpi screen.

>You should specify all dimensions in your UI using the dp unit, with one exception: When specifying 
font sizes, use the unit sp, which is based on the dp unit but also scaled according to 
the users’ preferences for font sizes.

48 dp == 48 px on MDPI
> The short answer is “the 48-dp rhythm.” Basically, 48 dp translates to about 9 mm on the screen, 
which is within the suitable range for objects on a touchscreen that users need to interact 
with using their fingers.


Text Size

Four standard font size: 
These are micro (12 sp), small (14 sp), medium (18 sp), and large (22 sp).


## Chapter 5: Android User Interface Operations



#### Designing Custom Views (自定组件)
------------------------------

```plain
extends View 

View.onAttachedToWindow() //  call when add to window
View.onMeasure() //
View.onLayout() // calculate layout
View.onDraw() // begin actual drawing
View.onDetachedFromWindow() 

```


####  Multi-Touching: pass
------------------------------


## Chapter 6: Services and Background Tasks



>Simplified, a Service is either started or stopped, which makes it much easier to handle than the more 
complicated lifecycle of an Activity. All you really need to remember is to create expensive objects
in onCreate() and do all cleanup in onDestroy()


onCreate & onDestroy 仍然在主线程执行, 可以用 AsyncTask 来处理这一部分





```plain
Service.onCreate
Service.onDestroy # both executed on main thread

Service.onStartCommand():Int # 返回值控制系统如果终止了service 再次启动此service的行为

  START_STICKY : 
    it signals that you want your Service to be restarted if the system shuts it down for some reason
    However, when the system restarts the Service, the onStartCommand() is called with the Intent parameters 
    set to null, so you have to take care of this in your code

  START_NOT_STICKY:
    your Service won’t restart after the system shuts it down.

  START_REDELIVER_INTENT:
    works like START_STICKY, except that the original Intent is redelivered when the system restarts your Service.


Staying Alive:
  Service.startForeground() 
    # prevent system from stop this service when this service when swtich to background
    # dont do this unless really necessary, (may wasting system resource)

  Service.stopForeground()



Stopping Services:

  * end by system due to lack resource
  * end by user
  * (in case of local bind service) Context.unbindService() is called
    >The exception is if you also call Service. startForeground() in your Service to keep it 
    alive after the last client disconnects, which is why it’s important to call Service.stopForeground() properly.

  * (in case of started by Context.startService()) 
    >then the only way to ensure that your Service is stopped is by calling either 
    Service.stopSelf() or Context.stopService()





ways of create/start service:
  extends Service: 


  Local Binding Services: ( start by Context.bindService() )
    这种方式需要注意的是 在activity中 onResume onPause 方法中绑定和解除绑定对应的service资源

    这种方式通过实现 IBinder 的方式, 最终 Activty 使用的时候需要继承 ServiceConnection, 然后通过实现此接口(ServiceConnection)的
    方法拿到 service 实例

    这种方式的好处就是可以拿到对应service实例, 并可以调用其实现的方法, 比如inject进度通知.


  IntentService:
    这种方式一般比继承service的方式要好 由于Google的封装, 你需要实现的接口也少, 也不需要关心 Service.onCreate Service.onDestroy
    在 main thread 上执行的问题， 而且过多的任务会有一个队列catch住

```


#### Parallel Execution
------------------------------
书中介绍了一种继承service实现的一次性多线程处理任务的service


#### Communicating with Services
------------------------------

* BroadcastReceiver
    >The drawback to this solution is that you’re limited to what an Intent can carry. 
    Also, you cannot use this solution for multiple, fast updates between the IntentService 
    and your Activity, such as progress updates, because doing so will choke the system. 


* Local bind service with callback


## Chapter 7: Android IPC


two ways of IPC:

* the Binder IPC (Inter-Process Communication).
* existing solution in the Linux kernel named dbu



#### Binder
------------------------------

>The Binder isn’t just used by Services, it also handles all communication between Android 
components and the Android system.

why service dont block main thread:
>Communication using the Binder follows the client-server model. Clients use a client-side
proxy to handle the communication with the kernel driver. On the server-side, the Binder 
framework maintains a number of Binder threads. The kernel driver delivers the messages 
from the client-side proxy to the receiving object using one of the Binder threads on the 
server-side. This is important to remember because when you receive calls to a 
Service through the Binder, they will not be executed on the main thread of your application.
That way, a client to a remote Service cannot block the Service application’s main thread.


Binder client vs Binder Server:
client call IBinder.transact(), server will receive data from Binder.onTransact()



Binder address:
>Clients that want to communicate with a Service or other component query the ServiceManager,
implicitly through the Intent resolution, to receive the Binder address.


Binder Transactions:

Parcel:
>They are used as simple data containers for the data you want to include in the transaction.

you can create a Parcel by implement Parcelable interface


Link to Death:

>Another feature of the Binder in Android is that it allows clients to be notified 
when a Service is terminated

```plain
Designing APIs:

做plugin使用

Provide a remote api service 
  Interface Definition Language (IDL) :
  Android Interface Definition Language (AIDL):
    注意这里定义的方法不会阻塞main thread 原因看上面-Binder

    >First, for all non-primitive parameters, you need to specify one of three directional types:
    in, out, or inout. The in type indicates that they are used only for input and that your 
    client won’t see any changes that the Service does to this object. The out type indicates 
    that the input object contains no relevant data but will be populated with data by the Service 
    that’s relevant in the response from the method. The inout type is a combination of both types.
    It’s very important to use only the type that’s needed because there’s a cost associated with each type.

    >Another thing to remember is that for all custom classes used in communication, 
    you need to create an AIDL file that declares your class as a Parcelable.


    characters of AIDL file
      once created you can not change exsiting method defiantion without break backward compatibility
      This method for versioning is one of the drawbacks with using AIDL files.




    tips:

      handle versioning (without breaking backward compatibility)


  Messenger:
    就是通过创建 local binder 的方式, 将 Messenger class wrap 到 binder 之上，然后通过消息的类别, payload, 回调三个参数进行通讯



  Wrapping APIs with Library Projects：
    将 AIDL 定义的接口通过子工程Libary的方式用java wrap你定义的接口, 然后将此lib打包发布.好处就是可以隐藏细节，
    并可以使用@deprecated annotation


  Securing Remote APIs:

    The important areas are shown in bold. First, you need to set the attribute android:exported to true.

    If you don’t include an intent-filter, the Service is only for internal use (addressed through its
    component name), and it won’t be exported. 

    If you’re exporting a Service, the most important part is to set up permissions.



```


## Chapter 8: Mastering BroadcastReceivers and Configuration Changes


#### BroadcastReceivers
------------------------------

create a BroadcastReceiver:

>BroadcastReceivers can also be registered programmatically within Activities and Services. 
Some broadcast Intents can only be registered programmatically, and some only 
work if you declare them in your manifest. 

register by:
  xml mainifest

  programmatically:
    dynamic, register only your acitivty is running(thus consume less resource), remmber to unregister it onPause


Local BroadcastReceivers:

>If you want to send and receive broadcast only within your own application’s process, 
consider using the LocalBroadcastManager instead of the more generic Context.sendBroadcast() method.
This approach is more efficient because no cross-process management is included and you don’t have to 
consider the security issues normally involved with broadcasts.


Normal and Ordered Broadcasts:

* Normal:没有顺序
  sticky broadcast: nomarl braodcast 的一个变种, 当 Intent（消息）被发送成功之后, 消息不会立即消失, 
  方便下个组件监听这个事件仍会得到此分类的通知
  directed broadcasts: 可以指定获得的组件名称

* Ordered: 有顺序, 任何一个环节可以阻止通知的下沉


System Events:

network, screen lock, configuration....




## Chapter 9: Data Storage and Serialization Techniques

```plain
two types of Persistence:
  Preference files are stored in an XML format and managed by the SharedPreferences class. 
    创建的方式有细微的差别 用的时候注意下, 还有注意最终的结果是以xml的形式保存的

    Every preference file has a single instance within the same process

  SQLite databases have a more complex API and are usually wrapped in a ContentProvider component.

    High-Performance ContentProviders:

data serialization:
  json:
  Google Protocol Buffers:

    advantages:
      The main advantage of protobuf is that it consumes much less memory and is faster to read and 
      write than JSON. Protobuf objects are also immutable, which is useful when you want to ensure
      that the values of the object remain the same during its lifetime.

    how it works:
      The schema defines a number of messages, where each message has a number of name-value pair fields. 
      Each field can be either one of the supported built-in primitive data types, an enum, or another message. 
      You can also specify whether a field is required or optional as well as some other parameters. 
      Once your protobuf schema is complete, you use the protobuf tools to generate the Java code for your data. 
      The generated Java classes can now be used for reading and writing protobuf data in a convenient way.

  Application Data Backup:

```


## Chapter 10: Writing Automated Tests

Using TDD involves a number of tools and techniques. First, you need a unit-testing framework for writing
the automated tests. This framework, which is included in the Android APIs, is the main focus of this chapter. 
Second, you need a continuous integration and build server. This server automatically builds your 
application and performs all the automated tests for every change in the code. Finally, you need a 
code-coverage tool that tells you how much of the application’s code is really being tested.



AndroidTestCase:
  without android component life cycle

Testing Activities:
  ActivityUnitTestCase
  ActivityInstrumentationTestCase2



## Chapter 12: Secure Android Applications


Signatures and Keys:

uring normal development, you will use the auto-generated debug key for signing your application 
(which is done automatically by the Gradle build system or the IDE). When you publish your 
application on the Google Play Store, be sure to use a unique key that you generate manually 
using the keytool application.

the only reasons for using the same key for multiple applications is when the applications 
need to access each other’s data directly, or you have defined permissions with the protection level signature.

When you generate a new key, the keytool asks you for a password. If you don’t provide a password, 
anyone with access to the keystore file can create a properly signed application. Thus it’s highly
 recommended that you use a unique password for each keystore file.

Android Permissions:
normal: 用户会接受到通知
dangeous: 用户会在安装时提示
signature: 需要同一个证书签名
signatureOrSystem and system: 手机制造商的权限，一般程序无法达到





## Chapter 14: Native Code and JNI

## Chapter 15: The Hidden Android APIs

## Chapter 17: Networking, Web Service, and Remote APIs


Never perform network calls on the main thread.
Always do networking on Android from a Service and avoid performing network calls from within an Activity
  (exception is Auth2)

I recommend using either a callback interface or the LocalBroadcastManager to communicate network results
between the Activity and the Service. Another way is to store the results from a network call directly 
into a ContentProvider and let the provider notify registered clients about the changes to any data, 


Volley (vs OkHttp?)

OkHttp and SPDY
If you choose to use SPDY as your communication protocol, I recommend that you choose the 
third-party library called OkHttp, which is developed by Square, Inc.

When you create a new instance of the OkHttpClient, it will set up everything you need, 
such as connection polling and response cache

WebSocket:

  If you need to transmit a large file between the client and server, stick to standard HTTP instead. 

Facebook SDK for Android:


Finding Online Web Services and APIs:

* http://www.programmableweb.com/
* http://apis.io/
* https://www.mashape.com/


Often, it may be cheaper to pay for an existing service than try to implement one from scratch, unless that’s your core business.


Server Side Notifcation:

* Power Efficient Network Polling:
* Server-Side Push (through: like SMS, or it can be a regular TCP socket with a long keep-alive.)
* Web Sockets for Server-Side Push
   and you need to adjust the timeout on the socket on both the server and the client for this method to be efficient.





Designing Custom Views


The Manifest Element
------------------------------



extras:

    SystemClock.sleep(50);


    不管是 service, acitivty 都需要在结束的时候处理好资源的释放



todos:
    how to write a android module
        You can set up a library project from the Android Studio IDE. 
        Simply create a new module and choose Android Library as the module type. 


    setting up a continuous integration system 


    http://developer.android.com/design/index.html

    ConcurrentLinkedQueue

    how to build layout with android

    Review Context and Intent

    call ui related func off main thread, what consequence may it be

    Docker setting up a jenkins for android








Developer Options on Android Devices






