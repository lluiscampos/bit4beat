package main

import (
	"fmt"
	"os"

	"github.com/lluiscampos/bit4beat.back/store/filestore"
)

func main() {
	doMain(os.Args)
}

func doMain(args []string) {
	fmt.Println("getting record")
	store := filestore.NewStore("/dummy/dir")
	r, err := store.GetRecord(42)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(r.ID)
}
