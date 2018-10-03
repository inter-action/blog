

# notes
* arduino IDE 有很多实例可以参考. 还有官方网站中的tutorial.



# [Arduino 入门教程](https://www.youtube.com/watch?v=d8_xXNcGYgo&list=PLGs0VKk2DiYx6CMdOQR_hmJ2NbB4mZQn-), 这个系列适合入门.


Serial.print/println 时序的作用, 用于input/output交互
Serial.available() === 0/parseInt ...etc

* Ohm's Law : `V = I*R`
* 串联电路中电流是相等的, 电压是各个电阻组件分布的 Sum(Voltage)= Vr1 (voltage in resister r1) + Vr2 + Vr3 ...etc
* potentiometer/电位器:
  作用就是通过旋钮来控制3个节点中2个的电压, 从而达到读数的目的 V(a, b) = Va/Va+Vb * V(总电压)

```
int pin = A0 // 这种写法有点怪
```

analogRead 必须使用 A0-A5 pin 去读
digitalWrite 和 analogWrite 的pin也不一样, 需要注意下

LESSON 15: Arduino Color Sensor and RGB LED
```
pulseIn(pin, LOW) // return 0 to 102400
```

18 节的课看不太懂, 用pulseIn测试距离, 没法理解, 总觉的哪里有问题, 先跳过, 后面再看

lesson 23, 24 are skipped



LESSON 28: Tutorial for Programming Software Interrupts on Arduino
这节就是说如何通过 interrupt 来之行多任务, 通过引入第三方库来实现的

LESSON 30: Advanced Software Interrupt Techniques for Reading Serial Data on Arduino
https://github.com/PaulStoffregen/TimerOne , 通过这个库可以暂停 interrupt timer.



# [Arduino and Python](https://www.youtube.com/watch?v=95w4sx_zoB8&list=PLGs0VKk2DiYylFUUMMv9WiL3x3tpscDUQ)
这个 playlist 主要是将如何用python和arduino进行通信, 通过serial port. 对于我来说没有太大价值.

Arduino and Python LESSON 2: Installing the Software and Libraries:

talk to serial port: need `port name & boaut rate` to read info from programming liberaries.

通信模块可以用xbees Radio. 这货通信的方式就是通过Serial 去通信的. 和Arduino 的 Serial 
console 是一样的


Arduino with Python LESSON 16: Simple Client Server Model Over Ethernet


# Arduino, 这个视频 Playlist 还高阶些, 推荐看这个系列的.

## [Arduino DC Motor Control Tutorial](https://www.youtube.com/watch?v=sOz41WQF7wE&list=PLrLOfmf1dB-gNMG2RBjhx1V51bKhtHI4E&index=13)

这节讲了如何控制一个dc motor, 需要注意的点是, 电机在运作的时候会产生一个back voltage, 会损伤 arduino, 所以需要通过二极管的电路连接去cancel back voltage. 注意二极管的短路.

<img src="./assets/dc motor.png" width="600">


## [Arduino H Bridge DC Motor Control Tutorial](https://www.youtube.com/watch?v=GumqesVRKyk&list=PLrLOfmf1dB-gNMG2RBjhx1V51bKhtHI4E&index=14)

通过H-Bridge来控制dc motor 的转速和方向


## 21 [Connecting a Sharp Distance Sensor to the Arduino Uno](https://www.youtube.com/playlist?list=PLrLOfmf1dB-gNMG2RBjhx1V51bKhtHI4E) 
Arduino 的 ananlog 最大的的值是2^10, 因为registor度数的最大的是10bits, 默认是用5v去量, 你可以用 `analogReference` 去改写这个最大值. arduino 内部读取数据的电阻是32k.

<img src="./assets/ref voltage.png" width="600">


## 22 [Arduino Shift Register Scanning LEDs effect](https://www.youtube.com/watch?v=s97TUARBc2U&list=PLrLOfmf1dB-gNMG2RBjhx1V51bKhtHI4E&index=22)

这节讲了如何利用 Shift Register 加上 `shiftOut` 命令去控制多个led 的显示与关闭. 虽然看了一遍还是没有看懂如何利用 Shift Register. 找了其他视频.

```processing
// RCLk, SER, SRCLK, shift rigister 引脚, SER,表示接受数据的引脚
digitalWrite(RCLk, LOW); // 通知 shift register 接受数据
shiftOut(SER, SRCLK, MSBFIRST, SEQ[i]); // 传递数据
digitalWrite(RCLk, HIGH); // 通知 shift regsiter 数据传递完毕
```


