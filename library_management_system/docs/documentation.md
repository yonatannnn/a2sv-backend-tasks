
Library Management System

Overview
This project is a simple Library Management System implemented in Go. It allows users to add, remove, borrow, and return books, as well as manage library members.

Project Structure
```
library_management_system/
├── controllers/
│   └── library_controller.go
├── services/
│   └── library_service.go
├── models/
│   └── book.go
│   └── member.go
└── main.go
```

Controllers

LibraryController
- Manages library operations like adding books, adding members, borrowing books, and returning books.
- Methods:
  - `AddBook()`
  - `AddMember()`
  - `RemoveBook()`
  - `BorrowBook()`
  - `ReturnBook()`
  - `DisplayAvailableBooks()`
  - `DisplayBorrowedBooks(memberId int)`

Services

LibraryManager
- Interface for library operations.
- Methods:
  - `AddBook(book Book)`
  - `AddMember(member Member)`
  - `RemoveBook(bookID int)`
  - `BorrowBook(bookID int, memberID int) error`
  - `ReturnBook(bookID int, memberID int) error`
  - `ListAvailableBooks() []Book`
  - `ListBorrowedBooks(memberID int) []Book`

Library
- Implementation of LibraryManager.
- Manages book and member data using maps.

Models

Book
- Struct with fields: `ID`, `Title`, `Author`, `Status`.

Member
- Struct with fields: `ID`, `Name`, `BorrowedBooks`.

Main
- Entry point of the application.
- Provides a menu for users to interact with the system.

Usage
Run the application using the command:

go run main.go

Follow the on-screen instructions to manage books and members.
