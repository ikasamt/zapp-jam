package app

// +jam ../clefs/is.go
// +jam ../clefs/some.go User Item
type Company struct {
	ID   int
	Name string
}

// +jam ../clefs/is.go
// +jam ../clefs/diff.go
type Item struct {
	ID        int
	Name      string
	CompanyID int
}
