# Pattern matching

## Pin operator

### Definition

`^` is used when you want to match and not bind to a variable.

### Examples

```quartz
>>> let one = 1
>>> let ^one = 2
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
>>> let '[1, sec, 3 | rest] = '[1, 2, 3, 4, 5]
>>> sec
2
>>> rest
'[4, 5]
```

```quartz
>>> let '[1, 2 | rest] = '[1, 2]
>>> rest
'[]
```

```quartz
>>> let '[1, 2, 3 | rest] = '[1, 2]
** (MatchError) no match of right hand side value: '[1, 2]
```

```quartz
>>> let '[1, 2, third] = '[1, 2, 3]
>>> third
3
```

```quartz
>>> let '[1, 2, third] = '[1, 2]
** (MatchError) no match of right hand side value: '[1, 2]
```

## Tuple pattern matching

### Definition

|Pattern|Description|
|---|---|
|'(`elem`, ...)|Exactly matches the given elements.|

### Examples

```quartz
>>> let '(one, 2, 3) = '(1, 2, 3)
>>> one
1
```

```quartz
>>> let '(1, 2, 3) = '()
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
>>> let '{r: r, g: g, b: b} = '{r: 255, g: 255, b: 255}
>>> r
255
```

```quartz
>>> let '{color: _} = '{color: 'red, length: 12} 
```

## Binary pattern matching

`'<>`
