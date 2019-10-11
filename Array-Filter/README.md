# Filtering a slice/array
Golang has no built-in function(s) to filter a slice/array. This lesson will teach you two different ways to filter.

The first is to create a separate slice/array and add the filtered elements to it. This works great, but doesn't use memory as efficiently as the second way.

The second way shows you how to filter by altering the original slice/array in place without making another copy to hold the filtered elements. But, it has a side-effect that I'll tell you about which may help you decide which approach you will use.

[Watch](https://egghead.io/lessons/go-filter-an-array-in-go) this lesson at Egghead.io.