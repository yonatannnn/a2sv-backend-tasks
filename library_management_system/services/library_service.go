package library_service

import (
	"errors"
	"fmt"
	. "library_management_system/models"
)

type library struct {
	books   map[int]Book
	members map[int]Member
}

type LibraryManager interface {
	AddBook(book Book)
	AddMember(member Member)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []Book
	ListBorrowedBooks(memberID int) []Book
}

func CreateLibrary() LibraryManager {
	return &library{
		books:   make(map[int]Book),
		members: make(map[int]Member),
	}
}

func (l *library) AddBook(book Book) {
	_, ok := l.books[book.ID]
	if ok {
		fmt.Println("Book already exists")
	} else {
		l.books[book.ID] = book
	}
}

func (l *library) AddMember(member Member) {
	_, ok := l.members[member.ID]
	if ok {
		fmt.Println("Member already exists")
	} else {
		l.members[member.ID] = member
	}
}

func (l *library) RemoveBook(bookID int) {
	_, ok := l.books[bookID]
	if ok {
		delete(l.books, bookID)
	} else {
		fmt.Println("Book does not exist")
	}
}

func (l *library) BorrowBook(bookID int, memberID int) error {
	book, ok := l.books[bookID]
	if !ok {
		return errors.New("book not found")
	}

	if book.Status == "Borrowed" {
		return errors.New("book already borrowed")
	}

	member, ok := l.members[memberID]
	if !ok {
		return errors.New("member not found")
	}

	book.Status = "Borrowed"
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.books[bookID] = book
	l.members[memberID] = member
	return nil
}

func (l *library) ReturnBook(bookID int, memberID int) error {
	book, ok := l.books[bookID]
	if !ok {
		return errors.New("book not found")
	}

	if book.Status != "Borrowed" {
		return errors.New("the book is not borrowed before!")
	}

	member, ok := l.members[memberID]
	if !ok {
		return errors.New("member not found!")
	}

	for i, borrowedBook := range member.BorrowedBooks {
		if borrowedBook.ID == book.ID {
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			book.Status = "Available"
			l.books[bookID] = book 
			l.members[memberID] = member 
			return nil
		}
	}

	return errors.New("this member didn't borrow this book")
}

func (l *library) ListAvailableBooks() []Book {
	var availableBooks []Book
	for _, book := range l.books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func (l *library) ListBorrowedBooks(memberID int) []Book {
	var borrowedBooks []Book
	member, ok := l.members[memberID]
	if !ok {
		fmt.Println("Member not found")
		return borrowedBooks
	} else {
		borrowedBooks = member.BorrowedBooks
	}
	return borrowedBooks
}
