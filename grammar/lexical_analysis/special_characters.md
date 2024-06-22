# Comments

## Space

### Regex

```regexp
(\x20)+
```

### Diagram

```mermaid
graph LR;
    A((S0))
    B(((S1)))
    style B fill: green

    A & B -- SPACE --> B
```

## Tab

### Regex

```regexp
(\x09)+
```

### Diagram

```mermaid
graph LR;
    A((S0))
    B(((S1)))
    style B fill: green

    A & B -- TAB --> B
```

## Newline

### Regex

```regexp
(\r\n|\n\r|\r|\n)
```

### Diagram

```mermaid
graph LR;
    A((S0))
    B(((S1)))
    C(((S2)))
    D(((S3)))
    E(((S4)))
    style B fill: green
    style C fill: green
    style D fill: green
    style E fill: green

    A -- CR --> B
    A -- LF --> C
    B -- LF --> D
    C -- CR --> E
```
