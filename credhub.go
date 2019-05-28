package main

import (
	"github.com/dpb587/go-slack-topic-bot/message"
	"github.com/dpb587/go-slack-topic-bot/message/pairist"
	"github.com/martys34/go-slack-topic-bot/slack"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	pivotalTeam := pairist.PeopleInRole{
		Team: "therealslimcredhub",
		Role: "Bat-person",
		People: map[string]string{
			"Andrew":   "U8RH9A30R",
			"Mark":     "U05661PTK",
			"Marty":    "UBXEUQN6P",
			"Josh":     "U1V6C2UKW",
			"Tom":      "UG262F4EB",
			"Victoria": "U6SUTRCKB",
		},
	}
	pivotalPM := []string {"<@U0DFF9JLB>", "<@U27926NR3>"}
	createChannelMessage("PIVOTAL",
		"see pinned messages for helpful links", pivotalTeam, pivotalPM)

	cloudFoundryTeam := pairist.PeopleInRole{
		Team: "therealslimcredhub",
		Role: "Bat-person",
		People: map[string]string{
			"Andrew":   "U8TDZ8VU3",
			"Mark":     "U02SQ5CJW",
			"Marty":    "UC1H82QF8",
			"Josh":     "U1YKRGMDZ",
			"Tom":      "UG4JCSF5H",
			"Victoria": "U6W2F82B1",
		},
	}

	cloudFoundryPM := []string {"<@UDFK4K0KT>", "<@UHPMJCXGC>"}
	createChannelMessage("CLOUDFOUNDRY",
		"Please include your CredHub logs in case of Errors", cloudFoundryTeam, cloudFoundryPM)

	interruptPair, _ := pivotalTeam.Message()
	githubReminder := "gentle reminder to check GitHub issues 😊"
	err := slack.SendMessage(interruptPair, githubReminder)

	if time.Now().Weekday() == time.Weekday(6) {
		err = slack.SendMessage(interruptPair,
			"don't forget to spin the feedback wheel! :fidgetspinner: \n https://tinyurl.com/credhubfeedback")
	}

	if err != nil {
		log.Panicf("ERROR: %v", err)
	}
}

func createChannelMessage(workspace, firstLine string, team pairist.PeopleInRole, PM []string) {
	pmString := strings.Join(PM, ", ")

	msg, err := message.Join(
		" | ",
		message.Literal(firstLine),
		message.Prefix(
			"interrupt: ",
			message.Conditional(
				pairist.WorkingHours("09:00", "18:00", "America/New_York"),
				message.Join(
					" ",
					team,
					message.Literal("| break glass: `@credhub-team` |"),
					message.Literal("PM: " + pmString),
				),
			),
		),
	).Message()

	if err != nil {
		log.Panicf("ERROR: %v", err)
	}

	log.Printf("DEBUG: expected message: %s", msg)

	err = slack.UpdateChannelTopic(os.Getenv(workspace + "_SLACK_CHANNEL"), os.Getenv(workspace + "_SLACK_TOKEN"), msg)
	if err != nil {
		log.Panicf("ERROR: %v", err)
	}

}


