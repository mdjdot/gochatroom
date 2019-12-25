package common

// User .
type User struct {
	UserID   int    `json:"userID,omitempty"`
	UserPWD  string `json:"userPWD,omitempty"`
	UserName string `json:"userName,omitempty"`
}
