Basic Library System in Go

A simple command-line library management system built in Golang.

This project demonstrates:
  Structs and methods in Go
  Working with slices to store data
  Using helper functions for ID generation
  Organizing code into multiple .go files

Features
  Add books with title, author, and description
  Add members with name and date of birth
  Generate unique IDs using timestamp-based helpers
  View all books and members in the library


Project Structure

basic_library_system/
│── main.go        # Entry point
│── library.go     # Library struct, Book & Member structs, and related methods
│── utils.go       # Utility functions for ID generation


How to Run
# Clone repository
git clone https://github.com/yourusername/basic_library_system.git
cd basic_library_system

# Run the program
go run .
