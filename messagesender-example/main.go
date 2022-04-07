package main

import "fmt"

// Declare a Telegram type which satisfies the MessageSender interface(String(string) string)
type Telegram struct {
}

func (telegram Telegram) Send(message string) string {
	return message + " mesajını telegram ile gönderdim"
}

// Declare a Discord type which satisfies the MessageSender interface(String(string) string)
type Discord struct {
}

func (discord Discord) Send(message string) string {
	return message + " mesajını telegram ile gönderdim"
}

// Declare a Slack type which satisfies the MessageSender interface(String(string) string)
type Slack struct {
}

func (slack Slack) Send(message string) string {
	return message + " mesajını slack ile gönderdim"
}

// Declare MessageSender Interface
type MessageSender interface {
	Send(string) string
}

// Declare a messageWriter function which takes any object that satisfies
// the MessageSender interface as a parameter.
func messageWriter(ms MessageSender, message string) {
	fmt.Println(ms.Send(message))
}

func main() {
	//Initialize a Telegram, Discord and Slack object and pass them to MessageWriter()
	telegram := Telegram{}
	discord := Discord{}
	slack := Slack{}

	messageWriter(telegram, "Selam Dostum")
	messageWriter(discord, "Selam Kanka")
	messageWriter(slack, "Selam Bro")
}
