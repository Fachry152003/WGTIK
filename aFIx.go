package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const hPrioritas int = 100000
const hEksekutif int = 50000
const hEkonomi int = 30000

type penumpang struct {
	id string
	nama string
	usia int
	kelas string
	kursi string
	kr string
	fixed bool
	ci string
	nmOrtu string
}


type kereta struct {
	prioritas [9][4]penumpang
	eksekutif [12][4]penumpang
	ekonomi [14][4]penumpang
	nPrioritas int
	nEksekutif int
	nEkonomi int
}

const  maxPrioritas int = 36
const  maxEksekutif int = 48
const  maxEkonomi int = 56

var bayi int
var noResv int = 1
var TEMP [1000]penumpang
var train kereta

func batas() {
	fmt.Println("_______________________________________")
}

func menu(pilih *int) {
	fmt.Println("\nSELAMAT DATANG DI IFONE TRAIN STATION \nSILAHKAN PILIH OPTION BERIKUT\n1. Lakukan reservasi\n2. Cek Harga\n3. Ubah Tempat\n4. Tampilkan Tempat Kosong\n5. Mencari Orang Tua\n6. Data bayi\n7. Data Seluruh Penumpang\n8. Tampilkan Okupansi\n9. Baca Data Dari File\n10. Cari lansia\n0. Exit (selesai)")

	fmt.Scan(&*pilih)
}

func reservasi(){
	var temp penumpang
	var nTiket int
	var t int = 0

	for temp.kelas != "prioritas" && temp.kelas != "eksekutif" && temp.kelas != "ekonomi"{
		fmt.Print("Masukkan kelas tiket (prioritas/eksekutif/ekonomi) : ")
		fmt.Scan(&temp.kelas)
	}

	for nTiket > 10 || nTiket < 1 {
		fmt.Print("Masukkan jumlah tiket (maksimal 10 tiket) : ")
		fmt.Scan(&nTiket)
	}
//     var x int = 0

	for t < nTiket {
		fmt.Println("SILAHKAN MASUKAN ORANG ke -", t+1)

		fmt.Print("nama: ")
		fmt.Scan(&temp.nama)

		fmt.Print("usia: ")
		fmt.Scan(&temp.usia)

        if temp.usia >= 18{
            fmt.Print("id: ")
            fmt.Scan(&temp.id)
        }

        if temp.usia >2 && temp.usia <18{
            fmt.Print("id orang tua: ")
            fmt.Scan(&temp.id)
        }

		if temp.usia <= 2{
            fmt.Print("id orang tua: ")
            fmt.Scan(&temp.id)
        }

        var cek bool = false
		if temp.kelas == "prioritas" {	
			for i := 0; i < 9 && cek == false ; i++ {
				for j := 0; j < 4 && cek == false ;j++ {
					if train.prioritas[i][j].nama == "" {
						if j == 0{
							temp.kursi = "A"
						}else if j == 1{
							temp.kursi ="B"
						}else if j == 2{
							temp.kursi = "C"
						}else {
							temp.kursi = "D"
						}

						TEMP[train.nPrioritas+train.nEksekutif+train.nEkonomi] = temp
						train.prioritas[i][j] = temp
						train.nPrioritas++
						t+=1
						cek = true
						
					}
				}
				fmt.Println("")
			}
			fmt.Println("Kelas Prioritas")
			for i := 0; i < 9; i++ {
				for j := 0; j < 4; j++ {
					if train.prioritas[i][j].nama == "" {
						fmt.Print("[]", " ")
					} else {
						fmt.Print(train.prioritas[i][j].nama, " ")
					}
				}
				fmt.Println("")
			}
		}

		if temp.kelas == "eksekutif" {	
			for i := 0; i < 12 && cek == false ; i++ {
				for j := 0; j < 4 && cek == false ;j++ {
					if train.eksekutif[i][j].nama == "" {
						if j == 0{
							temp.kursi = "A"
						}else if j == 1{
							temp.kursi ="B"
						}else if j == 2{
							temp.kursi = "C"
						}else {
							temp.kursi = "D"
						}

						TEMP[train.nPrioritas+train.nEksekutif+train.nEkonomi] = temp
						train.eksekutif[i][j] = temp
						train.nEksekutif++
						t+=1
						cek = true
						
					}
				}
				fmt.Println("")
			}
			fmt.Println("Kelas Eksekutif")
			for i := 0; i < 12; i++ {
				for j := 0; j < 4; j++ {
					if train.eksekutif[i][j].nama == "" {
						fmt.Print("[]", " ")
					} else {
						fmt.Print(train.eksekutif[i][j].nama, " ")
					}
				}
				fmt.Println("")
			}
		}		

		if temp.kelas == "ekonomi" {	
			for i := 0; i < 14 && cek == false ; i++ {
				for j := 0; j < 4 && cek == false ;j++ {
					if train.ekonomi[i][j].nama == "" {
						if j == 0{
							temp.kursi = "A"
						}else if j == 1{
							temp.kursi ="B"
						}else if j == 2{
							temp.kursi = "C"
						}else {
							temp.kursi = "D"
						}

						TEMP[train.nPrioritas+train.nEksekutif+train.nEkonomi] = temp
						train.ekonomi[i][j] = temp
						train.nEkonomi++
						t+=1
						cek = true
						
					}
				}
				fmt.Println("")
			}
			fmt.Println("Kelas Ekonomi")
			for i := 0; i < 14; i++ {
				for j := 0; j < 4; j++ {
					if train.ekonomi[i][j].nama == "" {
						fmt.Print("[]", " ")
					} else {
						fmt.Print(train.ekonomi[i][j].nama, " ")
					}
				}
				fmt.Println("")
			}
		
		}	
		
	}
	cekHarga(temp.kelas)
}

