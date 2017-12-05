package korkort

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
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
	gorm.Model
	OriginalID  int    `gorm:"unique"`
	Content     string `gorm:"type:text"`
	image       string
	Explanation string `gorm:"type:text"`
}

type Choice struct {
	gorm.Model
	QuestionOriginalID int
	Content            string `gorm:"type:text"`
	Status             bool
}
