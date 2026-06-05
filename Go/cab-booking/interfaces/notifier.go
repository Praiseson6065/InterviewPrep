package interfaces

type Notifier interface{
	Notify(userId string,message string)error

}