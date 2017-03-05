


## decorator:
// https://www.typescriptlang.org/docs/handbook/decorators.html

type Decorator: (target, ...args)=>void

* decorator run at runtime
* decorator is one kind of way of doing language proxy.

//todo:
* decorator evaluation:
  * Parameter Decorators, followed by Method, Accessor, or Property Decorators are applied for each instance member.
  * Parameter Decorators, followed by Method, Accessor, or Property Decorators are applied for each static member.
  * Parameter Decorators are applied for the constructor.
  * Class Decorators are applied for the class.


* decorator type:
  * class decorator: take class constructor(class function) as its only argument.
  ```typescript
  function sealed(constructor: Function) {
    Object.seal(constructor);
    Object.seal(constructor.prototype);
  }
  ```

  * Method Decorators: The decorator is applied to the Property Descriptor for the method, and can be used to observe, modify, or replace a method definition. 
    * arguments: ?

  * Accessor Decorators: 
    * arguments:
      * 1st:
        * if static member, constructor class 
        * else prototype of the class
      * 2st: name of the number
      * 3st: The Property Descriptor for the member.

  * Property Decorators:
    * arguments:
      * 1st:
        * if static member, constructor class 
        * else prototype of the class
      * 2st: name of the number
    * @return: If the property decorator returns a value, it will be used as the Property Descriptor for the member.

  * Parameter Decorators:
    * arguments:
      * 1st:
        * if static member, constructor class 
        * else prototype of the class
      * 2st: name of the number
      * 3st: The ordinal index of the parameter in the function’s parameter list.

    * @return: The return value of the parameter decorator is ignored.

      
* use case