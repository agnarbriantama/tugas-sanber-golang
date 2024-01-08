package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

type ApplyJob struct {
    Id_Apply      uint           `json:"id_apply" gorm:"primaryKey"`
    User_Id		  uint   		 `json:"user_id"`
	JobID  		  uint   		 `json:"job_id"`
    Status        string         `json:"status"`
    CreatedAt     time.Time      `json:"created_at" gorm:"-"`
    UpdatedAt     time.Time      `json:"updated_at" gorm:"-"`
    DeletedAt     gorm.DeletedAt `json:"delete_at" gorm:"index;" swaggertype:"string" format:"date-time"`
  
}

// Value implements the driver.Valuer interface
func (a ApplyJob) Value() (driver.Value, error) {
    return json.Marshal(a)
}

// Scan implements the sql.Scanner interface
func (a *ApplyJob) Scan(value interface{}) error {
	data, ok := value.([]byte)
	if !ok {
		return errors.New("failed to unmarshal ApplyJob data")
	}
	return json.Unmarshal(data, &a)
}
