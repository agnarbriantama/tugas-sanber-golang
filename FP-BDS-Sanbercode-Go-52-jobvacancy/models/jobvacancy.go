package models

import (
	"time"

	"gorm.io/gorm"
)

type Jobvacancy struct {
	Id_Jobvacancy uint           `json:"id_jobvacancy" gorm:"primaryKey"`
	Title         string         `form:"title"`
	CompanyName   string         `form:"company_name"`
	CompanyDesc   string         `form:"company_desc"`
	CompanySalary int         	 `form:"company_salary"`
	CompanyStatus int        	 `form:"company_status"`
	ApplyJob      []ApplyJob     `json:"applyjob" gorm:"foreignKey:JobID"`
	CreatedAt     time.Time      `json:"created_at" gorm:"-"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"-"`
	DeletedAt     gorm.DeletedAt `json:"delete_at" gorm:"index;" swaggertype:"string" format:"date-time"`
}

