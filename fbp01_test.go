package gofbp

import (
	"log"
	"math"
	"testing"

	"github.com/jllopis/gofbp/components"
)

func Test01(t *testing.T) {
	// --- define network ---
	network := NewNetwork()

	sender := network.DefProc(&components.Sender{})
	copier := network.DefProc(&components.Copier{})
	recvr := network.DefProc(&components.Recvr{})

	network.Initialize(sender, "COUNT", math.Pow(10.0, 5.0))
	network.Connect(sender, "OUT", copier, "IN", 5)
	network.Connect(copier, "OUT", recvr, "IN", 5)

	// --- run ---
	network.Run(&RuntimeOpts{trace: true}, func() { log.Info("FINISHED") })
}
