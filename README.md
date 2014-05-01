sortslice
=========

Sorting slices with less lines

This package allows you to write:

<code>
  a := []int{1,2,34}
  
  sortslice.Any(a)
</code>

It's as fast as the stdlib when the type inside the slide is primitive, and 6 times slower for types that are of primitive kind.

For other types ... it panics :-)
