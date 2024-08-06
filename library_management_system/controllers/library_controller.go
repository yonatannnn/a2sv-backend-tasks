package controllers

import (
	"bufio"
	"fmt"
	"library_management_system/models"
	"library_management_system/services"
	"os"
	"strconv"
	"strings"
)

type LibraryController struct {
	library library_service.LibraryManager
}

func CreateController(library library_service.LibraryManager) *LibraryController {
	return &LibraryController{library: library}
}

func (lc *LibraryController) AddBook() {
	reader := bufio.NewReader(os.Stdin)
	var bookID int
	var title string
	var author string

	var err error

	for {
		fmt.Print("Enter Book ID: ")
		StringBookID, _ := reader.ReadString('\n')
		bookID, err = strconv.Atoi(strings.TrimSpace(StringBookID))
		if err == nil && bookID > 0 {
			break
		}
		fmt.Println("Invalid input. Please enter a valid positive integer for Book ID.")
	}

	for {
		fmt.Print("Enter Book Title: ")
		title, _ = reader.ReadString('\n')
		title = strings.TrimSpace(title)
		if title != "" {
			break
		}
		fmt.Println("Invalid input. Book title cannot be empty. Please enter a valid title.")
	}

	for {
		fmt.Print("Enter Book Author: ")
		author, _ = reader.ReadString('\n')
		author = strings.TrimSpace(author)
		if author != "" {
			break
		}
		fmt.Println("Invalid input. Book author cannot be empty. Please enter a valid author name.")
	}

	book := models.Book{
		ID:     bookID,
		Title:  title,
		Author: author,
		Status: "Available",
	}
	lc.library.AddBook(book)
	fmt.Println("Book added successfully.")
}

func (lc *LibraryController) AddMember() {
	reader := bufio.NewReader(os.Stdin)
	var memberID int
	var name string

	var err error

	for {
		fmt.Print("Enter Member ID: ")
		StringMemberID, _ := reader.ReadString('\n')
		memberID, err = strconv.Atoi(strings.TrimSpace(StringMemberID))
		if err == nil && memberID > 0 {
			break
		}
		fmt.Println("Invalid input. Please enter a valid positive integer for Member ID.")
	}

	for {
		fmt.Print("Enter Member Name: ")
		name, _ = reader.ReadString('\n')
		name = strings.TrimSpace(name)
		if name != "" {
			break
		}
		fmt.Println("Invalid input. Member name cannot be empty. Please enter a valid name.")
	}

	member := models.Member{
		ID:           memberID,
		Name:         name,
		BorrowedBooks: []models.Book{},
	}
	lc.library.AddMember(member)
	fmt.Println("Member added successfully.")
}

func (lc *LibraryController) RemoveBook() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Book ID to remove: ")
	StringBookID, _ := reader.ReadString('\n')
	bookID, err := strconv.Atoi(strings.TrimSpace(StringBookID))
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid positive integer for Book ID.")
		return
	}
	lc.library.RemoveBook(bookID)
	fmt.Println("Book removed successfully.")
}

func (lc *LibraryController) BorrowBook() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Book ID to borrow: ")
	StringBookID, _ := reader.ReadString('\n')
	bookID, err := strconv.Atoi(strings.TrimSpace(StringBookID))
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid positive integer for Book ID.")
		return
	}

	fmt.Print("Enter Member ID: ")
	StringMemberID, _ := reader.ReadString('\n')
	memberID, err := strconv.Atoi(strings.TrimSpace(StringMemberID))
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid positive integer for Member ID.")
		return
	}

	err = lc.library.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Book borrowed successfully.")
	}
}

func (lc *LibraryController) ReturnBook() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Book ID to return: ")
	StringBookID, _ := reader.ReadString('\n')
	bookID, err := strconv.Atoi(strings.TrimSpace(StringBookID))
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid positive integer for Book ID.")
		return
	}

	fmt.Print("Enter Member ID: ")
	StringMemberID, _ := reader.ReadString('\n')
	memberID, err := strconv.Atoi(strings.TrimSpace(StringMemberID))
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid positive integer for Member ID.")
		return
	}

	err = lc.library.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Book returned successfully.")
	}
}

func (lc *LibraryController) DisplayAvailableBooks() {
	books := lc.library.ListAvailableBooks()
	fmt.Println("Available Books:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}

func (lc *LibraryController) DisplayBorrowedBooks() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Member ID to view borrowed books: ")
	StringMemberID, _ := reader.ReadString('\n')
	memberID, err := strconv.Atoi(strings.TrimSpace(StringMemberID))
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid positive integer for Member ID.")
		return
	}
	books := lc.library.ListBorrowedBooks(memberID)
	fmt.Printf("Borrowed Books by Member %d:\n", memberID)
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}
