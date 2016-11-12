package main

import (
	"net/http"
	"fmt"
)

func unitsearch(args []string) (string) {

	unitID := ""
	if len(args) > 0 {
		unitID = args[0]
	} else {
		return("no UnitID given")
	}

	res, err := http.Get("http://masterunitlist.info/Unit/Card/" + unitID + "?skill=4")
	if err != nil {
		return "Error: " + fmt.Sprintf("%v", err)
	}

	return res.Request.URL.String()
}
