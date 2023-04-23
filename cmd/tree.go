package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	levelFlag []bool
	fileCount,
	dirCount int
)

const (
	space  = "   "
	line   = "│  "
	last   = "└─ "
	middle = "├─ "
)

func ShowPathTree(path string, subDepth int) {
	levelFlag = make([]bool, subDepth)
	fmt.Println("<.>")
	walk(path, 0)
	str := fmt.Sprintf("\nThis Path have %d dirs,%d files.", dirCount, fileCount)
	fmt.Println(str)

}

func walk(dir string, level int) {
	levelFlag[level] = true
	if files, err := os.ReadDir(dir); err == nil {
		for fi := range files {
			absFile := filepath.Join(dir, files[fi].Name())
			isLast := fi == len(files)-1
			levelFlag[level] = !isLast
			showLine(level, isLast, files[fi])
			if files[fi].IsDir() {
				walk(absFile, level+1)
			}
		}
	} else {
		panic(err)
	}
}

func showLine(level int, isLast bool, info fs.DirEntry) {
	preFix := buildPrefix(level)
	outTemp, out := "%s%s%s", ""
	fName := info.Name()
	if info.IsDir() {
		fName = fmt.Sprintf("<%s>", fName)
		dirCount++
	} else {
		fileCount++
	}
	if isLast {
		out = fmt.Sprintf(outTemp, preFix, last, fName)
	} else {
		out = fmt.Sprintf(outTemp, preFix, middle, fName)
	}
	fmt.Println(out)
}

func buildPrefix(level int) string {
	result := ""
	for idx := 0; idx < level; idx++ {
		if levelFlag[idx] {
			result += line
		} else {
			result += space
		}
	}
	return result
}

var treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "File Trees.",
	Long:  `File Trees.`,
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		} else {
			// fmt.Println(dir)
			ShowPathTree(dir, 100)
		}
	},
}

func init() {
	rootCmd.AddCommand(treeCmd)
}
