package entities

import (
	"fmt"
)

// ProposerObj - struct for proposer
type ProposerObj struct {
	id            int
	proposeValue  Any
	proposeNumber int
	receives      chan Message
	acceptors     []chan Message
	learners      []chan Message
}

// Proposer - constructor
func Proposer(id int, val Any, receives chan Message,
	acceptors []chan Message, learners []chan Message) *ProposerObj {
	proposer := new(ProposerObj)
	proposer.id = id
	proposer.proposeValue = val
	proposer.proposeNumber = 0
	proposer.receives = receives
	proposer.acceptors = acceptors
	proposer.learners = learners
	return proposer
}

func broadcast(group []chan Message, msg Message) {
	for _, s := range group {
		s <- msg
	}
}

func (p *ProposerObj) prepare() {
	p.proposeNumber++
	msg := PrepareMessage(p.id, p.proposeNumber)
	fmt.Printf("Proposer %v: sending Prepare\n", p.id)
	broadcast(p.acceptors, msg)
}

func (p *ProposerObj) acceptRequest() {
	msg := AcceptRequestMessage(p.id, p.proposeNumber, p.proposeValue)
	fmt.Printf("Proposer %v: sending Accept request\n", p.id)
	broadcast(p.acceptors, msg)
}

// Run - runs the proposer
func (p *ProposerObj) Run() {
	fmt.Printf("Proposer %v: running\n", p.id)

	for true {
		// Prepare-Promise
		p.prepare()
		responded := make(map[int]bool)
		max := p.proposeNumber
		for len(responded) < len(p.acceptors)/2+1 {
			msg := <-p.receives
			switch msg.msgType {
			case Promise:
				responded[msg.sender] = true
				if msg.proposeNumber > max {
					p.proposeValue = msg.proposeValue
					max = msg.proposeNumber

				}
			default:
			}
		}

		// Accept
		p.acceptRequest()
		responded = make(map[int]bool)
		max = p.proposeNumber

		for len(responded) >= len(p.acceptors)/2+1 {
			msg := <-p.receives
			switch msg.msgType {
			case Accepted:
				fmt.Printf("Proposer %v\n", p.id)
				responded[msg.sender] = true
				if msg.proposeNumber > max {
					max = msg.proposeNumber

				}
			default:
			}

		}

		if p.proposeNumber == max {
			break
		}
		p.proposeNumber = max
	}
	p.learn()

}

// learn - learn the chosen value
func (p *ProposerObj) learn() {
	msg := LearnMessage(p.id, p.proposeValue)
	fmt.Printf("Proposer %v: sending Chosen\n", p.id)
	broadcast(p.learners, msg)
}
