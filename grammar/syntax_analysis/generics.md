# Generics

```quartz
fn func<T>(t: T) -> T = ...
```

```quartz
class Cls<T> {
    let t: T = ...

    fn func_1<T>() -> T = ...
    fn func_2<T, F>(t: T) -> F = ... // F needs to be specified or is deduced automatically when calling the function. T is passed from the class context.
}
```

```quartz
type arr_1<T> = Array<T>
type arr_2 = Array<Integer>
```
