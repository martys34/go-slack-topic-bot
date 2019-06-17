package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/dpb587/go-slack-topic-bot/message"
	"github.com/dpb587/go-slack-topic-bot/message/pairist"
	"github.com/martys34/go-slack-topic-bot/slack"
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

	interruptPair, _ := pivotalTeam.Message()

	if time.Now().Weekday() == time.Weekday(5) {
		err := slack.SendMessage(interruptPair,
			"don't forget to spin the feedback wheel! :fidgetspinner:\nhttps://tinyurl.com/credhubfeedback")

		if err != nil {
			log.Panicf("ERROR: %v", err)
		}
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
					message.Literal("PM: "+pmString),
				),
			),
		),
	).Message()

	if err != nil {
		log.Panicf("ERROR: %v", err)
	}

	log.Printf("DEBUG: expected message: %s", msg)

	err = slack.UpdateChannelTopic(os.Getenv(workspace+"_SLACK_CHANNEL"), os.Getenv(workspace+"_SLACK_TOKEN"), msg)
	if err != nil {
		log.Panicf("ERROR: %v", err)
	}

}
