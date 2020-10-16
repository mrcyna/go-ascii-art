package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	args := os.Args[1:]

	sentence := strings.Join(args, " ")

	border := strings.Repeat("-", len(sentence)+2)

	fmt.Println(`
       ////\\\\
       |      |
      @  O  O  @                    +` + border + `+
       |  ~   |         \__         | ` + sentence + ` |
        \ -- /          |\ |        +` + border + `+
      ___|  |___        | \|
     /          \      /|__|
    /            \    / /
   /  /| .  . |\  \  / /
  /  / |      | \  \/ /
 <  <  |      |  \   /
  \  \ |  .   |   \_/
   \  \|______|
     \_|______|
       |      |
       |  |   |
       |  |   |
       |__|___|
       |  |  |
       (  (  |
       |  |  |
       |  |  |
      _|  |  |
  cccC_Cccc___)
 `)
}
