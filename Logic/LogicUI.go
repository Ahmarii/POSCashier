package Logic

import (
	"time"

	"cogentcore.org/core/colors"
	"cogentcore.org/core/core"
	"cogentcore.org/core/styles"
)

func (m *MANAGER) UILogicSetup() {
	m.ServiceStatusLogic()
}

func (m *MANAGER) UILogicLoop() {
	for range time.Tick(time.Second) {
		m.ui.AsyncUpdateElement(&m.ui.RFID_StatusFrame.WidgetBase)
		m.ui.AsyncUpdateElement(&m.ui.RFID_Chooser.WidgetBase)
	}
}

func (m *MANAGER) ServiceStatusLogic() {
	RFID_TextStatus := m.ui.RFID_StatusFrame.ChildByName("TextFrame").AsTree().ChildByName("TextStatus").(*core.Text)
	RFID_BarStatus := m.ui.RFID_StatusFrame.ChildByName("StatusBarFrame").AsTree().ChildByName("BarStatus").(*core.Frame)
	RFID_Chooser := m.ui.RFID_StatusFrame.ChildByName("StatusBarFrame").AsTree().ChildByName("Chooser").(*core.Chooser)
	m.ui.RFID_StatusFrame.Updater(func() {
		if m.logic.RFIDstatus {
			RFID_TextStatus.SetText("ONLINE")
			RFID_BarStatus.Styler(func(s *styles.Style) { s.Background = colors.Scheme.Success.OnContainer })
		} else {
			RFID_TextStatus.SetText("OFFLINE")
			RFID_BarStatus.Styler(func(s *styles.Style) { s.Background = colors.Scheme.Error.OnContainer })

		}
	})

	RFID_Chooser.Updater(func() {
		if len(m.logic.PortNames) > 0 {
			println(12312)
			RFID_Chooser.Items = make([]core.ChooserItem, len(m.logic.PortNames))
			for i, port := range m.logic.PortNames {
				RFID_Chooser.Items[i] = core.ChooserItem{Value: port.Name}
				//println(i, port)
			}
		} else {
			RFID_Chooser.SetPlaceholder("No Ports")
			RFID_Chooser.Items = make([]core.ChooserItem, 0)
		}

	})
}
