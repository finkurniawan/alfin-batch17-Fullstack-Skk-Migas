package main

import (
	"fmt"
	"sort"
	"strings"
)

type CharCounterData struct {
	Char  rune
	Count int
}

func CariFrequencyDanPengurutan(data []string) string {
	//	cari frequency
	// membuat penyimpan data menggunakan map
	//	saya menggunakan tipe data rune untuk menandakan data bukan hanya sekedar number biasa / mewakili suatu huruf
	wordFrequency := make(map[rune]int)

	//	 melakukan looping setiap huruf untuk menghitung frequencynya
	for _, word := range data {
		for _, char := range word {
			// saya melakukan increment setiap char yang sama
			wordFrequency[char]++
		}
	}

	// saya membuat slice untuk nantinya di urutkan, alasannya karena map tidak menyimpan urutan bisa menyebabkan urutan yang salah jika tidak di ubah ke slice terlebih dahulu
	charData := make([]CharCounterData, 0, len(wordFrequency))
	for char, count := range wordFrequency {
		charData = append(charData, CharCounterData{char, count})
	}

	// melakukan sorting sesuai kriteria yang di minta
	sort.Slice(charData, func(i, j int) bool {

		// jika ada yang sama return sesuai ASC
		if charData[i].Count == charData[j].Count {

			// urutkan berdasarkan i adalah huruf besar  / param ke 1
			if (charData[i].Char >= 'A' && charData[i].Char <= 'Z') && (charData[j].Char >= 'a' && charData[j].Char <= 'z') {
				return true
			}

			// tidak urutkan jika j yang huruf besar / param ke 2
			if (charData[j].Char >= 'A' && charData[j].Char <= 'Z') && (charData[i].Char >= 'a' && charData[i].Char <= 'z') {
				return false
			}

			// urutkan Char secara ASC
			return charData[i].Char < charData[j].Char
		}

		// urutkan sesuai data terbanyak
		return charData[i].Count > charData[j].Count
	})

	result := strings.Builder{}
	// menggabungkan setiap huruf
	for _, char := range charData {
		result.WriteRune(char.Char)
	}

	// merubah data menjadi string
	return result.String()
}

func main() {
	// saya menyiapkan data yang di perlukan
	input1 := []string{"Abc", "bCd"}
	input2 := []string{"Oke", "One"}
	input3 := []string{"Pendanaan", "Terproteksi", "Untuk", "Dampak", "Berarti"}

	// menguji kode sesuai kriteria yang di minta beserta outputnya
	fmt.Println(CariFrequencyDanPengurutan(input1)) // Output: bACcd
	fmt.Println(CariFrequencyDanPengurutan(input2)) // Output: Oekn
	fmt.Println(CariFrequencyDanPengurutan(input3)) // Output: aenrktipBDPTUdmosu
}
