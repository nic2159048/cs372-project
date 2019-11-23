# ![GOlang](https://golang.org/lib/godoc/images/go-logo-blue.svg)
### Introduction
[GOlang](https://golang.org/) (a.k.a. Go) is relatively new langauge (2009) created by google. This document gives a brief introduction to Go, focusing mostly on Go's characteristics and classification as a programming language.

## History and Current Status
Go was first released in November of 2009. Afterwards the Go team at Google continued to work on developing the project and building a community of interested developers. Go's development has been shaped by the community's feedback. As a result Go is a more intuitive language, very centralized, and substantially self-branded language (Go sounds like Gopher which has become a mascot for the language). 

Go is [currently on version](https://golang.org/project/#go1) 1.13 (as of November 11, 2019). The standard library is still expanding and the development community is being careful to make smart choices for Go's future.

Another emerging hot technology is Web Assembly. Go's compiler can compile to Web Assembly, so this might be another adoption point.

## Paradigm
Go truly escapes the imperative paradigm with built-in concurrency via goroutines and channels, but Go doesn't really fit in with the Object Oriented languages either. Unlike the common OO languages Go doesn't have a class, but it has interfaces which allow one to supply a list of functions that should exist to act on a type.

## Typing System
Go's typing system enables it to be consistent and secure without thrawrting bitwise-wizardry.

### Basic Types
Go's `int`, `uint`, `float`, and `bool` types are simmilar to those in the C-family, but Go also has specific-size varients like ECMAScript's TypedArrays (e.g. `int32`, `float64`). Go also has built-in `complex` and `string` types simmilar to those in C++. For representing unicode characters Go has `rune`, an alias for `uint32`.

### Composite Types
A composite type is defined based on relationships with basic and composite types. Like C, Go has `pointers`, `structs`, and `arrays`. Go also has `slice` types which are a dynamic containers, and `map` types which are (hashed) associative arrays. Go also has channel types for synchronization among goroutines and interfaces.

In the future GO will have custom-type generics.

 - <cite>[Type System Overview](https://go101.org/article/type-system-overview.html)</cite>

## Control Structures
## Semantics
## Desirable Language Characteristics
### Go Is FAST
Built-in concurrency, an awesome compiler, and other features make Go a fast (i.e. high-performance) language.
<cite>[Achieving Concurrency in Go](https://medium.com/rungo/achieving-concurrency-in-go-3f84cbf870ca)</cite>
<cite>[Five Things that Make Go Fast](https://dave.cheney.net/2014/06/07/five-things-that-make-go-fast)</cite>

### Uniformity & Minimalism
Another common fault in OO languages is verbosity. Consider this Java program:
```Java
package helloworld;

public class HelloWorld {

    public static void main(String[] args) {
        System.out.println("Hello World!");
    }

}
```
 - <cite>[Java's HelloWorld tutorial](https://docs.oracle.com/javase/tutorial/getStarted/cupojava/index.html)</cite>

And compare it to this:
```GO
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}
```
 - <cite>[GO's Hello World widget](https://golang.org/)</cite>

In general Go attempts to be as minimal as possible to avoid unneeded characters, semicolons, and types (indeed, the type is only needed if it can't be discerend by other means). While minimalism sometimes has negative consequences (e.g. less readable, requireing knowledge of sub-standard acronyms and abbreviations) it can also be wonderful.

Furthermore, the verbose vs. minimalist approaches can be seen in each language's hello-world tutorials.

### category 3
### category 4
## Data Abstractions
## Syntax
### Type and Declaration
Most C-family languages use a syntax for type and declaration that is similar to a normal expression, but with types explicit. This is most clear for function pointers: the declaration `int (*fp)(int a, int b);` is meant to look like the expression `fp(a,b);` which evaluates to an int.

However, the designers of Go decided to use a seperate syntax for types and declarations in order to improve orthogonality and consistency. This is especially important when dealing with higher-order functions (with multiple returns). Consider the C function taking and returning a function pointer: 
```C
int (*(*fp)(int (*)(int, int), int))(int, int);
```
And the Go equivilent: 
```GO
f func(func(int,int) int, int) func(int, int) int
```
One of goals of the seperate syntax was to make a GO declaration left->right readable. 
The above can be read: 
`f is a func taking a func taking two ints and two ints and returning a function taking two ints and returning an int.` 
It's still a mouthful--but it's far less obsfucated if you ask me!

 - <cite>[Go's Declaration Syntax](https://blog.golang.org/gos-declaration-syntax)</cite>
 - <cite>[Go's Specification (with Extended Backus-Naur Form)](https://golang.org/ref/spec#Notation)</cite>
