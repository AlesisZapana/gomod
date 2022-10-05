package main

import "fmt"

type INotificationFactory interface {
	EnviarNotificacion()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

type SMSNotification struct {
}

func (SMSNotification) EnviarNotificacion() {
	fmt.Println("Enviar Notificación por SMS")
}

func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

type SMSNotificationSender struct {
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "TWilio"
}

type EmailNotification struct {
}

func (EmailNotification) EnviarNotificacion() {
	fmt.Println("Enviando Notificacion por email")
}

type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "SES"
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

func getNotificationFactory(notificationType string) (INotificationFactory, error) {
	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	}

	if notificationType == "Email" {
		return &EmailNotification{}, nil
	}

	return nil, fmt.Errorf("No Notification type")
}

func enviarNotificacion(f INotificationFactory) {
	f.EnviarNotificacion()
}

func getMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("Email")

	enviarNotificacion(smsFactory)
	enviarNotificacion(emailFactory)

	getMethod(smsFactory)
	getMethod(emailFactory)
}
