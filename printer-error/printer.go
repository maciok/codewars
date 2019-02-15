package print_kata

import "fmt"

func PrinterError(s string) string {
	errs := 0
	for _, c := range s {

		if 'm' < c {
			errs++
		}
	}

	return fmt.Sprintf("%d/%d", errs, len(s))
}
