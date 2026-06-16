package main

import "fmt"

const NMAX int = 999

type data struct {
	room, name string
	watt, time int
}

type tabData [NMAX]data

func main() {
	var c1, c2, search string
	var d tabData
	var n int = 0

	isiData(&d, &n)

	fmt.Printf("Welcome\n")
	fmt.Printf("FEATURES :\n")
	fmt.Printf("A. Edit\nB. Search\nC. Sort\nD. Show Statistics\nE. Stop\n")
	fmt.Printf("Option : ")

	for c1 != "E" {
		fmt.Scan(&c1)
		switch c1 {
		case "A":
			fmt.Printf("EDIT :\nA. Add Data\nB. Change Data\nC. Delete Data\nD. Back\nOption : ")
			for c2 != "D" {
				fmt.Scan(&c2)
				switch c2 {
				case "A":
					addData(&d, &n)
					callMenu(1)
				case "B":
					changeData(&d, n)
					callMenu(1)
				case "C":
					deleteData(&d, &n)
					callMenu(1)
				case "D":
					fmt.Printf("A. Edit\nB. Search\nC. Sort\nD. Show Statistics\nE. Stop\nOption : ")
				default:
					fmt.Printf("Option doesnt exists. Please re-enter option : ")
				}
			}
		case "B":
			fmt.Printf("Search by:\nA. Name\nB. Room\nOption : ")
			fmt.Scan(&c2)
			fmt.Printf("Search for: ")
			fmt.Scan(&search)
			if c2 == "A" {
				searchDataByName(search, d, n)
			} else {
				searchDataByRoom(search, d, n)
			}
			callMenu(3)
			callMenu(0)
		case "C":
			callMenu(2)
			for c2 != "C" {
				fmt.Scan(&c2)
				switch c2 {
				case "A":
					sortEnergy(&d, n)
					callMenu(2)
				case "B":
					sortAlphabeticalName(&d, n)
					callMenu(2)
				case "C":
					fmt.Printf("A. Edit\nB. Search\nC. Sort\nD. Show Statistics\nE. Stop\nOption : ")
				default:
					fmt.Printf("Option doesnt exists. Please re-enter option : ")
				}
			}
		case "D":
			showData(d, n)
			showStatistics(d, n)
			callMenu(3)
			callMenu(0)
			fmt.Printf("Option : ")
		case "E":
			fmt.Printf("Goodbye!")
		default:
			fmt.Printf("Option doesnt exists. Please re-enter option : ")
		}
	}
}

//placeholder data for easy testing
func isiData(d *tabData, n *int) {
	d[0].name = "Rice_Cooker"
	d[0].room = "Kitchen"
	d[0].watt = 360
	d[0].time = 720

	d[1].name = "Refrigerator"
	d[1].room = "Kitchen"
	d[1].watt = 150
	d[1].time = 1440

	d[2].name = "Microwave"
	d[2].room = "Kitchen"
	d[2].watt = 1200
	d[2].time = 30

	d[3].name = "Television"
	d[3].room = "Living_Room"
	d[3].watt = 100
	d[3].time = 300

	d[4].name = "Air_Conditioner"
	d[4].room = "Bedroom"
	d[4].watt = 1500
	d[4].time = 480

	d[5].name = "Computer"
	d[5].room = "Study_Room"
	d[5].watt = 65
	d[5].time = 360

	d[6].name = "Ceiling_Fan"
	d[6].room = "Bedroom"
	d[6].watt = 75
	d[6].time = 600

	d[7].name = "Washing_Machine"
	d[7].room = "Laundry_Room"
	d[7].watt = 500
	d[7].time = 90

	d[8].name = "Water_Heater"
	d[8].room = "Bathroom"
	d[8].watt = 2000
	d[8].time = 60

	d[9].name = "Vacuum_Cleaner"
	d[9].room = "Living_Room"
	d[9].watt = 800
	d[9].time = 45
	*n = 10
}

//Menu shenanigans
func callMenu(menu int) {
	if menu == 0 {
		fmt.Printf("\nA. Edit\nB. Search\nC. Sort\nD. Show Statistics\nE. Stop\nOption : ")
	} else if menu == 1 {
		fmt.Printf("EDIT :\nA. Add Data\nB. Change Data\nC. Delete Data\nD. Back\nOption : ")
	} else if menu == 2 {
		fmt.Printf("Sort by :\nA. Highest Energy Consumption\nB. Alphabetical Order\nC. Back\nOption : ")
	} else if menu == 3 {
		var s string
		fmt.Print("Type Any to Go Back : ")
		fmt.Scan(&s)
		for s == " " {
			fmt.Print("Option Invalid. Type Any to Go Back : ")
			fmt.Scan(&s)
		}
	}
}

