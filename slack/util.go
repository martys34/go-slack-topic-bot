package slack

import (
	"github.com/nlopes/slack"
	"github.com/pkg/errors"
	"log"
	"os"
	"regexp"
	"strings"
)

func UpdateChannelTopic(channel, token, msg string) error {
	api := slack.New(token)

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

func SendMessage(interruptPair string, message string) error {
	api := slack.New(os.Getenv("PIVOTAL_SLACK_TOKEN"))

	msgOptions := slack.MsgOptionCompose(
		slack.MsgOptionText(interruptPair + " " + message, false),
		slack.MsgOptionAsUser(true))

	trimmedInterruptPair := regexp.MustCompile("[^a-zA-Z0-9 ]*").ReplaceAllString(interruptPair, "")
	people := strings.Fields(trimmedInterruptPair)

	for person := range people {
		_, err := api.PostEphemeral(os.Getenv("PIVOTAL_SLACK_CHANNEL"), people[person], msgOptions)

		if err != nil {
			return errors.Wrap(err, "sending github reminder")
		}
	}

	return nil
}
