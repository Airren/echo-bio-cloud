package dal

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/airren/echo-bio-backend/model"
)

var db *gorm.DB

func InitMySQL() error {
	var (
		err error
	)
	dsn := "root:1q2w3e4r%T@tcp(124.223.99.93:3306)/echo_bio?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}
	//db.SingularTable(true)
	db = db.Debug()
	err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Algorithm{})
	err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Job{})
	err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Parameter{})
	if err != nil {
		return err
	}

	//db.Model(&model.Order{}).AddIndex("idx_task_name", "name")
	//db.Model(&model.Order{}).AddIndex("idx_task_name_claimer_operator", "name", "claimer", "operator")
	//db.Model(&model.Order{}).AddIndex("idx_task_priority", "priority")
	//db.Model(&model.Order{}).AddIndex("idx_task_status", "status")
	//db.Model(&model.Order{}).AddIndex("idx_task_group_id", "group_id")
	//db.Model(&model.Order{}).AddIndex("idx_task_created_at", "created_at")

	return nil
}
