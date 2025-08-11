package main

type Book struct {
	ID          string
	Title       string
	Author      string
	Description string
}

type Member struct {
	ID   string
	Name string
	DOB  string
}

type Library struct {
	Books   []Book
	Members []Member
}

func (l *Library) AddBook(title, author, description string) {
	id := generateBookID(GenerateTimestampBasedID())
	book := Book{
		ID:          id,
		Title:       title,
		Author:      author,
		Description: description,
	}
	l.Books = append(l.Books, book)
}

func (l *Library) AddMember(name, dob string) {
	id := generateMemberID(GenerateTimestampBasedID())
	member := Member{
		ID:   id,
		Name: name,
		DOB:  dob,
	}
	l.Members = append(l.Members, member)
}

func (l *Library) GetBooks() []Book {
	return l.Books
}

func (l *Library) GetMembers() []Member {
	return l.Members
}
