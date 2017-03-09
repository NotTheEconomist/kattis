package main

import (
	"fmt"
	"time"
)

var (
	SCHOOLDATE_CUTOFF = time.Date(2009, 12, 31, 23, 59, 56, 999, time.Local)
	DOB_CUTOFF        = time.Date(1990, 12, 31, 23, 59, 56, 999, time.Local)
	COURSE_CUTOFF     = 40
)

func main() {
	var N int // number of test cases
	fmt.Scanln(&N)

	for n := 0; n < N; n++ {
		var name, sschooldate, sdob string
		var numCourses int
		fmt.Scanln(&name, &sschooldate, &sdob, &numCourses)
		schooldate, err := time.Parse("2006/01/02", sschooldate)
		if err != nil {
			panic("Bad school date")
		}
		dob, err := time.Parse("2006/01/02", sdob)
		if err != nil {
			panic("Bad DOB")
		}
		switch {
		case schooldate.After(SCHOOLDATE_CUTOFF):
			fmt.Printf("%s eligible\n", name)
		case dob.After(DOB_CUTOFF):
			fmt.Printf("%s eligible\n", name)
		case numCourses > COURSE_CUTOFF:
			fmt.Printf("%s ineligible\n", name)
		default:
			fmt.Printf("%s coach petitions\n", name)
		}
	}
}
