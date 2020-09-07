package hadlers

import (
	"encoding/json"
	"net/http"
	"restapi/application/usecase"
	"restapi/domain/books"
)

//Get all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	books, _ := usecase.GetAllBooks()
	json.NewEncoder(w).Encode(books)
}

//Get book find id
//func GetBook(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	params := mux.Vars(r) //Get params
//
//	//Loop through books and infd with id
//	for _, item := range repositories.Books {
//		if item.ID.String() == params["id"] {
//			json.NewEncoder(w).Encode(item)
//			return
//		}
//	}
//
//	json.NewEncoder(w).Encode(&books.Book{})
//}

//Create new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bookRequest books.Book
	_ = json.NewDecoder(r.Body).Decode(&bookRequest)
	book, _ := bookRequest.NewBook()
	result, _ := usecase.CreateNewBook(book)

	json.NewEncoder(w).Encode(result)
}

//Update book
//func UpdateBook(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	params := mux.Vars(r) //Get params
//
//	for index, item := range repositories.Books {
//		if item.ID.String() == params["id"] {
//			repositories.Books = append(repositories.Books[:index], repositories.Books[index+1:]...)
//			var book books.Book
//			_ = json.NewDecoder(r.Body).Decode(&book)
//			book.ID, _ = uuid.NewUUID() // Mock uuid -not safe
//			repositories.Books = append(repositories.Books, book)
//			json.NewEncoder(w).Encode(book)
//			return
//		}
//	}
//	json.NewEncoder(w).Encode(repositories.Books)
//}

//Delete book
//func DeleteBook(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	params := mux.Vars(r) //Get params
//
//	for index, item := range repositories.Books {
//		if item.ID.String() == params["id"] {
//			repositories.Books = append(repositories.Books[:index], repositories.Books[index+1:]...)
//			break
//		}
//	}
//	json.NewEncoder(w).Encode(repositories.Books)
//}
