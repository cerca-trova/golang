package main

import "fmt"

type Filter interface {
	About() string
	Process([]int) []int
}

type uniqueFilter struct{}

func (uf uniqueFilter) About() string {
	return "remove duplicate numbers"
}
func (uf uniqueFilter) Process(s []int) []int {
	slice_new := make([]int, 0, len(s)) // create a new slice and return it
	pusheds := make(map[int]bool)

	for _, i := range s {
		if !pusheds[i] {
			pusheds[i] = true
			slice_new = append(slice_new, i)
		}
	}
	return slice_new
}

type multiFilter int

func (mf multiFilter) About() string {
	return fmt.Sprintf("keep multiples of %v", mf)
}

func (mf multiFilter) Process(s []int) []int {
	slice_fllterd := make([]int, 0, len(s))
	for _, i := range s {
		if i%int(mf) == 0 {
			slice_fllterd = append(slice_fllterd, i)
		}
	}

	return slice_fllterd
}

func filterAndPrint(fltr Filter, unfilterd []int) []int {
	filtered := fltr.Process(unfilterd)
	fmt.Println(fltr.About()+":\n\t", filtered)
	return filtered
}

func main() {
	numbers := []int{12, 7, 21, 12, 12, 26, 25, 21, 30}
	fmt.Println("before filtering:\n\t", numbers)

	filters := []Filter{
		uniqueFilter{},
		multiFilter(2),
		multiFilter(3),
	}

	for _, fltr := range filters {
		numbers = filterAndPrint(fltr, numbers)
	}
}
