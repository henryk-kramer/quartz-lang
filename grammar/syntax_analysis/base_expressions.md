# Base expressions

## Boolean expressions

|Expression|Description|
|---|---|
|! `op`|Logical NOT on boolean|
|`op` && `op`|Logical AND on boolean (short-circuit evaluation)|
|`op` \|\| `op`|Logical OR on boolean (short-circuit evaluation)|

## Integer expressions

|Expression|Description|
|---|---|
|not `op`|Logical NOT on each bit|
|`op` and `op`|Logical AND on each bit|
|`op` or `op`|Logical OR on each bit|
|`op` xor `op`|Logical exclusive OR on each bit|
|`op` shl `num`|Logical shift bits to the left|
|`op` shr `num`|Logical shift bits to the right|
|`op` ashr `num`|Arithmetic shift bits to the right|
|`op` cshl `num`|Circular shift bits to the left|
|`op` cshr `num`|Circular shift bits to the right|

## Number expressions

|Expression|Description|
|---|---|
|+`op`|No effect but only works on numbers|
|-`op`|Invert the sign of a number|
|`op` + `op`|Add two numbers|
|`op` - `op`|Subtract two numbers|
|`op` * `op`|Multiply two numbers|
|`op` / `op`|Divide two numbers|
|`op` < `op`|Less than comparison|
|`op` > `op`|Greater than comparison|
|`op` <= `op`|Less than or equal comparison|
|`op` >= `op`|Greater than or equal comparison|

## List expressions

|Expression|Description|
|---|---|
|\[`elem`, ...\]|Returns a list with the given elements|
|`list`\[`num`\]|Returns the element at index `x`|
|`list`\[`num`\]\[`num`\]|Returns the element of a nested list. Special case of `list`\[`x`\]|
|\[`elem`, ... ~ `list`\]|Returns the `list` with the given elements prepended|

## Map expressions

|Expression|Description|
|---|---|
|&\[`key` -> `val`, ...\]|Returns a map with the given key/value pairs|
|`map`\[`x`\]|Returns the value at key `x`|
|`map`\[`x`\]\[`y`\]|Returns the value of a nested map. Special case of `map`\[`x`\]|
|&\[`key` -> `val`, ... ~ `map`\]|Returns the `map` with the given key/value pairs added/updated|

## Struct (Tuple) expressions

|Expression|Description|
|---|---|
|&(`key`: `val`, ...)|Returns a struct with the given key/value pairs|
|&(`elem`, ...)|Returns a struct with the index as key and the elements as values|
|`struct`.`x`|Returns the value at key `x`|
|`struct`.`x`.`y`|Returns the value of a nested struct. Special case of `struct`.`x`|
|&(`key`: `val`, ... ~ `struct`)|Returns the `struct` with the given key/value pairs added/updated|

## Any type expressions

|Expression|Description|
|---|---|
|`op` == `op`|Equality check|
|`op` != `op`|Inequality check|
|(`expr`)|Grouping/Raise precedence|
|`op` ?? `default`|If `op` is nil return the `default` value (short-circuit evaluation). Only works on nillable types.|