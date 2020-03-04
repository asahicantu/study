// +build !solution

package singlepaxos

// Learner represents a learner as defined by the single-decree Paxos
// algorithm.
type Learner struct {
	//TODO(student): Task 2 and 3 - algorithm and distributed implementation
	// Add needed fields
	id        int
	nrOfNodes int

	learned map[int]Value
	val     Value
	rnd     Round

	learnIn chan Learn

	valueOut chan<- Value

	stop chan struct{}
}

// NewLearner returns a new single-decree Paxos learner. It takes the
// following arguments:
//
// id: The id of the node running this instance of a Paxos learner.
//
// nrOfNodes: The total number of Paxos nodes.
//
// valueOut: A send only channel used to send values that has been learned,
// i.e. decided by the Paxos nodes.
func NewLearner(id int, nrOfNodes int, valueOut chan<- Value) *Learner {
	//TODO(student): Task 2 and 3 - algorithm and distributed implementation
	return &Learner{
		id:        id,
		nrOfNodes: nrOfNodes,

		learned: make(map[int]Value),
		val:     ZeroValue,
		rnd:     NoRound,

		learnIn: make(chan Learn, 10),

		valueOut: valueOut,

		stop: make(chan struct{}),
	}
}

// Start starts l's main run loop as a separate goroutine. The main run loop
// handles incoming learn messages.
func (l *Learner) Start() {
	go func() {
		for {
			select {
			case lrn := <-l.learnIn:
				val, output := l.handleLearn(lrn)
				if output {
					l.valueOut <- val
				}
			case <-l.stop:
				return
			}
			//TODO(student): Task 3 - distributed implementation
		}
	}()
}

// Stop stops l's main run loop.
func (l *Learner) Stop() {
	//TODO(student): Task 3 - distributed implementation
	l.stop <- struct{}{}
}

// DeliverLearn delivers learn lrn to learner l.
func (l *Learner) DeliverLearn(lrn Learn) {
	//TODO(student): Task 3 - distributed implementation
	l.learnIn <- lrn
}

// Internal: handleLearn processes learn lrn according to the single-decree
// Paxos algorithm. If handling the learn results in learner l emitting a
// corresponding decided value, then output will be true and val contain the
// decided value. If handleLearn returns false as output, then val will have
// its zero value.
func (l *Learner) handleLearn(learn Learn) (val Value, output bool) {
	//TODO(student): Task 2 - algorithm implementation
	if learn.Rnd >= l.rnd {
		if (l.val != ZeroValue && learn.Val != l.val) || (learn.Rnd > l.rnd) {
			l.resetLearned()
			l.rnd = learn.Rnd
			l.val = learn.Val
		}
		l.learned[learn.From] = learn.Val
	}

	if l.consensus() {
		l.resetLearned()
		return l.val, true
	}
	return ZeroValue, false
}

//TODO(student): Add any other unexported methods needed.
func (l *Learner) consensus() bool {
	return len(l.learned) > l.nrOfNodes/2
}

func (l *Learner) resetLearned() {
	l.learned = make(map[int]Value)
}
