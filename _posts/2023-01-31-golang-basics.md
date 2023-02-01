---
title: A Frustrating Start
date: 2023-01-31
---
##Long-winded Details
I started learning Golang by using ChatGPT, which provides a nice interactive environment. First, I installed golang on my Mac by downloading a dmg package from here: https://go.dev/dl/ and installing it.

I learned about a basic <a href="https://github.com/kate-holdener/learningGolang/blob/main/code/basics/helloWorld.go">Hello World</a> program. There is a 1-step and a 2-step way to build and run a Golang program. The 1-step process is with <code>go run fileName.go</code>. I am guessing this compiles and runs the code (although, I am not seeing where the compiled binary is). 

The 2-step process is:
1. <code>go build fileName.go</code>
2. <code>./fileName</code>

My problem witht he 2-step process is that there isn't a good way to exclude the compiled binary from getting committed to the git repository, since it doesn't have any special extension. This is something I want to dig into deeper later.

Next, I wanted to learn about working with multiple files in Golang. Looks like code is organized in packages (so similar to Java, but the package name does not need to match the directory name, but it's still a good idea to name packages after the directorys where they are contained). Here, I ran into my first big frustration. I wanted to write a basic example where I can call a function defined in one file from another file. I create a file <a href="https://github.com/kate-holdener/learningGolang/blob/main/code/basics/getName.go">getName.go</a> and defined it in a package called 'input'. Then, I attempted to import the 'input' package from <a href="https://github.com/kate-holdener/learningGolang/blob/main/code/basics/greetByName.go">greetByName.go</a> file, and got an error:
```
package input is not in GOROOT (/usr/local/go/src/input)
```
I attempted to use a relative import, but that was a failure as well. Digging more into this problem I learned that there are two modes of handling dependencies in golang: module mode and GOPATH mode. Module mode is the newer and preferred way of handling dependencies and here's what ChatGPT had to say about it: "This mode is based on Go modules, which are a new way of managing dependencies in Go. In module mode, Go uses a file named go.mod to manage dependencies, and all imports must use absolute import paths. This mode is the recommended way to manage dependencies in Go and is designed to be simple and straightforward."

I am a little stumped, because I really don't want to use absolute import paths in my code. I must be misunderstanding something here. I will now dig into the module mode to understand it better.

...

Turns out, I needed to define my code (which is in the 'basics' directory) as a "module with:
```
go mod init basics
```
Then, I can import code in sub-directories. For example, I have an 'input' sub-directory, with <a href="https://github.com/kate-holdener/learningGolang/blob/main/code/basics/input/getName.go">getName.go</a> file, which is part of an <code>input</code> package. I can import that code with:
```
import basics/input
```

I was curious about the connection between the directory name and package name. So I kept the 'input' as the directory name and changed the package to userInput. Turns out, the 'import' statement refers to the directory (not the package). When using code from a package, I need to mention the package name explicitly. <a href="https://golangbyexample.com/package-folder-name-golang/">This resource</a> was very instrumental to my understanding of this finer point.

##Summary
When starting a golang project, create a directory and initialize the project with:
``` 
go mod init <project_name>
```
The main entry point into the program is the main function, which must be in a file containing: <code>package main</code> at the top. You can add other functions to the same file or import them from other files. To import code from other files, add: <code>import "project_name/<dir_name>"</code>. This will import the packages contained in <dir_name>. In the <dir_name> directory, add .go files with your code, and give them a package name at the top with: <code>package package_name</code>. Then (once imported), you can reference the functions you wrote with <package_name>.<function_name>.

