package database

import (
	"github.com/stretchr/testify/assert"
	"github.com/ylascombe/go-api/models"
	"testing"
)

var (
	email = "email@localhost"
)

func TestInsert(t *testing.T) {

	db := NewDBDriver()
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&models.ApiUser{})

	// Create
	db.Create(&models.ApiUser{Firstname: "Yohan", Lastname: "Test", Email: email})

	// Read
	var user models.ApiUser
	db.First(&user, "Firstname = ?", "Yohan") // find product with FirstName Yohan

	// Update - update product's price to 2000
	db.Model(&user).Update("Lastname", "updated")

	db.Model(&user).Update("SshPublicKey", "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDSFWtiCOxrn9Yupr11111111111111111111111111111111111111111111111111ISDaQEyhywd123v9k7qceH123BIyJvzxq8hUEAvrx1239zoiKMdn4Wu5NrGAxVeOACnFbJ4Vibs1KTUADHXQaPHjDw8czfVvzeaJvct0tJjj8PhsXNoyMWajx+kPyrXxURXkkgwtCI1DJ2222222222222222222222222222222222222223333333333333333333+3333333333444444444555555555555555555555555555556666666666666666666666666888 ylascombe")

	// Delete - delete product
	db.Delete(&user)

}

func TestAutoMigrateDB(t *testing.T) {
	db := NewDBDriver()
	defer db.Close()

	AutoMigrateDB(db)
}

// Force to remove test user
func TestTearDown(t *testing.T) {
	// remove user in order to not change initial state
	db := NewDBDriver()
	defer db.Close()
	//db.Delete(user)
	res := db.Exec("delete from api_users where email = ?", email).Error
	assert.Nil(t, res)
}
