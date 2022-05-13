package main

import (
	"flag"
	"fmt"
	tmp "abc"
	tmp_2 "abc/ss/tmp_2"
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
	tmp.Run2()
	tmp_2.Run3()
}
