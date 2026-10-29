# The Mulang Programming Language

Why Mulang?

- Why not?

Mulang is a relatively small programming language (implemented in ~3500 loc) that began as my side project.

Since the language is small, here is the syntax:

### Assign a variable
```mu
thoos x = 43;
```

We currently support the following data types:
- Integer
- Floats
- Strings
- Arrays
- Boolean
- Hashmaps

Integers, floats, and strings are defined the usual way. Please note that we do not yet support scientific notation for floats.

### Functions

Defining a function is as easy as
```mu
thoos add = kaam_gar(x, y) {
    patha x + y;
};
add(3, 4);
```

Yes, functions should be assigned to a variable by writing out a function expression. Mulang supports only one return value per function which is returned using the `patha` keyword.
Since functions are first class citizens, you can easily pass functions into another function.

### Arrays
Arrays are defined the usual way.
```mu
thoos myArr = [10, 20, 30];
```
Mulang supports all types of data structures in an array like Javascript does (our inspiration)
```mu
thoos myArr = [1, "newstring", 3.14];
```

Arrays can be indexed and modified as you would do in the English version of this language(JS).

### Boolean
The two truth values in Mulang are:
```mu
thoos yes = sacho;
thoos no = jhut;
```

### Conditionals
The vanilla if / else-if / else statement (with blocks) is supported. You can chain as many `nabhae` branches as you need, with an optional `natra` fallback at the end.
```mu
$ No alternative $
yedi(condition) {
    ...
}

$ With alternative $
yedi(condition) {
    ...
} nabhae(condition 2) {
    ...
} nabhae (condition 3) {
    ...
} natra {
    ...
}
```

### Comments
Comments should begin and end with `$` as shown above.

### Loops
Mulang supports two kinds of loops. The `jaba_samma` loop (your traditional `while` loop), and `ghuma` loop (also called the `for` loop traditionally).

```mu
thoos i = 0;
thoos sum = 0;
jaba_samma(i < 100) {
    sum = sum + i;
    i = i + 1;
}
```

```mu
thoos i = 0;
thoos sum = 0;
ghuma(i = 0; i < 42; i = i + 1) {
    sum = sum + i;
}
```

### Hashmaps
Mulang also supports hashmaps. The keys of the hashmap should be strings. Value can be anything.

```mu
thoos country_codes = { "NP": "+977", "IN": "+91" };
```

As with arrays, hashmaps can be indexed.

### Builtins
We support a few builtin functions as of now:

#### `lambai`
Applicable to:
- Arrays
- Hashmaps
- Strings

Returns the length (integer) in each case.

Use case
```mu
thoos x = [1,2,3];
lambai(x);
```

#### `khaad`
Applicable to:
- Arrays

Python's `append` equivalent. The first argument is the array, the second is the element to append.
```mu
thoos myarr = [1,2,3];
khaad(myarr, 4);
```

#### `udaa`
The first argument should be an array, the second argument is index, and is optional.
If supplied, removes the element at the given index, and returns the removed object. If not supplied, does the same thing to the last element of the array

#### `bhan`
Your `print()` equivalent.
```mu
bhan("Hello, world!");
```

#### `abs`
Applicable to floats and intgers. Calcualtes the absolute value
```mu
thoos x = -42;
thoos y = abs(x);
```

Please check out the `example-programs` to know more.

Please note that the language is in the pre-alpha stage. You may encounter bugs. We encourage you to report any issues you find. 

Please help us grow by contributing! If you wish to add more features, please write corresponding tests as well!

Next step:
- [ ] Improve error messages

Build using

```bash
make
```

Run the REPL

```bash
./build/mu
```
Or, run the interpreter on a file

```bash
./build/mu example.mu
```
