package models

type Team struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Name  string `gorm:"type:varchar(255);uniqueIndex;not null"`
	Users []User `gorm:"foreignKey:TeamID;constraint:OnDelete:SET NULL"`
}
