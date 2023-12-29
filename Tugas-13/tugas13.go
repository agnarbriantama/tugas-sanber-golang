package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
)
type NilaiMahasiswa struct {
	ID           uint
	Nama         string
	MataKuliah   string
	IndeksNilai  string
	Nilai        uint
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}
var mu sync.Mutex

func main() {
	http.HandleFunc("/post-nilai", handlePostNilaiMahasiswa)
	http.HandleFunc("/get-nilai", handleGetNilaiMahasiswa)

	// Menjalankan server di port 8080
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", basicAuthMiddleware(http.DefaultServeMux))
}
// Middleware untuk Basic Auth
func basicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if username == "admin" && password == "admin" {
			next.ServeHTTP(w, r)
			return
		}

		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}

// Fungsi untuk menangani rute POST /nilai
func handlePostNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Membaca data dari body JSON
	var input struct {
		Nama       string `json:"nama"`
		MataKuliah string `json:"mata_kuliah"`
		Nilai      uint   `json:"nilai"`
	}

	contentType := r.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "application/json") {
		// Membaca data dalam bentuk JSON
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&input); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
	} else if strings.HasPrefix(contentType, "multipart/form-data") {
		// Membaca data dalam bentuk formData
		err := r.ParseMultipartForm(10 << 20) // 10 MB, sesuaikan dengan kebutuhan
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}

		// Mengambil nilai dari form-data
		input.Nama = r.FormValue("nama")
		input.MataKuliah = r.FormValue("mata_kuliah")
		nilaiStr := r.FormValue("nilai")

		// Mengonversi nilai menjadi tipe uint
		nilai, err := strconv.ParseUint(nilaiStr, 10, 32)
		if err != nil {
			http.Error(w, "Invalid Nilai", http.StatusBadRequest)
			return
		}
		input.Nilai = uint(nilai)
	} else {
		http.Error(w, "Unsupported Content Type", http.StatusUnsupportedMediaType)
		return
	}

	// Proses pengolahan nilai dan indeks nilai
	indeksNilai := calculateIndeksNilai(input.Nilai)

	// Mengunci akses ke nilaiNilaiMahasiswa untuk mencegah race condition
	mu.Lock()
	defer mu.Unlock()

	// Menambahkan data NilaiMahasiswa ke slice
	newID := uint(len(nilaiNilaiMahasiswa) + 1)
	newNilaiMahasiswa := NilaiMahasiswa{
		ID:           newID,
		Nama:         input.Nama,
		MataKuliah:   input.MataKuliah,
		Nilai:        input.Nilai,
		IndeksNilai:  indeksNilai,
		
		
	}

	nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, newNilaiMahasiswa)

	// Mengembalikan response
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newNilaiMahasiswa)
}


// Fungsi untuk menangani rute GET /nilai
func handleGetNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Mengunci akses ke nilaiNilaiMahasiswa untuk mencegah race condition
	mu.Lock()
	defer mu.Unlock()

	// Mengembalikan response dengan semua data NilaiMahasiswa
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nilaiNilaiMahasiswa)
}

// Fungsi untuk menghitung indeks nilai berdasarkan nilai
func calculateIndeksNilai(nilai uint) string {
	switch {
	case nilai >= 80:
		return "A"
	case nilai >= 70:
		return "B"
	case nilai >= 60:
		return "C"
	case nilai >= 50:
		return "D"
	default:
		return "E"
	}
}
