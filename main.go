package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/yookoala/realpath"
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/php7"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/visitor"
)

func main() {
	var nodes node.Node
	var comments comment.Comments
	var positions position.Positions

	usePhp5 := flag.Bool("php5", false, "use PHP5 parser")
	flag.Parse()

	for _, path := range flag.Args() {
		real, err := realpath.Realpath(path)
		checkErr(err)

		err = filepath.Walk(real, func(path string, f os.FileInfo, err error) error {
			if !f.IsDir() && filepath.Ext(path) == ".php" {
				fmt.Printf("==> %s\n", path)

				src, _ := os.Open(string(path))
				if *usePhp5 {
					nodes, comments, positions = php5.Parse(src, path)
				} else {
					nodes, comments, positions = php7.Parse(src, path)
				}

				visitor := visitor.Dumper{
					Indent:    "  | ",
					Comments:  comments,
					Positions: positions,
				}
				nodes.Walk(visitor)
			}
			return nil
		})
		checkErr(err)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
