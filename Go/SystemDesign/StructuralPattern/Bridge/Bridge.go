package main

import "fmt"


type MessageSender interface{
	SendMessage(content string)
}

type SMSSender struct{

}

func (sender *SMSSender) SendMessage(content string){
	fmt.Println("Sending message through SMS :",content)
}

type EmailSender struct{

}

func (sender *EmailSender) SendMessage(content string){
	fmt.Println("Sending message through Email :",content)
}

type Message struct{
	sender MessageSender
	content string
}

func NewMessage(content string,sender MessageSender) *Message{
	return &Message{
		sender: sender,
		content: content,
	}
}

func (msg *Message) Send(){
	msg.sender.SendMessage(msg.content)
	fmt.Println("Message Sent Sucessfully")
}


type TextMessage struct{
	Message
}

func NewTextMessage(sender MessageSender, content string) *TextMessage {
	return &TextMessage{
		Message: Message{
			sender:  sender,
			content: content,
		},
	}
}


func (m *TextMessage) Send() {
	m.sender.SendMessage(m.content)
}

type UrgentMessage struct {
	Message
}

func NewUrgentMessage(sender MessageSender, content string) *UrgentMessage {
	return &UrgentMessage{
		Message: Message{
			sender:  sender,
			content: content,
		},
	}
}

func (m *UrgentMessage) Send() {
	m.sender.SendMessage("[URGENT] " + m.content)
}


func main() {

	email := &EmailSender{}
	sms := &SMSSender{}

	m1 := NewTextMessage(
		email,
		"Hello there",
	)

	m2 := NewUrgentMessage(
		sms,
		"Server is down",
	)

	m1.Send()
	m2.Send()
}