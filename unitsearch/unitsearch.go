package asunitsearch

import (
	"github.com/go-chat-bot/bot"
	"net/http"
)

func unitsearch(command *bot.Cmd) (string, error) {

	// searchstring := command.RawArgs
	// unitID := "2366"
	unitID := command.RawArgs

	res, err := http.Get("http://masterunitlist.info/Unit/Card/" + unitID + "?skill=4")
	if err != nil {
		return "", err
	}

	return res.Request.URL.String(), nil
}

func init() {
	bot.RegisterCommand(
		"search",
		"searches for a unit in the masterunitlist.info",
		"unitname - the unit you search for",
		unitsearch)
	bot.RegisterCommand(
		"search@ASHelperBot",
		"searches for a unit in the masterunitlist.info",
		"unitname - the unit you search for",
		unitsearch)
}
