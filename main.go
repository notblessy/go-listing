/*
Copyright © 2023 BLESSY <frederichblessy@gmail.com>

*/
package main

import (
	"github.com/notblessy/go-listing/db"
)

func main() {
	psql := db.InitDB()
	defer db.CloseDB(psql)
}
