package database

import (
	"fmt"
	"log"
	"os"
	"sayeed1999/social-connect-golang-api/config"
	"sayeed1999/social-connect-golang-api/models"
	"strconv"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	Db *gorm.DB
}

var DB DBInstance

// Connect function
func Connect() {
	dbConfig := config.GetConfig().DATABASE

	// because our config function returns a string, we are parsing our str to int here
	port, err := strconv.ParseUint(dbConfig.PORT, 10, 32)
	if err != nil {
		fmt.Println("Error parsing port str to int")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", dbConfig.HOST, dbConfig.USER, dbConfig.PASSWORD, dbConfig.NAME, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("Connected to database")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations on database")

	if err = db.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Comment{},
	); err != nil {
		log.Fatal(err)
	}

	// Seed data
	seedDatabase(db)

	DB = DBInstance{
		Db: db,
	}
}

func seedDatabase(db *gorm.DB) {
	log.Println("Seeding database")

	users := []models.User{
		{Name: "User I", BaseModel: models.BaseModel{ID: uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d471")}},
		{Name: "User II", BaseModel: models.BaseModel{ID: uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d472")}},
		{Name: "User III", BaseModel: models.BaseModel{ID: uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d473")}},
	}

	for _, user := range users {
		if err := db.FirstOrCreate(&user, models.User{Name: user.Name}).Error; err != nil {
			log.Printf("Failed to seed user %s: %v\n", user.Name, err)
		}
	}

	posts := []models.Post{
		{Body: "Post I", UserID: users[0].ID, BaseModel: models.BaseModel{ID: uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d471")}},
		{Body: "Post II", UserID: users[1].ID, BaseModel: models.BaseModel{ID: uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d472")}},
		{Body: "Post III", UserID: users[2].ID, BaseModel: models.BaseModel{ID: uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d473")}},
		{Body: "Post IV", UserID: users[0].ID, BaseModel: models.BaseModel{ID: uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d474")}},
		{Body: "Post V", UserID: users[1].ID, BaseModel: models.BaseModel{ID: uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d475")}},
		{Body: "Post VI", UserID: users[2].ID, BaseModel: models.BaseModel{ID: uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d476")}},
	}

	for _, post := range posts {
		if err := db.FirstOrCreate(&post, models.Post{Body: post.Body, UserID: post.UserID}).Error; err != nil {
			log.Printf("Failed to seed post %s: %v\n", post.Body, err)
		}
	}
}
