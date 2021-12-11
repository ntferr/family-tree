package models

type Relationship struct {
	ID        string `json:"id" gorm:"primaryKey"`
	PersonAID string `json:"person_a_id" gorm:"colunm:PersonAID;not null;type:varchar(70)"`
	PersonBID string `json:"person_b_id" gorm:"colunm:PersonBID;not null;type:varchar(70)"`
	TypeID    int    `json:"type_id"`
}
