In Go, variables are used to store data we have different types eg:
> String: which stores text
> int: that stores integers such as 123 or -123
> bool: stores values with two states true or false.
> float32: this stores floating point numbers, and decimals
You can declare them in two main ways: using `var` (e.g., `var age int = 20`) or the shorter `:=` syntax (e.g., `age := 20`). Go is strongly typed, which means every variable has a specific type, but the language can often figure out the type for you automatically.

Variable names should be clear and meaningful, and they must start with a letter or an underscore, a variable name cannot start with a digit, and it must not contain spaces. Once declared, you can change their values unless they are constants (`const`). Go also has a special rule: if you create a variable and don’t use it, the program will give an error—this helps keep your code clean and avoids unnecessary variables.
