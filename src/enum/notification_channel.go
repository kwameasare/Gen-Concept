package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type NotificationChannel int

const (
	Email NotificationChannel = iota
	SMS
	PushNotification
	InApp
	VoiceCall
	Slack
	MicrosoftTeams
	WhatsApp
	Telegram
	Webhook
	PagerDuty
	Discord
	Signal
	Line
	WeChat
	Viber
	Messenger
	Fax
)

func (n NotificationChannel) String() string {
	names := [...]string{
		"EMAIL",
		"SMS",
		"PUSH_NOTIFICATION",
		"IN_APP",
		"VOICE_CALL",
		"SLACK",
		"MICROSOFT_TEAMS",
		"WHATSAPP",
		"TELEGRAM",
		"WEBHOOK",
		"PAGERDUTY",
		"DISCORD",
		"SIGNAL",
		"LINE",
		"WECHAT",
		"VIBER",
		"MESSENGER",
		"FAX",
	}
	if n < Email || int(n) >= len(names) {
		return "UNKNOWN"
	}
	return names[n]
}

func (n NotificationChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.String())
}

func (n *NotificationChannel) UnmarshalJSON(data []byte) error {
	var channelStr string
	if err := json.Unmarshal(data, &channelStr); err != nil {
		return err
	}

	switch channelStr {
	case "EMAIL":
		*n = Email
	case "SMS":
		*n = SMS
	case "PUSH_NOTIFICATION":
		*n = PushNotification
	case "IN_APP":
		*n = InApp
	case "VOICE_CALL":
		*n = VoiceCall
	case "SLACK":
		*n = Slack
	case "MICROSOFT_TEAMS":
		*n = MicrosoftTeams
	case "WHATSAPP":
		*n = WhatsApp
	case "TELEGRAM":
		*n = Telegram
	case "WEBHOOK":
		*n = Webhook
	case "PAGERDUTY":
		*n = PagerDuty
	case "DISCORD":
		*n = Discord
	case "SIGNAL":
		*n = Signal
	case "LINE":
		*n = Line
	case "WECHAT":
		*n = WeChat
	case "VIBER":
		*n = Viber
	case "MESSENGER":
		*n = Messenger
	case "FAX":
		*n = Fax
	default:
		return fmt.Errorf("invalid NotificationChannel: %s", channelStr)
	}

	return nil
}

func (n NotificationChannel) Value() (driver.Value, error) {
	return n.String(), nil
}

func (n *NotificationChannel) Scan(value interface{}) error {
	if value == nil {
		*n = Email
		return nil
	}

	var channelStr string

	switch v := value.(type) {
	case string:
		channelStr = v
	case []byte:
		channelStr = string(v)
	default:
		return fmt.Errorf("unsupported Scan type for NotificationChannel: %T", value)
	}

	switch channelStr {
	case "EMAIL":
		*n = Email
	case "SMS":
		*n = SMS
	case "PUSH_NOTIFICATION":
		*n = PushNotification
	case "IN_APP":
		*n = InApp
	case "VOICE_CALL":
		*n = VoiceCall
	case "SLACK":
		*n = Slack
	case "MICROSOFT_TEAMS":
		*n = MicrosoftTeams
	case "WHATSAPP":
		*n = WhatsApp
	case "TELEGRAM":
		*n = Telegram
	case "WEBHOOK":
		*n = Webhook
	case "PAGERDUTY":
		*n = PagerDuty
	case "DISCORD":
		*n = Discord
	case "SIGNAL":
		*n = Signal
	case "LINE":
		*n = Line
	case "WECHAT":
		*n = WeChat
	case "VIBER":
		*n = Viber
	case "MESSENGER":
		*n = Messenger
	case "FAX":
		*n = Fax
	default:
		return fmt.Errorf("invalid NotificationChannel: %s", channelStr)
	}

	return nil
}
