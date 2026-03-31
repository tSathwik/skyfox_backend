package user

type User struct {
	ID        string      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	Age       int       `gorm:"not null" json:"age"`
}


func (User) TableName() string {
	return "users"
}