func printKereta() {
	fmt.Println("Kelas Prioritas")
	var PrioKosong int = 0 
	for i := 0; i < 9; i++ {
		for j := 0; j < 4; j++ {
			if train.prioritas[i][j].nama == "" {
				fmt.Print("[]", " ")
				PrioKosong += 1
			} else {
				fmt.Print(train.prioritas[i][j].nama, " ")
			}
		}
		fmt.Println("")
	}
	fmt.Println("Terdapat",PrioKosong,"Kursi Kosong")
	fmt.Println("")
	fmt.Println(`
	=================
	`)
	
	fmt.Println("Kelas Eksekutif")
	var EkseKosong int = 0
	for i := 0; i < 12; i++ {
		for j := 0; j < 4; j++ {
			if train.eksekutif[i][j].nama == "" {
				fmt.Print("[]", " ")
				EkseKosong +=1
			} else {
				fmt.Print(train.eksekutif[i][j].nama, " ")
			}
		}
		fmt.Println("")
	}
	fmt.Println("Tedapat",EkseKosong,"Kursi Kosong")
	fmt.Println("")
	fmt.Println(`
	=================
	`)
	
	fmt.Println("Kelas Ekonomi")
	var EkoKosong int = 0
	for i := 0; i < 14; i++ {
		for j := 0; j < 4; j++ {
			if train.ekonomi[i][j].nama == "" {
				fmt.Print("[]", " ")
				EkoKosong += 1
			} else {
				fmt.Print(train.ekonomi[i][j].nama, " ")
			}
		}
		fmt.Println("")
	}
	fmt.Println("Tedapat",EkoKosong,"Kursi Kosong")
	fmt.Println("")
	fmt.Println(`
	=================
	`)
	
}

// func cekHarga(){
// 	var train kereta
// 	var Ekodewasa int
// 	var Eksdewasa int
// 	var Pridewasa int
// 	var Ekoanak int
// 	var Eksanak int
// 	var Prianak int
// 	var hargaEko float64
// 	var hargaPri float64
// 	var hargaEks float64

// 	Ekodewasa= 0
// 	Eksdewasa = 0
// 	Pridewasa = 0
// 	Eksanak = 0
// 	Ekoanak = 0
// 	Prianak = 0
// 	harga = 0

// 	for i:=0 ; i<14;i++{
// 		for j:=0 ;j<4;j++{
// 			if train.ekonomi[i][j].usia >= 18{
// 				Ekodewasa += 1
// 			}
// 		}	
// 	} 
// 	for i:=0 ; i<14;i++{
// 		for j:=0 ;j<4;j++{
// 			if train.prioritas[i][j].usia >= 18{
// 				Pridewasa += 1
// 			}
// 		}	
// 	}
// 	for i:=0 ; i<14;i++{
// 		for j:=0 ;j<4;j++{
// 			if train.eksekutif[i][j].usia >= 18{
// 				Eksdewasa += 1
// 			}
// 		}	
// 	}  

