package db

import (
	"fmt"

	"github.com/SebastianLozanoValle/gambituser/models"
	"github.com/SebastianLozanoValle/gambituser/tools"
	_ "github.com/go-sql-driver/mysql"
)

func SingUp(sig models.SingUp) error {
	fmt.Println("Comienza Registro")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	sentencia := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('" + sig.UserEmail + "','" + sig.UserUUID + "','" + tools.FechaMySQL() + "')"
	fmt.Println(sentencia)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("SignUp > Ejecución Exitosa")
	return nil
}
