# fo-experiments
This repo is used to play with the Go generic implementation at 
[albrow/fo](https://github.com/albrow/fo).

I used [data structures and algorithms](https://github.com/floyernick/Data-Structures-and-Algorithms) 
do these experiments.

## Reverse Array
This algorithm simply reverses the elements of an array. I used it to
reverse integer, string, and float arrays.

Output:
```
$ fo run ReverseArray.fo
[3 2 1]
[c b a]
[3.3 2.2 1.1]
$
```

*Observation* this generic algorithm works well because the needed operators are defined in the Go language for both numeric types and for strings. Without operator overloading, it would not work for more complex types. For example, a "Point" type having an x and y coordinate cannot be used since "<" is not defined. This short-coming could be addressed by supplying a "less than" function that the generic algorithm could use.
