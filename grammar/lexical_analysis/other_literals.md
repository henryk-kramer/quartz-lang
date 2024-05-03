# Other literals

## Boolean literals

### Regex

```regexp
(true|false)
```

### Examples

```quartz
true
false
```

## Symbol literal

### Regex

```regexp
'([a-zA-Z0-9_]|[^\x00-\x7F])+
```

### Diagram

```mermaid
graph LR;
    A((S0))
    B((S1))
    C(((S2)))
    style B fill: red
    style C fill: green

    A -- QUOTE --> B
    B -- UNDERLINE --> C
    B -- a..z,A..Z,0..9 --> C
    B -- \x80.. --> C
```

### Examples

```quartz
'Symbol
'日本語
'€
```
