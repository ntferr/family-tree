package models

type Person struct {
	ID   string `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"column:Name;not null;type:varchar(70)"`
}
