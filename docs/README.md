# Go - Form To Struct

Golang module that allows you to persist form data into structs for storage, sanitation and database operations.

## Motivation
I was recently working on a CRUD app that required receiving user information as form data on the backend. I wanted to easily persist the data to multiple *different* structs at once rather than having to type convert and manually assign each field. I decided to just write this package myself, and use it as an opportunity to learn more about Go, its type system and reflection.

## Usage
The module currently supports two functions, Fts() and Mfts(). Fts() will persist form data into a single struct, whereas Mfts() will persist form data into multiple structs specified by the user. 

Ex. 
```go

var str struct{
    firstName string
    lastName  string
    address   string
    nums      []int
}

http.handleFunc("/index", func(w http.ResponseWriter, r *http.Request){
    r.ParseForm
    //r.Form => [ firstName: [James], lastName: [Bond], address: [25 James Bond Rd], nums: [1, 2,3]]
    var t str

    e := Fts(r.Form, &t)
    if (e != nil){
        log.fatal(e)
    }

    //t => {James Bond 25 James Bond Rd [1,2,3]}
    
})

```

## Installation

```
go get -u github.com/michelcodes123/gofts

```

## Supported Types 

The module supports types that would most commonly be received as form data for an app. Other types are ignored.
- int - int64
- uint - uint64
- float32, float64
- bool 
- string
- slice (as long as the slice is of one of the above types)
