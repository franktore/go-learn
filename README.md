# go-learn
Go or GoLang is a programming language originally developed by Google that uses high-level syntax similar to scripting languages.
## why go?
future-proof

easy to learn

self-containing

highly scalable

easy to maintain

high performance

statically compiled -> fast

concurrency is free -> fast

great career opportunities

designed for problem-solving

## get started
* download and install go [here](https://go.dev/doc/install)
* set GOPATH=/home/(somewhere)/golang
* write some code

### packages and modules
* packages form the basic building blocks of a Go program
* packages make for re-usable units
* modules are collections of packages
* to make a package executable it must be named `main`
* folder architecture is important
    something like this is quite common
    ```
    src
    └───project
    │   │   go.mod      | 
    │   │   go.sum      | the main
    │   │   helper.go   | package
    |   |   myapp.go    | 
    |   └───pkg
    │       └───greetings
    │       |   │   hello.go    |
    │       |   │   goodbye.go  | the greetings
    │       |   │   ...         | package
    │       | ....
    |
    └───other-project
        │   go.mod
        │   go.sum
        |   other.go
        |   ...
    ```


you create a module with the `go mod init myapp` command
this will create a `go.mod` file. This is where all module dependencies will be listed.

If you want your packages downloadable you should give it a name like `github.com/yourname/myapp`


# resources
* https://go.dev/doc/install
* https://go.dev/doc/tutorial/getting-started