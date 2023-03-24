package main

import (
	"fmt"
	"os"
	"strconv"
)

type biodata struct {
  nama    string
  alamat  string
  Pekerjaan   string
  alasan string
}

func main() {
  args := os.Args
  
  if len(args) == 2 {

    absen, err := strconv.Atoi(args[1])
    if err != nil {
        panic(err)
    }
  
    // membuat slice sepanjang 5 data agar manampilkan 5 data endpoint
    teman := []biodata{
        biodata{
          nama:    "Ayam Geprek",
          alamat: "Mawar No.56, Surabaya",
          Pekerjaan: "Guru",
          alasan: "Ingin Memperdalam Ilmu Pemrograman",
        },
        biodata{
          nama:    "Burger Jamur",
          alamat: "Mawar No.56, Surabaya",
          Pekerjaan: "Siswa",
          alasan: "Ingin Paham dan Menguasai Bahasa Go",
        },
        biodata{
          nama:    "Gorengan Gorengan",
          alamat: "Mawar No.56, Surabaya",
          Pekerjaan: "Karyawan",
          alasan: "Ingin Belajar mengenai Golang",
        },
        biodata{
          nama:    "Rendang Padang",
          alamat: "Mawar No.56, Surabaya",
          Pekerjaan: "Penjual",
          alasan: "Hoby baru Belajar Golang",
        },
        biodata{
          nama:    "Selat Solo",
          alamat: "Mawar No.56, Surabaya",
          Pekerjaan: "Pembeli",
          alasan: "Ingin Belajar Tantangan Bahasa Go",
        },
    }

    // menampilkan data teman dikelas dari angka absen yang dimasukkan
    data := teman[absen-1]
    fmt.Printf("Nama \t\t: %s\nAlamat \t\t: %s\nPekerjaan \t: %s\nAlasan \t\t: %s\n", data.nama, data.alamat, data.Pekerjaan, data.alasan)

  } else {
    fmt.Println("Masukkan angka absen teman kalian dalam kelas")
  }

}