# fo-experiments
This repo is used to play with the Go generic implementation at 
[albrow/fo](https://github.com/albrow/fo).

I used [data structures and algorithms](https://github.com/floyernick/Data-Structures-and-Algorithms) 
do these experiments.

## Reverse Array
This algorithm simply reverses the elements of an array. I used it to
reverse integer, string, and float arrays. I expected this one to be easy to port to `fo` and it was. Below shows how to execute the code. I do not show this or the output in subsequent algorithms.

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

*Observation #1* This port raised a key problem that I have seen discussed elsewhere, namely, dealing with error handling in a generic method. Typically, in Go, you return an error like so:
```
    return 0, false // integers
    return "", false // for strings
    return 0.0, false // for floats
    return nil, false // for structs
```
Go does not have a keyword that represents *any* zero value. Thus `nil` can't be used, which was my first "natural" attempt. To work around this I added a method to the `Node` type called `Value()` which returns the data value in the node. I can't, however, provide an error to the caller. This new method plus careful code to avoid empty `Node`s mitigated this issue - at least for this particular case. It would nicer to address this as part of a generics implementation. For example, suppose a new keyword named `zerovalue` was introduced, which is instantiated at compile time to a zero for integers, 0.0 for floats, "" for strings, and nil for structs. Then an error return would be simply `return zerovalue, false`.

*Observation #2* Once the generic port was done, it was trivial to use it for integers, floats, and strings. Worked perfectly and easy to understand and read.

## Bubble Sort
This uses a simple bubble sort agorighm to order a slice of type T.

Without operator overloading, the sort code cannot use a `>` operation to test two slice elements. To overcome this, I added a new first argument which is a function to use as the comparator. To wit:
```
func bubbleSort[T](gt func(T,T) bool, array []T) {
	swapCount := 1
	for swapCount > 0 {
		swapCount = 0
		for itemIndex := 0; itemIndex < len(array)-1; itemIndex++ {
            // use supplied function to evaluate two elements in slice
			if gt(array[itemIndex], array[itemIndex+1]) { 
				array[itemIndex], array[itemIndex+1] = array[itemIndex+1], array[itemIndex]
				swapCount += 1
			}
		}
	}
}
```

Then to use, I call as follows (for floats):
```
	ra := []float64{4.4, 2.2, 3.3}
	gtf := func(x,y float64) bool { return x > y }
	bubbleSort[float64](gtf, ra)
```

*Observation* Most in the Go community find operator overloading to be less than satisfying, even dangerous. Providing a function to use instead compromises code readability, but not overly so, in my opinion.

## Circular Buffer
This is similar to `container/ring` in the standard library, but is missing a number of the methods available `container/ring`. The original algorithm used an fixed size array. I changed this to a slice and added a constructor to allow the user to specify the size of the circular buffer.

*Observation* This is the first time I have needed a generic constructor. Fairly straightforward and it looks like this:
```
func New[T](sz int) *CircularBuffer[T] {
	cdata := make([]T,sz)
	cb := &CircularBuffer[T]{data:cdata, pointer: 0, size:sz }
	return cb
}
```

## Binary Search
For this test, I used both `int`s and a custom type `Point`:
```
type Point struct {
    x,y float64
}
```

I modified the original so that the binary search function was a method of a comparator type that provided the definitions of equality and less than. For the less than operation, the provided function compares the distance from the origin of the two Points. Here is the type definition:
```
type Comparators[T] struct {
    equals func(T,T) bool
    lessthan func(T,T) bool
}
```

*Observation* This worked well and I easily extended this to provide other search methods as well. There a number of search algorithms in the source repo and I implemented these:
- Binary Search
- Linear Search
- Ternary Search