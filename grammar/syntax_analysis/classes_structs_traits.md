# Classes and structs

Classes and structs contain both data and functionality. The difference to object oriented programming languages like Java is that the function/method is not called directly on the object but rather is the object passed into the function.

```Java
class Example {
    public void doSomething() { ... }
}

var ex = new Example();

ex.doSomething(); // function
```

```quartz
class Example {
    pub fn do_something(self) -> nil { ... }
}

let ex = 'Example{}

Example::do_something(ex)
```

## Classes

Every class member is private by default which means it cannot be accessed outside of the class context.

A class starts with the keyword `class` followed by the name and curly braces in between which the members are defined.

```quartz
class ClassName { ... }
```

Class members can types, constants, variables and functions. Functions must specify the class itself as its first parameter by writing `self`. In the function body self can then be used to access other class members of the object passed into `self`.

```quartz
class ClassName {
    type TypeName = ...
    const const_name: ... = ...
    let var_name: ... = ...

    fn func_name(self, ...) -> ... = ...
}
```

If a member should be accessible from outside the class, it can be prepended by the `pub` keyword.

```quartz
class ClassName {
    pub let var_name: ... = ...
}
```

Accessing constants and variables can be done by using a dot on the object itself. The return the value they hold.

```quartz
let obj = 'Object{}
obj.const_name
obj.var_name
```

Accessing types and functions on the other hand can be done by using `::` on the class name.

```quartz
namespace::ClassName::TypeName

let obj = 'Object{}
namespace::ClassName::function_name(obj, ...)
```

## Structs

Structs are essentially the same as classes with the exception that everything is public by default and can be made private by the `priv` keyword.

```quartz
struct StructName {
    let var_name: ... = ...

    fn func_name(self, ...) -> ... = ...
    priv fn func_name_2(self, ...) -> ... = ...
}
```

## Traits

A trait is a component which describes function signatures that must be implemented by classes and structs that use that trait. Every function described by a trait is public.

A trait can be defined by the `trait` keyword followed by the trait name and curly braces which contain the function head definitions. As classes trait functions must specify `self` as its first parameter as well.

```quartz
trait TraitName { 
    fn trait_func(self, ...) -> ...
}
```

To implement a trait for a struct or class the `impl` keyword can be used. Implementing multiple traits can be achieved by using `impl` multiple times.

```quartz
impl ClassName : TraitName1 {
    fn trait_func(self, ...) -> ... = ...
}

impl ClassName : TraitName2 {
    fn trait_func(self, ...) -> ... = ...
}
```

Calling a function on an object works the same as for classes and structs. The only difference is that an instance of a class or struct that implements that trait is passed to the trait function.

```quartz
trait TraitName {
    fn get_value(self) -> u8
}

class ClassName {
    let value: u8
}

impl ClassName : TraitName {
    fn get_value(self) -> u8 = self.value
}

let obj = 'ClassName{value: 255}
TraitName::get_value(obj) // returns 255
```

Traits are especially useful if we want to group objects that are instances of different classes. It is then possible to call the same functions on all of them.

```quartz
trait Animal {
    fn make_sound(self) -> str
}

class Cat {
    pub fn new() -> Cat = 'Cat{}
}

impl Cat : Animal {
    fn make_sound(self) -> str = "Miau"
}

class Dog {
    pub fn new() -> Dog = 'Dog{}
}

impl Dog : Animal {
    fn make_sound(self) -> str = "Wuff"
}

let animals: '[Animal] = '[Cat::new(), Dog::new()]
Animal::make_sound(animals[0]) // returns "Miau"
Animal::make_sound(animals[1]) // returns "Wuff"
```