// 	for i:=0 ; i<14;i++{
// 		for j:=0 ;j<4;j++{
// 			if train.eksekutif[i][j].usia < 18 && train.eksekutif[i][j].usia >2{
// 				Eksanak += 1
// 			}
// 		}	
// 	} 

// 	for i:=0 ; i<14;i++{
// 		for j:=0 ;j<4;j++{
// 			if train.ekonomi[i][j].usia < 18 && train.ekonomi[i][j].usia >2{
// 				Ekoanak += 1
// 			}
// 		}	
// 	} 
	 
// 	for i:=0 ; i<14;i++{
// 		for j:=0 ;j<4;j++{
// 			if train.prioritas[i][j].usia < 18 && train.prioritas[i][j].usia >2{
// 				Prianak += 1
// 			}
// 		}	
// 	} 
	
// 	if Ekoanak + Ekodewasa % 4==0 {

// 	}
// 	hargaPri = float64(hPrioritas)*float64(Prianak)*0.5 + float64(hPrioritas)*float64(Prianak)*0.5
// 	hargaEko = float64(hPriorits)*float64(Prianak)*0.5
// } 


// func cekharga(){
// 	var nD = 10
// 	var nA = 3

// 	// var Dx float64 = 1
// 	// var Ax float64 = 0.5

// 	var gs int = 0
// 	// var bayar = 0


// 	var tot int = nD+nA
	
// 	if tot/4 >=1 {
// 		gs +=1
// 	} else if tot/3 >=1{
// 		gs+=1
// 	}
	
// }

// func cariparent(x string, y int) bool {

// 	var found = false
	
// 	for z:=0; (!found && z<y);{
// 		if x == a[z].ni {
// 			fmt.Println("======DATA  Orang Tua======")
// 			fmt.Print("Nomor Identitas		: ")
// 			fmt.Println(aTEMP[z].id)
// 			fmt.Print("Nama			: ")
// 			fmt.Println(a[z].n
//	ama)
// 			fmt.Print("Usia			: ")
// 			fmt.Println(a[z].usia)
// 			fmt.Print("Kelas Tiket		: ")
// 			fmt.Println(a[z].kelas)
// 			fmt.Println("=========================")
// 			found = true
// 			return true
// 		} else {
// 			z++
// 		}
// 	}
// 	fmt.Println("Maaf, Nomor induk orang tua tidak ditemukan.")
// 	fmt.Scanln()
// 	return false
// }

func sortUmur() {
	fmt.Println("Nama anak-anak atau bayi yang diurutkan berdasarkan umur :")
	var dataSort [2022]penumpang
	var nmax int
	for i := 0; i < train.nPrioritas+train.nEkonomi+train.nEksekutif+bayi; i++ {
		if TEMP[i].usia < 99999 {
			dataSort[nmax] = TEMP[i]
			nmax++
		}

	}

	for i := 0; i < nmax; i++ {
		min := i
		j := i + 1
		for j < nmax {
			if dataSort[min].usia > dataSort[j].usia {
				min = j
			}
			j++
		}
		temp := dataSort[i]
		dataSort[i] = dataSort[min]
		dataSort[min] = temp
	}
	for i := 0; i < nmax; i++ {
		fmt.Println("Nama : ", dataSort[i].nama, "| Usia : ", dataSort[i].usia,"|Kelas : ", dataSort[i].kelas,"| Kursi :", strconv.Itoa(i+1)+dataSort[i].kursi)
	}
}

// func checkIn() {
// 	var kr string
// 	var checkIn bool
// 	fmt.Print("Masukkan kode reservasi : ")
// 	fmt.Scan(&kr)
// 	fmt.Println("")

// 	if kr[2] == 'i' {
// 		fmt.Println("Kereta Kelas Prioritas")
// 	}
// 	if kr[2] == 'o' {
// 		fmt.Println("Kereta Kelas Ekonomi")
// 	}
// 	if kr[2] == 's' {
// 		fmt.Println("Kereta Kelas Eksekutif")
// 	}


