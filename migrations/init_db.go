package migrations

import (
	"prot/db"
)
import . "prot/models"

func InitDB()  {
	db.MysqlDB.AutoMigrate(&User{})
	db.MysqlDB.AutoMigrate(&Article{})
	db.MysqlDB.AutoMigrate(&Tag{})

	InitTag()
}

func InitTag() {
	tags := []Tag {
		Tag{Code: "wenzhang", Name: "文章"},
		Tag{Code: "shuoshuo", Name: "说说"},
		Tag{Code: "riji", Name: "日记"},
		Tag{Code: "shenghuo", Name: "生活"},
		Tag{Code: "ganwu", Name: "感悟"},
		Tag{Code: "it", Name: "IT"},
		Tag{Code: "go", Name: "Go"},
		Tag{Code: "javascript", Name: "JavaScript"},
		Tag{Code: "c", Name: "C"},
		Tag{Code: "rust", Name: "Rust"},
		Tag{Code: "java", Name: "Java"},
		Tag{Code: "python", Name: "Python"},
		Tag{Code: "typescript", Name: "TypeScript"},
		Tag{Code: "frontend", Name: "前端"},
		Tag{Code: "backend", Name: "后端"},
		Tag{Code: "mysql", Name: "MySql"},
		Tag{Code: "nosql", Name: "NoSql"},
	}

	for _, val := range tags {
		db.MysqlDB.Create(&val)
	}
}
