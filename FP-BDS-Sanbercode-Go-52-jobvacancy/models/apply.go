package models

type Apply struct {
	IDApply       int    `json:"id_apply"`
	ID            int    `json:"id"`
	IDUser        int    `json:"id_user"`
	JobID         int    `json:"job_id"`
	JobTitle      string `json:"job_title"`
	CompanyName   string `json:"company_name"`
	CompanyDesc   string `json:"company_desc"`
	CompanySalary int    `json:"company_salary"`
	CompanyStatus string `json:"company_status"`
	UserID        int    `json:"user_id"`
	Username      string `json:"username"`
	Status        string `json:"status_lamaran"`
}
