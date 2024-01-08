package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
    Id_User  uint   `json:"id_user" gorm:"primaryKey"`
    Email    string `json:"email" form:"email"`
    Password string `json:"password" form:"password"`
    Username string `json:"username" form:"username"`
    Role     string `json:"role" form:"role"`
    ApplyJobs []ApplyJob     `json:"applyjobs" gorm:"foreignKey:User_Id"`
    CreatedAt    time.Time      `json:"created_at" gorm:"-"`
    UpdatedAt    time.Time      `json:"updated_at" gorm:"-"`
    DeletedAt gorm.DeletedAt `json:"delete_at" gorm:"index;" swaggertype:"string" format:"date-time"`

}
