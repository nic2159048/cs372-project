# GOlang
### Introduction
[GOlang](https://golang.org/) (a.k.a. GO) is relatively new langauge (2009) created by google. This document gives a brief introduction to GO, focusing mostly on GO's characteristics and classification as a programming language.

## History and Current Status
GO was first released in November of 2009. Afterwards the GO team at Google continued to work on developing the project and building a community of interested developers. GO's development has been shaped by the community's feedback. As a result GO is a more intuitive language, very centralized, and substantially self-branded language (GO sounds like GOpher which has become a mascot for the language). 

GO is [currently on version](https://golang.org/project/#go1) 1.13 (as of November 11, 2019).

## Paradigm
GO is a Object-Oriented (OO) programming language; but unlike most OO languages GO truly escapes the imperative paradigm with built-in concurrency. 

## Typing System
## Control Structures
## Semantics
## Desirable Language Characteristics
### Go Is FAST
https://medium.com/rungo/achieving-concurrency-in-go-3f84cbf870ca
https://dave.cheney.net/2014/06/07/five-things-that-make-go-fast
Built-in concurrency, an awesome compiler, and other features make GO a very fast language.

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
[Java's HelloWorld tutorial](https://docs.oracle.com/javase/tutorial/getStarted/cupojava/index.html)

And compare it to this:
```GO
package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
}
```
[GO's Hello World widget](https://golang.org/)

In general GO attempts to be as minimal as possible to avoid unneeded characters, semicolons, and types (indeed, the type is only needed if it can't be discerend by other means). While minimalism sometimes has negative consequences (e.g. less readable, requireing knowledge of sub-standard acronyms and abbreviations) it can also be wonderful.

Furthermore, the verbose vs. minimalist approaches can be seen in each language's hello-world tutorials.

### category 3
### category 4
## Data Abstractions
## Syntax
[GO's notation is defined in Extended Backus-Naur Form (EBNF)](https://golang.org/ref/spec#Notation)
