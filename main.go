package main

import (
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("a([a-z]+)e", "apple")
	fmt.Println(match)

	r := regexp.MustCompile("a([a-z]+)e")
	ms := r.MatchString("apple")
	fmt.Println(ms)

	// s := "/view/test"
	r2 := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
	fs := r2.FindString("/view/test")
	fmt.Println(fs)
	fss := r2.FindStringSubmatch("/view/test")
	fmt.Println(fss)
	fmt.Println(fss[0])
	fmt.Println(fss[1])
	fmt.Println(fss[2])
	fss2 := r2.FindStringSubmatch("/edit/test")
	fmt.Println(fss2)
	fmt.Println(fss2[0])
	fmt.Println(fss2[1])
	fmt.Println(fss2[2])
	fss3 := r2.FindStringSubmatch("/save/test")
	fmt.Println(fss3)
	fmt.Println(fss3[0])
	fmt.Println(fss3[1])
	fmt.Println(fss3[2])

}
