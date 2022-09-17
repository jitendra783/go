package Models


type User struct {

	UserId     int `json:"userid"`
	UserName   string `json:"username"`
	Age int `json:"age"`
    Dob string `json:"dob"`
		
}

func (b *User) TableName() string {
	return "user"
}
