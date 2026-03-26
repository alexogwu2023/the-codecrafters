In Go, a map is a built-in data type that stores data in key–value pairs. It’s similar to a dictionary, where each unique key is used to retrieve a corresponding value. For example, you can use a map to store names as keys and their ages as values, making it easy to look up information quickly.

You create a map using the make function or a map literal, and you access or update values by referencing their keys (e.g., m["key"]). If you try to access a key that doesn’t exist, Go returns the zero value of the value type. Maps are unordered, meaning the data is not stored in any specific sequence.

Maps are widely used for counting, lookups, and organizing data efficiently.