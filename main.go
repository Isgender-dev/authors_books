package main

import (
	"log"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"github.com/jackc/pgx/v5"
)

type AuthorsResponse struct {
	AuthorId  int    `json:"author_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BirthYear int    `json:"birth_year"`
}

type AuthorCreateRequest struct {
	AuthorId  int    `json:"author_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BirthYear int    `json:"birth_year"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

type BooksResponse struct {
	BookId          int             `json:"book_id"`
	Title           string          `json:"title"`
	Genre           string          `json:"genre"`
	PublicationYear int             `json:"publication_year"`
	AuthorId        int             `json:"author_id"`
	AuthorsResponse AuthorsResponse `json:"author"`
}

type BookCreateRequest struct {
	BookId          int    `json:"book_id"`
	Title           string `json:"title"`
	Genre           string `json:"genre"`
	PublicationYear int    `json:"publication_year"`
}

func authors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rows, err := db.Query(context.Background(), "SELECT * from authors")
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{err.Error(), "404"})
		return
	}
	defer rows.Close()
	var list []AuthorsResponse
	for rows.Next() {
		var res AuthorsResponse
		err = rows.Scan(&res.AuthorId, &res.FirstName, &res.LastName, &res.BirthYear)
		if err != nil {
			json.NewEncoder(w).Encode(ErrorResponse{err.Error(), "404"})
			return
		}
		list = append(list, res)
	}
	json.NewEncoder(w).Encode(list)
}

func createAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		json.NewEncoder(w).Encode(ErrorResponse{"Only POST method is working", "400"})
		return
	}
	var req AuthorCreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{err.Error(), "404"})
		return
	}
_, err = db.Exec(
	context.Background(),
	"INSERT INTO authors (author_id, first_name, last_name, birth_year) VALUES ($1, $2, $3, $4)",
	req.AuthorId, req.FirstName, req.LastName, req.BirthYear,
)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{err.Error(), "404"})
		return
	}
	json.NewEncoder(w).Encode(true)
}

func updateAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "PUT" {
		json.NewEncoder(w).Encode(ErrorResponse{"Only PUT method is allowed", "400"})
		return
	}
	var req AuthorCreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{err.Error(), "404"})
		return
	}
		_, err = db.Exec(context.Background(), "UPDATE authors SET first_name = $1, last_name = $2, birth_year = $3 WHERE author_id = $4", req.FirstName, req.LastName, req.BirthYear, req.AuthorId)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{err.Error(), "404"})
		return
	}
	json.NewEncoder(w).Encode(true)
}

func deleteAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "DELETE" {
		json.NewEncoder(w).Encode(ErrorResponse{"Only DELETE method is working", "404"})
		return
	}
	var req AuthorCreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{err.Error(), "404"})
		return
	}
		_, err = db.Exec(context.Background(), "DELETE FROM authors WHERE author_id = $1", req.AuthorId)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{err.Error(), "404"})
		return
	}
	json.NewEncoder(w).Encode(true)
}

func books(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rows, err := db.Query(context.Background(), "SELECT b.book_id, b.title, b.genre, b.publication_year, b.author_id, a.author_id, a.first_name, a.last_name, a.birth_year FROM books b JOIN authors a ON b.author_id = a.author_id")
	if err != nil {
http.Error(w, err.Error(), http.StatusInternalServerError)
return
	}
	defer rows.Close()
	var list []BooksResponse
	for rows.Next() {
		var res BooksResponse
		err = rows.Scan(&res.BookId, &res.Title, &res.Genre, &res.PublicationYear, &res.AuthorId, &res.AuthorsResponse.AuthorId, &res.AuthorsResponse.FirstName, &res.AuthorsResponse.LastName, &res.AuthorsResponse.BirthYear)
		if err != nil {
	continue
}
		list = append(list, res)
	}
	json.NewEncoder(w).Encode(list)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		json.NewEncoder(w).Encode(ErrorResponse{"Only POST method is working", "400"})
		return
	}
	var req BookCreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{err.Error(), "404"})
		return
	}
	_, err = db.Exec(context.Background(), "INSERT INTO books (book_id, title, genre, publication_year) VALUES ($1, $2, $3, $4)", req.BookId, req.Title, req.Genre, req.PublicationYear,)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{err.Error(), "404"})
		return
	}
	json.NewEncoder(w).Encode(true)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "PUT" {
		json.NewEncoder(w).Encode(ErrorResponse{"Only PUT method is allowed", "400"})
		return
	}
	var req BookCreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{err.Error(), "404"})
		return
	}
	_, err = db.Exec(context.Background(), "UPDATE books SET title = $1, genre = $2, publication_year = $3 WHERE book_id = $4", req.Title, req.Genre, req.PublicationYear, req.BookId)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{err.Error(), "404"})
		return
	}
	json.NewEncoder(w).Encode(true)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "DELETE" {
		json.NewEncoder(w).Encode(ErrorResponse{"Only DELETE method is working", "404"})
		return
	}
	var req BookCreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{err.Error(), "404"})
		return
	}
_, err = db.Exec(
	context.Background(),
	"DELETE FROM books WHERE book_id = $1",
	req.BookId,
)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponse{err.Error(), "404"})
		return
	}
	json.NewEncoder(w).Encode(true)
}

var db *pgx.Conn

func main() {
	db = connectDB("postgres://user1:12345@localhost:5432/library_db")
	defer db.Close(context.Background())
	http.HandleFunc("/authors", authors)
	http.HandleFunc("/authors/create", createAuthor)
	http.HandleFunc("/authors/update", updateAuthor)
	http.HandleFunc("/authors/delete", deleteAuthor)
	http.HandleFunc("/books", books)
	http.HandleFunc("/books/create", createBook)
	http.HandleFunc("/books/update", updateBook)
	http.HandleFunc("/books/delete", deleteBook)
	//http.ListenAndServe(":8008", nil)
	log.Fatal(http.ListenAndServe("0.0.0.0:8008", nil))
}

func connectDB(config string) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}
