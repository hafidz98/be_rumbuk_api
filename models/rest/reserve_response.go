package rest

import (
	// "strings"
	"time"
	//"github.com/hafidz98/be_rumbuk_api/utils"
)

//type ct utils.CustomTime

type ReserveResponse struct {
	ReserveID  int       `json:"id"`
	BookDate   time.Time `json:"booking_date"`
	StudentID  string    `json:"student_id"`
	Activity   string    `json:"activity"`
	Room       *RoomData `json:"room"`
	Status     string    `json:"status"`
	StatusText string    `json:"status_text,omitempty"`
}

type RoomData struct {
	RoomName     string `json:"room"`
	FloorName    string `json:"floor"`
	BuildingName string `json:"building"`
	Capacity     int    `json:"capacity"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`
}

// type CustomDate time.Time

// // Implement Marshaler and Unmarshaler interface
// func (j *CustomDate) UnmarshalJSON(b []byte) error {
//     s := strings.Trim(string(b), "\"")
//     t, err := time.Parse("2006-01-02", s)
//     if err != nil {
//         return err
//     }
//     *j = CustomDate(t)
//     return nil
// }

// func (j CustomDate) MarshalJSON() ([]byte, error) {
//     return []byte("\"" + time.Time(j).Format("2006-01-02") + "\""), nil
// }

// // Maybe a Format function for printing your date
// func (j CustomDate) Format(s string) string {
//     t := time.Time(j)
//     return t.Format(s)
// }
