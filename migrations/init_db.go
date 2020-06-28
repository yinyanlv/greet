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
		Tag{ID: "wenzhang", Name: "文章"},
		Tag{ID: "shuoshuo", Name: "说说"},
		Tag{ID: "riji", Name: "日记"},
		Tag{ID: "shenghuo", Name: "生活"},
		Tag{ID: "ganwu", Name: "感悟"},
		Tag{ID: "it", Name: "IT"},
		Tag{ID: "go", Name: "Go"},
		Tag{ID: "javascript", Name: "JavaScript"},
		Tag{ID: "c", Name: "C"},
		Tag{ID: "rust", Name: "Rust"},
		Tag{ID: "java", Name: "Java"},
		Tag{ID: "python", Name: "Python"},
		Tag{ID: "typescript", Name: "TypeScript"},
		Tag{ID: "frontend", Name: "前端"},
		Tag{ID: "backend", Name: "后端"},
		Tag{ID: "mysql", Name: "MySql"},
		Tag{ID: "nosql", Name: "NoSql"},
	}

	for _, val := range tags {
		db.MysqlDB.Create(&val)
	}
}
