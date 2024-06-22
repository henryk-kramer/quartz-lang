# Control flow

## If-else expression

The if-else expression checks if a condition holds true and evaluates the first block (then-block) in that case; If not it evaluates the second block (else-block). The last expression in the evaluated block is the value of the of expression.

```quartz
if ... {...} else {...}
```

You do not need to specify an else-block. In the case that the condition does not hold true `nil` is returned.

```quartz
if ... {...} 
```

## Cond expression

The cond expression is similiar to the if expression but accepts multiple conditions before it evaluates the else block. You can think of it as nested if-else expressions.

```quartz
cond {
    ... -> ...
    ... -> ...
    else -> ... 
}
```

As in the if-else expression you do not need to specify `else`.

```quartz
cond {
    ... -> ...
    ... -> ...
}
```

## Case expression

The case expression is used to match a value against multiple given patterns. The matching logic is the same for as for normal pattern matching. The first matching block is executed and returned as value.

```quartz
case ... {
    ... -> ...
    ... -> ...
}
```

## Return (early) statement

A return statement can be used to exit a function at any time. This can prevent complex nesting and help make the code clearer.

A return statement consist of the `return` keyword and the value that should be returned. When no value is given `nil` is returned.

```quartz
fn func() -> str = {
    if ... {
        return "early return"
    }

    "normal return"
}
```

You should use a return statement with caution though. Since it breaks the normal control flow, too many or too deeply nested return statements can make the logic harder to read and reason about. Early returns are mostly useful when you need to validate input through many different rules.
