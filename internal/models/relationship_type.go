package models

type RelationshipType struct {
	ID   int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Type string `json:"type" gorm:"colunm:Type;not null;type:varchar(70)"`
}
