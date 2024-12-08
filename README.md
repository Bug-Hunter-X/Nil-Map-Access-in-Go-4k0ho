# Nil Map Access in Go

This example demonstrates a common, yet subtle issue in Go: accessing keys in a nil map.

In many other languages, attempting to access a key in a nil map would result in a runtime error (e.g., a `NullPointerException`).  Go, however, handles this differently.  Instead of panicking, it returns the zero value of the map's value type.

This behavior can lead to hard-to-debug issues if not carefully considered.

## The Bug

The `bug.go` file contains code that attempts to access a key in an uninitialized map.

## The Solution

The `bugSolution.go` file provides a corrected version that explicitly checks if the map is nil before accessing a key.
