package config

import "fmt"

const (
	DBUsername = "root"
	DBPassword = ""
	DBHost     = "localhost"
	DBPort     = "3306"
	DBName     = "db_mahasiswa"
)

// GetDBConnectionInfo mengembalikan string koneksi ke database MySQL
func GetDBConnectionInfo() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DBUsername, DBPassword, DBHost, DBPort, DBName)
}
