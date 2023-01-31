---
title: I learned Golang basics
date: 2023-01-31
---

I started learning Golang by using ChatGPT, which provides a nice interactive environment. First, I installed golang on my Mac by downloading a dmg package from here: https://go.dev/dl/ and installing it.

I learned about a basic <a href="https://github.com/kate-holdener/learningGolang/blob/main/code/basics/helloWorld.go">Hello World</a> program. There is a 1-step and a 2-step way to build and run a Golang program. The 1-step process is with <code>go run fileName.go</code>. I am guessing this compiles and runs the code (although, I am not seeing where the compiled binary is). 

The 2-step process is:
1. <code>go build fileName.go</code>
2. </code>./fileName</code>

My problem witht he 2-step process is that there isn't a good way to exclude the compiled binary from getting committed to the git repository, since it doesn't have any special extension. This is something I want to dig into deeper later.

Next, I wanted to learn about working with multiple files in Golang. Looks like code is organized in packages (so similar to Java, but the package name does not need to match the directory name, but it's still a good idea to name packages after the directorys where they are contained). Here, I ran into my first big frustration. I wanted to write a basic example where I can call a function defined in one file from another file. I create a file <a href="https://github.com/kate-holdener/learningGolang/blob/main/code/basics/getName.go">getName.go</a> and defined it in a package called 'input'. Then, I attempted to import the 'input' package from <a href="https://github.com/kate-holdener/learningGolang/blob/main/code/basics/greetByName.go">greetByName.go</a> file, and got an error:
```
package input is not in GOROOT (/usr/local/go/src/input)
```
I attempted to use a relative import, but that was a failure as well. Digging more into this problem I learned that there are two modes of handling dependencies in golang: module mode and GOPATH mode. Module mode is the newer and preferred way of handling dependencies and here's what ChatGPT had to say about it: "This mode is based on Go modules, which are a new way of managing dependencies in Go. In module mode, Go uses a file named go.mod to manage dependencies, and all imports must use absolute import paths. This mode is the recommended way to manage dependencies in Go and is designed to be simple and straightforward."

I am a little stumped, because I really don't want to use absolute import paths in my code. I must be misunderstanding something here.