// 	for i := 0; i < N; i++ {
// 		for j := 0; j < 4; j++ {
// 			if train.prioritas[i][j].kr == kr {
// 				batas()
// 				fmt.Println("Nama :", train.proritas[i][j].nama, "|", "Nomor Identitas :", train.prioritas[i][j].id, "|", "Usia :", train.prioritas[i][j].usia,"Kelas : Prioritas")
// 				fmt.Println("No Kursi :", 4*i+j)
// 				if train.Prioritas[i][j].ci == "Sudah Check-in" {
// 					checkIn = true
// 				}
// 			}
// 		}
// 	}
// 	for i := 0; i < M; i++ {
// 		for j := 0; j < 4; j++ {
// 			if plane.ekonomi[i][j].kr == kr {
// 				batas()
// 				fmt.Println("Nama :", plane.ekonomi[i][j].nama, "|", "Nomor Paspor :", plane.ekonomi[i][j].np, "|", "Dietary :", plane.ekonomi[i][j].dietary)
// 				fmt.Println("No Kursi :", 6*i+j)
// 				if plane.ekonomi[i][j].ci == "Sudah Check-in" {
// 					checkIn = true
// 				}
// 			}
// 		}
// 	}
// 	var pilih int = 1
// 	for !checkIn {
// 		for pilih != 0 {
// 			fmt.Println("\nData penumpang belum melakukan check-in\nSilahkan melakukan check-in\nApabila sudah melakukan check-in, maka rombongan tersebut tidak dapat mengubah tempat duduk dan dietary\n1. Check-in\n2. Ganti kursi\n3. Ganti status dietary\n0. Exit (selesai)")
// 			fmt.Scan(&pilih)

// 			if pilih == 1 {
// 				if kr[0] == 'B' {
// 					for i := 0; i < N; i++ {
// 						for j := 0; j < 4; j++ {
// 							if plane.bisnis[i][j].kr == kr {
// 								plane.bisnis[i][j].ci = "Sudah Check-in"
// 								checkIn = true
// 							}
// 						}
// 					}
// 				} else if kr[0] == 'E' {
// 					for i := 0; i < N; i++ {
// 						for j := 0; j < 4; j++ {
// 							if plane.ekonomi[i][j].kr == kr {
// 								plane.ekonomi[i][j].ci = "Sudah Check-in"
// 								checkIn = true
// 							}
// 						}
// 					}
// 				}
// 				pilih = 0
// 				batas()
// 				fmt.Println("Terimakasih sudah melakukan check-in")

// 			} else if pilih == 2 {
// 				var np string
// 				var nama string
// 				fmt.Print("Masukkan nama anda : ")
// 				fmt.Scan(&nama)
// 				fmt.Print("Masukkan nomor passport : ")
// 				fmt.Scan(&np)
// 				gantiKursi(kr, np, nama, &plane)
// 			} else if pilih == 3 {
// 				var np string
// 				var nama string
// 				fmt.Print("Masukkan nama anda : ")
// 				fmt.Scan(&nama)
// 				fmt.Print("Masukkan nomor passport : ")
// 				fmt.Scan(&np)

// 				if kr[0] == 'B' {
// 					for i := 0; i < N; i++ {
// 						for j := 0; j < 4; j++ {
// 							if plane.bisnis[i][j].kr == kr && plane.bisnis[i][j].np == np && plane.bisnis[i][j].nama == nama {

// 								for plane.bisnis[i][j].dietary != "ya" && plane.bisnis[i][j].dietary != "bukan" {
// 									fmt.Print("Apakah anda vegetarian atau bukan? (ya/bukan) : ")
// 									fmt.Scan(&plane.bisnis[i][j].dietary)
// 								}

// 								if plane.bisnis[i][j].dietary == "ya" {
// 									plane.bisnis[i][j].dietary = "Vegetarian"
// 									for i := 0; i < plane.jumlah_bisnis+plane.jumlah_ekonomi+bayi; i++ {
// 										if DATA[i].nama == nama && DATA[i].np == np {
// 											DATA[i].dietary = "Vegetarian"
// 										}

