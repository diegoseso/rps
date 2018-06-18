package server

import (
	"encoding/json"
	"log"
)

func routeIncomming(rawMessage []byte){

	message := &Message{}
    err := json.Unmarshal(rawMessage, message)
    if err != nil{
	    log.Fatal("Message not accepted")
	}

	message.TypeValidator()
	message.ActionValidator()
	action , err := message.ActionDecoder()
	if err != nil{

	}
	//action.Execute();
}

type Message struct{
	Type string    `json:"type" mapstructure:"type"`
	Action []byte  `json:"action" mapstructure:"action"`
}

func( m *Message)TypeValidator()error{
	return nil
}

func( m *Message)ActionValidator()error{
	return nil
}

func( m *Message)ActionDecoder()(*action, error){
	return nil, nil
}

type action interface {
	Execute()([]byte, error)
}


