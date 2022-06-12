package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"os"
	"time"
)

type User struct {
	gorm.Model
	Name  string `gorm:"not null"`
	Email string `gorm:"not null;unique"`
}

func main() {
	port := "3000"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{Logger: newLogger})
	if err != nil {
		panic(err)
	}
	db.Migrator().DropTable(User{})
	db.AutoMigrate(User{})

	user := User{
		Name:  "Bob",
		Email: "bob@bob.com",
	}

	db.Create(&user)

	var newUser User
	email := db.Where("email = ?", "bob@bob.com")
	err = email.First(&newUser).Error
	if err != nil {
		panic(err)
	}
	fmt.Println("newUser", newUser)

	r := mux.NewRouter()
	fmt.Println("Starting the development server on port" + port)
	http.ListenAndServe(":"+port, r)
}
