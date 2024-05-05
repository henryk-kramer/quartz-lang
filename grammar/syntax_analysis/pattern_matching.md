# Pattern matching

## List pattern matching

|Pattern|Description|
|---|---|
|\[`elem`, ... ~ `rest`\]|Matches the first x elements and binds a list only containing the unmatched elements to `rest`. Elements can be skipped by using a muted identifier `_`|
|\[`elem`, ...\]|Matches all the given elements. Elements can be skipped by using a muted identifier `_`|

## Map pattern matching

|Pattern|Description|
|---|---|
|&\[`key` -> `val`, ... ~ `rest`\]|Matches the given key/value pairs and binds a map only containing the unmatched key/value pairs to `rest`. It is possible to only match `key` by specifying `val` as muted identifier `_` or by leaving out '-> `val`' completely|
|&\[`key` -> `val`, ...\]|Matches all the given key/value pairs. It is possible to only match `key` by specifying `val` as muted identifier `_` or by leaving out '-> `val`' completely|

## Struct (Tuple) pattern matching

|Pattern|Description|
|---|---|
|&(`key`: `val`, ...)|Matches the given key/value pairs|
|&(`elem`, ...)|The key is the index and the value is the element. It then matches the key/value pairs: `&(r, g, b)` is equivalent to `&(0: r, 1: g, 2: b)`|
