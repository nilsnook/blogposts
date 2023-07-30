Title: A Tour of Go 1.16's `io/fs` package
Description: The most anticipated addition to the Go's standard library, the new `io/fs` and `testing/fs` package.
Tags: go,io,fs,testfs
---
Go's `io.Reader` and `io.Writer` interfaces, along with `os.File` and its analogs, go a long way in abstracting common operations on opened files. However, until now there hasn't been a great story for abstracting entire filesystem.

Why might you want to do this? Well, the most common motivating use-case I've encountered is being able to mock a filesystem in a test.

Abstracting the filesystem in tests can prevent tests from being distributed by side effects, and provides a more relialbe way to setup test data. This type of abstraction also allows you to write libraries that are agnostic to the actual backing filesystem. With an interface, no one knows you're a cloud blob store.
