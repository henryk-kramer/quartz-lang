# Pattern matching

## Pin operator

### Definition

`^` is used when you want to match to match and not bind to a variable.

### Examples

```quartz
>>> let one = 1
>>> ^one = 2
** (MatchError) no match of right hand side value: 2
>>> one
1
```

## List pattern matching

### Definition

|Pattern|Description|
|---|---|
|'\[`elem`, ... \| `rest`\]|Matches the first x elements and binds a list only containing the unmatched elements to `rest`.|
|'\[`elem`, ...\]|Exactly matches the given elements.|

### Examples

```quartz
>>> '[1, let sec, 3 | let rest] = '[1, 2, 3, 4, 5]
>>> sec
2
>>> rest
'[4, 5]
```

```quartz
>>> '[1, 2 | let rest] = '[1, 2]
>>> rest
'[]
```

```quartz
>>> '[1, 2, 3 | let rest] = '[1, 2]
** (MatchError) no match of right hand side value: '[1, 2]
```

```quartz
>>> '[1, 2, let third] = '[1, 2, 3]
>>> third
3
```

```quartz
>>> '[1, 2, let third] = '[1, 2]
** (MatchError) no match of right hand side value: '[1, 2]
```

## Tuple pattern matching

### Definition

|Pattern|Description|
|---|---|
|'(`elem`, ...)|Exactly matches the given elements.|

### Examples

```quartz
>>> '(let one, 2, 3) = '(1, 2, 3)
>>> one
1
```

```quartz
>>> '(1, 2, 3) = '()
** (MatchError) no match of right hand side value: '()
```

## Struct pattern matching

### Definition

|Pattern|Description|
|---|---|
|'{`key`: `value`, ...}|Matches the given key/value pairs.|

The `key` must be known at compile time.

### Examples

```quartz
>>> '{r: let r, g: let g, b: let b} = '{r: 255, g: 255, b: 255}
>>> r
255
```

```quartz
>>> '{color: _} = '{color: 'red, length: 12} 
```

## Named struct pattern matching

|Pattern|Description|
|---|---|
|'`struct`{`key`: `value`, ...}|Matches the given values for the given keys and also matches all other keys on that struct.|

```quartz
>>> struct color = '{r: u8, g: u8, b: u8}
>>> 'color{r: let r} = '{r: 255, g: 255, b: 255}
>>> r
255
```

```quartz
>>> struct color = '{r: u8, g: u8, b: u8}
>>> 'color{r: 255} = '{r: 255, g: 255, b: 255} 
```

```quartz
>>> struct color = '{r: u8, g: u8, b: u8}
>>> 'color{} = '{r: 255}
** (MatchError) no match of right hand side value: '{r: 255}
```

## Binary pattern matching

`'<>`
