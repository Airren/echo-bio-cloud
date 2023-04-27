package dal

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/airren/echo-bio-backend/config"
	"github.com/airren/echo-bio-backend/model"
)

var db *gorm.DB

func InitMySQL() error {
	var (
		err error
	)
	dns := config.Conf.MysqlURI
	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		return err
	}
	//db.SingularTable(true)
	db = db.Debug()
	err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Algorithm{})
	err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.AnalysisJob{})
	err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.AlgoParameter{})
	err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.File{})
	err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.AlgoGroup{})
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
