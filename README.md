# GoCRUD

This project is for me to learn more about the go language and also to create a relatively generic crud app I can reuse later.

I made this in an ubuntu VM but you should be able to apply to this basically any platform (windows does environmental variables slightly differently). 

###Requirements

* GoLang installed (https://github.com/golang/go/wiki/Ubuntu)
* Ubuntu (or any platform really)
* mysql (or another sql but you may need to mess with the project a bit)

### Installation

You need to set your GOPATH to the root folder of this repo and also
add the bin folder to your path. If you've used go before and/or you've setup your GOPATH, you can skip everything after the first step.

```sh
$ git clone https://github.com/SyedWasiHaider/GoCRUD.git
$ cd GoCRUD
$ mkdir pkg
$ mkdir bin
$ pwd
/some/path/you/need/to/copy/GoCRUD
$ export GOPATH=/some/path/you/need/to/copy/GoCRUD
$ export PATH=$PATH:$GOPATH/bin
```

Next make sure you have all the packages this project depends on.

```sh
cd crud
go get github.com/go-sql-driver/mysql
go get -u github.com/jinzhu/gorm
go get github.com/gorilla/mux
```
Run the program:
```sh
go run *.go
```

### Usage

If everything went well, you should have a local webserver running at
localhost:1234. You can make some requests using postman or your REST client of choice.

### TODO

 - Add oauth 
 - Write Tests
