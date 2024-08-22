package confiq

import (
    "fmt"
    "os"
    "github.com/joho/godotenv"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func Database() *gorm.DB {
	if err := godotenv.Load(); err != nil {
        panic("Error loading .env file")
    }
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    name := os.Getenv("DB_NAME")
    charset := os.Getenv("DB_CHARSET")
    parseTime := os.Getenv("DB_PARSE_TIME")
    loc := os.Getenv("DB_LOC")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
        user, password, host, port, name, charset, parseTime, loc)

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic(err)
    }
    return db
}
