package user

import (
	"errors"
	"strconv"

	"github.com/george-e-shaw-iv/beta/pkg/database"
	"github.com/george-e-shaw-iv/beta/pkg/encryption"
)

type User struct {
	Roll       int
	FirstName  string
	MiddleName string
	LastName   string
	Suffix     string
	Positions  string
	password   []byte
	Secret     string
}

func init() {
	db, err := database.Open(database.DB_MAIN)
	if err != nil {
		panic("Error opening database for initial checks")
	}
	defer db.Close()

	if count, _ := db.Count(database.BUCKET_USERS); count == 0 {
		initUser := User{
			Roll:       1,
			FirstName:  "John",
			MiddleName: "Riley",
			LastName:   "Knox",
			Suffix:     "",
			Positions:  "admin,member",
			Secret:     "",
		}

		initUser.password, err = encryption.HashPassword([]byte("root"))
		if err != nil {
			panic("Error inserting default user")
		}

		err = db.Put(database.BUCKET_USERS, []byte("1"), initUser)

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

func (u *User) Authenticate(password []byte) error {
	err := encryption.CheckPassword(u.password, password)
	if err != nil {
		return errors.New(string(u.password))
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
