In Go, the `switch` statement is used to compare a value against multiple cases, making your code cleaner than writing many `if-else` statements. For example:


switch day {
case "Monday":
    fmt.Println("Start of the week")
case "Friday":
    fmt.Println("End of the week")
default:
    fmt.Println("Midweek")
}
```

Go automatically breaks after a matching case. The `default` case runs if none of the other cases match. This makes `switch` useful for handling multiple possibilities in an organized way.
