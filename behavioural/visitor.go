package behavioural

// Define the Visitable interface
type Visitable interface {
	Accept(Visitor)
}

// Define the Visitor interface
type Visitor interface {
	VisitBook(Book)
	VisitMagazine(Magazine)
}

// Define the Book struct
type Book struct {
	Title  string
	Author string
}

// Implement the Accept method for the Book struct
func (b Book) Accept(v Visitor) {
	v.VisitBook(b)
}

// Define the Magazine struct
type Magazine struct {
	Title     string
	Publisher string
}

// Implement the Accept method for the Magazine struct
func (m Magazine) Accept(v Visitor) {
	v.VisitMagazine(m)
}

// Define a concrete implementation of the Visitor interface
type InventoryVisitor struct {
	NumBooks     int
	NumMagazines int
}

// Implement the Visit method for the InventoryVisitor struct for Book struct
func (iv *InventoryVisitor) VisitBook(b Book) {
	iv.NumBooks++
}

// Implement the Visit method for the InventoryVisitor struct for Magazine struct
func (iv *InventoryVisitor) VisitMagazine(m Magazine) {
	iv.NumMagazines++
}
