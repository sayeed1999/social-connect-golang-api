package database

import (
	"fmt"
	"log"
	"sayeed1999/social-connect-golang-api/config"
	"sayeed1999/social-connect-golang-api/models"
	"strconv"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database interface defines the methods that our database implementation must have.
// This allows us to mock the database for testing.
type Database interface {
	Instance() *gorm.DB
	Connect() error
	Close() error
}

type database struct {
	db *gorm.DB
}

// NewDatabase creates a new instance of a Database.
func NewDatabase() Database {
	return &database{}
}

// Instance returns the gorm.DB instance.
func (d *database) Instance() *gorm.DB {
	return d.db
}

// Connect initializes the database connection.
func (d *database) Connect() error {
	dbConfig := config.GetConfig().DATABASE

	// Parse port from string to uint
	port, err := strconv.ParseUint(dbConfig.PORT, 10, 32)
	if err != nil {
		return fmt.Errorf("error parsing port: %w", err)
	}

	// Connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		dbConfig.HOST, dbConfig.USER, dbConfig.PASSWORD, dbConfig.NAME, port)

	// Open connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// Log connection success
	log.Println("Connected to database")

	d.db = db
	log.Println("Running migrations on database")

	// Run migrations
	if err := d.db.AutoMigrate(
		// Add your models here
		&models.User{},
		&models.Post{},
		&models.Comment{},
	); err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	// Seed data
	if err := seedDatabase(d.db); err != nil {
		return fmt.Errorf("database seeding failed: %w", err)
	}

	return nil
}

// Close cleans up database resources.
func (d *database) Close() error {
	sqlDB, err := d.db.DB()
	if err != nil {
		return fmt.Errorf("failed to retrieve underlying SQL DB: %w", err)
	}
	return sqlDB.Close()
}

// Seed initial data into the database.
func seedDatabase(db *gorm.DB) error {
	log.Println("Seeding database")

	// Add seeding logic here. Example:
	users := []models.User{
		{Name: "User I", BaseModel: models.BaseModel{ID: uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d471")}},
		{Name: "User II", BaseModel: models.BaseModel{ID: uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d472")}},
		{Name: "User III", BaseModel: models.BaseModel{ID: uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d473")}},
	}

	for _, user := range users {
		if err := db.FirstOrCreate(&user, models.User{Name: user.Name}).Error; err != nil {
			log.Printf("Failed to seed user %s: %v\n", user.Name, err)
			return err
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
			return err
		}
	}

	return nil
}
