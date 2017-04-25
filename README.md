# Go dep example

This application is an example of [go dep](https://github.com/golang/dep) usage, a dependency management tool for Go language.

This example is an adaptation of the article [Using go dep as a project maintainer](https://hackernoon.com/using-go-dep-as-a-project-maintainer-641d1f3006d7).


## Setup 

First, to install go dep, you should ensure that you have have one recent version of Go installed (I'm using 1.8.1). 

Next, if you don't have Go Dep installed you need to get it: 

``` shell
$ go get -u github.com/golang/dep/cmd/dep
```

You should ensure we have `$GOPATH/bin` in the path. If not, add to your environment. This steps is dependent on your operating system.

To check if is correctly installed and configured the path:

``` shell
$ dep
```  


### The application appllication

This example application is a simple shell which will echoes the commands you write. The `exit` command will exit. It uses the `github.com/chzyer/readline` package. This will be the one we need configure.

The application source code: 

``` go
package main

import (
	"fmt"

	"github.com/chzyer/readline"
)

func main() {
	rl, err := readline.New("go-dep-example> ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil || line == "exit" {
			break
		}
		fmt.Printf("received command: %s\n", line)
	}
}
```


## Dependency management

In order to vendor the package imported in line 6, I executed this commands in the application directory:

``` shell
$ dep init 
$ dep ensure github.com/chzyer/readline@1.4
```

The last line is not estrictly required, but if you need to pin the packakge to a `readline` to a version use it.

After that the directory looks like the following:

```shell
.
├── Gopkg.lock
├── Gopkg.toml
├── main.go
├── README.md
└── vendor/
   └── github.com/
      └── chzyer/
         └── readline
``` 

Now you can `go build` or `go install` the program and execute it.


### Deal with project maintance

The first question should be: *what sould be pushed to repository?*. My first answer is, and keep in mind that I'm not an expert, once the dependencies are defined and solved you only need the `Gopkg.toml` file in the repo. Then, when you need to download or update dependencies, only need to send the `ensure` command.

``shell
$ dep ensure -update
``

But, as the `vendor/` directory is part of the project, you could also put all your dependencies in your repopsitory.

So, you have two main approaches:

  - Keep your project **as thin as** as possible. 
  - Create a **full independent** project.
