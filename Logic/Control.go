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

	m.data.DataInit()
	m.ui.InitUI()
	m.LogicUISetup()

	go m.LogicUILoop()
	go m.data.FetchAllProduct()
	go m.logic.LogicAsync()

	m.ui.RunUI()
}
