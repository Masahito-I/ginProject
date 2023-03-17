// models/user.go

package models

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetUser(id string) (*User, error) {
	// get user data from database
	// ...

	// if user == nil {
	// 	return nil, fmt.Errorf("user not found")
	// }

	// return user, nil
	return nil, nil
}

func CreateUser(user User) error {
	// save user data to database
	// ...

	return nil
}
