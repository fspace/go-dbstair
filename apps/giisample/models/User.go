package models

import (
	//"github.com/go-ozzo/ozzo-validation"
	"database/sql"
	"gopkg.in/guregu/null.v3"
	"time"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

type User struct {
	//   Id          int    `json:"id"` // int32

	Id                 int         `json:"id"  `                   //
	Username           string      `json:"username"  `             //
	Email              string      `json:"email"  `                //
	PasswordHash       string      `json:"password_hash"  `        //
	AuthKey            string      `json:"auth_key"  `             //
	ConfirmedAt        null.Int    `json:"confirmed_at"  `         //
	UnconfirmedEmail   null.String `json:"unconfirmed_email"  `    //
	BlockedAt          null.Int    `json:"blocked_at"  `           //
	RegistrationIp     null.String `json:"registration_ip"  `      //
	Flags              int         `json:"flags"  `                //
	Status             int         `json:"status"  `               //
	PasswordResetToken null.String `json:"password_reset_token"  ` //
	CreatedAt          int         `json:"created_at"  `           //
	UpdatedAt          int         `json:"updated_at"  `           //

	// 依赖 Repo  可以用来做一些唯一性检测的验证约束
	// repo UserRepo `json:"-" form:",omitempty"` //
}

// TableName sets the insert table name for this struct type
func (model *User) TableName() string {
	return "user"
}
