package kafka

import (
	"math/rand"
	"strconv"
	"sync/atomic"
	"time"
)

// TicketIDAtomicCounter is a counter that adds when each time a function is called
var TicketIDAtomicCounter int64

// TicketIDGenerator is a random unique ticket ID generator, output is a string
func TicketIDGenerator(prefix string) Ticket {
	now := 10000000 + int(time.Now().UnixNano())%89999999
	atomic.AddInt64(&TicketIDAtomicCounter, 1)
	atomicCounter := 10000 + int(TicketIDAtomicCounter)%89999
	randomNumber := 10000 + rand.Intn(89999)
	truelyRandNumber := prefix + strconv.Itoa(atomicCounter) + strconv.Itoa(now) + strconv.Itoa(randomNumber)
	ticket := Ticket(truelyRandNumber)
	return ticket
}
