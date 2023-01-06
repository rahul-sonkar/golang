package main
import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Need atleast one argument.")
		return
	}

	file := args[1]
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)
	for _,dir := range pathSplit {
		fullPath := filepath.Join(dir,file)
		// Does it exits?
		fileInfo,err := os.Stat(fullPath)
		if err == nil {
			mode := fileInfo.Mode()
			// It is a regular file?
			if mode.IsRegular() {
				// It is a executable?
				if mode&0111 != 0 {
					fmt.Println(fullPath)
				}
			}
		}
	}
}