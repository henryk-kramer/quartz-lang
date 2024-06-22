# Generics

Generics are a way to write functions for multiple types at the same time without repeating code.

## Function generics

The easiest way to use generics is in functions. In diamond brackets you can provide multiple generic type parameters which can be used as placeholders for real types in the parameter list and the function body.

```quartz
fn func<T>(t: T) -> T = ...
```

```quartz
fn print<T>(element: T) -> nil = {
    io::print(element)
}

print<str>("Print me")  // explicit type specification
print("Print me")       // automatic type deduction
```

## Class/struct generics

It is also possible to have generic classes and structs. This works the same except that you specify the generic parameters after the class name.

```quartz
class ClassName<T> {
    let t: T = ...

    fn func(self) -> T = ...
}
```

The functions of a class can have extra generic parameters.

```quartz
class ClassName<T> {
    fn func<F>(self, f: F) -> T = ...
}

let cls = 'ClassName<u32>{}
ClassName::func<str>(cls, "Generic")
```

## Trait generics

Traits behave the same as classes and structs when it comes to generics.

```quartz
trait TraitName<T> {
    fn func(self, t: T) -> T
}
```

When using generic parameters in the implementation, you need to specify them directly after the `impl` keyword.

```quartz
impl<T> ClassName<T> : TraitName<T> {...}
```

When you want to implement a trait for any type you can use the generic paramter for that.

```quartz
impl<C, T> C : TraitName<T> {...}
```

Below is a more complex example how a list implementation could look like.

```quartz
trait List<T> {
    fn add(self, elem: T) -> self
    fn get(self, idx: u32) -> T
}

class ArrayList<T> {
    let list: '[T] = '[]

    pub fn new<T>() -> ArrayList<T> = 'ArrayList<T>{}
}

impl<T, F> ArrayList<T> : List<T> { // The generic parameters after the impl are used when instantiating the new ArrayList class. See the 'new' function.
    pub fn add(self, elem: T) -> self = ...
    pub fn get(self, idx: u32) -> T = ...
}

let list = ArrayList::new<u32, u32>()
```

## Types and generics

When working with generics in types you can either define them directly or pass the generic paramater.

```quartz
type Arr1<T> = Array<T>
type Arr2 = Array<Integer>
```

## Bounds

You can restrict what types can be passed into a generic parameter by using bounds.

```quartz
fn func<T: Number>(a: T, b: T) -> T = ...
```

Note that it is only possible to specify a single type as a bound. This makes it possible to use functions on the generic paramter that can normally be used on the bound.