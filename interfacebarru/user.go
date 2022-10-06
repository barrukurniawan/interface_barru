package service

type User struct {
	Nama string
}

type UserService struct {
	db []*User
}

type UserInterface interface {
	Register(u *User) string
	GetUser() []*User
}

func NewUserService(db []*User) UserInterface {
	//
	return &UserService{
		db: db,
	}
}

func (u *UserService) Register(user *User) string {
	u.db = append(u.db, user)
	return user.Nama + " success register user"
}

func (u *UserService) GetUser() []*User {
	return u.db
}
