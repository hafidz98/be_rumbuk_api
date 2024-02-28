package domain

type AvailableRoom struct {
	Building Building
	Floor    Floor
	Room     Room
	TimeSlot TimeSlot
	Reserved *bool
}
