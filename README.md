## Go For Beginners

This is a set of example programs in Go (golang) to learn Go. The Go language is often referred to as golang to help searches.

### How to run ?
* Install GO (https://golang.org/doc/install)
* Setup GOPATH (https://github.com/golang/go/wiki/SettingGOPATH)
* move to GOPATH : cd $GOPATH
* Build : go build go-playbook
* Run : ./go-playbook

### INDEX : TODO
* package : main.go , package.go
* import : main.go
* exported-name : main.go
* function, named return values : package go-playbook/function
* function values : function-values.go
* variables , short variable declarations: main.go
* constants : package go-playbook/function
* for : control/control.go
* switch : switch.go
* defer : defer.go
* pointer : pointer.go
* array : array.go
* array slice : array_slice.go , making-slice.go , slice_literals.go, slice-len-cap.go
* struct : struct.go
* struct literals : struct-literals.go
* interface : interface.go , interface-values.go , interfaces-are-satisfied-implicitly.go
* empty interface : empty-interface.go
* method : method.go
* range : range.go
* stringers : stringers.go
* error : error.go
* type switch : type-assertions.go
* concurrency : package concurrency

### Advanced
* oops : object-oriented-implementations.go
* http : http.go
* web : web-app-base.go
* time and date : time-date.go
* jason : json.go

### ORM
* This package is sample with GORM and sqlite
* install GORM and go-sqlite3 (go get -u github.com/jinzhu/gorm , go get github.com/mattn/go-sqlite3)
* Refer the orm/init.go for the implementation

## Contribute

corrections and any contributions are encouraged, please submit a pull request with your change or an issue for a bug or fix.

## Resources

This set of examples assumes a certain level of programming experience and is intended for someone learning the Go language and not someone new to programming altogether.

The official site has a [Tour of Go](http://tour.golang.org/) which is an interactive walk through, another good introduction to the language.

### Other references
* https://gobyexample.com/
* https://golang.org/doc/
* https://github.com/mkaz/working-with-go/tree/master/euler
* https://github.com/GoesToEleven/GolangTraining
* [A curated list of awesome Go frameworks, libraries and software](https://github.com/avelino/awesome-go)
* [Web Frameworks](https://github.com/mingrammer/go-web-framework-stars)
* [Go ORM](https://github.com/jinzhu/gorm)
* [Go dependency management tool](https://github.com/golang/dep)
* [Gophers Slack Channel](http://gophers.slack.com/messages/awesome)
