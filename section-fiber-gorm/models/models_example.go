package models

import (
	"gorm.io/gorm"
)

type Registration struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,username_validate"`
	Password string `json:"password" validate:"required,min=6,max=20"`
	LineId   string `json:"lineId" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Business string `json:"business" validate:"required"`
	WebName  string `json:"webName" validate:"required,min=2,max=30,web_validate"`
}

type Dogs struct {
	gorm.Model
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
}

type DogsRes struct {
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
	Type  string `json:"type"`
}

type ResultData struct {
	Data        []DogsRes `json:"data"`
	Name        string    `json:"name"`
	Count       int       `json:"count"`
	Sum_red     int       `json:"sum_red"`
	Sum_green   int       `json:"sum_green"`
	Sum_pink    int       `json:"sum_pink"`
	Sum_noColor int       `json:"sum_nocolor"`
}

type Companies struct {
	gorm.Model
	Name           string  `json:"name"`
	Address        string  `json:"address"`
	RegistrationAt string  `json:"registrationAt"`
	Value          float64 `json:"value"`
	Type           string  `json:"type"`
}

type UserProfiles struct {
	gorm.Model
	EmployeeID string `gorm:"unique" json:"employee_id"`
	Name       string `json:"name" validate:"required"`
	LastName   string `json:"lastname" validate:"required"`
	Birthday   string `json:"birthday" validate:"required"`
	Age        int    `json:"age" validate:"required,min=18"`
	Email      string `json:"email" validate:"required,email"`
	Tel        string `json:"tel" validate:"required"`
}

type UserProfileAgesResult struct {
	Data          []UserProfileResult `json:"data"`
	Name          string              `json:"name"`
	Count         int                 `json:"count"`
	SumGenZ       int                 `json:"sum_genz"`
	SumGenY       int                 `json:"sum_geny"`
	SumGenX       int                 `json:"sum_genx"`
	SumBabyBoomer int                 `json:"sum_babyboomer"`
	SumGI         int                 `json:"sum_gi"`
}

type UserProfileResult struct {
	EmployeeID string `json:"employee_id"`
	Name       string `json:"name"`
	LastName   string `json:"lastname"`
	Birthday   string `json:"birthday"`
	Age        int    `json:"age"`
	Email      string `json:"email"`
	Tel        string `json:"tel"`
	Gen        string `json:"gen"`
}
