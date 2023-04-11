# myLib CLI

myLib CLI is a command line interface program that allows you to manage a personal book collection.
It allows you to:
- Add and manage books into the system, including some basic information about those
  books (title, author, published date, edition, description, genre)
- Create and manage collections of books
- List all books, all collections and filter book lists by author, genre or a range of
  publication dates.


* [Author commands](#author-commands)
  * [List of sub commands](#list-of-sub-commands)
  * [Create an author](#create-an-author)
  * [Update an author](#update-an-author)
  * [Delete an author](#delete-an-author)
* [Genre commands](#genre-commands)
  * [List of sub commands](#list-of-sub-commands-1)
  * [Create a genre](#create-a-genre)
  * [Rename a genre](#rename-a-genre)
  * [Delete a genre](#delete-a-genre)
* [Book commands](#book-commands)
  * [List of sub commands](#list-of-sub-commands-2)
  * [Create a book](#create-a-book)
  * [Update a book](#update-a-book)
  * [Delete a book](#delete-a-book)
  * [List books](#list-books)
  * [Add a book to a collection(s)](#add-a-book-to-a-collection)
  * [Remove a book from a collection](#remove-a-book-from-a-collection)
  * [Assign Genre to a book](#assign-a-genre-to-a-book)
  * [Unassign a Genre from a book](#unassign-a-genre-from-a-book)
* [Collection commands](#collection-commands)
  * [List of sub commands](#list-of-sub-commands-3)
  * [Create a collection](#create-a-collection)
  * [Rename a collection](#rename-a-collection)
  * [Delete a collection](#delete-a-collection)
  * [List collections](#list-collections)
  * [Add a book to a collection](#add-a-book-to-a-collection)
  * [Remove a book from a collection](#remove-a-book-from-a-collection)


## Usage

```shell
$ myLib command [flags]
```

### Available commands
```
  author	Manage authors
  book		Manage books
  collection	Manage collections of books
  genre		Manage genres
  help          Help about any command
  version       Display the version of the program
```
> Use  "myLib [command] --help" for more information about a command.

### Flags
```
  -h, --help		Print help
```
---

# Author commands

## List of sub commands

### Description
Manage authors 

### Usage
```shell
$ myLib author [flags] 
$ myLib author [command]
```

### Available Commands
```
  create	Create an author
  update	Update an author
  delete	Delete an author
```

### Flags
```
  -h, --help		Print help
 ```

> Use  "myLib book [command] --help" for more information about a command.
---

## Create an author

### Usage
```shell
$ myLib author create <lastname> [OPTIONS]
```

### Flags
```
  -f, --firstname	Set the firstname (or initials) of the author 
  -m, --middlename      Set the middlename of the author
  -h, --help		Print help
```

### Examples

Create the author Tolkien with the firstname (or initial) J.R.R.
```shell
 $ myLib author create Tolkien --firstname=J.R.R.
 Author Tolkien created.
```

### Notes
 * An author is identified by its first name, middle name and last name. Once an author is created, attempting to create 
another author with the same first name, middle name and last name will display an error message.
---

## Update an author

### Usage:
```shell
$ myLib author update <lastname> [flags]
 ```
### Flags
```
  -l, --lastname	Set the last name of the author
  -f, --firstname	Set the first name of the author
  -m, --middlename      Set the middle name of the author
  -h, --help		Print help
```
 
### Examples:
Update the author Tolkien and set the first name to "John"
```shell
$ myLib author update Tolkien -f John
Author Tolkien updated.
```
Update the author "sartre" and set the name to "Sartre" and the first name to Jean Paul
```shell  
$ myLib author update sartre --lastname=Sartre --firstname="Jean Paul" 
Author Sartre updated.
```

### Notes: 
If two authors exist with the same last name, the program will prompt you 
to choose which author to update:

```shell
$ myLib author update Tolkien -f J.R.R
Two authors exist with last name Tolkien:
[1] John Tolkien
[2] Christopher Tolkien 
Choose which one you want to update or enter c to cancel. 
```
---

## Delete an author

### Usage
```shell
$ myLib author delete <lastname> [flags]
```

### Flags
```
  -h, --help		Print help
```

### Examples
Delete the author Tolkien
```shell
$ myLib author delete Tolkien
Author Tolkien deleted.
```

### Notes
If two authors exist with the same last name, the program will prompt you
to choose which author to delete:
```shell 
$ myLib author delete Tolkien
Two authors exist with last name Tolkien:
[1] John Tolkien
[2] Christopher Tolkien 
Choose which one you want to delete or enter c to cancel.
```
---

# Genre commands

## List of sub commands

### Description 
Manage genres

### Usage
```shell
$ myLib genre [flags] 
$ myLib genre [command]
```

### Available commands
```
  create	Create a genre
  rename	Rename a genre
  delete	Delete a genre
```

### Flags
``` 
  -h, --help		Print help
```

> Use  "myLib genre [command] --help" for more information about a command.
---

## Create a genre

### Usage
```shell
$ myLib genre create <name> [flags]
``` 

### Flags
```
  -h, --help		Print help
```

### Examples
Create the genre Fantasy
```shell
$ myLib genre create Fantasy
Genre Fantasy created.
```
### Notes
* A genre is identified by its name. Once a genre is created, attempting to create
  another genre with the same name will display an error message.
---

## Rename a genre

### Usage
```shell
$ myLib genre rename <name> <newName> [flags]
```

### Flags
```
  -h, --help		Print help
```

### Examples
Rename the genre "SF" to "Science Fiction"
```shell
$ myLib genre rename SF "Science Fiction"
Genre SF renamed as Science Fiction.
```
---

## Delete a genre

### Usage
```shell
$ myLib genre delete <name> [flags]
```

### Flags
```
  -h, --help		Print help
```

### Examples
Delete the genre Fantasy
```shell
$ myLib genre delete Fanatsy
Genre Fantasy deleted.
```
---

# Book commands

## List of sub commands

### Description
Manage books

### Usage
```shell
$ myLib book [command]
$ myLib book [flags]
```

### Available Commands
```
  create	Create a book
  update	Update a book
  delete	Delete a book
  list          List all the books
  add           Add a book to a collection 
  remove        Remove a book from a collection 
  assign        Assign a genre to a book 
  unassign      Unassign a genre from a book 
```

### Flags
```shell
  -h, --help		Print help
 ```

> Use  "myLib book [command] --help" for more information about a command.
---

## Create a book

### Usage
```shell
$ myLib book create <title> [flags]
```

### Flags
```
  -a, --authors	     Set the authors of the book by lastname, separated by commas 
  -y, --year         Set the publication year
  -e, --edition      Set the edition 
  -d, --descrition   Set a descrition
  -g, --genres       Set the genres of the book, seperated by commas
  -h, --help         Print help
```

### Examples

Create a book with title The Hobbit
```shell
$ myLib book create "The Hobbit"
Book The Hobbit created.
```

Create a book with title The Hobbit with author Tolkien
```shell
$ myLib book create "The Hobbit" -a Tolkien 
Book The Hobbit created.
```

Create a book with all the information possible
```shell
$ myLib book create "The Hobbit" \
-a Tolkien -y 1937 \ 
-d "Also called There and Back Again" \
-g Fantasy, Children \
-e "First edition"

Book The Hobbit created.
```

### Notes
A book is identified by its title, authors and edition. If you create a book with the identical title, authors and edition, the program will display an error message. 

### Special cases
#### Non-existent resources
 
If you create a book with an author or a genre that does not exist in the database, the program will prompt you with the possibility to create the missing information. 

```shell 
$ myLib create book "Les Hobbit" -a Tolkien
The author Tolkien doesn't exist. Do you want to create it? [y,n]
```

If you select "no" the book creation will be cancelled. If you select "yes" an author with name Tolkien will be created before the book is created. 

#### Ambiguous author
If you attempt to create a book with an ambiguous author name, the program will prompt
you to choose an author:  
```shell
$ myLib create book "The Hobbit" -a Tolkien
Two authors exists with last name Tolkien:
[1] J.R.R Tolkien
[2] Christopher Tolkien 
Choose which one you want to use or enter c to cancel. 
```
---

## Update a book

### Usage
```shell
$ myLib book update <title> [flags]
 ```
### Flags
```
  -t, --title        Set the title of the book 
  -a, --authors	     Set the authors of the book by lastname, separated by commas 
  -y, --year         Set the publication year
  -e, --edition      Set the edition 
  -d, --descrition   Set a descrition
  -g, --genres       Set the genres of the book, seperated by commas
  -h, --help         Print help
```

### Examples
Update the book The Hobbit and set the author to Tolkien
```shell
$ myLib book update "The Hobbit" -a Tolkien
Book The Hobbit updated.
```
Update the book "The hobbit" and set the title to "The Hobbit"
```shell  
$ myLib book update "The hobbit" -t  "The Hobbit"
Book The Hobbit updated
```
### Notes

#### Non-existing resource

If you update a book with an author or a genre that does not exist, 
the program will prompt you with the possibility to create the missing information.

```shell 
$ myLib update book "The Hobbit" -g Children
The genre Children doesn't exist. Do you want to create it? [y,n]
```

If you select "no", the book update will be cancelled. 
If you select "yes" a genre with the name Children will be created before the book is updated.

#### Ambiguous author
If you attempt to create a book with an ambiguous author name, the program will prompt
you to choose an author:
```shell
$ myLib create book "The Hobbit" -a Tolkien
Two authors exists with last name Tolkien:
[1] J.R.R Tolkien
[2] Christopher Tolkien 
Choose which one you want to use or enter c to cancel. 
```
---

## Delete a book

### Usage
```shell
$ myLib book delete <title> [flags] 
```

### Flags
```
  -h, --help		Print help
```

### Examples
Delete the book The Hobbit
```shell
$ myLib book delete "The Hobbit" 
Book The Hobbit deleted.
```

### Notes
If two books exist with the same title, the program will prompt you to choose which book to delete:

```shell 
$ myLib book delete The Hobbit
2 books exist with the title The Hobbit:
[1] The Hobbit by J.R.R Tolkien - 1937 - First Edition
[2] The Hobbit by J.R.R Tolkien - 1937 - Second Edition 
Choose which one you want to delete or enter c to cancel. 
```
---

## List books

### Usage
```shell
$ myLib book list [flags]
```
### Flags
```
  -a, --author	     Filter by author 
  -g, --genre        Filter by genre
  -f, --from         Filter by start of range (year->)
  -t, --to           Filter by end of range (->year)
  -h, --help	     Print help
```

### Examples

List all the books in the library
```shell
$ myLib book list
- Oryx and Crake    Margaret Atwood   2003    Dystopian fiction
- Ubik              Philip K Dick     1969    Science fiction
...
```
List all the books in the library by the author Tolkien
```shell 
$ myLib book list -a Tolkien
- The Hobbit                    J.R.R Tolkien     1937    Fantasy
- The Fellowship of the Ring    J.R.R Tolkien     1954    Fantasy   
...
```

List all the books in the library from the year 2000 to the present
```shell 
$ myLib book list -f 2000
- Book1    Author1    2000    Fantasy
- Book2    Author2    2010    Novel, Science Fiction   
...
```

List all the books in the library from 1990 to 1999
```shell 
$ myLib book list -f 1990 -t 1999
- Book1    Author1    1992    Fantasy
- Book2    Author2    1998    Novel   
...
```

### Notes

#### Ambiguous author
If you attempt to list books by an ambiguous author name, the program will prompt you
to choose the author(s) to list:
```shell
$ myLib list book -a Tolkien
Two authors exist with last name Tolkien:
[1] J.R.R Tolkien
[2] Christopher Tolkien 
Choose which one you want to use or enter c to cancel.  
```
---

## Add a book to a collection

### Usage
```shell
$ myLib book add <title> <collection> [flags] 
```

### Flags
```
  -h, --help		Print help
```

### Examples
Add the book Ubik to the collection "My favorites"
```shell
$ myLib book insert Ubik "My favorites"
The book Ubik has been added to the collection My favorites.
```

### Notes
#### Non-existent collection
If you try to add a book to a non-existent collection, the program will prompt you
with the possibility to create the collection. 

```shell
$ myLib book add Ubik "My favorites"
The collection "My favorites" doesn't exist, do you want to create it? [y/n]
```
If you select "no" the command will be canceled. If you select "yes" the collection 
will be created and the book added to the collection.

#### Ambiguous book title
If you try to add a book with an ambiguous title, the program will prompt you 
to choose the book to add to the collection:
```shell
$ myLib book add Ubik "My favorites"
Two books exist with same title:
[1] Ubik  Philip K Dick 
[2] Ubik  
Choose which one you want to use or enter c to cancel.  
```
---

## Remove a book from a collection

### Usage
```shell
$ myLib book remove <title> <collection> [flags] 
```

### Flags
```
  -h, --help		Print help
```

### Examples
Remove the book Ubik from the collection My favorites
```shell
$ myLib book remove Ubik "My favorites"
The book Ubik has been removed from the collection "My favorites".
```

## Assign a genre to a book

### Usage
```shell
$ myLib book assign <title> <genre> [flags] 
```

### Flags
```
  -h, --help		Print help
```

### Examples
Assign the genre "Science Fiction" to  the book Ubik
```shell
$ myLib book assign Ubik "Science Fiction"
The genre Science Fiction been assigned to the book Ubik.
```

## Unassign a genre from a book

### Usage
```shell
$ myLib book unassign <title> <genre> [flags] 
```

### Flags
```
  -h, --help		Print help
```

### Examples
Unassign the genre "Science Fiction" from the book Ubik
```shell
$ myLib book unassign Ubik "Science Fiction"
The genre Science Fiction been unassigned from the book Ubik.
```

# Collection commands

## List of sub commands
### Description
Manage collections

### Usage
```shell
$ myLib collection [command] 
$ myLib collection [flags] 
```

### Available Commands
```
  create	Create a collection
  rename	Rename a collection
  delete	Delete a collection
  list          List all collections
  add           Add book to collection
  remove        Remove book from collection
```
### Flags
```
  -h, --help		Print help
```

> Use  "myLib collection [command] --help" for more information about a command.
---

## Create a collection

### Usage
```shell
$ myLib collection create <name> [flags]
``` 

### Flags
```
  -h, --help		Print help
```

### Examples
Create the collection "My favorites"
```shell
$ myLib collection create "My favorites"
Collection My favorites created.
```
---

## Rename a collection

### Usage
```shell
$ myLib collection rename <name> <newName> [flags]
```

### Flags
```
  -h, --help		Print help
```

### Examples
Rename the collection "My collection" to "My favorites"
```shell
$ myLib collection rename "My collection" "My favorites"
Collection My collection renamed to My favorites
```
---

## Delete a collection

### Usage
```shell
$ myLib collection delete <name> [flags]
```

### Flags
```
  -h, --help		Print help
```

### Examples
Delete the collection My Favorites
```shell
$ myLib collection delete "My Favorites" 
Collection "My favorites" deleted
```
---

### Notes
#### Deleting a collection that contains books
If you try to delete a collection that contains books, the program will prompt you
to confirm the deletion.

```shell
$ myLib collection delete "My favorites"
The collection "My favorites" contains 5 books. 
Do you still want to delete the collection? (The books won't be deleted) [y/n]
```

## List collections

### Usage
```shell
$ myLib collection list [flags]
```
### Flags
```
  -h, --help		Print help
```
### Examples

List all the collections 
```shell
$ myLib collection list
- My Favorites
- Study books
- Technical books
...
```

---

## Add a book to a collection

### Usage
```shell
$ myLib collection add <collection> <title> [flags]
```

### Flags
```
  -h, --help		Print help
```

### Examples
Add the book Ubik to the collection My favorites
```shell
$ myLib collection add "My favorites" Ubik
The book Ubik has been added to the collection "My favorites".
```

### Notes
#### Non-existent collection
If you try to add a book to a non-existent collection, the program will prompt you
with the possibility to create the collection.

```shell
$ myLib collection add "My favorites" Ubik
The collection "My favorites" doesn't exist, do you want to create it? [y/n]
```

If you select "no" the command will be canceled. If you select "yes" the collection 
will be created and the book added to the collection.

#### Ambiguous book title
If you try to add a book with an ambiguous title, the program will prompt you
 to choose the book to add to the collection:
```shell
$ myLib book add Ubik "My favorites"
2 books exist with same title:
[1] Ubik  Philip K Dick 
[2] Ubik  
Choose which one you want to use or enter c to cancel.  
```

## Remove a book from a collection

### Usage
```shell
$ myLib collection remove <collection> <title> [flags]
```

### Flags
```
  -h, --help		Print help
```

### Examples
Remove the book Ubik from the collection My favorites
```shell
$ myLib collection remove "My favorites" Ubik
The book Ubik has been removed from the collection "My favorites".
```
