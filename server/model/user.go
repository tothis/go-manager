package model

type User struct {
	Model
	Name string `json:"userName"`
	Age  uint8
}

// 手动设置表名称
/*func (u User) TableName() string {
	return "sys_user"
}*/
