package Logic

import (
	"go.bug.st/serial/enumerator"
)

type LogicManager struct {
	RFIDstatus     bool
	PortNames      []*enumerator.PortDetails
	PortNamesIndex int
}

func (logic *LogicManager) LogicAsync() {
	go logic.RFID()
}
