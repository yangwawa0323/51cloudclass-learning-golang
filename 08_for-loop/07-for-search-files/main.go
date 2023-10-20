// javascript Array.map( (index, element) )
// golang for   range

package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	root := "c:\\windows\\system32"

	fileSystem := os.DirFS(root)

	matches, _ := fs.Glob(fileSystem, "*.dll")

	// matches : [ xboxgipsynthetic.dll,
	//             xmlfilter.dll  ,
	// 			   xmllite.dll,...
	//           ]

	// range ->  index, element
	for index, file := range matches {
		fmt.Println(index, ": ", file)
	}

}
