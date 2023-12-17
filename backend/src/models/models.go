package models

import "time"

type User struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_lame"`
	Password  string `json:"password"`
}

type Meeting struct {
	ID        int64     `json:"id"`
	HostID    int64     `json:"host_id"`
	CreatorID int64     `json:"creator_id"`
	Title     string    `json:"title"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}
