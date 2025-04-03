package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		i1, i2, i3 int
		b1, b2, b3 bool
		s1, s2, s3 string
		n1, n2, n3 int
		err        error
	)
	r := strings.NewReader("5 true gophers")
	n1, err = fmt.Fscanf(r, "%d %t %s", &i1, &b1, &s1)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Fscanf 1 : %v\n", err)
	}
	fmt.Printf("%d,%t,%q,%d\n", i1, b1, s1, n1)
	r = strings.NewReader("5\ntrue gophers")
	n2, err = fmt.Fscanf(r, "%d %t %s", &i2, &b2, &s2)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Fscanf 2: %v\n", err)
	}
	fmt.Printf("%d,%t,%q,%d\n", i2, b2, s2, n2)

	r = strings.NewReader("5 true\ngophers")
	n3, err = fmt.Fscanf(r, "%d %t %s", &i3, &b3, &s3)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Fscanf 3: %v\n", err)
	}
	fmt.Printf("%d,%t,%q,%d\n", i3, b3, s3, n3)
}
