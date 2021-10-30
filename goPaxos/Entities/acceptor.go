package entities

import "fmt"

// AccepterObj - struct for accepter
type AccepterObj struct {
	id                  int
	acceptProposeNumber int
	acceptProposeValue  Any
	maxProposeNumber    int
	receives            chan Message
	proposers           []chan Message
}

// Accepter - constructor
func Accepter(id int, receives chan Message, proposers []chan Message) *AccepterObj {
	a := new(AccepterObj)
	a.id = id
	a.acceptProposeNumber = 0
	a.acceptProposeValue = nil
	a.maxProposeNumber = -1
	a.receives = receives
	a.proposers = proposers
	return a
}

// Run - the acceptor
func (a *AccepterObj) Run() {
	fmt.Printf("Acceptor %v: started\n", a.id)

	for {
		msg := <-a.receives
		switch msg.msgType {
		// Prepare-Promise
		case Prepare:
			if msg.proposeNumber > a.maxProposeNumber {
				a.maxProposeNumber = msg.proposeNumber
				a.proposers[msg.sender] <- PromiseMessage(a.id, a.acceptProposeNumber, a.acceptProposeValue)
				fmt.Printf("Acceptor %v: sending Promise\n", a.id)
			}

		// Accept-Accepted
		case Accept:
			if msg.proposeNumber >= a.maxProposeNumber {
				a.maxProposeNumber = msg.proposeNumber
				a.acceptProposeNumber = msg.proposeNumber
				a.acceptProposeValue = msg.proposeValue

			}
			a.proposers[msg.sender] <- AcceptedMessage(a.id, a.maxProposeNumber)

			fmt.Printf("Acceptor %v: sending Accepted\n", a.id)
		default:
		}
	}

}
