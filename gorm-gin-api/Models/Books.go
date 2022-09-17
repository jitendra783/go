package Models

import (
	"example/Config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllUser(b *[]User) (err error) {
	if err = Config.DB.Find(b).Error; err != nil {
		return err
	}
	return nil
}

func AddNewUser(b *User) (err error) {
	if err = Config.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

func GetOneUser(b *User, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

func PutOneUser(b *User, id string) (err error) {
	fmt.Println(b)
	Config.DB.Save(b)
	return nil
}

func DeleteUser(b *User, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(b)
	return nil
}
