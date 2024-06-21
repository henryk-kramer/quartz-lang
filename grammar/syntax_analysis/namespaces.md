# Namespaces

A namespace is a way to organize your code and group specific functions, classes, constants and types.

## Defining namespaces

The namespace statement is optional but has to be the first statement in each file if used. If no namespace is specified the global namespace is used.

```quartz
namespace some::specific::name
```

It is possible to define the same namespace in multiple files so they can access each others functions, classes, constants and types directly.

## Exporting

Each function, class, constant or type prepended with a `pub` is exported automatically.

```quartz
namespace foo::bar

pub const const_name_1 = ...    // exported
const const_name_2 = ...        // not exported
```

## Importing

Everything from the global namespace is imported automatically. All other namespaces can be imported by using the `import` keyword. 

```quartz
import foo::bar::foobar
import foo::bar
```

It is possible to only import specific elements from a namespace.

```quartz
import [
    func_name/3, class_name, type_name, const_name
] from foo::bar
```

When you want to reference namespaces or individual elements from a namespace by a different name you can use the `as` keyword. Only simple identifiers are allowed.

```quartz
import foo::bar as foobar
import [foo/1 as bar] from foo::bar
```

If you do not import a namespace you can access exported elements by its full name.

## External packages

⚠️ Need to write ⚠️
