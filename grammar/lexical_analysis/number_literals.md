# Number literals

## Boolean notation

### Regex

```regexp
0b[0-1](([0-1]|_)*[0-1])?
```

### Diagram

```mermaid
graph LR;
    A((S<sub>0</sub>))
    B((S<sub>1</sub>))
    C((S<sub>2</sub>))
    D(((S<sub>3</sub>)))
    E((S<sub>4</sub>))
    style C fill: red
    style D fill: green
    style E fill: red

    A -- 0 --> B
    B -- b --> C
    C -- 0..1 --> D
    D -- 0..1 --> D
    D -- UNDERLINE --> E
    E -- UNDERLINE --> E
    E -- 0..1 --> D
```

### Examples

```quartz
0b0
0b01
0b0001_0000
0b1010_0101__1000_1111
```

## Octal notation

### Regex

```regexp
0o[0-7](([0-7]|_)*[0-7])?
```

### Diagram

```mermaid
graph LR;
    A((S<sub>0</sub>))
    B((S<sub>1</sub>))
    C((S<sub>2</sub>))
    D(((S<sub>3</sub>)))
    E((S<sub>4</sub>))
    style C fill: red
    style D fill: green
    style E fill: red

    A -- 0 --> B
    B -- o --> C
    C -- 0..7 --> D
    D -- 0..7 --> D
    D -- UNDERLINE --> E
    E -- UNDERLINE --> E
    E -- 0..7 --> D
```

### Examples

```quartz
0o0
0o07
0o777_777
0o777_000__777_000
```

## Decimal notation

### Regex

```regexp
0b[0-9](([0-9]|_)*[0-9])?
```

### Diagram

```mermaid
graph LR;
    A((S<sub>0</sub>))
    B((S<sub>1</sub>))
    C((S<sub>2</sub>))
    D(((S<sub>3</sub>)))
    E((S<sub>4</sub>))
    style C fill: red
    style D fill: green
    style E fill: red

    A -- 0 --> B
    B -- d --> C
    C -- 0..9 --> D
    D -- 0..9 --> D
    D -- UNDERLINE --> E
    E -- UNDERLINE --> E
    E -- 0..9 --> D
```

### Examples

```quartz
0d0
0d99
0d1_234_567_890
0d99999__99999
```

## Hexadecimal notation

### Regex

```regexp
0b[0-9a-f](([0-9a-f]|_)*[0-9a-f])?
```

### Diagram

```mermaid
graph LR;
    A((S<sub>0</sub>))
    B((S<sub>1</sub>))
    C((S<sub>2</sub>))
    D(((S<sub>3</sub>)))
    E((S<sub>4</sub>))
    style C fill: red
    style D fill: green
    style E fill: red

    A -- 0 --> B
    B -- x --> C
    C -- 0..9,a-f --> D
    D -- 0..9,a-f --> D
    D -- UNDERLINE --> E
    E -- UNDERLINE --> E
    E -- 0..9,a-f --> D
```

### Examples

```quartz
0x0
0xff
0x70_ff
0xff_ff__ff_ff
```

## Scientific notation

### Regex

```regexp
[0-9]+(\.[0-9]+)?(e(\+|-)?[0-9]+)?
```

### Diagram

```mermaid
graph LR;
    A((S<sub>0</sub>))
    B(((S<sub>1</sub>)))
    C((S<sub>2</sub>))
    D((S<sub>3</sub>))
    E(((S<sub>4</sub>)))
    F((S<sub>5</sub>))
    G((S<sub>6</sub>))
    H((S<sub>7</sub>))
    I(((S<sub>8</sub>)))
    J((S<sub>9</sub>))
    style B fill: green
    style C fill: red
    style D fill: red
    style E fill: green
    style F fill: red
    style G fill: red
    style H fill: red
    style I fill: green
    style J fill: red


    A & B & C-- 0..9 --> B
    B -- UNDERLINE --> C

    B -- DOT --> D
    D & E & F -- 0..9 --> E
    E -- UNDERLINE --> F

    B & E -- e --> G
    G -- +,- --> H
    G & H & I & J -- 0..9 --> I
    I -- UNDERLINE --> J
```

### Examples

```quartz
0
0.1
2.08e12
23e2
14e-6
```
