package slack

type IncomingWebhook struct {
	Text string `json:text`
}

func GetLink(path string, text string) string {
	return "<" + path + "|" + text + ">"
}
