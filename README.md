# Go - Form To Struct

Golang module that allows you to persist form data into structs for storage, sanitation and database operations.

## Motivation
I was recently working on a CRUD app that required receiving user information as form data on the backend. I wanted to easily persist the data to multiple *different* structs at once rather than having to type convert and manually assign each field. I decided to just write this package myself, and use it as an opportunity to learn more about Go, its type system and reflection.

## Usage


## Installation

## Supported Types 

The module supports types that would most commonly be received as form data for an app.
- int - int64
- uint - uint64
- float32, float64
- bool 
- string
- slice (as long as the slice is of one of the above types)
