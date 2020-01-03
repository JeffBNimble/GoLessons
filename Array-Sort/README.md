# Sorting in Go
There are numerous ways to sort slices in Go.

The simplest way uses some basic, built-in functions from Go's **sort** package to allow you to easily sort slices of primitives, that being **Strings**, **Ints** and **Float64s**. Using these sort functions, we don't have control over the sort order as the data will be sorted in ascending alphabetical (for strings) or ascending numerical (for Ints and Float64s) order.

Depending upon what you want to do, you can use some of the simple sorting functions in the **sort** package.

In these lessons, we will take a quick look at the numerous ways to sorts slices, including:
* Sorting slices of primitives (Strings, Ints and Float64s)
* Sorting slices of primitives or structs using a **Less Func** (less function)

## How to run the examples
This project contains examples for two lessons:
* Sorting primitives => sort_primitives.go
* Sorting slices => sort_slices.go

Change the function that is being called in the main func of main.go.

To run the sorting primitives example, change main,go to call sortPrimitives().

To run the sorting slices example, change main.go to call sortSlices().
