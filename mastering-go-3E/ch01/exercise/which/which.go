package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Uses: which.go argument")
		return
	}

	for i := 1; i < len(args); i++ {
		file := args[i]

		paths := os.Getenv("PATH")
		pathSplit := filepath.SplitList(paths)
		fmt.Println(file,":")
		for _,v := range pathSplit {
			path := path.Join(v,file)
			f,err := os.Stat(path)
			if err == nil {
				mode := f.Mode()

				if mode.IsRegular() {
					if mode&0111 != 0 {
						fmt.Println("   ",path)
					}
				}

			}
		}
		fmt.Println()
	}
}