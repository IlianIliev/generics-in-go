## Warning

The code from this project has been copied to https://github.com/IlianIliev/go-topics where you can also find examples and videos about other topics.

This repo may be removed in the future.


---
# Generics in Go


## Overview

Generics have been available in Go since version 1.18 (15 March 2022) but two and half years later they are still a bit of mystery and are avoided by many developers.
Here I will try to explain what they are and to give some examples of how to use them.


They allow you to create generic functions and types.
Generics are a way to parameterize types and functions so that they can be used with different types.

A screencast of how I refactor the initial state to use generics is available at https://youtu.be/9G8-lzl7Q9Y


### Notes

- The examples below are for Go 1.18 or newer.
- This is in no way a production ready code. To focus on the problem I am trying to illustrate, I have simplified the rest of the code as much as possible so you may notice missing things like ignoring errors in some cases etc.
