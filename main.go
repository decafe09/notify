package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

var version = "1.0.0"

// Slack is basic struct
type Slack struct {
	Username    string       `json:"username"`
	IconEmoji   string       `json:"icon_emoji"`
	Channel     string       `json:"channel"`
	Mrkdwn      bool         `json:"mrkdwn"`
	Attachments []Attachment `json:"attachments"`
}

// Attachment is attachment slack message.
type Attachment struct {
	Color    string   `json:"color"`
	MrkdwnIn []string `json:"mrkdwn_in"`
	Fields   []Field  `json:"fields"`
}

// Field is part of attachment.
type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

func main() {

	slackWebhookURL := os.Getenv("SLACK_WEBHOOK_URL")
	if slackWebhookURL == "" {
		fmt.Println("env SLACK_WEBHOOK_URL is not set")
		return
	}

	var title, level, slackChannel string
	var showVersion, showHelp bool
	flag.StringVar(&title, "t", "", "set title")
	flag.StringVar(&level, "l", "#dddddd", "set level [good|warning|danger]")
	flag.StringVar(&slackChannel, "c", "", "notify channel")
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&showHelp, "h", false, "show help")
	flag.Parse()

	if len(os.Args) < 2 || showHelp {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "command | %s -t [title] -l [level] -c [channel]\n", os.Args[0])
		flag.PrintDefaults()
		return
	}

	if showVersion {
		fmt.Println("version:", version)
		return
	}

	if terminal.IsTerminal(0) {
		fmt.Fprintln(os.Stderr, "use pipe like. -h show help")
		return
	}

	if slackChannel == "" {
		fmt.Fprintln(os.Stderr, "channel is not selected. -h show help")
		return
	}

	content := make([]byte, 0, 1024)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		content = append(content, scanner.Text()...)
		content = append(content, "\n"...)
	}

	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
		return
	}

	field := Field{
		Title: title,
		Value: "```" + string(content) + "```",
		Short: false,
	}

	attachment := Attachment{
		Color:    level,
		MrkdwnIn: []string{"fields"},
		Fields:   []Field{field},
	}

	slack := Slack{
		Username:    "notify",
		IconEmoji:   ":loudspeaker:",
		Channel:     "#" + slackChannel,
		Mrkdwn:      true,
		Attachments: []Attachment{attachment},
	}

	params, err := json.Marshal(slack)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := http.PostForm(slackWebhookURL, url.Values{"payload": {string(params)}})
	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	fmt.Println(string(body))
}
