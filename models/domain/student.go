package domain

//domain model representasi dari tabel dalam basis data

type Student struct {
	ID          int
	StudentID   string
	Name        string
	Gender      string
	BatchYear   int
	Major       string
	Faculty     string
	PhoneNumber string
	Email       string
	Password    string
}
