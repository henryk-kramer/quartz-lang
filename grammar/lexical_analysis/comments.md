# Comments

## Single line comment

### Regex

```regexp
\/\/.*(\r\n|\n\r|\r|\n)
```

### Diagram

```mermaid
graph LR;
    A((S0));
    B((S1));
    C(((S2)));
    D(((S3)));
    E(((S4)));
    F(((S5)));
    G(((S6)));
    style C fill: green;
    style D fill: green;
    style E fill: green;
    style F fill: green;
    style G fill: green;

    A -- / --> B;
    B -- / --> C;
    C -- ELSE --> C;
    C -- CR --> D;
    C -- LF --> E;
    D -- LF --> F;
    E -- CR --> G;
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
    A((S0));
    B((S1));
    C((S2));
    D((S3));
    E(((S4)));
    style C fill: red;
    style D fill: red;
    style E fill: green;


    A -- / --> B;
    B -- * --> C;
    C -- ELSE --> C;
    C & D-- * --> D;
    D -- / --> E;
    D -- ELSE --> C;
```

### Examples

```quartz
/**/

/* Multi line comment */

/*
 * Multi line comment
 */
```
