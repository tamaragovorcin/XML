package saga

import (
	"encoding/json"
)

const (
	UserChannel    string = "UserChannel"
	InteractionChannel    string = "InteractionChannel"
	ShippingChannel string = "ShippingChannel"
	ReplyChannel    string = "ReplyChannel"
	ServiceInteraction    string = "Interaction"
	ServiceUser    string = "User"
	ServiceShipping string = "Shipping"
	ActionStart     string = "Start"
	ActionDone      string = "DoneMsg"
	ActionError     string = "ErrorMsg"
	ActionRollback  string = "RollbackMsg"
	ActionDonee      string = "Done"
)

type Message struct {
	Service       string         `json:"service"` // na koji se naredni salje od orkestratora
	SenderService string         `json:"sender_service"` // poslao por na orkestrator
	Action        string         `json:"action"` //actionstart,rollback,done,error
	User           string    `json:"userId"`
	User2 		string			`json:"userId2"`
	Ok            bool           `json:"ok"`
}
func (m Message) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}
