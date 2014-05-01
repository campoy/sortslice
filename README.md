sortslice
=========

Sorting slices with less lines

This package allows you to write:

```
	a := []int{1, 2, 3, 4}
	sortslice.Any(a)
```

It's as fast as the stdlib when the type inside the slide is primitive, and 6 times slower for types that are of primitive kind.

For other types ... it panics :-)
