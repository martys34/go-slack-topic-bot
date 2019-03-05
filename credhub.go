package main

import (
	"github.com/dpb587/go-slack-topic-bot/message"
	"github.com/dpb587/go-slack-topic-bot/message/pairist"
	"github.com/martys34/go-slack-topic-bot/slack"
	"log"
	"os"
)

func main() {
	team := pairist.PeopleInRole{
		Team: "therealslimcredhub",
		Role: "Bat-person",
		People: map[string]string{
			"Andrew":   "U8RH9A30R",
			"Mark":     "U05661PTK",
			"Marty":    "UBXEUQN6P",
			"Walter":   "U1L7Q0LDR",
			"Frances":  "UF5N1RPQB",
			"Josh":     "U1V6C2UKW",
			"Tom":      "UG262F4EB",
			"Victoria": "U6SUTRCKB",
		},
	}

	msg, err := message.Join(
		" | ",
		message.Literal("see pinned messages for helpful links"),
		message.Prefix(
			"interrupt: ",
			message.Conditional(
				pairist.WorkingHours("09:00", "18:00", "America/New_York"),
				message.Join(
					" ",
					team,
					message.Literal("| break glass: `@credhub-team` |"),
					message.Literal("PMs: <@U0DFF9JLB>, <@U27926NR3>"),
				),
			),
		),
	).Message()

	if err != nil {
		log.Panicf("ERROR: %v", err)
	}

	log.Printf("DEBUG: expected message: %s", msg)

	err = slack.UpdateChannelTopic(os.Getenv("SLACK_CHANNEL"), msg)
	if err != nil {
		log.Panicf("ERROR: %v", err)
	}

	interruptPair, _ := team.Message()

	err = slack.SendGitHubReminder(interruptPair)

	if err != nil {
		log.Panicf("ERROR: %v", err)
	}
}


