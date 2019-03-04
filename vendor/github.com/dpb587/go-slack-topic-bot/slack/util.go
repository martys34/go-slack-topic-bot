package slack

import (
	"log"
	"os"

	"github.com/nlopes/slack"
	"github.com/pkg/errors"
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
