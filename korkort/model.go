package korkort

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"time"
)

func GetDB() *gorm.DB {
	var err error

	dbFile, _ := Cfg.GetValue("app", "database")
	DB, err := gorm.Open("sqlite3", dbFile)

	if err != nil {
		log.Fatalf("Fail to load db(%s): %v", dbFile, err)
	}

	return DB
}

type Question struct {
	ID          int
	Content     string `gorm:"type:text"`
	Explanation string `gorm:"type:text"`
	Image       string
	OriginalID  int `gorm:"unique"`
	Category    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Choices     []Choice
}

type Choice struct {
	ID         int
	QuestionID int
	Content    string `gorm:"type:text"`
	Status     int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
