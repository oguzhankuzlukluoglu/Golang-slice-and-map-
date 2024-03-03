package main

import "fmt"

func main() {
	myMap1 := map[string]int{
		"ahmet": 49,
		"ayşe":  85,
		"veli":  74,
		"canan": 42,
	}
	fmt.Println(myMap1)
	fmt.Println(myMap1["ahmet"])

	myArr := [5]string{"tahran", "belgrad", "roma", "tiflis", "moskova"}
	for index, value := range myArr {
		fmt.Println(index, value)
	}
	myslc := []int{1, 2, 3, 4, 5}
	for i, v := range myslc {
		fmt.Println(i, v)
	}
	myslc = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	slc1 := myslc[2:4]
	fmt.Println(slc1)
	slc2 := myslc[6:8]
	fmt.Println(slc2)
	slc3 := append(slc1, slc2...)
	fmt.Println(slc3)
	myMap := map[string][]string{
		"yaşar kemal":     {"ince memed", "yer demir gök bakır"},
		"sabahattin ali":  {"kuyucaklı yusuf", "kürk mantolu madonna", "değirmen"},
		"haruki murakami": {"1Q84", "dans dans dans", "kumandanı öldürmek"},
	}
	fmt.Println(myMap)
	fmt.Println(myMap["sabahattin ali"])
	fmt.Println(myMap["haruki murakami"][1])
	for key, value := range myMap {
		fmt.Println("yazarımız:", key)
		fmt.Println("kitapları aşağıdadır:")
		for index, value := range value {
			fmt.Println("\t", index+1, value)
		}
	}
}
