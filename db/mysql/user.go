package mysql

import (
	"time"
)

type User struct {
	Id          int      `xorm:"not null pk autoincr comment('主键') INT(11)"`
	Name       string    `xorm:"not null index default '' comment('手机号') VARCHAR(255)"`
	Avator     string    `xorm:"not null default '' comment('头像') VARCHAR(500)"`
	Status     int        `xorm:"not null default 0 comment('状态') INT(11)"`
	Extend     string      `xorm:"not null default '' comment('扩展') VARCHAR(255)"`
    Score      int        `xorm:"not null default 0 comment('分数') INT(11)"`

	CreatedAt   time.Time `xorm:"not null created comment('创建时间')"`
	UpdatedAt   time.Time `xorm:"not null updated comment('更新时间')"`
	DeletedAt   time.Time `xorm:"deleted comment('删除时间')"`
}

func (this *User) TableName() string {
	return "user"
}