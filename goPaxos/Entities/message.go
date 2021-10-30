package entities

// Any - define any type
type Any interface{}

// MessageType - classify msg
type MessageType int

// Enums
const (
	Prepare MessageType = iota
	Promise
	Accept
	Accepted
	Chosen
)

// Message - between entites
type Message struct {
	msgType       MessageType
	sender        int
	proposeNumber int
	proposeValue  Any
}

// PrepareMessage - creates prepare
func PrepareMessage(from, pn int) Message {
	return Message{msgType: Prepare, sender: from, proposeNumber: pn}
}

// PromiseMessage - creates promise
func PromiseMessage(from, pNum int, pValue Any) Message {
	return Message{msgType: Promise, sender: from, proposeNumber: pNum, proposeValue: pValue}
}

// AcceptRequestMessage - creates accept from proposer
func AcceptRequestMessage(from int, pNum int, pValue Any) Message {
	return Message{msgType: Accept, sender: from, proposeNumber: pNum, proposeValue: pValue}
}

// AcceptedMessage - saves accepted message
func AcceptedMessage(from, pNum int) Message {
	return Message{msgType: Accepted, sender: from, proposeNumber: pNum}
}

// LearnMessage - to broadcast to listeners
func LearnMessage(from int, pVal Any) Message {
	return Message{msgType: Chosen, sender: from, proposeValue: pVal}
}
