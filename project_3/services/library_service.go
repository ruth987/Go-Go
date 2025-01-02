package services

import (
	"errors"
	"project_3/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int) error
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) ([]models.Book, error)
	AddMember(member models.Member)
	GetMember(memberID int) (models.Member, error)
}

type Library struct {
	Books   map[int]models.Book
	Members map[int]models.Member
}

func NewLibrary() *Library {
	return &Library{
		Books:   make(map[int]models.Book),
		Members: make(map[int]models.Member),
	}
}

func (l *Library) AddBook(book models.Book) {
	l.Books[book.ID] = book
}

func (l *Library) RemoveBook(bookID int) error {
	if _, exists := l.Books[bookID]; !exists {
		return errors.New("book not found")
	}
	delete(l.Books, bookID)
	return nil
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, exists := l.Books[bookID]
	if !exists {
		return errors.New("book not found")
	}
	member, exists := l.Members[memberID]
	if !exists {
		return errors.New("member not found")
	}

	book.Status = "Borrowed"
	l.Books[bookID] = book
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.Members[memberID] = member
	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	book, exists := l.Books[bookID]
	if !exists {
		return errors.New("book not found")
	}
	member, exists := l.Members[memberID]
	if !exists {
		return errors.New("member not found")
	}

	found := false
	newBorrowedBooks := make([]models.Book, 0)
	for _, b := range member.BorrowedBooks {
		if b.ID == bookID {
			found = true
		} else {
			newBorrowedBooks = append(newBorrowedBooks, b)
		}
	}

	if !found {
		return errors.New("book was not borrowed by this member")
	}

	member.BorrowedBooks = newBorrowedBooks
	book.Status = "Available"

	l.Books[bookID] = book
	l.Members[memberID] = member

	return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
	availableBooks := make([]models.Book, 0)
	for _, book := range l.Books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func (l *Library) ListBorrowedBooks(memberID int) ([]models.Book, error) {
	member, exists := l.Members[memberID]
	if !exists {
		return nil, errors.New("member not found")
	}
	return member.BorrowedBooks, nil
}

func (l *Library) GetMember(memberID int) (models.Member, error) {
	member, exists := l.Members[memberID]
	if !exists {
		return models.Member{}, errors.New("member not found")
	}
	return member, nil
}

func (l *Library) AddMember(member models.Member) {
	l.Members[member.ID] = member
}
