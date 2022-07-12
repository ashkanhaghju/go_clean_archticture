package entity

type User struct {
	ID   uint64 `gorm:"id,primarykey,unique"`
	Name string `gorm:"name"`
}

func (u User) TableName() string {
	return "user"
}
