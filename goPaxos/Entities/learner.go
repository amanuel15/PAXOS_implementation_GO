package entities

import "fmt"

// LearnerObj - struct for learner
type LearnerObj struct {
	id       int
	receives chan Message
}

// Learner - constructor
func Learner(id int, receives chan Message) *LearnerObj {
	l := new(LearnerObj)
	l.id = id
	l.receives = receives
	return l
}

// Run the learner
func (l *LearnerObj) Run() Any {
	var v Any = -1
	for v == -1 {
		msg := <-l.receives
		switch msg.msgType {
		case Chosen:
			v = msg.proposeValue
		default:
		}
	}
	fmt.Printf("Learner %v: Chosen %v\n", l.id, v)
	return v
}