// 									}
// 								} else if plane.bisnis[i][j].dietary == "bukan" {
// 									plane.bisnis[i][j].dietary = "non-Vegetarian"
// 									for i := 0; i < plane.jumlah_bisnis+plane.jumlah_ekonomi+bayi; i++ {
// 										if DATA[i].nama == nama && DATA[i].np == np {
// 											DATA[i].dietary = "Vegetarian"
// 										}

// 									}
// 								}
// 							}
// 						}
// 					}
// 				} else if kr[0] == 'E' {
// 					for i := 0; i < M; i++ {
// 						for j := 0; j < 4; j++ {
// 							if plane.ekonomi[i][j].kr == kr && plane.ekonomi[i][j].np == np && plane.ekonomi[i][j].nama == nama {
// 								for plane.ekonomi[i][j].dietary != "ya" && plane.ekonomi[i][j].dietary != "bukan" {
// 									fmt.Print("Apakah anda vegetarian atau bukan? (ya/bukan) : ")
// 									fmt.Scan(&plane.ekonomi[i][j].dietary)
// 								}
// 								if plane.ekonomi[i][j].dietary == "ya" {
// 									plane.ekonomi[i][j].dietary = "Vegetarian"
// 									for i := 0; i < plane.jumlah_bisnis+plane.jumlah_ekonomi+bayi; i++ {
// 										if DATA[i].nama == nama && DATA[i].np == np {
// 											DATA[i].dietary = "Vegetarian"
// 										}

// 									}
// 								} else if plane.ekonomi[i][j].dietary == "bukan" {
// 									plane.ekonomi[i][j].dietary = "non-Vegetarian"
// 									for i := 0; i < plane.jumlah_bisnis+plane.jumlah_ekonomi+bayi; i++ {
// 										if DATA[i].nama == nama && DATA[i].np == np {
// 											DATA[i].dietary = "non-Vegetarian"
// 										}

// 									}
// 								}
// 							}
// 						}
// 					}
// 				}

// 			}

// 		}
// 		if pilih == 0 {
// 			break
// 		}
// 	}
// 	if checkIn {
// 		fmt.Println("\nRombongan dengan kode reservasi", kr, "sudah melakukan check-in")
// 		fmt.Println()
// 	}
// }

func cariNamaOrtu() {
	fmt.Println("Data orang tua yang membawa anak :")
	for i := 0; i < train.nPrioritas+train.nEkonomi+train.nEksekutif+bayi; i++ {
		if TEMP[i].nmOrtu != "" {
			fmt.Println(TEMP[i].nmOrtu)
		}

	}
}

// func TempatBayi(){
// 	fmt.Println("Kelas Prioritas")
// 	for i := 0; i < 9; i++ {
// 		for j := 0; j < 4; j++ {
// 			if train.prioritas[i][j].usia > 2 {
// 				fmt.Print("[]", " ")
// 			} else {
// 				fmt.Print(train.prioritas[i][j].nama, " ")
// 			}
// 		}
// 		fmt.Println("")
// 	}
	
// 	fmt.Println(`
// 	=================
// 	`)
	
// 	fmt.Println("Kelas Eksekutif")
// 	for i := 0; i < 12; i++ {
// 		for j := 0; j < 4; j++ {
// 			if train.eksekutif[i][j].usia >2  {
// 				fmt.Print("[]", " ")
// 			} else {
// 				fmt.Print(train.eksekutif[i][j].nama, " ")
// 			}
// 		}
// 		fmt.Println("")
// 	}
	
// 	fmt.Println(`
// 	=================
// 	`)
	
// 	fmt.Println("Kelas Ekonomi")
// 	for i := 0; i < 14; i++ {
// 		for j := 0; j < 4; j++ {
// 			if train.ekonomi[i][j].usia > 2 {
// 				fmt.Print("[]", " ")
// 			} else {
// 				fmt.Print(train.ekonomi[i][j].nama, " ")
// 			}
// 		}
// 		fmt.Println("")
// 	}
	
// 	fmt.Println(`
// 	=================
// 	`)
	
// }

