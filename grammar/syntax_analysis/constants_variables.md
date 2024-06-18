# Constants and variables

## Constants

Constants are variables that can only be bound once. A constant definition consists of the `const` keyword, a name, a type (optional) and an expression or block.

```quartz
const const_name_1 = 45
const const_name_2: u32 = 45
```

## Variables

Variables are bound the same way as constant but with the `let` keyword. 

```quartz
let var_name_1 = 45
let var_name_2: u32 = 45
```

The main difference between constants and variables is that variables can be rebound. Either with or without checking the type.

```quartz
let a: u32 = 45         // binding 45 to 'a'
let a: str = "first"    // rebinding "first" to 'a' without type checking
a := "second"           // rebinding "second" to 'a' with type checking
a := 46                 // trying to rebind 46 to 'a' throws compilation error due to type mismatch
```
