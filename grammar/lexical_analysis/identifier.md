# Identifier

## Normal identifier

### Regex

```regexp
([a-zA-Z0-9]|[^\x00-\x7F])([a-zA-Z0-9_]|[^\x00-\x7F])*
```

### Diagram

```mermaid
graph LR;
    A((S0))
    B(((S1)))
    style B fill: green

    A -- a..z,A..Z,0..9<br>\x80.. --> B
    B -- UNDERLINE<br>a..z,A..Z,0..9<br>\x80.. --> B
```

### Examples

```quartz
Variable
some_var
日本語
```

## Muted identifier

### Regex

```regexp
_([a-zA-Z0-9_]|[^\x00-\x7F])*
```

### Diagram

```mermaid
graph LR;
    A((S0))
    B(((S1)))
    style B fill: green

    A -- UNDERLINE --> B
    B -- UNDERLINE<br>a..z,A..Z,0..9<br>\x80.. --> B
```

### Examples

```quartz
_
_some_var
_日本語
___
```
