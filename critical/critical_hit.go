package ascritical

import (
	"fmt"
	"github.com/go-chat-bot/bot"
	"math/rand"
	//"regexp"
	"os"
	"strings"
)

//const (
//pattern = "(\\!|\\/)(critical|kritisch).*"
//)

//var (
//re = regexp.MustCompile(pattern)
//)

func criticalRoll(command *bot.Cmd) (string, error) {
	//if !re.MatchString(command.Raw) {
	//return "", nil
	//}

	// the following are the critical hit tables

	crit_mech := map[int]string{
		0:  "BattleMech",
		2:  "Ammo Hit",
		3:  "Engine Hit",
		4:  "Fire Control Hit",
		5:  "No Critical Hit",
		6:  "Weapon Hit",
		7:  "MP Hit",
		8:  "Weapon Hit",
		9:  "No Critical Hit",
		10: "Fire Control Hit",
		11: "Engine Hit",
		12: "Unit Destroyed",
	}

	crit_vehicle := map[int]string{
		0:  "Vehicle",
		2:  "Ammo Hit",
		3:  "",
		4:  "Fire Control Hit",
		5:  "Fire Control Hit",
		6:  "No Critical Hit",
		7:  "No Critical Hit",
		8:  "No Critical Hit",
		9:  "Weapon Hit",
		10: "Weapon Hit",
		11: "Crew Killed",
		12: "Engine Hit",
	}

	crit_protomech := map[int]string{
		0:  "ProtoMech",
		2:  "",
		3:  "",
		4:  "",
		5:  "",
		6:  "",
		7:  "",
		8:  "",
		9:  "",
		10: "",
		11: "",
		12: "",
	}

	crit_aerospace := map[int]string{
		0:  "Aerospace",
		2:  "Fuel Hit",
		3:  "Fire Control Hit",
		4:  "Engine Hit",
		5:  "Weapon Hit",
		6:  "No Critical Hit",
		7:  "No Critical Hit",
		8:  "No Critical Hit",
		9:  "Weapon Hit",
		10: "Engine Hit",
		11: "Fire Control Hit",
		12: "Crew Killed",
	}

	crit_dropship := map[int]string{
		0:  "Dropship",
		2:  "",
		3:  "",
		4:  "",
		5:  "",
		6:  "",
		7:  "",
		8:  "",
		9:  "",
		10: "",
		11: "",
		12: "",
	}

	crit_empty := map[int]string{
		0:  "None",
		2:  "",
		3:  "",
		4:  "",
		5:  "",
		6:  "",
		7:  "",
		8:  "",
		9:  "",
		10: "",
		11: "",
		12: "",
	}

	d1 := rand.Intn(5) + 1
	d2 := rand.Intn(5) + 1

	// roll 2D6
	roll := d1 + d2
	crit_hit := ""
	unitType := ""
	args := strings.Split(command.Raw, " ")
	if len(args) > 1 {
		unitType = args[1]
	}

	switch unitType {
	case "", "BattleMech", "Battlemech", "battlemech", "mech", "Mech", "bm", "BM":
		crit_hit = crit_mech[roll]
		unitType = crit_mech[0]
	case "CV", "Vehicle", "SV", "Fahrzeug", "cv":
		crit_hit = crit_vehicle[roll]
		unitType = crit_vehicle[0]
	case "ProtoMech", "Protomech", "protomech", "PM":
		crit_hit = crit_protomech[roll]
		unitType = crit_protomech[0]
	case "Aerospace", "aerospace", "Plane", "plane", "luft", "luft/raum", "jäger", "AS", "as":
		crit_hit = crit_aerospace[roll]
		unitType = crit_aerospace[0]
	case "Dropship", "dropship", "landungsschiff", "Landungsschiff", "ds", "DS":
		crit_hit = crit_dropship[roll]
		unitType = crit_dropship[0]
	default:
		crit_hit = crit_empty[roll]
		unitType = crit_empty[0]
	}

	return fmt.Sprintf("[ (%d + %d) = %d ] [ Type: %s ]\n----%s----", d1, d2, roll, unitType, crit_hit), nil
}

func init() {
	bot.RegisterCommand(
		"critical@"+os.Getenv("TELEGRAM_BOTNAME"),
		"Roll for a critical hit for a unittype",
		"unit-type",
		criticalRoll)
	bot.RegisterCommand(
		"critical",
		"Roll for a critical hit for a unittype",
		"unit-type",
		criticalRoll)
}