//Adds New Data Input
func addData(d *tabData, n *int) {
	fmt.Printf("Device name: ")
	fmt.Scan(&d[*n].name)
	fmt.Printf("Location: ")
	fmt.Scan(&d[*n].room)
	fmt.Printf("Power: ")
	fmt.Scan(&d[*n].watt)
	fmt.Printf("Duration (in minutes): ")
	fmt.Scan(&d[*n].time)
	*n++
	fmt.Println("Data Added!")
	callMenu(3)
}

//Replaces Data
func changeData(d *tabData, n int) {
	fmt.Println("Which data do you want you change?")
	showData(*d, n)
	var s string
	fmt.Scan(&s)
	searchDataByName(s, *d, n)
	fmt.Println("Data Changed!")
	callMenu(3)
}

//Deletes Single Data
func deleteData(d *tabData, n *int) {
	var x int
	switch {
	case *n <= 0:
		fmt.Printf("Data still empty!")
	default:
		showData(*d, *n)
		fmt.Printf("Which number to delete : ")
		fmt.Scan(&x)
		for i := x - 1; i < *n; i++ {
			d[i] = d[i+1]
		}
		*n--
	}
}

//Searchs Data By Name
func searchDataByName(s string, d tabData, n int) {
	var ketemu data
	ketemu.name = "Belum"
	for i := 0; i < n; i++ {
		if d[i].name == s {
			ketemu = d[i]
		}
	}
	if ketemu.name == "Belum" {
		fmt.Println("Data not found!")
	} else {
		printSingles(ketemu)
	}
}

//Searchs Data By Room
func searchDataByRoom(s string, d tabData, n int) {
	sortAlphabeticalRoom(&d, n)
	var l, r, mid, ketemu int
	ketemu = -1
	l = 0
	r = n - 1
	for l <= r {
		mid = (l + r) / 2
		if d[mid].room == s {
			ketemu = mid
		}
		if d[mid].room < s {
			l = mid + 1
		}
		if d[mid].room > s {
			r = mid - 1
		}
	}
	if ketemu == -1 {
		fmt.Println("Data not found!")
	} else {
		printSingles(d[ketemu])
	}
}

//sorts data Ascending by Usage
func sortEnergy(d *tabData, n int) {
	var idx int
	var t data
	i := 1
	for i <= n-1 {
		idx = i - 1
		j := i
		for j < n {
			if d[idx].watt < d[j].watt {
				idx = j
			}
			j = j + 1
		}
		t = d[idx]
		d[idx] = d[i-1]
		d[i-1] = t
		i = i + 1
	}

}

//sorts data lexicographically by room
func sortAlphabeticalRoom(d *tabData, n int) {
	var idx_min int
	var temp data
	for i := 0; i < n-1; i++ {
		idx_min = i
		for j := i + 1; j < n; j++ {
			if d[j].room > d[idx_min].room {
				idx_min = j
			}
		}
		temp = d[i]
		d[i] = d[idx_min]
		d[idx_min] = temp
	}
}

//sorts data lexicographically by name
func sortAlphabeticalName(d *tabData, n int) {
	var idx_min int
	var temp data
	for i := 0; i < n-1; i++ {
		idx_min = i
		for j := i + 1; j < n; j++ {
			if d[j].name > d[idx_min].name {
				idx_min = j
			}
		}
		temp = d[i]
		d[i] = d[idx_min]
		d[idx_min] = temp
	}
}

//Prints whole array
func showData(d tabData, n int) {
	fmt.Printf("No | %-16s | %-16s | Power | Duration (in minutes)\n", "Name", "Room")
	for i := 0; i < n; i++ {
		fmt.Printf("%-3d| %-16s | %-16s | %-5d | %-5d\n", i+1, d[i].name, d[i].room, d[i].watt, d[i].time)
	}
}

//Printing a single data instead of the whole array
func printSingles(a data) {
	fmt.Println("Name : ", a.name)
	fmt.Println("Room : ", a.room)
	fmt.Println("Power : ", a.watt)
	fmt.Println("Duration (in minutes) : ", a.time)
}

func showStatistics(d tabData, n int) {
	var total, total1 int
	for i := 0; i < n; i++ {
		total += d[i].time
		total1 += d[i].watt
	}
	fmt.Println("Average Usage Time :", total)
	fmt.Println("Average Energy Consumed :", total1)
}
