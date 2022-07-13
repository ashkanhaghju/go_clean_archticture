package entity

type User struct {
	ID   uint64 `gorm:"id,primarykey,unique" bson:"id,omitempty"`
	Name string `gorm:"name" bson:"name,omitempty"`
}

func (u User) TableName() string {
	return "user"
}
