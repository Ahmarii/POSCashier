package Logic

import (
	DATA "CorgentPos/Data"
	UI "CorgentPos/UI"
)

type MANAGER struct {
	data  DATA.DataManager
	ui    UI.UIManager
	logic LogicManager
}

func (m *MANAGER) InitProcess() {
	m.data = DATA.DataManager{}
	m.ui = UI.UIManager{}

	m.ui.InitUI()

	m.UILogicSetup()
	go m.UILogicLoop()
	// go func() {
	// 	time.Sleep(time.Second)
	// 	m.logic.RFIDstatus = true
	// 	m.logic.PortNames = append(m.logic.PortNames, &enumerator.PortDetails{Name: "GG"})
	// }()
	m.ui.RunUI()
}
