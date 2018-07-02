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

*Observation* This generic algorithm works well because the needed operators are defined in the Go language for both numeric types and for strings. Without operator overloading, it would not work for more complex types. For example, a "Point" type having an x and y coordinate cannot be used since "<" is not defined. This short-coming could be addressed by supplying a "less than" function that the generic algorithm could use.

## Doubly Linked List
This creates a list that can be traversed head to tail or tail to head.

*Observation One* This port raised a key problem that I have seen discussed elsewhere, namely, dealing with error handling in a generic method. Typically, in Go, you return an error like so:
```
    return 0, false // integers
    return "", false // for strings
    return 0.0, false // for floats
    return nil, false // for structs
```
Go does not have a keyword that represents the zero value. Thus `nil` can't be used, which was my first "natural" attempt. To work around this I added a method to the `Node` type called `Value()` which returns the data value in the node. I can't, however, provide an error to the caller. This new method plus careful code to avoid empty `Node`s mitigated this issue - at least for this particular case.

*Observation Two* Once the generic port was done, it was trivial to use it for integers, floats, and strings. Worked perfectly and easy to understand and read.