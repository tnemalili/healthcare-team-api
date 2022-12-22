package models

import "time"

type OTPRequest struct {
	Message   string `json:"message" example:"Your OTP is: 123456"`
	Recipient string `json:"recipient" example:"+27731234567"`
	Channel   string `json:"channel" example:"SMS|EMAIL"`
}

type HealthStatus struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type UserType struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Role        *Role  `json:"role"`
}

type Role struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type User struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Surname      string    `json:"surname"`
	Email        string    `json:"email"`
	CompanyID    string    `json:"company_id"`
	MobileNumber string    `json:"mobile_number"`
	UserType     *UserType `json:"user_type"`
}

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Dare struct {
	ID       string    `json:"id"`
	Username string    `json:"username" binding:"required"`
	Password string    `json:"password" binding:"required"`
	Created  time.Time `json:"created"`
}

type GenericMessage struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type Credential struct {
	Username  string     `json:"username" binding:"required"`
	Password  string     `json:"password" binding:"required"`
	LastLogin *time.Time `json:"last_login"`
	Modified  *time.Time `json:"modified"`
	Created   *time.Time `json:"created"`
}

func NewDare(id string, username string, otp string) *Dare {

	now := time.Now()
	return &Dare{ID: id, Created: now, Username: username, Password: otp}
}

func NewCred(username string, password string) *Credential {

	now := time.Now()
	return &Credential{
		Username:  username,
		Password:  password,
		Created:   &now,
		Modified:  &now,
		LastLogin: nil,
	}
}
