package models

type Hospital struct {
	ID   int    `gorm:"primaryKey;column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

func (Hospital) TableName() string {
	return "hospitals"
}
