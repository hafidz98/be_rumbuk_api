package domain

//domain representasi dari tabel dalam basis data

type Staff struct {
	ID       int
	StaffID  string
	Name     string
	Role     string
	Status   string
	Email    string
	Password string
}

type StaffFilter struct {
	StaffID string
}
