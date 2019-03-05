package slack

import (
	"github.com/nlopes/slack"
	"github.com/pkg/errors"
	"log"
	"os"
	"regexp"
	"strings"
)

func UpdateChannelTopic(channel, msg string) error {
	api := slack.New(os.Getenv("SLACK_TOKEN"))

	channelInfo, err := api.GetChannelInfo(channel)
	if err != nil {
		return errors.Wrap(err, "getting channel info")
	}

	log.Printf("DEBUG: current topic: %s", channelInfo.Topic.Value)

	if channelInfo.Topic.Value == msg {
		log.Printf("DEBUG: no change needed")

		return nil
	}

	newTopic, err := api.SetChannelTopic(channel, msg)
	if err != nil {
		return errors.Wrap(err, "setting topic")
	}

	log.Printf("INFO: updated topic: %s", newTopic)

	return nil
}

func SendGitHubReminder(interruptPair string) error {
	api := slack.New(os.Getenv("SLACK_TOKEN"))

	msgOptions := slack.MsgOptionCompose(
		slack.MsgOptionText(interruptPair + " gentle reminder to check GitHub issues 😊", false),
		slack.MsgOptionAsUser(true))

	trimmedInterruptPair := regexp.MustCompile("[^a-zA-Z0-9 ]*").ReplaceAllString(interruptPair, "")
	people := strings.Fields(trimmedInterruptPair)

	for person := range people {
		_, err := api.PostEphemeral(os.Getenv("SLACK_CHANNEL"), people[person], msgOptions)

		if err != nil {
			return errors.Wrap(err, "sending github reminder")
		}
	}

	return nil
}
