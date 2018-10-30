package blogs

import (
	"blogs-service/common"

	"github.com/jinzhu/gorm"
)

//dbMigrate db的迁移，没有创建的过的表格会直接创建
func DbMigrate() {
	db := common.DBBegin()
	defer db.DBCommit()

	db.AutoMigrate(UserInfo{})
}

// Model 自己封装的
/*type Model struct {
    ID        uint `gorm:"primary_key"`
    CreatedAt time.Time
    UpdatedAt time.Time
}
*/

// user info
type UserInfo struct {
	gorm.Model
	UserName string `gorm:"type:varchar(30);not null" json:"user_name"`
	Password string `gorm:"type:varchar(30);not null" json:"password"`
	Sex      int64  `json:"sex"`
	TelNum   string `json:"tel_num"`
}
