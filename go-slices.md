In Go, a slice is a flexible version of an array. Unlike arrays, slices don’t have a fixed size—you can grow or shrink them as needed. You can create a slice like this: `nums := []int{1, 2, 3}`, and you can add more elements using `append(nums, 4)`..

Slices also let you work with parts of data easily. For example, `nums[0:2]` gives you a portion of the slice (from index 0 up to, but not including, index 2). Internally, slices reference an underlying array, but you don’t usually need to worry about that as a beginner—they’re mainly just a convenient way to handle lists of data.
