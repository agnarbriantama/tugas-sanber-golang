package models

type Jobvacancy struct {
	ID            int    `json:"id"`
	Title         string `form:"title"`
	CompanyName   string `form:"company_name"`
	CompanyDesc   string `form:"company_desc"`
	CompanySalary string `form:"company_salary"`
	CompanyStatus string `form:"company_status"`
}