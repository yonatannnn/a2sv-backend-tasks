package main

import (
	"fmt"
	"library_management_system/controllers"
	"library_management_system/services"
)

func main() {
	library := library_service.CreateLibrary()
	controller := controllers.CreateController(library)

	for {
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. Add Member")
		fmt.Println("3. Remove Book")
		fmt.Println("4. Borrow Book")
		fmt.Println("5. Return Book")
		fmt.Println("6. Display Available Books")
		fmt.Println("7. Display Borrowed Books")
		fmt.Println("8. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			controller.AddBook()
		case 2:
			controller.AddMember()
		case 3:
			controller.RemoveBook()
		case 4:
			controller.BorrowBook()
		case 5:
			controller.ReturnBook()
		case 6:
			controller.DisplayAvailableBooks()
		case 7:
			controller.DisplayBorrowedBooks()
		case 8:
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
