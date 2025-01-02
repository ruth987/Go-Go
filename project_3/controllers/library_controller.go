package controllers

import (
	"bufio"
	"fmt"
	"os"
	"project_3/models"
	"project_3/services"
	"strconv"
	"strings"
)

type LibraryController struct {
	library *services.Library
	reader  *bufio.Reader
}

func NewLibraryController(library *services.Library) *LibraryController {
	return &LibraryController{
		library: library,
		reader:  bufio.NewReader((os.Stdin)),
	}
}

func (lc *LibraryController) getInput(prompt string) string {
	fmt.Print(prompt)
	input, _ := lc.reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func (lc *LibraryController) AddBook() {
	fmt.Println("\n=== Add New Book ===")

	id, _ := strconv.Atoi(lc.getInput("Enter Book ID: "))
	title := lc.getInput("Enter Title: ")
	author := lc.getInput(("Enter Author: "))

	book := models.NewBook(id, title, author)
	lc.library.AddBook(book)
	fmt.Println("Book added successfully!")
}

func (lc *LibraryController) RemoveBook() {
	fmt.Println("\n=== Remove Book ===")

	id, _ := strconv.Atoi(lc.getInput("Enter Book ID to remove: "))

	err := lc.library.RemoveBook(id)
	if err != nil {
		fmt.Printf("Error removing book: %v\n", err)
		return
	}
	fmt.Println("Book removed successfully!")
}

func (lc *LibraryController) AddMember() {
	fmt.Println("\n=== Add New Member ===")

	id, _ := strconv.Atoi(lc.getInput("Enter Member ID: "))
	name := lc.getInput("Enter Member Name: ")

	member := models.NewMember(id, name)
	lc.library.AddMember(member)
	fmt.Println("Member added successfully!")
}

func (lc *LibraryController) BorrowBook() {
	fmt.Println("\n=== Borrow Book ===")

	memberID, _ := strconv.Atoi(lc.getInput("Enter Member ID: "))
	bookID, _ := strconv.Atoi(lc.getInput("Enter Book ID to borrow: "))

	err := lc.library.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Printf("Error borrowing book: %v\n", err)
		return
	}
	fmt.Println("Book borrowed successfully!")
}
func (lc *LibraryController) ReturnBook() {
	fmt.Println("\n=== Return Book ===")

	memberID, _ := strconv.Atoi(lc.getInput("Enter Member ID: "))
	bookID, _ := strconv.Atoi(lc.getInput("Enter Book ID: "))

	err := lc.library.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("Book returned successfully!")
}

func (lc *LibraryController) ListAvailableBooks() {
	fmt.Println("\n=== Available Books ===")

	books := lc.library.ListAvailableBooks()
	if len(books) == 0 {
		fmt.Println("No books available.")
		return
	}

	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}

func (lc *LibraryController) ListBorrowedBooks() {
	fmt.Println("\n=== Borrowed Books ===")

	memberID, _ := strconv.Atoi(lc.getInput("Enter Member ID: "))

	books, err := lc.library.ListBorrowedBooks(memberID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(books) == 0 {
		fmt.Println("No books borrowed by this member.")
		return
	}

	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}
