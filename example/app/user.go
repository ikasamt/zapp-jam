package app

// +jam /clefs/example/is.go
// +jam ValidatePresenceOf Name,CompanyID
// +jam Ngram Name,CompanyID
type User struct {
	ID        int
	Name      string
	CompanyID int
}
