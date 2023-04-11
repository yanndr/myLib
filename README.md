# myLib program

## Description

myLib is a program divided into two projects, a REST API and a CLI client.

The program allows you to manage a personal book collection as follows:
- Add and manage books into the system, including some basic information about those
  books (title, author, published date, edition, description, genre)
- Create and manage collections of books
- List all books, all collections and filter book lists by author, genre or a range of
  publication dates.

## Documentation
* Step 1: [Usage_Step1.md](/Usage_Step1.md)
* Step 2: [RestApi_Step2.md](/RestApi_Step2.md)
* Step 3: [Database_Step3.md](/Database_Step3.md)

## Installation
To install the programs, run:

For the API
```shell
make api
```

For the client:
```shell
make client
```

### Run the REST API
```shell
myLibApi
```
### Run the Client
```shell
myLib
```

## Notes
For this assignment, I only implemented a slice of the program: author management. 
The client will allow you to create, update and delete authors.  

The REST API provides all the methods to create, update and delete authors. 
It also lists authors and filters the results by the author last name. 