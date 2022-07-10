package signUP_IN

import (
	"database/sql"
	"errors"
	"fmt"
)

func userExist(Email string, db *sql.DB) (int, bool) {
	var id = -1

	result, err := db.Query("SELECT user_id FROM user_info WHERE email=?", Email)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
		return id, false
	}
	for result.Next() {
		err = result.Scan(&id)
	}
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
		return -2, false
	}
	if id == -1 {
		return id, false
	}
	return id, true
}
func SignUp(db *sql.DB) error {
	var Email string
	var Password string

	fmt.Println("email:")
	fmt.Scanln(&Email)
	fmt.Println("password")
	fmt.Scanln(&Password)

	var id = -1

	result, err := db.Query("SELECT user_id FROM user_info WHERE email=?", Email)
	if err != nil {
		panic(err)
	}
	for result.Next() {
		err = result.Scan(&id)
	}
	if id != -1 {
		fmt.Printf("User %s already exists! \nplease signin \n", Email)
		return errors.New("user already exists")
	}
	_, err = db.Query(`INSERT INTO user_info  SET  first_name ='X',last_name='X', email = ?, password =? ,mobile='X',address1='X',address2='X'`, Email, Password)

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return nil

}
func SignIn(db *sql.DB) (int, error) {
	var Email string
	var Password string
	// get email, password and then check in database

	fmt.Println("email:")
	fmt.Scanln(&Email)
	fmt.Println("password")
	fmt.Scanln(&Password)

	var id = -1
	var id2 = -1
	var exist bool
	id, exist = userExist(Email, db)
	if exist == true {
		//check password
		result, err := db.Query("SELECT user_id FROM user_info WHERE password=?", Password)
		if err != nil {
			panic(err)
		}
		for result.Next() {
			err = result.Scan(&id2)
		}
		if err != nil {
			panic(err)
		}
		if id != id2 {
			fmt.Printf("password is incorrect!\ntry again.")
			return -1, nil
		}
		return id, nil
	}

	fmt.Printf("User %s does not exist! please signUp \n", Email)
	return id, errors.New("please signup")

}
