package handlers

type User struct {
	ID             string `json:"id"`
	Email          string `json:"email"`
	Verified_email bool   `json:"verified_email"`
	Picture        string `json:"picture"`
}

/*
func save() {
	u := &User{}

	json.Unmarshal(data, u)

	fmt.Println(u.ID)
	fmt.Println(u.Email)
	fmt.Println(u.Verified_email)
	fmt.Println(u.Picture)
	fmt.Println(oauthState)
}
*/
