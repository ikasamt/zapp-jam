package app

// +jam /clefs/example/is.go
// +jam /clefs/example/some.go User Item
// +jam ValidatePresenceOf Name
// +jam Setter Name
type Company struct {
	ID   int
	Name string
}

// +jam /clefs/example/is.go
// +jam /clefs/example/diff.go
// +jam ValidatePresenceOf Name,CompanyID
type Item struct {
	ID        int
	Name      string
	CompanyID int
}
