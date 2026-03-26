## GO Syntax
A complete Go file consist of the following parts: 
  Package declaration
  Import packages
  Functions
  Statements and expressions

Package declaration
In Go, every file starts with a package line. A package is like a folder that groups related code together. This tells Go: “this is the main program.”

Import packages
Sometimes your program needs extra features (like printing text or manipulating strings). You bring these in using import. 

Functions
Functions are blocks of code that do a specific task. You define them using func.

Go (Go) is designed to be simple and easy to read, especially for beginners. Every Go program starts with `package main` and a `main()` function—this is where your code begins running. Go is strict about types (like numbers, strings, etc.), so you always know what kind of data you’re working with. It also uses simple structures like `if` for decisions and `for` for loops (it only has one loop type, but it’s flexible).

One nice thing about Go is that it avoids complicated features to keep things clean. You don’t need brackets around conditions, but you must always use `{}` for blocks of code. Instead of hiding errors, Go makes you handle them directly—functions can return an error, and you check it yourself. Go also has a special feature called goroutines (`go`) that lets you run tasks at the same time, which is useful for building fast programs.
