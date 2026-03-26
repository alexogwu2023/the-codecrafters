In Go, an array is a fixed-size collection of multiple variables that all have the same data type. You define an array by specifying its size and type, like this: `var nums [3]int`, which creates an array that can hold 3 integers. You can also initialize it with values: `nums := [3]int{1, 2, 3}`. Each element is accessed using its index, starting from 0 (e.g., `nums[0]` gives the first value).

Arrays in Go have a fixed length, meaning you cannot add or remove elements after creation. Because of this, it is adviced that you use slices (which are more flexible), but arrays are still useful when you know the exact number of items you need to store.

In Go the index of an array always starts at 0 and to remove or change an element of an array you have to reference its index 