func TempatBayi(){
	

	fmt.Println("Kelas Prioritas")
	for i:= 0;i<train.nPrioritas;i++{
		if TEMP[i].usia<=2 && (TEMP[i].kelas == "Prioritas" || TEMP[i].kelas == "prioritas") {
			fmt.Println("Nama : ", TEMP[i].nama, "| Usia : ", TEMP[i].usia, "| Kursi :", strconv.Itoa(i+1)+TEMP[i].kursi)
		}
	}
	batas()
	fmt.Println("Kelas Ekonomi")
	for i:= 0;i<train.nEkonomi;i++{
		if TEMP[i].usia<=2 && (TEMP[i].kelas == "Ekonomi" || TEMP[i].kelas == "ekonomi") {
			fmt.Println("Nama : ", TEMP[i].nama, "| Usia : ", TEMP[i].usia,"| Kursi :", strconv.Itoa(i+1)+TEMP[i].kursi)
		}
	}	
	batas()
	fmt.Println("Kelas Eksekutif")
	for i:= 0;i<train.nEksekutif;i++{
		if TEMP[i].usia<=82 && (TEMP[i].kelas == "Eksekutif" || TEMP[i].kelas == "eksekutif") {
			fmt.Println("Nama : ", TEMP[i].nama, "| Usia : ", TEMP[i].usia,"| Kursi :", strconv.Itoa(i+1)+TEMP[i].kursi)
		}
	}
	
}
func createFile() {

	var dataPrint [200]string
	for i := 0; i < train.nPrioritas+train.nEksekutif+train.nEkonomi+bayi; i++ {
		if TEMP[i].usia < 2 {
			dataPrint[i] = string("Nama :" + " " + TEMP[i].nama + " " + "|" + " " + "Usia :" + " " + strconv.Itoa(TEMP[i].usia) + " " + "|" + " " + "Nomor Identitas :" + " " + TEMP[i].id + "\n" + "Nama orang tua :" + " " + TEMP[i].nmOrtu + " " + "|" + " " + "Kode reservasi :" + " " + TEMP[i].kr + "\n")
		} else if TEMP[i].usia < 18 {
			dataPrint[i] = string("Nama :" + " " + TEMP[i].nama + " " + "|" + " " + "Usia :" + " " + strconv.Itoa(TEMP[i].usia) + " " + "|" + " " + "Nomor Identitas :" + " " + TEMP[i].id + " "  + " "+ "\n" + "Nama orang tua :" + " " + TEMP[i].nmOrtu + " " + "|" + " " + "Kode reservasi :" + " " + TEMP[i].kr + "\n")
		} else {
			dataPrint[i] = string("Nama :" + " " + TEMP[i].nama + " " + "|" + " " + "Usia :" + " " + strconv.Itoa(TEMP[i].usia) + " " + "|" + " " + "Nomor Identitas :" + " " + TEMP[i].id + " " + "|" + " " + "\n" + "Kode reservasi :" + " " + TEMP[i].kr + "\n")
		}
	}

	f, err := os.Create("DataPenumpang.txt")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	for _, v := range dataPrint {
		fmt.Fprintln(f, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file .txt berisi data penumpang sudah berhasil dibuat dengan nama file DataPenumpang.txt")
}

//prosedur untuk memasukan data penumpang dari file teks atau csv
func reservasiScan(nama, usia, id, ortu, kr string, count int) {

	var data penumpang

	data.ci = "Belum check-in"
	data.nama = nama
	data.id = id
	data.kr = kr

	umur, err := strconv.Atoi(usia)
	if err != nil {
		fmt.Println(err)
		return
	}
	data.usia = umur
	data.nmOrtu = ortu

	if data.kr[2] == 'i' {
		data.kelas = "Prioritas"
	} else if data.kr[2] == 'o' {
		data.kelas = "Ekonomi"
	} else if data.kr[2] == 's'{
		data.kelas = "Eksekutif"
	}

	TEMP[train.nPrioritas+train.nEksekutif+train.nEkonomi+bayi] = data

	if data.usia < 2 {
		bayi++
	}
	if data.kelas == "Ekonomi" {
		var stop bool = false
		for i := 0; i < 14 && !stop; i++ {
			for j := 0; j < 4 && !stop; j++ {
				if train.ekonomi[i][j].id == "" && data.usia >= 2 {
					train.ekonomi[i][j] = data
					train.nEkonomi++
					stop = true
				}

			}
		}
	} else if data.kelas == "Prioritas" {
		var stop bool = false
		for i := 0; i < 9 && !stop; i++ {
			for j := 0; j < 4 && !stop; j++ {
				if train.prioritas[i][j].id == "" && data.usia >= 2 {
					train.prioritas[i][j] = data
					train.nPrioritas++
					stop = true
				}
			}
		}
	}	else if data.kelas == "Eksekutif" {
		var stop bool = false
		for i := 0; i < 12 && !stop; i++ {
			for j := 0; j < 4 && !stop; j++ {
				if train.eksekutif[i][j].id == "" && data.usia >= 2 {
					train.eksekutif[i][j] = data
					train.nEksekutif++
					stop = true
				}
			}
		}
	}
}

//prosedur untuk membaca file csv
func scanData() {
	file, err := os.Open("Data.csv")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	count := 0
	items := make([]string, 6)

	for scanner.Scan() {
		line := scanner.Text()
		items = strings.Split(line, ",")
		reservasiScan(items[0], items[1], items[2], items[3], items[4], count)

	}

	fmt.Print("Data penumpang berhasil diinput\n")
}

func cekHarga(a string){

	
	// const prioritas float64 	= 150000
	// const ekonomi 	float64 = 50000

	const reguler float64 = 1
	const kid float64 = 0.5

	var bayar float64 = 0
	var bayarNormal float64 = 0
	var bayarGS float64 = 0
	var nGS float64 = 0
	var tempGS int = 0
	var tempGSarr [4]string
	if a == "eksekutif" || a == "Eksekutif"{
		for i := 0; i < 12; i++ {
			for j:= 0; j< 4;j++{
				if train.eksekutif[i][j].usia > 18 {
					bayarNormal += float64(hEksekutif) * reguler
					bayar += float64(hEksekutif) * reguler
				}else if train.eksekutif[i][j].usia <18 {
					bayarNormal += float64(hEksekutif) * kid
					bayar += float64(hEksekutif) * kid
				}
				tempGSarr[tempGS] = train.eksekutif[i][j].id
				tempGS += 1
				if tempGS == 4 {
					nGS += 1
					tempGS = 0

					for j := 0; j < 2; j++ {
						bayarGS += float64(hEksekutif)* reguler
					}
					if bayarGS < bayar {
						bayar = bayarGS
					} else {
						bayar = bayarNormal
						nGS -= 1
					}
					bayarNormal = 0
				}
				if tempGS == 3 && i+1 == len(train.eksekutif) {
					for k := 0; k < 2; j++ {
						bayarGS += float64(hEksekutif) * reguler
					}
					nGS += 1
				if bayarGS < bayar {
					bayar = bayarGS
				} else {
					bayar = bayarNormal
					nGS -= 1
				}

					bayarNormal = 0
		
				}	
			}		
		}
	}	
	fmt.Println(bayar)
	fmt.Println(nGS) //jumlah group save
}

func okupansi() {
	var pilih string
	fmt.Print("Pilih Gerbong :")
	fmt.Scan(&pilih)
	if pilih == "prioritas" || pilih == "Prioritas"  {
		fmt.Print((float64(train.nPrioritas) / float64(36)) * float64(100))
	} else if pilih == "ekonomi" || pilih == "Ekonomi" {
		fmt.Print((float64(train.nEkonomi) / float64(56)) * float64(100))
	} else if pilih == "eksekutif" || pilih == "Eksekutif" {
		fmt.Print((float64(train.nEksekutif) / float64(48)) * float64(100))
	}

}
func cariLansia() {
	for i := 0; i < train.nPrioritas+train.nEksekutif+train.nEkonomi+bayi; i++ {
		if TEMP[i].usia >= 65 {
			fmt.Println("Data lansia :")
			fmt.Println("Nama :", TEMP[i].nama, "Usia :", TEMP[i].usia, "tahun")
		}

	}
}
func main() {

	var running bool = true

	for running {

		var pilih int
		menu(&pilih)

		switch (pilih) {
		case 1 :
			reservasi()
		case 2 :
			//cekHarga()
		case 3 :
			//ubahTempat()
		case 4 :
			printKereta()
		case 5 :
			cariNamaOrtu()
		case 6 : 
			TempatBayi()
		case 7 :
			sortUmur()
		case 8 :
			okupansi()
		case 9 :
			scanData()
		case 10 : 
			cariLansia()
		default :
			fmt.Print("exit")
			running = false
		}

		fmt.Scanln()

		
	}
}