package runner

import (
	"fmt"

	Entity "../Entities"
)

// Network - struct for network
type Network struct {
	proposers []*Entity.ProposerObj
	acceptors []*Entity.AccepterObj
	learners  []*Entity.LearnerObj
}

func makeChannels(n int) []chan Entity.Message {
	chans := make([]chan Entity.Message, n)
	for i := range chans {
		chans[i] = make(chan Entity.Message, 1024)
	}
	return chans
}

// CreateNetwork - creates a Paxos network
func CreateNetwork(nProposers, nAcceptors, nLearners int, vs []Entity.Any) *Network {
	cProposers := makeChannels(nProposers)
	cAcceptors := makeChannels(nAcceptors)
	cLearners := makeChannels(nLearners)

	n := new(Network)
	n.proposers = make([]*Entity.ProposerObj, nProposers)
	n.acceptors = make([]*Entity.AccepterObj, nAcceptors)
	n.learners = make([]*Entity.LearnerObj, nLearners)

	for i := range n.proposers {
		n.proposers[i] = Entity.Proposer(i, vs[i], cProposers[i], cAcceptors, cLearners)
	}

	for i := range n.acceptors {
		n.acceptors[i] = Entity.Accepter(i, cAcceptors[i], cProposers)
	}

	for i := range n.learners {
		n.learners[i] = Entity.Learner(i, cLearners[i])
	}

	return n
}

// RunInMain - run the network
func RunInMain(n *Network) {

	for _, a := range n.acceptors {
		go a.Run()
	}
	for _, p := range n.proposers {
		go p.Run()
	}
	if n.learners[0].Run() != n.learners[1].Run() {
		fmt.Print("Error")
	}

}
