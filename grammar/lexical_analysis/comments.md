# Comments

## Single line comment

### Regex

```regexp
\/\/.*(\r\n|\r|\n|)
```

### Diagram

```mermaid
graph LR;
    A((S<sub>0</sub>))
    B((S<sub>1</sub>))
    C(((S<sub>2</sub>)))
    D(((S<sub>3</sub>)))
    style C fill: green
    style D fill: green

    A -- / --> B
    B -- / --> C
    C -- ELSE --> C
    C -- NEWLINE --> D
```

### Examples

```quartz
//
// Single line comment
... // Single line comment
```

## Multi line comment

### Regex

```regexp
\/\*[\s\S]*\*\/
```

### Diagram

```mermaid
graph LR;
    A((S<sub>0</sub>))
    B((S<sub>1</sub>))
    C((S<sub>2</sub>))
    D((S<sub>3</sub>))
    E(((S<sub>4</sub>)))
    style C fill: red
    style D fill: red
    style E fill: green


    A -- / --> B
    B -- * --> C
    C -- ELSE --> C
    C & D-- * --> D
    D -- / --> E
    D -- ELSE --> C
```

### Examples

```quartz
/**/

/* Multi line comment */

/*
 * Multi line comment
 */
```
