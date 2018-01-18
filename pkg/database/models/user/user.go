package user

import (
	"github.com/george-e-shaw-iv/beta/pkg/database"
	"strconv"
	"github.com/george-e-shaw-iv/beta/pkg/encryption"
	"errors"
)

type User struct {
	Roll int
	FirstName string
	MiddleName string
	LastName string
	Suffix string
	Positions string
	password string
	Secret string
}

func init() {
	db, err := database.Open(database.DB_MAIN)
	if err != nil {
		panic("Error opening database for initial checks")
	}
	defer db.Close()

	if count, _ := db.Count(database.BUCKET_USERS); count == 0 {
		p, err := encryption.HashPassword("password")
		if err != nil {
			panic("Error inserting default user")
		}

		err = db.Put(database.BUCKET_USERS, []byte("1"), User{
			Roll: 1,
			FirstName: "John",
			MiddleName: "Riley",
			LastName: "Knox",
			Suffix: "",
			Positions: "admin,member",
			password: p,
			Secret: "",
		})

		if err != nil {
			panic("Error inserting default user")
		}
	}
}

func Fetch(roll int) (*User, error) {
	var u User

	db, err := database.Open(database.DB_MAIN)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.Get(database.BUCKET_USERS, []byte(strconv.Itoa(roll)), &u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func New(u User) error {
	db, err := database.Open(database.DB_MAIN)
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Put(database.BUCKET_USERS, []byte(strconv.Itoa(u.Roll)), u)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) Authenticate(password string) error {
	err := encryption.CheckPassword([]byte(u.password), []byte(password))
	if err != nil {
		return errors.New("unable to authenticate user")
	}

	return u.setSecret(encryption.RandomString(16))
}

func (u *User) setSecret(secret string) error {
	db, err := database.Open(database.DB_MAIN)
	if err != nil {
		return err
	}
	defer db.Close()

	u.Secret = secret
	err = db.Put(database.BUCKET_USERS, []byte(strconv.Itoa(u.Roll)), u)
	if err != nil {
		return err
	}

	return nil
}
