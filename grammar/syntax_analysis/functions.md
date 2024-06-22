# Functions

In Quartz functions are the main way to work with data. A basic function consists of the `fn` keyword, the function name, a parameter list, a return type and the function body.

```quartz
fn function_name(...) -> ... = ...
```

Each paramater in the parameter list consists of a name and a type. A function almost always has at least one parameter although it is not necessary. The parameter can then be referenced in the function body.

```quartz
fn no_param_func() -> ... = ...
fn one_param_func(a: u32) -> ... = ...
fn two_param_func(a: u32, b: u32) -> ... = ...
```

Functions can be called via the function name and some arguments.

```quartz
func()              // calls a function with no parameters
func(255)           // calls a function with one parameter
func(255, 255, ...) // calls a function with multiple parameters
```

In the same scope the combination of the function name and the parameter list types cannot be defined twice.

```quartz
fn func(a: u32) -> ... = ...
fn func(b: u32) -> ... = ...    // compilation error
fn func(c: str) -> ... = ...    // works fine
```

If you define two functions with the same name that take different number types (e.g. `u32` and `u8`) all compile-time numbers will be cast into their smallest representation and therefore the function with the smaller type will be taken. Since it is hard to find the problem in such a case, it is advised to avoid overriding number types if possible.

```quartz
fn func(n: u32) -> ... = ...
fn func(n: u8) -> ... = ...

func(255) // calls func(n: u8)
```

A function can have any type as return type. The last expression in the body must then return a value of that type or else a compilation error is thrown. There is one special return type called `nil`. A function with that return type returns always `nil` and not the last expression. The function body can either be an expression or a block with an expression at the end.

```quartz
fn func_1() -> nil = 5  // returns nil
fn func_2() -> u32 = 5  // returns 5 as u32
fn func_3() -> nil = {  // returns nil
    ...
    5
}
fn func_4() -> u32 = {  // returns 5 as u32
    ...
    5
}
```
