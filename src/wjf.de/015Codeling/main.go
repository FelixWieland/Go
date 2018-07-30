package main

import (
	"fmt"

	"wjf.de/015Codeling/lib/live"

	dtbs "wjf.de/015Codeling/lib/database"
)

func main() {
	fmt.Print(dtbs.Export)
	fmt.Print(live.Demo)
}