视频没看懂, 看了官方文档才看懂 :(, [
Serial to Parallel Shifting-Out with a 74HC595](https://www.arduino.cc/en/Tutorial/ShiftOut)


## 27 [Arduino SPI DigiPot Control Tutorial](https://www.youtube.com/watch?v=pp36Q5i08HE&index=27&list=PLrLOfmf1dB-gNMG2RBjhx1V51bKhtHI4E)

这个视频也没怎么看懂, 但大致的意思就是通过 arduino 来控制 Digital Potentiometer, 来动态增加或减少电阻值. 从而达到电压/电流控制.


[AD5171 Digital Potentiometer
](https://www.arduino.cc/en/Tutorial/DigitalPotentiometer), 看完这个链接之后, 大致明白了. I2C 的通讯原理和方式, I2C 需要两根线, 一根是时钟, 一根是data. 方式是找到 device address, 然后先发送 instruction 指令, 再发送对应指令的值.


```cpp
#include <Wire.h>

void setup() {
  Wire.begin(); // join i2c bus (address optional for master)
}

byte val = 0;

void loop() {
  Wire.beginTransmission(44); // transmit to device #44 (0x2c)
  // device address is specified in datasheet
  Wire.write(byte(0x00));            // sends instruction byte
  Wire.write(val);             // sends potentiometer value byte
  Wire.endTransmission();     // stop transmitting

  val++;        // increment value
  if (val == 64) { // if reached 64th position (max)
    val = 0;    // start over from lowest value
  }
  delay(500);
}
```


## 30 [Arduino Fading LED Interrupt Sketch tutorial](https://www.youtube.com/watch?v=Yotl88ieYuo&list=PLrLOfmf1dB-gNMG2RBjhx1V51bKhtHI4E&index=30)

这节讲解了如何通过硬件的方式来控制interrupt. 即每次按钮按下, 代码从正常的loop跳到对应的函数去执行.

## 31 [Arduino and Hardware Debouncing tutorial](https://www.youtube.com/watch?v=9UKM0vlHGkI&list=PLrLOfmf1dB-gNMG2RBjhx1V51bKhtHI4E&index=31)

这节讲的很有深度. 讲了电容的作用(用来解决按钮在按下的过程中, 按钮的物理特性造成的电压抖动, 用电容去解决). 还讲了 schmitt trigger 组件, 这个组件是用来解决增加电容之后, 电压不能sharply rise or fall 导致的问题. 

电容的加入电容之后的图片: 

<img src="./assets/Capacitance usage.png" width="600">

schmitt trigger:

<img src="./assets/schmitt trigger.png" width="600">

final result:

<img src="./assets/final circuit of hardware debouncing.png" width="600">


* 电容:
  * Q = CV, Q 代表 charge, C 是电容的值, V 电伏(Voltage). 
  * 我对 charge 还不是很理解. 后面需要看看
  * [Capacitance | Circuits | Physics | Khan Academy](https://www.youtube.com/watch?v=ngOC4eUQl8Y), 这里讲解了为啥 Q=CV 这个公式, 看了个大致懂, 最后面的算式推导没看懂 :(.
  * [Capacitors and Capacitance: Capacitor physics and circuit operation](https://www.youtube.com/watch?v=f_MZNsEqyQw), 这个视频对电容的解释就非常形象了. 从物理现象去解释的.



## 41 [Arduino I2C Tutorial](https://www.youtube.com/watch?v=vZr9LEIWOsA&list=PLrLOfmf1dB-gNMG2RBjhx1V51bKhtHI4E&index=41)

讲了I2C protocol 的使用方式, 方法. 需要注意的是SDA, SCL pin都需要接上拉电阻(pull up resistor).

<img src="./assets/i2c flowchart.png" width="600">



# Links

* [how to read a risistor](https://www.youtube.com/watch?v=GLD7AgAYqwA)
* [reference](https://www.arduino.cc/reference/en/language/functions/advanced-io/pulsein/)
* [arduino tutorial , offcial website, 分类很详细](https://www.arduino.cc/en/Tutorial/BuiltInExamples)
* [project hub, 官方的项目demo, 别人做的](https://create.arduino.cc/projecthub)


# todo:
* what is charge.
