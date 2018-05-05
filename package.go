package main

//The first statement in a Go source file must be package name
//where name is the package's default name for imports.
//All files in a package must use the same name.

//Go's convention is that the package name is the last element of the import path:
//the package imported as "crypto/rot13" should be named rot13.

//Executable commands must always use package main.

//Remote packages

//An import path can describe how to obtain the package source code using a revision control system such as Git or Mercurial.
//The go tool uses this property to automatically fetch packages from remote repositories.
//For instance, the examples described in this document are also kept in a Git repository hosted at GitHub github.com/golang/example.
//If you include the repository URL in the package's import path, go get will fetch, build, and install it automatically:

//$ go get github.com/golang/example/hello
