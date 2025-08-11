package main

import "fmt"

func main() {
	fmt.Println("Library System Initialized...")

	lib := Library{}

	lib.AddBook("The Great Gatsby", "F. Scott Fitzgerald", "A novel set in the 1920s about the American dream.")
	lib.AddBook("Clean Code", "Robert C. Martin", "A handbook of agile software craftsmanship.")

	lib.AddMember("Kanishka", "1999-05-21")
	lib.AddMember("John Doe", "1985-02-15")

	fmt.Println("\nBooks in Library:")
	for _, b := range lib.GetBooks() {
		fmt.Printf("%s - %s by %s\n", b.ID, b.Title, b.Author)
	}

	fmt.Println("\nMembers in Library:")
	for _, m := range lib.GetMembers() {
		fmt.Printf("%s - %s (DOB: %s)\n", m.ID, m.Name, m.DOB)
	}
}
