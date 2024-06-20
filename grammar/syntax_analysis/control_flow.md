# Control flow

## If

```quartz
if ... {...} else {...}
```

## Cond

```quartz
cond {
    ... -> ...
    ... -> ...
    else -> ... 
}
```

## Case

```quartz
case ... {
    ... -> ...
    ... -> ...
}
```

## Return

```quartz
fn func() -> str = {
    if(...) {
        return "early return"
    }

    "normal return"
}
```

```quartz
/*
 * CUST0X01 = Premium
 * VIP0001  = Premium
 * CUST0001 = Not premium
 */

fn is_premium_1() -> bool = {
    let number = get_customer_number()

    if(String::starts_with(number, "VIP")) {
        true
    } else {
        let number = String::substring(number, 5, 1)

        if(number == "X") {
            true
        } else {
            false
        }
    }
}

// Early return
fn is_premium_2() -> bool = {
    let number = get_customer_number()

    if(String::starts_with(number, "VIP")) {
        return true
    }

    let number = String::substring(number, 5, 1)

    if(number == "X") {
        return true
    }

    false
}
```