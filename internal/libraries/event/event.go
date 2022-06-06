package event

// Name ..
// All custom events names must be of this Name type.
type Name string

// LoanLogs ..
type LoanLogs struct {
	Request  interface{}
	Response interface{}
}

// Event ..
// All custom event types must satisfy this Event interface.
type Event interface {
	Handle()
}
