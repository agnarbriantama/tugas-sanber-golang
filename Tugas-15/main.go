// main.go
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB

// Struktur untuk Mahasiswa
type Mahasiswa struct {
	ID    int       `json:"id"`
	Nama  string    `json:"nama"`
	Nilai []Nilai   `json:"nilai,omitempty"`
}

// Struktur untuk Mata Kuliah
type MataKuliah struct {
	ID    int       `json:"id"`
	Nama  string    `json:"nama"`
	Nilai []Nilai   `json:"nilai,omitempty"`
}

// Struktur untuk Nilai
type Nilai struct {
	ID           int       `json:"id"`
	Indeks       string    `json:"indeks"`
	Skor         int       `json:"skor"`
	MahasiswaID  int       `json:"mahasiswa_id"`
	MataKuliahID int       `json:"mata_kuliah_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func main() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/db_mahasiswa")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := mux.NewRouter()

	// Rute API untuk CRUD Mahasiswa
	r.HandleFunc("/mahasiswa", GetAllMahasiswa).Methods("GET")
	r.HandleFunc("/mahasiswa/{id}", GetMahasiswa).Methods("GET")
	r.HandleFunc("/mahasiswa", CreateMahasiswa).Methods("POST")
	r.HandleFunc("/mahasiswa/{id}", UpdateMahasiswa).Methods("PUT")
	r.HandleFunc("/mahasiswa/{id}", DeleteMahasiswa).Methods("DELETE")

	// Rute API untuk CRUD Mata Kuliah
	r.HandleFunc("/mata_kuliah", GetAllMataKuliah).Methods("GET")
	r.HandleFunc("/mata_kuliah/{id}", GetMataKuliah).Methods("GET")
	r.HandleFunc("/mata_kuliah", CreateMataKuliah).Methods("POST")
	r.HandleFunc("/mata_kuliah/{id}", UpdateMataKuliah).Methods("PUT")
	r.HandleFunc("/mata_kuliah/{id}", DeleteMataKuliah).Methods("DELETE")

	// Rute API untuk CRUD Nilai
	r.HandleFunc("/nilai", GetAllNilai).Methods("GET")
	r.HandleFunc("/nilai/{id}", GetNilai).Methods("GET")
	r.HandleFunc("/nilai", CreateNilai).Methods("POST")
	r.HandleFunc("/nilai/{id}", UpdateNilai).Methods("PUT")
	r.HandleFunc("/nilai/{id}", DeleteNilai).Methods("DELETE")

	fmt.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func GetAllMahasiswa(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT * FROM mahasiswa")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var mahasiswas []Mahasiswa

    for rows.Next() {
        var m Mahasiswa
        if err := rows.Scan(&m.ID, &m.Nama); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        mahasiswas = append(mahasiswas, m)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(mahasiswas)
}

// Fungsi untuk mendapatkan data Mahasiswa berdasarkan ID
func GetMahasiswa(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    rows, err := db.Query("SELECT * FROM mahasiswa WHERE id = ?", id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var m Mahasiswa

    if rows.Next() {
        if err := rows.Scan(&m.ID, &m.Nama); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    } else {
        http.Error(w, "Not Found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(m)
}

// Fungsi untuk menambah data Mahasiswa
func CreateMahasiswa(w http.ResponseWriter, r *http.Request) {
    var m Mahasiswa
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&m); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    result, err := db.Exec("INSERT INTO mahasiswa (nama) VALUES (?)", m.Nama)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    id, _ := result.LastInsertId()

    m.ID = int(id)

    w.WriteHeader(http.StatusCreated)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(m)
}

// Fungsi untuk mengupdate data Mahasiswa
func UpdateMahasiswa(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    var m Mahasiswa
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&m); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    _, err := db.Exec("UPDATE mahasiswa SET nama = ? WHERE id = ?", m.Nama, id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    m.ID, _ = strconv.Atoi(id)

    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(m)
}

// Fungsi untuk menghapus data Mahasiswa
func DeleteMahasiswa(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    _, err := db.Exec("DELETE FROM mahasiswa WHERE id = ?", id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{"message": "Deleted successfully"})
}

// Fungsi untuk mendapatkan semua data Mata Kuliah
func GetAllMataKuliah(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT * FROM mata_kuliah")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var mataKuliahs []MataKuliah

    for rows.Next() {
        var mk MataKuliah
        if err := rows.Scan(&mk.ID, &mk.Nama); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        mataKuliahs = append(mataKuliahs, mk)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(mataKuliahs)
}

// Fungsi untuk mendapatkan data Mata Kuliah berdasarkan ID
func GetMataKuliah(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    rows, err := db.Query("SELECT * FROM mata_kuliah WHERE id = ?", id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var mk MataKuliah

    if rows.Next() {
        if err := rows.Scan(&mk.ID, &mk.Nama); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    } else {
        http.Error(w, "Not Found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(mk)
}

// Fungsi untuk menambah data Mata Kuliah
func CreateMataKuliah(w http.ResponseWriter, r *http.Request) {
    var mk MataKuliah
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&mk); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    result, err := db.Exec("INSERT INTO mata_kuliah (nama) VALUES (?)", mk.Nama)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    id, _ := result.LastInsertId()

    mk.ID = int(id)

    w.WriteHeader(http.StatusCreated)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(mk)
}

// Fungsi untuk mengupdate data Mata Kuliah
func UpdateMataKuliah(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    var mk MataKuliah
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&mk); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    _, err := db.Exec("UPDATE mata_kuliah SET nama = ? WHERE id = ?", mk.Nama, id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    mk.ID, _ = strconv.Atoi(id)

    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(mk)
}

// Fungsi untuk menghapus data Mata Kuliah
func DeleteMataKuliah(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    _, err := db.Exec("DELETE FROM mata_kuliah WHERE id = ?", id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{"message": "Deleted successfully"})
}

// Fungsi untuk mendapatkan semua data Nilai
func GetAllNilai(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT * FROM nilai")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var nilais []Nilai

    for rows.Next() {
        var n Nilai
        if err := rows.Scan(&n.ID, &n.Indeks, &n.Skor, &n.MahasiswaID, &n.MataKuliahID, &n.CreatedAt, &n.UpdatedAt); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        nilais = append(nilais, n)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(nilais)
}

// Fungsi untuk mendapatkan data Nilai berdasarkan ID
func GetNilai(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    rows, err := db.Query("SELECT * FROM nilai WHERE id = ?", id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var n Nilai

    if rows.Next() {
        if err := rows.Scan(&n.ID, &n.Indeks, &n.Skor, &n.MahasiswaID, &n.MataKuliahID, &n.CreatedAt, &n.UpdatedAt); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    } else {
        http.Error(w, "Not Found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(n)
}

// Fungsi untuk menambah data Nilai
func CreateNilai(w http.ResponseWriter, r *http.Request) {
    var n Nilai
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&n); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    if n.Skor > 100 {
        http.Error(w, "Skor cannot be more than 100", http.StatusBadRequest)
        return
    }

    result, err := db.Exec(
        "INSERT INTO nilai (indeks, skor, mahasiswa_id, mata_kuliah_id, created_at, updated_at) VALUES (?, ?, ?, ?, NOW(), NOW())",
        n.Indeks, n.Skor, n.MahasiswaID, n.MataKuliahID,
    )
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    id, _ := result.LastInsertId()

    n.ID = int(id)

    w.WriteHeader(http.StatusCreated)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(n)
}

// Fungsi untuk mengupdate data Nilai
func UpdateNilai(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    var n Nilai
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&n); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    if n.Skor > 100 {
        http.Error(w, "Skor cannot be more than 100", http.StatusBadRequest)
        return
    }

    _, err := db.Exec(
        "UPDATE nilai SET indeks = ?, skor = ?, mahasiswa_id = ?, mata_kuliah_id = ?, updated_at = NOW() WHERE id = ?",
        n.Indeks, n.Skor, n.MahasiswaID, n.MataKuliahID, id,
    )
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    n.ID, _ = strconv.Atoi(id)

    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(n)
}

// Fungsi untuk menghapus data Nilai
func DeleteNilai(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    _, err := db.Exec("DELETE FROM nilai WHERE id = ?", id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{"message": "Deleted successfully"})
}




