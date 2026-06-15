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
			fmt.Printf("Search by:\nA. Name\nB.Room\nOption: ")
			fmt.Scan(&c2)
			fmt.Printf("Search for: ")
			fmt.Scan(&search)
			searchData(search)
		case "C":
			callMenu(2)
			for c2 != "C" {
				fmt.Scan(&c2)
				switch c2 {
				case "A":
					sortEnergy(&d, n)
					callMenu(2)
				case "B":
					sortAlphabetical()
					callMenu(2)
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

func isiData(d *tabData, n *int) {
	d[0].name = "Rice_Cooker"
	d[0].room = "Kitchen"
	d[0].watt = 360
	d[0].time = 720
	*n++
}

func callMenu(menu int) {
	if menu == 0 {
		fmt.Printf("\nA. Edit\nB. Search\nC. Sort\nD. Show Statistics\nE. Stop\n")
	} else if menu == 1 {
		fmt.Printf("EDIT :\nA. Add Data\nB. Change Data\nC. Delete Data\nD. Back\nOption : ")
	} else if menu == 2 {
		fmt.Printf("Sort by :\nA. Highest Energy Consumption\nB. Alphabetical Order\nC. Back\nOption : ")
	}
}

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
}

func changeData(d *tabData, n int) {
	fmt.Println("Which data do you want you change?")
	showData(*d, n)
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

func showData(d tabData, n int) {
	fmt.Printf("No | %-16s | %-16s | Power | Duration (in minutes)\n", "Name", "Room")
	for i := 0; i < n; i++ {
		fmt.Printf("%-3d| %-16s | %-16s | %-5d | %-5d\n", i+1, d[i].name, d[i].room, d[i].watt, d[i].time)
	}
}

func showStatistics(d *tabData, n *int) {

}
