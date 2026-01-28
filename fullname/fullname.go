package fullname

type User struct {
    FirstName string
    LastName  string
}

// FullName returns firstname and lastname of the user
func (u User) FullName() string {
    return u.FirstName + " " + u.LastName
}