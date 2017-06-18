package alert

import (
	"log"
	"fmt"
	"strings"

	"github.com/nlopes/slack"
)

const (
	// action is used for slack attachment action.
	actionSelect = "select"
	actionStart  = "start"
	actionCancel = "cancel"

	ChannelID = "C5VDCLP98"
	TeamID    = "T29BGV9CM"
	RobotID   = "U5W5JJ8K0"
)

var SL *SlackListener

type SlackListener struct {
	Client *slack.Client

	BotID     string
	ChannelID string
}

// ListenAndResponse listens slack events and response
// particular messages. It replies by slack message button.
func (s *SlackListener) ListenAndResponse() {
	rtm := s.Client.NewRTM()

	// Start listening slack events
	go rtm.ManageConnection()

	// Handle slack events
	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			if err := s.handleMessageEvent(ev); err != nil {
				log.Printf("[ERROR] Failed to handle message: %s", err)
			}
		}
	}
}

// handleMessageEvent handles message events.
func (s *SlackListener) handleMessageEvent(ev *slack.MessageEvent) error {

	// Only response in specific channel. Ignore else.
	if ev.Channel != s.ChannelID {
		log.Printf("%s %s", ev.Channel, ev.Msg.Text)
		return nil
	}

	// Only response mention to bot. Ignore else.
	if !strings.HasPrefix(ev.Msg.Text, fmt.Sprintf("<@%s> ", s.BotID)) {
		return nil
	}

	// Parse message
	m := strings.Split(strings.TrimSpace(ev.Msg.Text), " ")[1:]
	if len(m) == 0 || m[0] != "hey" {
		return fmt.Errorf("invalid message")
	}

	// value is passed to message handler when request is approved.
	attachment := slack.Attachment{
		Text:       "What Can I help you?",
		Color:      "#f9a41b",
		CallbackID: "help",
		Actions: []slack.AttachmentAction{
			{
				Name: actionSelect,
				Type: "select",
				Options: []slack.AttachmentActionOption{
					{
						Text:  "Reload the caddy",
						Value: "reload-caddy",
					},
					{
						Text:  "Show the net stat",
						Value: "show-netstat",
					},
				},
			},

			{
				Name:  actionCancel,
				Text:  "Cancel",
				Type:  "button",
				Style: "danger",
			},
		},
	}

	params := slack.PostMessageParameters{
		Attachments: []slack.Attachment{
			attachment,
		},
	}

	if _, _, err := s.Client.PostMessage(ev.Channel, "", params); err != nil {
		return fmt.Errorf("failed to post message: %s", err)
	}

	return nil
}

func SendMessage(text string) {

	params := slack.PostMessageParameters{}
	attachment := slack.Attachment{
		Pretext: "Task result",
		Text:    "some text",
		// Uncomment the following part to send a field too
		/*
			Fields: []slack.AttachmentField{
				slack.AttachmentField{
					Title: "a",
					Value: "no",
				},
			},
		*/
	}

	params.Attachments = []slack.Attachment{attachment}
	channelID, timestamp, err := SL.Client.PostMessage(SL.ChannelID, text, params)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}
