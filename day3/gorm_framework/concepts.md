**GORM Cheat Sheet - Golang ORM Framework**

### 1. **Installation & Setup**
```sh
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite # Or MySQL/PostgreSQL
```

#### **Connect to a Database**
```go
db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
if err != nil {
    log.Fatal("Failed to connect to database", err)
}
```

---

### 2. **Defining Models & Auto Migration**
```go
type User struct {
    ID    uint   `gorm:"primaryKey"`
    Name  string
    Email string `gorm:"unique"`
    Age   int
}

db.AutoMigrate(&User{})
```

---

### 3. **CRUD Operations**
#### **Create**
```go
db.Create(&User{Name: "Alice", Email: "alice@example.com", Age: 25})
```

#### **Read**
```go
var user User
db.First(&user, 1) // Find user with ID 1
db.First(&user, "email = ?", "alice@example.com")
```

#### **Update**
```go
db.Model(&user).Update("Age", 30)
db.Model(&user).Updates(User{Name: "Alice Updated", Age: 28})
```

#### **Delete**
```go
db.Delete(&user)
```

---

### 4. **Relationships**
#### **One-to-One**
```go
type Profile struct {
    ID     uint
    UserID uint
    Bio    string
}

type User struct {
    ID      uint
    Name    string
    Profile Profile `gorm:"foreignKey:UserID"`
}
```

#### **One-to-Many**
```go
type Post struct {
    ID     uint
    UserID uint
    Title  string
}

type User struct {
    ID    uint
    Name  string
    Posts []Post `gorm:"foreignKey:UserID"`
}
```

#### **Many-to-Many**
```go
type Group struct {
    ID    uint
    Name  string
    Users []User `gorm:"many2many:user_groups"`
}
```

---

### 5. **Advanced Queries**
```go
db.Where("age > ?", 20).Find(&users)
db.Order("age desc").Limit(5).Find(&users)
db.Raw("SELECT * FROM users WHERE age > ?", 25).Scan(&users)
db.Exec("UPDATE users SET age = age + 1 WHERE age > ?", 25)
```

---

### 6. **Soft Deletes**
```go
type User struct {
    gorm.Model
    Name string
}
db.Delete(&user) // Soft delete
db.Unscoped().Delete(&user) // Hard delete
```

---

### 7. **Transactions**
```go
err := db.Transaction(func(tx *gorm.DB) error {
    if err := tx.Create(&user).Error; err != nil {
        return err
    }
    return nil
})
```

---

### 8. **Hooks (Lifecycle Callbacks)**
```go
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
    log.Println("Before creating user:", u.Name)
    return nil
}
```

---

### 9. **Performance Optimizations**
```go
db.Create(&users) // Batch Insert
db.Select("name", "email").Find(&users) // Select Specific Columns
type User struct {
    ID    uint
    Name  string `gorm:"index"`
    Email string `gorm:"uniqueIndex"`
}
```

---

### 10. **GORM Configurations**
```go
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
    SkipDefaultTransaction: true,
    PrepareStmt:            true,
    Logger:                 logger.Default.LogMode(logger.Silent),
})
```

---

This cheat sheet covers the major concepts of GORM for quick revision!

