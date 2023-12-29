package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"tugas14/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Struktur untuk NilaiMahasiswa
// Struktur untuk NilaiMahasiswa
type NilaiMahasiswa struct {
    ID          uint   `json:"id"`
    Nama        string `json:"nama"`
    MataKuliah  string `json:"mata_kuliah"`
    IndeksNilai string `json:"indeks_nilai"`
    Nilai       uint   `json:"nilai"`
    CreatedAt   string `json:"created_at"`
    UpdatedAt   string `json:"updated_at"`
}


var db *sql.DB

func main() {
    // Koneksi ke database MySQL
    var err error
    db, err = sql.Open("mysql", config.GetDBConnectionInfo())
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Inisialisasi router
    r := mux.NewRouter()

    // Rute API untuk CRUD Nilai Mahasiswa
    r.HandleFunc("/nilai-mahasiswa", GetAllNilaiMahasiswa).Methods("GET")
    r.HandleFunc("/nilai-mahasiswa/{id}", GetNilaiMahasiswa).Methods("GET")
    r.HandleFunc("/nilai-tambah", CreateNilaiMahasiswa).Methods("POST")
    r.HandleFunc("/nilai-edit/{id}", UpdateNilaiMahasiswa).Methods("PUT")
    r.HandleFunc("/nilai-delete/{id}", DeleteNilaiMahasiswa).Methods("DELETE")

    // Menjalankan server di port 8080
    fmt.Println("Server is running on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}

// Fungsi untuk mendapatkan semua data Nilai Mahasiswa
func GetAllNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {
    // Query ke database
    rows, err := db.Query("SELECT * FROM nilai_mahasiswa")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var nilaiMahasiswas []NilaiMahasiswa

    // Ambil data dari hasil query
    for rows.Next() {
        var nilaiMahasiswa NilaiMahasiswa
        if err := rows.Scan(
            &nilaiMahasiswa.ID,
            &nilaiMahasiswa.Nama,
            &nilaiMahasiswa.MataKuliah,
            &nilaiMahasiswa.IndeksNilai,
            &nilaiMahasiswa.Nilai,
            &nilaiMahasiswa.CreatedAt,
            &nilaiMahasiswa.UpdatedAt,
        ); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        nilaiMahasiswas = append(nilaiMahasiswas, nilaiMahasiswa)
    }

    // Mengembalikan response dalam format JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(nilaiMahasiswas)
}

// Fungsi untuk mendapatkan data Nilai Mahasiswa berdasarkan ID
func GetNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {
    // Mendapatkan nilai ID dari variabel path
    vars := mux.Vars(r)
    id := vars["id"]

    // Query ke database
    rows, err := db.Query("SELECT * FROM nilai_mahasiswa WHERE id = ?", id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var nilaiMahasiswa NilaiMahasiswa

    // Ambil data dari hasil query
    if rows.Next() {
        if err := rows.Scan(
            &nilaiMahasiswa.ID,
            &nilaiMahasiswa.Nama,
            &nilaiMahasiswa.MataKuliah,
            &nilaiMahasiswa.IndeksNilai,
            &nilaiMahasiswa.Nilai,
            &nilaiMahasiswa.CreatedAt,
            &nilaiMahasiswa.UpdatedAt,
        ); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    } else {
        http.Error(w, "Not Found", http.StatusNotFound)
        return
    }

    // Mengembalikan response dalam format JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(nilaiMahasiswa)
}

// Fungsi untuk menambah data Nilai Mahasiswa
func CreateNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {
    // Membaca data dari body JSON
    var input NilaiMahasiswa
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&input); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    // Menambahkan data ke database
    result, err := db.Exec(
        "INSERT INTO nilai_mahasiswa (nama, mata_kuliah, indeks_nilai, nilai, created_at, updated_at) VALUES (?, ?, ?, ?, NOW(), NOW())",
        input.Nama,
        input.MataKuliah,
        calculateIndeksNilai(input.Nilai),
        input.Nilai,
    )
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Mendapatkan ID dari data yang baru ditambahkan
    id, _ := result.LastInsertId()

    // Mengembalikan response dalam format JSON
    w.WriteHeader(http.StatusCreated)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{"id": id})
}

// Fungsi untuk mengupdate data Nilai Mahasiswa
func UpdateNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {
    // Mendapatkan nilai ID dari variabel path
    vars := mux.Vars(r)
    id := vars["id"]

    // Membaca data dari body JSON
    var input NilaiMahasiswa
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&input); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    // Menyimpan perubahan ke database
    _, err := db.Exec(
        "UPDATE nilai_mahasiswa SET nama = ?, mata_kuliah = ?, indeks_nilai = ?, nilai = ?, updated_at = NOW() WHERE id = ?",
        input.Nama,
        input.MataKuliah,
        calculateIndeksNilai(input.Nilai),
        input.Nilai,
        id,
    )
    
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Mengembalikan response dalam format JSON
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{"message": "Updated successfully"})
}

// Fungsi untuk menghapus data Nilai Mahasiswa
func DeleteNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {
    // Mendapatkan nilai ID dari variabel path
    vars := mux.Vars(r)
    id := vars["id"]

    // Menghapus data dari database
    _, err := db.Exec("DELETE FROM nilai_mahasiswa WHERE id = ?", id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Mengembalikan response dalam format JSON
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{"message": "Deleted successfully"})
}

// Fungsi untuk menghitung indeks nilai berdasarkan nilai
func calculateIndeksNilai(nilai uint) string {
    switch {
    case nilai >= 80 && nilai <= 100:
        return "A"
    case nilai >= 70 && nilai < 80:
        return "B"
    case nilai >= 60 && nilai < 70:
        return "C"
    case nilai >= 50 && nilai < 60:
        return "D"
    default:
        return "E"
    }
}

