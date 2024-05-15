# Pattern matching

## General

`^` is used to match against the value of a variable and does not bind to the variable. Example: `^some_var := 1` checks
if `1` can bind to `some_var` (meaning that `some_var` has the value of `1`).

## List pattern matching

### Definition

|Pattern|Description|
|---|---|
|'\[`elem`, ... \| `rest`\]|Matches the first x elements and binds a list only containing the unmatched elements to `rest`. Elements can be skipped by using a muted identifier `_`|
|'\[`elem`, ...\]|Matches all the given elements. Elements can be skipped by using a muted identifier `_`|

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
>>> '[1, 2, third] = '[1, 2, 3]
>>> third
3
```

```quartz
>>> '[1, 2, third] = '[1, 2]
** (MatchError) no match of right hand side value: '[1, 2]
```

## Tuple pattern matching

`'()`

## Struct pattern matching

`'{}`

## Named pattern matching

`'struct{}`

## Binary pattern matching

`'<>`
