package db

const (
	DBNAME = "contacts-menager"
	DBURI  = "mongodb://localhost:27017"
)

type Store struct {
	User UserDo
}
