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
			fmt.Printf("Search by:\nA. Name\nB.Room\nOption: ")
			fmt.Scan(&c2)
			fmt.Printf("Search for: ")
			fmt.Scan(&search)
			searchData(search)
		case "C":
			fmt.Printf("Sort by :\nA. Highest Energy Consumption\nB. Alphabetical Order\nC.Back\nOption : ")
			for c2 != "C" {
				fmt.Scan(&c2)
				switch c2 {
				case "A":
					sortEnergy()
				case "B":
					sortAlphabetical()
				case "C":
					fmt.Printf("A. Edit\nB. Search\nC. Sort\nD. Show Statistics\nE. Stop\nOption : ")
				default:
					fmt.Printf("Option doesnt exists. Please re-enter option : ")
				}
			}
		case "D":
			showData(d, n)
			callMenu(0)
			fmt.Printf("Option : ")
		case "E":
			fmt.Printf("Goodbye!")
		default:
			fmt.Printf("Option doesnt exists. Please re-enter option : ")
		}
	}
}

func callMenu(menu int) {
	if menu == 0 {
		fmt.Printf("\nA. Edit\nB. Search\nC. Sort\nD. Show Statistics\nE. Stop\n")
	} else if menu == 1 {
		fmt.Printf("EDIT :\nA. Add Data\nB. Change Data\nC. Delete Data\nD. Back\nOption : ")
	}
}

func addData(d *tabData, n *int) {
	fmt.Printf("Device name: ")
	fmt.Scan(&d[*n].name)
	fmt.Printf("Location: ")
	fmt.Scan(&d[*n].room)
	fmt.Printf("Power: ")
	fmt.Scan(&d[*n].watt)
	fmt.Printf("Duration: ")
	fmt.Scan(&d[*n].time)
	*n++
}

func changeData(d *tabData, n int) {
	fmt.Println("Which data do you want you change?")
	showData(*d,n)
	var s string
	fmt.Scan(&s)
	searchData(s)

	fmt.Printf("Data Changed!")
}

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

func searchData(s string) {
	fmt.Print(s)
}
func sortEnergy() {
	fmt.Printf("Data Sorted! (Energy)")
}

func sortAlphabetical() {
	fmt.Printf("Data Sorted Alphabetically!")
}

func showData(d tabData, n int) {
	fmt.Printf("No | %-16s | %-16s | Power | Duration\n", "Name", "Room")
	for i := 0; i < n; i++ {
		fmt.Printf("%-3d| %-16s | %-16s | %-5d | %-5d\n", i+1, d[i].name, d[i].room, d[i].watt, d[i].time)
	}
}

func showStatistics(d *tabData, n *int) {

}
