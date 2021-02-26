package main

import (
	"fmt"
	"os"

	"github.com/lluiscampos/bit4beat.back/api"
	"github.com/lluiscampos/bit4beat.back/store/filestore"
)

func main() {
	doMain(os.Args)
}

func doMain(args []string) {
	fmt.Println("getting record")
	store := filestore.NewStore("/tmp/dummy-dir")

	api := api.NewApi(store)
	api.Serve()
}
