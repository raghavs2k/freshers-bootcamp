package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User model
type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"unique"`
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"` // Soft delete
	Profile   Profile        `gorm:"foreignKey:UserID"`
	Posts     []Post         `gorm:"foreignKey:UserID"`
	Groups    []Group        `gorm:"many2many:user_groups"`
}

// Profile model (One-to-One relationship)
type Profile struct {
	ID     uint
	UserID uint
	Bio    string
}

// Post model (One-to-Many relationship)
type Post struct {
	ID     uint
	UserID uint
	Title  string
	Body   string
}

// Group model (Many-to-Many relationship)
type Group struct {
	ID    uint
	Name  string
	Users []User `gorm:"many2many:user_groups"`
}

func main() {
	// Connect to database (SQLite for simplicity)
	db, err := gorm.Open(sqlite.Open("gorm_demo.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	fmt.Println("✅ Connected to database!")

	// Drop tables if they exist
	fmt.Println("🔄 Dropping existing tables...")
	db.Migrator().DropTable(&User{}, &Profile{}, &Post{}, &Group{})

	// Auto Migrate models
	db.AutoMigrate(&User{}, &Profile{}, &Post{}, &Group{})
	fmt.Println("✅ Database schema migrated!")

	// Insert sample data
	user := User{
		Name:  "Alice",
		Email: "alice@example.com",
		Age:   25,
		Profile: Profile{
			Bio: "Software Engineer",
		},
		Posts: []Post{
			{Title: "First Post", Body: "This is my first post."},
			{Title: "Second Post", Body: "Learning GORM!"},
		},
		Groups: []Group{
			{Name: "Developers"},
			{Name: "Gophers"},
		},
	}

	db.Create(&user) // Insert user, profile, posts, and groups
	fmt.Println("✅ User, profile, posts, and groups inserted!")

	// Fetch user with relationships using Preload
	var fetchedUser User
	db.Preload("Profile").Preload("Posts").Preload("Groups").First(&fetchedUser, "email = ?", "alice@example.com")

	fmt.Println("\n🔹 Fetched User Details:")
	fmt.Printf("ID: %d, Name: %s, Email: %s, Age: %d\n", fetchedUser.ID, fetchedUser.Name, fetchedUser.Email, fetchedUser.Age)
	fmt.Printf("Profile Bio: %s\n", fetchedUser.Profile.Bio)
	fmt.Println("Posts:")
	for _, post := range fetchedUser.Posts {
		fmt.Printf(" - %s: %s\n", post.Title, post.Body)
	}
	fmt.Println("Groups:")
	for _, group := range fetchedUser.Groups {
		fmt.Printf(" - %s\n", group.Name)
	}

	// Update user data
	db.Model(&fetchedUser).Update("Age", 30)
	fmt.Println("\n✅ User age updated to 30!")

	// Soft delete user
	db.Delete(&fetchedUser)
	fmt.Println("✅ User soft deleted!")

	// Retrieve all users excluding soft deleted ones
	var activeUsers []User
	db.Find(&activeUsers)
	fmt.Println("\n🔹 Active Users (excluding deleted):")
	for _, u := range activeUsers {
		fmt.Printf("ID: %d, Name: %s\n", u.ID, u.Name)
	}

	// Retrieve users including soft deleted ones
	var allUsers []User
	db.Unscoped().Find(&allUsers)
	fmt.Println("\n🔹 All Users (including deleted):")
	for _, u := range allUsers {
		fmt.Printf("ID: %d, Name: %s (Deleted: %v)\n", u.ID, u.Name, u.DeletedAt.Valid)
	}

	// Transaction Example
	fmt.Println("\n🔹 Starting a transaction to insert Bob...")
	err = db.Transaction(func(tx *gorm.DB) error {
		var existingUser User
		if err := tx.Where("email = ?", "bob@example.com").First(&existingUser).Error; err == nil {
			fmt.Println("⚠️ Bob already exists, skipping insertion.")
			return nil
		}
		return tx.Create(&User{Name: "Bob", Email: "bob@example.com"}).Error
	})
	if err != nil {
		log.Println("Transaction failed:", err)
	} else {
		fmt.Println("✅ Transaction successful! Bob inserted.")
	}
}
