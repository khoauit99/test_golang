package main

import (
	"flag"
	"fmt"

	tmp_1 "github.com/khoauit99/test_golang"
	tmp_2 "github.com/khoauit99/test_golang/ss/tmp_2"
)

func usage() func() {
	return func() {
		// fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Println("hello world")
		// flag.PrintDefaults()
	}
}

// func usage() func() {
// 	return func() {
// 		fmt.Println("hello world")
// 		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
// 		// fmt.Fprintf(os.Stderr, "%s\n", os.Args[0])
// 		// fmt.Fprintf(os.Stderr, "Version %s, Git %s, Git SHA %s, BuildDate %s\n", version, GitRef, GitSHA, BuildDate)
// 		// fmt.Fprintln(os.Stderr, "Repo https://github.com/zph/moresql")
// 		// fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
// 		flag.PrintDefaults()
// 	}
// }

func main() {
	flag.Usage = usage()
	flag.Usage()
	tmp_1.Run2()
	tmp_2.Run3()
}
