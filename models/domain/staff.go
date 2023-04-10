package domain

//domain representasi dari tabel dalam basis data

type Staff struct {
	ID       int
	StaffID  string
	Name     string
	Role     string
	Email    string
	Password string
}
