package entity

// id, name, msisdn, email, birth_date, created_time, update_time_user_age
type User struct {
	UserID 		int 		`json:"user_id"`
	FullName 	string 		`json:"full_name"`
	Msisdn 		string 		`json:"msisdn"`
	UserEmail 	string 		`json:"user_email"`
	BirthDate 	string		`json:"birth_date"`
	UserAge		string		`json:"user_age"`
}

