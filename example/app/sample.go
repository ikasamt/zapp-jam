package app

// +jam ../clefs/is.go
// +jam ../clefs/some.go User Item
// +jam ValidatePresenceOf Name
// +jam Setter Name
type Company struct {
	ID   int
	Name string
}

// +jam ../clefs/is.go
// +jam ../clefs/diff.go
// +jam ValidatePresenceOf Name,CompanyID
type Item struct {
	ID        int
	Name      string
	CompanyID int
}
