package main

import "fmt"

type Data struct 
    {
    Nama, alamat string
    Nim      int
    }
func main() {
    multiA := [10]Data {
        {"Nasruddin","Jl. Hertasning",519012011},
		{"Rangga P", "Jl.Serigala", 519012012},
		{"Rezky A", "Jl.Mangerangi", 519012013},
		{"Chandra", "Jl.Perintis", 519012014},
		{"Fadel B", "Jl.Aroe Pala", 519012015},
		{"Saleh", "Sudiang", 519012016},
		{"Amira", "Rappocini", 519012017},
		{"Ayu", "Kerung-kerung", 519012018},
		{"Satar", "Pannampu", 519012019},
		{"Pijai", "Jl. A. Yani", 519012020},
	}
	
	fmt.Println("============================")
	fmt.Println("BIODATA MAHASISWA PANCASAKTI")

	for jumlah := 0; jumlah < 10; jumlah++ {
	 fmt.Println("===========================")
	 fmt.Printf("Nama \t: %v\n", multiA[jumlah].Nama)
	 fmt.Printf("Alamat \t: %v\n", multiA[jumlah].alamat)
	 fmt.Printf("NIM \t: %v\n", multiA[jumlah].Nim)
	}

	fmt.Println("============================")
	fmt.Println(" GUNAKAN DATA DENGAN BIJAK  ")
}