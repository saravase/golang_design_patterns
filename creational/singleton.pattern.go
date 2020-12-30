// Creation Pattern - Singleton

/**
1. When you want to use the same connection to a database to make every query
2. When you open a Secure Shell (SSH) connection to a server to do a few tasks,
and don't want to reopen the connection for each task
3. If you need to limit the access to some variable or space, you use a Singleton as
the door to this variable (we'll see in the following chapters that this is more
achievable in Go using channels anyway)
4. If you need to limit the number of calls to some places, you create a Singleton
instance to make the calls in the accepted window
*/

package creational

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type IUser interface {
	GetUser(id int64) (*User, error)
}

type manager struct {
	db *gorm.DB
}

func NewManager() *manager {
	dsn := "host=localhost user=primz password=primz@2207 dbname=primz port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("DB connection creation failed. Reason : %v\n", err)
	}

	return &manager{
		db: db,
	}
}

func (m *manager) GetUser(id int64) (*User, error) {
	user := new(User)
	r := m.db.Where("id = ?", id).Find(user)
	if err := r.Error; err != nil {
		log.Fatalf("User not found. Reason : %v\n", err)
		return nil, err
	}
	return user, nil
}
