package models

import (
	"github.com/go-gorp/gorp"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        int64  `db:"id" json:"id"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"password"`
	Name      string `db:"name" json:"name"`
	UpdatedAt int64  `db:"updated_at" json:"updatedAt"`
	CreatedAt int64  `db:"created_at" json:"createdAt"`
}

func (u *User) PreInsert(s gorp.SqlExecutor) error {
	u.CreatedAt = time.Now().UnixNano()
	u.UpdatedAt = u.CreatedAt
	pswd, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	u.Password = string(pswd)
	return nil
}

func (u *User) PreUpdate(s gorp.SqlExecutor) error {
	u.UpdatedAt = time.Now().UnixNano()
	pswd, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	u.Password = string(pswd)
	return nil
}
