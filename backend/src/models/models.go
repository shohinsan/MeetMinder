package models

type User struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_lame"`
	Password  string `json:"password"`
}

type Meeting struct {
	ID        int64  `json:"id"`
	HostID    int64  `json:"host_id"`
	Title     string `json:"title"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}
