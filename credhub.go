package main

import (
	"log"
	"os"

	"github.com/dpb587/go-slack-topic-bot/message"
	"github.com/dpb587/go-slack-topic-bot/message/pairist"
	"github.com/martys34/go-slack-topic-bot/slack"
)

func main() {
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

	msg, err := message.Join(
		" | ",
		message.Literal("Please include your CredHub logs in case of Errors"),
		message.Prefix(
			"interrupt: ",
			message.Conditional(
				pairist.WorkingHours("09:00", "18:00", "America/New_York"),
				message.Join(
					" ",
					cloudFoundryTeam,
					message.Literal("| break glass: `@credhub-team` |"),
					message.Literal("PM: <@UDFK4K0KT>, <@UHPMJCXGC>"),
				),
			),
		),
	).Message()

	if err != nil {
		log.Panicf("ERROR: %v", err)
	}

	log.Printf("DEBUG: expected message: %s", msg)

	err = slack.UpdateChannelTopic(os.Getenv("CLOUDFOUNDRY_SLACK_CHANNEL"), os.Getenv("CLOUDFOUNDRY_SLACK_TOKEN"), msg)
	if err != nil {
		log.Panicf("ERROR: %v", err)
	}
}
