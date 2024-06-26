# Expressions

## Boolean expressions

|Expression|Description|
|---|---|
|!`op`|Logical NOT on boolean|
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

### Definition

|Expression|Description|
|---|---|
|'\[`elem`, ...\]|Returns a list with the given elements.|
|`list`\[`num`\]|Returns the element at index `num`.|
|'\[`elem`, ... \| `list`\]|Returns the `list` with the given elements prepended.|

### Examples

```quartz
>>> '[1, 2, 3]
'[1, 2, 3]
```

```quartz
>>> let l = '[1, 2, 3]
>>> l[2]
3
```

```quartz
>>> let l = '[1, 2, 3]
>>> l[3]
** (OutOfBoundsError) index 3 is out of bounds for list '[1, 2, 3]
```

```quartz
>>> let l = '[[11, 12], [21, 22], [31, 32]]
>>> l[1][1]
22
```

```quartz
>>> let l = '[4, 5]
>>> '[1, 2, 3 | l]
'[1, 2, 3, 4, 5]
```

## Tuple expressions

### Definition

|Expression|Description|
|---|---|
|'(`elem`, ...)|Returns a tuple with the given elements.|
|`list`.`num`|Returns the element at index `num`.|

### Examples

```quartz
>>> '(1, 2, 3)
'(1, 2, 3)
```

```quartz
>>> let t = '(1, 2, 3)
>>> t.2
3
```

```quartz
>>> let t = '(1, 2, 3)
>>> t.3
** Compilation error: (OutOfBoundsError) index 3 is out of bounds for tuple '(1, 2, 3)
```

```quartz
>>> let t = '('(1, 2), 3, 4)
>>> t.0.1
2
```

## Struct expressions

### Definition

|Expression|Description|
|---|---|
|'{`key`: `val`, ...}|Returns a struct with the given key/value pairs.|
|`struct`.`key`|Returns the value at `key`.|

### Examples

```quartz
>>> let s = '{color: 'red}
>>> s.color
'red
```

```quartz
>>> let s = '{color: 'red}
>>> s.size
** Compilation error: (FieldDoesNotExist) field 'size' does not exist in struct '{color: 'red}
```

## Named struct expressions

### Definition

|Expression|Description|
|---|---|
|'`name`{`key`: `val`, ...}|Returns a struct `name` with the given key/value pairs. Every key that does not have a default value needs to be specified.|
|`struct`.`key`|Returns the value at `key`.|

### Examples

```quartz
>>> let s = 'color{r: 255, g: 255, b: 255}
>>> s.r
255
```

```quartz
>>> struct color = '{r: u8, g: u8, b: u8}
>>> let s = 'color{r: 255}
** (ConstructionError) all fields need to be initialized
```

```quartz
>>> let s = 'color{r: 255, g: 255, b: 255}
>>> s.size
** Compilation error: (FieldDoesNotExist) field 'size' does not exist in struct 'color{r: 255, g: 255, b: 255}
```

## Binary expressions

`'<>`

## Any type expressions

|Expression|Description|
|---|---|
|`op` == `op`|Equality check|
|`op` != `op`|Inequality check|
|(`expr`)|Grouping/Raise precedence|
|`op` ?? `default`|If `op` is nil return the `default` value (short-circuit evaluation). Only works on nillable types.|
