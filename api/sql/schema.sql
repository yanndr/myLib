CREATE TABLE authors
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    first_name  text                              NOT NULL DEFAULT '',
    last_name   text                              NOT NULL,
    middle_name text                              NOT NULL DEFAULT '',
    UNIQUE (first_name, last_name, middle_name)
);
CREATE INDEX authors_last_name_idx ON authors (last_name);

CREATE TABLE genres
(
    id   INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    name text collate nocase               NOT NULL,
    UNIQUE (name)
);
CREATE INDEX genres_name_idx ON genres (name);

CREATE TABLE books
(
    id             INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    title          text                              NOT NULL,
    published_date INTEGER,
    description    text,
    edition        text
);
CREATE INDEX books_title_id_idx ON books (title);
CREATE INDEX books_published_date_idx ON books (published_date);

CREATE TABLE collections
(
    id   INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    name text                              NOT NULL
);
CREATE INDEX collections_name_idx ON collections (name);

CREATE TABLE books_collections
(
    book_id       INTEGER NOT NULL,
    collection_id INTEGER NOT NULL,
    PRIMARY KEY (book_id, collection_id),
    FOREIGN KEY (book_id) REFERENCES books (id) ON DELETE CASCADE,
    FOREIGN KEY (collection_id) REFERENCES collections (id) ON DELETE CASCADE
);

CREATE TABLE books_genres
(
    book_id  INTEGER NOT NULL,
    genre_id INTEGER NOT NULL,
    PRIMARY KEY (book_id, genre_id),
    FOREIGN KEY (book_id) REFERENCES books (id) ON DELETE CASCADE,
    FOREIGN KEY (genre_id) REFERENCES genres (id) ON DELETE CASCADE
);

CREATE TABLE authors_books
(
    book_id   INTEGER NOT NULL,
    author_id INTEGER NOT NULL,
    PRIMARY KEY (book_id, author_id),
    FOREIGN KEY (book_id) REFERENCES books (id) ON DELETE CASCADE,
    FOREIGN KEY (author_id) REFERENCES authors (id) ON DELETE CASCADE
);