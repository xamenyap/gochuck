# Gochuck client
A client written in Go for getting random Chuck Norris facts provided by https://api.chucknorris.io/

### Installing
- Prerequisites: Go version 1.16 or higher
- Simply run: `go get github.com/xamenyap/gochuck` in your console

### Examples
```golang
// get a random chuck norris fact
fact, _ := gochuck.GetRandom()

// get fact categories
categories, _ := gochuck.GetCategories()

// get a random fact by a certain category
fact, _ := gochuck.GetRandomByCategory("food")

// search for a collection of facts by a particular query string
collection, _ := gochuck.GetByQuery("Circle of Life")
```

### License
MIT
