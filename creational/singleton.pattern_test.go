package creational

import (
	"log"
	"testing"
)

func TestSingltonPattern(t *testing.T) {
	var id int64 = 1
	m := NewManager()
	user, err := m.GetUser(id)
	log.Printf("User : %v\n", user)
	if err != nil {
		t.Error(err.Error())
	}
	if user.Name != "primz" {
		t.Error("Wrong user name")
	}
}
