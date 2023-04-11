# REST-API

## Description
This is the REST API used by the myLib client.
This document describes the structure of the API with examples.

## API Structure

* [`/`](#)
  * [`/v1`](#v1)
    * [`/v1/genres`](#v1genres)
      * [`/v1/genres/{name}`](#v1genresname)
        * [`/v1/genres/{name}/books`](#v1genresnamebooks)
    * [`/v1/authors`](#v1authors)
      * [`/v1/authors/{id}`](#v1authorsid)
        * [`/v1/authors/{id}/books`](#v1authorsidbooks)
    * [`/v1/books`](#v1books)
      * [`/v1/books/{id}`](#v1booksid)
        * [`/v1/books/{id}/genres`](#v1booksidgenres)
          * [`/v1/books/{id}/genres/{name}`](#v1booksidgenresname)
    * [`/v1/collections`](#v1collections)
      * [`/v1/collections/{name}`](#v1collectionsname)
        * [`/v1/collections/{name}/books`](#v1collectionsnamebooks)
          * [`/v1/collections/{name}/books/{id}`](#v1collectionsnamebooksid)


## Return values

There are four standard return types:

* No-content response
* Content response
* Created response
* Error response

### No-content response
This response is returned for update (PUT, PATCH) or delete (DELETE) operation.
```json
{
  "status_code": 200,
  "status": "Success"
}
```
HTTP code must be 200.

### Content response
This response is returned for GET operation
```json
{
  "status_code": 200,
  "status": "Success",
  "content": {}
}
```
Http code must be 200

### Created response
This response will be returned for creation operation (POST)
```json
{
  "status_code": 201,
  "status": "Success",
  "location": "/v1/books/1"
}
```
HTTP code must be 201

### Error response
This response is returned when an error occurs.
```json
{
  "status_code": 500,
  "status": "Error",
  "error_code": "ERR-NOT-FOUND",
  "error": "error message",
  "details": "the resource x was not found"
}
```
HTTP code must be one of 400, 404, 409, 500 or 501.

## API details
### `/`

#### GET
* Description: List of supported APIs
* Return: A content response with the list of supported API endpoint URLs.

Output:
```json
{
  "status_code": 200,
  "status": "Success",
  "content": [
   "/v1"
  ]
}
```
---

### `/v1`

#### GET
* Description: Information about the v1 API
* Return: A content response with information about the v1 API

Output:
```json
{
  "status_code": 200,
  "status": "Success",
  "content": {
    "api_version": "v1.0.0"
  }
}
```
---

### `/v1/genres`

#### GET
 * Description: List of book genres
 * Return: A content response with the list of book genres

Output:
```json
{
  "status_code": 200,
  "status": "Success",
  "content": [
    {"id":  1, "name": "Science Fiction"},
    {"id":  2, "name":  "Historical novel"}
  ]
}
```

#### POST
 * Description: Create a new book genre
 * Return: A created response with the location of the new genre

Input:
```json
{
  "name": "Fantasy"
}
```
Output:
```json
{
  "status_code": 201,
  "status": "Success",
  "location":"/v1/genres/Fantasy" 
}
```
---

### `/v1/genres/{name}`

#### GET
 * Description: Get a book genre by its name
 * Return: A content response with the book genre

Output:
```json
{
  "status_code": 200,
  "status": "Success",
  "content": {
    "id":  1, 
    "name": "Science Fiction"
  }
}
```

#### PUT
* Description: Update a book genre
* Return: A no-content response

Input:
```json
{
  "name": "Fantasy"
}
```

Output:
```json
{
  "status_code": 200,
  "status": "Success"
}
```

#### DELETE

* Description: Delete a genre
* Return: A no-content response

Output:
```json
{
  "status_code": 200,
  "status": "Success"
}
```
---

### `/v1/genres/{name}/books`
#### GET
* Description: List all the books of one genre
* Return: A content response with the List of all the books of the genre

Output:
```json
{
  "status_code": 200,
  "status": "Success",
  "content": [
    {
      "id": 1,
      "title": "book1",
      "published_date": 1937,
      "genres":[
        {"id":1,"name": "Fantasy"}
      ],
      "authors": [
        {"id": 1, "last_name": "Tolkien", "first_name": "J.R.R."}
      ],
      "edition": "1st Edition"
    },
    {
      "id": 2,
      "title": "book2",
      "published_date": 2005,
      "genres":[
        {"id":2,"name": "Novel"}
      ],
      "authors": [
        {"id": 2, "last_name": "Hugo", "first_name": "Victor"}
      ],
      "edition": "1st Edition"
    }
  ]
}
```
---

### `/v1/authors`

#### GET
* Description: List all the authors
* Return: A content response with the list of all authors

Output:
```json
{
  "status_code": 200,
  "status": "Success",
  "content": [
    {
      "id":  1, 
      "last_name": "Hugo", 
      "first_name": "Victor"
    },
    {
      "id":  2, 
      "last_name":  "Tolkien", 
      "first_name": "John", 
      "middle_name": "Ronald Reuel"}
  ]
}
```
##### Filtering
The authors list can be filtered by last name. To filter, add the `lastname` argument to the GET authors query:

Example:
```
/v1/authors?lastname=Hugo
```

#### POST
* Description: Create a new author
* Return: A created response with the location of the new resource

Input:
```json
{
  "last_name":  "Tolkien", 
  "first_name": "John",
  "middle_name": "Ronald Reuel"
}
```
Output:
```json
{
  "status_code": 201,
  "status": "Success",
  "location":"/v1/authors/1" 
}
```
---

### `/v1/authors/{id}`

#### GET
* Description: Get an author by its id
* Return: A content response with the author

Output:
```json
{
  "status_code": 200,
  "status": "Success",
  "content": {
    "id": 1,
    "last_name":  "Tolkien",
    "first_name": "John",
    "middle_name": "J.R.R"
  }
}
```

#### PUT
* Description: Update an author
* Return: A no-content response

Input:
```json
{
  "last_name":  "Tolkien",
  "first_name": "J.R.R",
  "middle_name": ""
}
```

Output:
```json
{
  "status_code": 200,
  "status": "Success"
}
```

#### PATCH
* Description: Partially update an author
* Return: A no-content response

Input:

Note that you don't need to update all the fields with a PATCH, but
here's an example of an input that updates all possible fields:
```json
{
  "last_name":"Tolkien",
  "modified_last_name": true,
  "first_name": "J.R.R",
  "modified_first_name": true,
  "middle_name":"",
  "modified_middle_name":true
}
```
Here is a more realistic example that updates only one field:
```json
{
  "first_name": "J.R.R",
  "modified_first_name": true
}
```
Output:
```json
{
  "status_code": 200,
  "status": "Success"
}
```

#### DELETE

* Description: Delete an author
* Return: A no-content response

Output:
```json
{
  "status_code": 200,
  "status": "Success"
}
```
---

### `/v1/authors/{id}/books`
#### GET
* Description: List all the books by one author
* Return: A content response with the list of all the books by the author

Output:
```json
{
  "status_code": 200,
  "status": "Success",
  "content": [
    {
      "id": 1,
      "title": "book1",
      "published_date": 1937,
      "genres":[
        {"id":1,"name": "Fantasy"}
      ],
      "authors": [
        {"id": 1, "last_name": "Tolkien", "first_name": "J.R.R."}
      ],
      "edition": "1st Edition"
    },
    {
      "id": 2,
      "title": "book2",
      "published_date": 1939,
      "genres":[
        {"id":2,"name": "Novel"}
      ],
      "authors": [
        {"id": 2, "last_name": "Tolkien", "first_name": "J.R.R."}
      ],
      "edition": "2nd Edition"
    }
  ]
}
```
---

### `/v1/books`

#### GET
* Description: List all the books
* Return: A content response with the list of all the books

Output:

```json
{
  "status_code": 200,
  "status": "Success",
  "content": [
    {
      "id": 1,
      "title": "book1",
      "published_date": 1936,
      "genres":[
        {"id":1,"name": "Fantasy"}
      ],
      "authors": [
        {"id": 1, "last_name": "Tolkien", "first_name": "J.R.R."}
      ],
      "edition": "1st Edition"
    },
    {
      "id": 2,
      "title": "book2",
      "published_date": 2005,
      "genres":[
        {"id":2,"name": "Novel"}
      ],
      "authors": [
        {"id": 2, "last_name": "Hugo", "first_name": "Victor"}
      ],
      "edition": "1st Edition"
    }
  ]
}
```
##### Filtering
The book list can be filtered by adding arguments to the GET authors query.

Table of available arguments:

| Arguments | Description                                  | Example                |
|-----------|----------------------------------------------|------------------------|
| authors   | A list of comma separated authors            | ?authors=Tolkien,Hugo  | 
| title     | A book title                                 | ?title=The%20Hobbit    | 
| genres    | A list of comma separated genres             | ?genres=Fantasy,SciFi  |
| edition   | An edition                                   | ?edition=1st%20Edition |
| from      | The date of the start of the range (year->)  | ?from=1936             |
| to        | The date of the end of the range (->year)    | ?to=1990               |

Examples with more than one argument:
```
/v1/books?title=The%20Hobbit&edition=First%20Edition

/v1/books?genres=Mystery,Fantasy&from=1950&to1970
```

#### POST
* Description: Create a new book
* Return: A created response with the location of the new book

Input:
```json
{
  "title": "book1",
  "published_date": 1936,
  "genres": [
    {
      "id": 1,
      "name": "Fantasy"
    }
  ],
  "authors": [
    {
      "id": 1,
      "last_name": "Tolkien",
      "first_name": "J.R.R."
    }
  ],
  "edition": "1st Edition"
}
```
Output:
```json
{
  "status_code": 201,
  "status": "Success",
  "location": "/v1/books/1"
}
```
---

### `/v1/books/{id}`

#### GET
* Description: Get a book by its id
* Return: A content response with the book

Output:
```json
{
  "id": 1,
  "title": "book1",
  "published_date": 1936,
  "genres": [
    {
      "id": 1,
      "name": "Fantasy"
    }
  ],
  "authors": [
    {
      "id": 1,
      "last_name": "Tolkien",
      "first_name": "J.R.R."
    }
  ],
  "edition": "1st Edition"
}
```

#### PUT
* Description: Update a book
* Return: A no-content response

Input:
```json
{
  "title": "book1",
  "published_date": 1936,
  "genres": [
    {
      "id": 1,
      "name": "Fantasy"
    }
  ],
  "authors": [
    {
      "id": 1,
      "last_name": "Tolkien",
      "first_name": "J.R.R."
    }
  ],
  "edition": "1st Edition"
}
```

Output:
```json
{
  "status_code": 201,
  "status": "Success"
}
```

#### PATCH
* Description: Partially update a book
* Return: A no-content response

Input:

Note that you don't need to update all the fields with a PATCH, but 
here's an example of an input that updates all possible fields:
```json
{
  "title": "book1",
  "published_date": 1936,
  "genres": [
    {
      "id": 1,
      "name": "Fantasy"
    }
  ],
  "authors": [
    {
      "id": 1,
      "last_name": "Tolkien",
      "first_name": "J.R.R."
    }
  ],
  "edition": "1st Edition",
  "modified_title":true,
  "modified_published_date":true,
  "modified_genres":true,
  "modified_authors":true,
  "modified_edition":true
}
```
Here is a more realistic example that updates only one field:

```json
{
  "edition": "1st Edition",
  "modified_edition":true
}
```

Output:
```json
{
  "status_code": 201,
  "status": "Success"
}
```

#### DELETE
* Description: Delete a book
* Return: A no-content response

Output:
```json
{
  "status_code": 201,
  "status": "Success"
}
```
---

### `/v1/books/{id}/genres`

#### GET
* Description: List all the genres assigned to a book
* Return: A content response with the list of all the genres assigned to a book

Output:
```json
{
  "status_code": 200,
  "status": "Success",
  "content": [
    {"id":  1, "name": "Science Fiction"},
    {"id":  2, "name":  "Historical novel"}
  ]
}
```
---

### `/v1/books/{id}/genres/{name}`

#### PUT
* Description: Assign a genre to a book
* Return: A no-content response

Output:
```json
{
  "status_code": 201,
  "status": "Success"
}
```

#### DELETE
* Description: Unassign a genre from a book
* Return: A no-content response

Output:
```json
{
  "status_code": 201,
  "status": "Success"
}
```
---

### `/v1/collections`

#### GET
* Description: List all the book collections
* Return: A content response with the list of all the book collections

Output:
```json
{
  "status_code": 200,
  "status": "Success",
  "content": [
    {"id":  1, "name": "My collection"},
    {"id":  2, "name":  "Technical books"}
  ]
}
```

#### POST
* Description: Create a new book collection
* Return: A created response with the location of the new collection

Input:
```json
{
  "name": "My collection"
}
```
Output:
```json
{
  "status_code": 201,
  "status": "Success",
  "location": "/v1/collections/My%20collection"
}
```

### `/v1/collections/{name}`

#### GET
* Description: Get a book collection
* Return: A content response with the book collection

Output:
```json
{
  "status_code": 201,
  "status": "Success",
  "content": {
    "id": 1,
    "name": "My collection"
  }
}
```

#### PUT
* Description: Update a book collection
* Return: A no-content response

Input:
```json
{
  "name": "My collection"
}
```
Output:
```json
{
  "status_code": 201,
  "status": "Success"
}
```

#### DELETE
* Description: Delete a collection
* Return: A no-content response


Output:
```json
{
  "status_code": 201,
  "status": "Success"
}
```

### `/v1/collections/{name}/books`

#### GET
* Description: List all the books in a collection
* Return: A content response with the list of all the books in the collection


Output:
```json
{
  "status_code": 201,
  "status": "Success",
  "content": [
    {
      "id": 1,
      "title": "book1",
      "published_date": 1936,
      "genres":[
        {"id":1,"name": "Fantasy"}
      ],
      "authors": [
        {"id": 1, "last_name": "Tolkien", "first_name": "J.R.R."}
      ],
      "edition": "1st Edition"
    },
    {
      "id": 2,
      "title": "book2",
      "published_date": 2005,
      "genres":[
        {"id":2,"name": "Novel"}
      ],
      "authors": [
        {"id": 2, "last_name": "Hugo", "first_name": "Victor"}
      ],
      "edition": "1st Edition"
    }
  ]
}
```

### `/v1/collections/{name}/books/{id}`

#### PUT
* Description: Add a book to a collection
* Return: A no-content response

Output:
```json
{
  "status_code": 201,
  "status": "Success"
}
```

#### DELETE
* Description: Remove a book from a collection
* Return: A no-content response

Output:
```json
{
  "status_code": 201,
  "status": "Success"
}
```