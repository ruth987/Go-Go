package main

import (
    "fmt"
    "os"
    "project_3/controllers"
    "project_3/services"
)

func displayMenu() {
    fmt.Println("\n=== Library Management System ===")
    fmt.Println("1. Add Book")
    fmt.Println("2. Remove Book")
    fmt.Println("3. Add Member")
    fmt.Println("4. Borrow Book")
    fmt.Println("5. Return Book")
    fmt.Println("6. List Available Books")
    fmt.Println("7. List Borrowed Books")
    fmt.Println("8. Exit")
    fmt.Print("Enter your choice: ")
}

func main() {
    library := services.NewLibrary()
    controller := controllers.NewLibraryController(library)

    for {
        displayMenu()
        var choice int
        fmt.Scanf("%d\n", &choice)

        switch choice {
        case 1:
            controller.AddBook()
        case 2:
            controller.RemoveBook()
        case 3:
            controller.AddMember()
        case 4:
            controller.BorrowBook()
        case 5:
            controller.ReturnBook()
        case 6:
            controller.ListAvailableBooks()
        case 7:
            controller.ListBorrowedBooks()
        case 8:
            fmt.Println("Thank you for using the Library Management System!")
            os.Exit(0)
        default:
            fmt.Println("Invalid choice. Please try again.")
        }
    }
}