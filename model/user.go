package model

type User struct {
	Id   int    `gorm:"primary_key;not null;autoIncrement" json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
