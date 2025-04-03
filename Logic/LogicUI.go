package Logic

import (
	"bytes"
	"image"
	"strconv"
	"time"

	"cogentcore.org/core/colors"
	"cogentcore.org/core/core"
	"cogentcore.org/core/cursors"
	"cogentcore.org/core/events"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/styles/abilities"
	"cogentcore.org/core/styles/units"
	"cogentcore.org/core/tree"
)

func (m *MANAGER) LogicUISetup() {
	m.ServiceComponetLogic()
	m.ItemComponetLogic()
	m.ActionComponetLogic()
}

func (m *MANAGER) LogicUILoop() {
	for range time.Tick(time.Second) {
		//RFID Async update
		m.ui.AsyncUpdateElement(&m.ui.RFID_StatusFrame.WidgetBase)
		m.ui.AsyncUpdateElement(&m.ui.RFID_Chooser.WidgetBase)
		if !m.logic.RFIDstatus {
			m.logic.PortNamesIndex = m.ui.RFID_Chooser.CurrentIndex
		}
	}
}

func (m *MANAGER) ServiceComponetLogic() {
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
			RFID_Chooser.Items = make([]core.ChooserItem, len(m.logic.PortNames))
			for i, port := range m.logic.PortNames {
				RFID_Chooser.Items[i] = core.ChooserItem{Value: port.Name}
			}
			if m.logic.RFIDstatus {
				RFID_Chooser.SetPlaceholder(m.logic.PortNames[m.logic.PortNamesIndex].Name)
			}
		} else {
			RFID_Chooser.SetPlaceholder("No Ports")
			RFID_Chooser.Items = make([]core.ChooserItem, 0)
		}

	})
}

func (m *MANAGER) ActionComponetLogic() {
	ItemsFrame := m.ui.ItemsFrame.Child(0).(*core.Frame)
	m.ui.ActionFrame.ChildByName("Refresh_Button").(*core.Button).OnClick(func(e events.Event) {
		ItemsFrame.DeleteChildren()
		m.data.DropItems()
		m.data.FetchAllProduct()
		ItemsFrame.Update()
	})
}

func (m *MANAGER) ItemComponetLogic() {
	ItemsFrame := m.ui.ItemsFrame.Child(0).(*core.Frame)

	ItemsFrame.Maker(func(p *tree.Plan) {
		tempItems := m.data.SortItems(m.data.Items, "name")

		// println("before build --------")
		// for item := range tempItems {
		// 	println(tempItems[item].ItemName)
		// }

		for id := range tempItems {
			// println("-------- building ")
			// println(tempItems[id].ItemName)
			// if id == 0 {
			// 	continue
			// }
			tree.AddAt(p, tempItems[id].ItemName, func(w *core.Frame) {
				//w.SetType(core.ButtonAction).SetText(m.data.Items[i].ItemName)
				w.Styler(func(s *styles.Style) {
					m.ui.SetSize(s, units.Pw(22), units.Dp(200))
					s.SetAbilities(true, abilities.Activatable, abilities.Focusable, abilities.Hoverable, abilities.DoubleClickable, abilities.TripleClickable)
					s.Cursor = cursors.Pointer
					s.Background = colors.Scheme.Secondary.Container
					s.Border.Width.Set(units.Dp(3))
					s.Border.Radius = styles.BorderRadiusExtraSmall
					s.Padding = styles.NewSideValues(units.Px(3), units.Px(3), units.Px(3), units.Px(3))
					s.Margin = styles.NewSideValues(units.Px(5))
					s.Align.Content = styles.Center
					s.Align.Items = styles.Center
				})

				photoFrame := m.ui.NewFrameZero(w, units.Pw(50), units.Ph(100))
				photoFrame.Styler(func(s *styles.Style) {
					s.Align.Content = styles.Center
					s.Align.Items = styles.Center
					s.Justify.Content = styles.Center
					s.Justify.Items = styles.Center
				})

				photoF := m.ui.NewFrameZero(photoFrame, units.Dp(120), units.Dp(120))
				photoF.Styler(func(s *styles.Style) {
					s.Background = colors.Scheme.Secondary.OnContainer
					s.Border.Width.Set(units.Dp(1))
					s.Margin = styles.NewSideValues(units.Px(5), units.Px(5), units.Px(5), units.Px(5))
					s.Align.Content = styles.Center
					s.Align.Items = styles.Center
				})

				photo := core.NewImage(photoF)
				photo.Styler(func(s *styles.Style) {
					s.Background = colors.Scheme.Secondary.OnContainer
					m.ui.SetZero(s)
					m.ui.SetSize(s, units.Pw(100), units.Ph(100))
				})
				//Decode Image
				imageTemp, _, _ := image.Decode(bytes.NewReader(m.data.ImageCache[tempItems[id].ItemImagePath]))
				photo.SetImage(imageTemp)

				textFrame := m.ui.NewFrameZero(w, units.Pw(50), units.Ph(100))
				textFrame.SetName("Info")
				textFrame.Styler(func(s *styles.Style) {
					s.Direction = styles.Column
					s.Display = styles.Flex
					s.Justify.Content = styles.Center
					s.Justify.Items = styles.Center
				})

				text := core.NewText(textFrame)
				text.SetText("ItemID: " + strconv.Itoa(tempItems[id].ItemID))
				text.Styler(func(s *styles.Style) {
					s.Cursor = cursors.Pointer
					s.SetAbilities(false, abilities.Selectable, abilities.DoubleClickable)
				})

				text0 := core.NewText(textFrame)
				text0.SetText(tempItems[id].ItemName)
				text0.Styler(func(s *styles.Style) {
					s.Cursor = cursors.Pointer
					s.SetAbilities(false, abilities.Selectable, abilities.DoubleClickable)
				})

				text1 := core.NewText(textFrame)
				text1.SetName("Stock")
				text1.SetText("Stock: " + strconv.Itoa(tempItems[id].ItemStock))
				text1.Styler(func(s *styles.Style) {
					s.Cursor = cursors.Pointer
					s.SetAbilities(false, abilities.Selectable, abilities.DoubleClickable)
				})

				text2 := core.NewText(textFrame)
				text2.SetText("Price: " + strconv.Itoa(tempItems[id].ItemPrice))
				text2.Styler(func(s *styles.Style) {
					s.Cursor = cursors.Pointer
					s.SetAbilities(false, abilities.Selectable, abilities.DoubleClickable)
				})

				w.OnClick(func(e events.Event) {
					m.data.AddCart(tempItems[id].ItemID, 1)
				})
				w.OnDoubleClick(func(e events.Event) {
					w.Send(events.Click, e)
				})
				w.On(events.TripleClick, func(e events.Event) {
					w.Send(events.Click, e)
				})
				//println("Pop")
			})

			// println("-------- after building ")
			// after := p.Children
			// for x := range after {
			// 	println(after[x].Name)
			// }
		}

	})
}
