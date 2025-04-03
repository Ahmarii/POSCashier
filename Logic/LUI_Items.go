package Logic

import (
	"bytes"
	"image"
	"strconv"

	"cogentcore.org/core/colors"
	"cogentcore.org/core/core"
	"cogentcore.org/core/cursors"
	"cogentcore.org/core/events"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/styles/abilities"
	"cogentcore.org/core/styles/units"
	"cogentcore.org/core/tree"
)

func (m *MANAGER) ItemsComponetLogic() {
	ItemsFrame := m.ui.ItemsFrame.Child(0).(*core.Frame)
	ItemsFrame.Maker(func(p *tree.Plan) {
		tempItems := m.data.SortItems(m.data.Items, "name")
		for id := range tempItems {
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

				photoLayout := m.ui.NewFrameZero(w, units.Pw(50), units.Ph(100))
				photoLayout.Styler(func(s *styles.Style) {
					s.Align.Content = styles.Center
					s.Align.Items = styles.Center
					s.Justify.Content = styles.Center
					s.Justify.Items = styles.Center
				})

				photoFrame := m.ui.NewFrameZero(photoLayout, units.Dp(120), units.Dp(120))
				photoFrame.Styler(func(s *styles.Style) {
					s.Background = colors.Scheme.Secondary.OnContainer
					s.Border.Width.Set(units.Dp(1))
					s.Margin = styles.NewSideValues(units.Px(5), units.Px(5), units.Px(5), units.Px(5))
					s.Align.Content = styles.Center
					s.Align.Items = styles.Center
				})

				photo := core.NewImage(photoFrame)
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

				ItemID := core.NewText(textFrame)
				ItemID.SetText("ItemID: " + strconv.Itoa(tempItems[id].ItemID))
				ItemID.Styler(func(s *styles.Style) {
					s.Cursor = cursors.Pointer
					s.SetAbilities(false, abilities.Selectable, abilities.DoubleClickable)
				})

				ItemName := core.NewText(textFrame)
				ItemName.SetText(tempItems[id].ItemName)
				ItemName.Styler(func(s *styles.Style) {
					s.Cursor = cursors.Pointer
					s.SetAbilities(false, abilities.Selectable, abilities.DoubleClickable)
				})

				ItemStock := core.NewText(textFrame)
				ItemStock.SetName("Stock")
				ItemStock.SetText("Stock: " + strconv.Itoa(tempItems[id].ItemStock))
				ItemStock.Styler(func(s *styles.Style) {
					s.Cursor = cursors.Pointer
					s.SetAbilities(false, abilities.Selectable, abilities.DoubleClickable)
				})

				ItemPrice := core.NewText(textFrame)
				ItemPrice.SetText("Price: " + strconv.Itoa(tempItems[id].ItemPrice))
				ItemPrice.Styler(func(s *styles.Style) {
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
			})
		}
	})
}
