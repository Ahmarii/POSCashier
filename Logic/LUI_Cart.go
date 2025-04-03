package Logic

import (
	"bytes"
	"image"
	"strconv"

	"cogentcore.org/core/colors"
	"cogentcore.org/core/core"
	"cogentcore.org/core/cursors"
	"cogentcore.org/core/events"
	"cogentcore.org/core/icons"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/styles/abilities"
	"cogentcore.org/core/styles/units"
	"cogentcore.org/core/tree"
)

func (m *MANAGER) UpdateCart() {
	CartFrame := m.ui.CartFrame.Child(0).(*core.Frame)
	CartFrame.DeleteChildren()
	CartFrame.Update()
}
func (m *MANAGER) CartComponetLogic() {
	CartFrame := m.ui.CartFrame.Child(0).(*core.Frame)

	CartFrame.Maker(func(p *tree.Plan) {
		// Get the sorted ids
		tempOrders := m.data.SortOrder(m.data.Orders, "id")

		for tempId := range tempOrders {
			ItemID := tempOrders[tempId].OrderItemID
			ItemPrice := m.data.Items[ItemID].ItemPrice
			ItemName := m.data.Items[ItemID].ItemName
			OrderAmount := m.data.Orders[ItemID].OrderAmount

			println(tempOrders[tempId].OrderItemName) // Print the order name
			println(ItemID)                           // Print the item ID

			Text := ItemName + " " + strconv.Itoa(ItemPrice) + "x" + strconv.Itoa(OrderAmount)

			tree.AddAt(p, ItemName, func(w *core.Frame) {
				//w.SetType(core.ButtonAction)
				//w.SetText(Text) // Show item name and amount
				w.Styler(func(s *styles.Style) {
					s.Display = styles.Flex
					s.Background = colors.Scheme.Secondary.Container
					s.Border.Width.Set(units.Dp(1))
					s.Border.Radius.Zero()
					s.Min.Set(units.Pw(95), units.Dp(80))
					s.Padding = styles.NewSideValues(units.Px(0), units.Px(0), units.Px(0), units.Px(0))
					s.Margin = styles.NewSideValues(units.Px(5), units.Px(5), units.Px(0), units.Px(5))
					s.Align.Items = styles.Center
					s.Justify.Content = styles.SpaceEvenly
					s.Justify.Items = styles.SpaceEvenly
				})

				photo := core.NewImage(w)
				photo.Styler(func(s *styles.Style) {
					s.Background = colors.Scheme.Secondary.OnContainer
					s.Border.Width.Set(units.Dp(1))
					s.Border.Radius.Zero()
					s.Min.Set(units.Dp(60), units.Dp(60))
					s.Max.Set(units.Dp(60), units.Dp(60))

					s.Padding = styles.NewSideValues(units.Px(0), units.Px(0), units.Px(0), units.Px(0))
					s.Margin = styles.NewSideValues(units.Px(5), units.Px(5), units.Px(0), units.Px(5))
				})
				imageTemp, _, _ := image.Decode(bytes.NewReader(m.data.ImageCache[m.data.Items[ItemID].ItemImagePath]))
				photo.SetImage(imageTemp)

				textFrame := core.NewFrame(w)
				textFrame.Styler(func(s *styles.Style) {
					s.Background = colors.Scheme.Secondary.On
					s.Border.Width.Set(units.Dp(1))
					s.Border.Radius.Zero()
					s.Min.Set(units.Dp(160), units.Dp(20))
					s.Padding = styles.NewSideValues(units.Px(0), units.Px(2), units.Px(0), units.Px(2))
					s.Margin = styles.NewSideValues(units.Px(5), units.Px(5), units.Px(0), units.Px(5))
				})
				textFrame.SetName("TextF")
				text := core.NewText(textFrame)
				text.SetName("Text")
				text.SetText(Text)
				text.Styler(func(s *styles.Style) {
					s.Cursor = cursors.Pointer
					s.SetAbilities(false, abilities.Selectable, abilities.DoubleClickable)
				})

				increaseBTN := core.NewButton(w)
				increaseBTN.Styler(func(s *styles.Style) {
					s.Border.Width.Set(units.Dp(1))
					s.Border.Radius.Zero()
					m.ui.SetSize(s, units.Dp(60), units.Dp(60))
					s.Padding = styles.NewSideValues(units.Px(0), units.Px(0), units.Px(0), units.Px(0))
					s.Margin = styles.NewSideValues(units.Px(5), units.Px(5), units.Px(0), units.Px(5))
				})
				increaseBTN.SetIcon(icons.Add)
				increaseBTN.SetType(core.ButtonAction)

				increaseBTN.OnClick(func(e events.Event) {
					//println("x" + strconv.Itoa(order.OrderAmount))
					m.data.UpdateOrder(ItemID, ItemName, OrderAmount+1)
					text.Text = ItemName + " " + strconv.Itoa(ItemPrice) + "x" + strconv.Itoa(OrderAmount)
					m.UpdateCart()
				})

				decreaseBTN := core.NewButton(w)
				decreaseBTN.Styler(func(s *styles.Style) {
					//s.Background = colors.Scheme.Secondary.OnContainer
					s.Border.Width.Set(units.Dp(1))
					s.Border.Radius.Zero()
					m.ui.SetSize(s, units.Dp(60), units.Dp(60))
					s.Padding = styles.NewSideValues(units.Px(0), units.Px(0), units.Px(0), units.Px(0))
					s.Margin = styles.NewSideValues(units.Px(5), units.Px(5), units.Px(0), units.Px(5))
				})

				decreaseBTN.SetIcon(icons.Remove)
				decreaseBTN.SetType(core.ButtonAction)
				decreaseBTN.OnClick(func(e events.Event) {
					//println("x" + strconv.Itoa(order.OrderAmount))
					m.data.UpdateOrder(ItemID, ItemName, OrderAmount-1)
					if m.data.Orders[ItemID].OrderAmount == 0 {
						if m.data.RemoveOrder(ItemID) { // Use id directly
							println("Order removed successfully")
						} else {
							println("Failed to remove order")
						}
					}
					m.UpdateCart()
				})

				deleteBTN := core.NewButton(w)

				deleteBTN.Styler(func(s *styles.Style) {
					//s.Background = colors.Scheme.Secondary.OnContainer
					s.Border.Width.Set(units.Dp(1))
					s.Border.Radius.Zero()
					m.ui.SetSize(s, units.Dp(60), units.Dp(60))
					s.Padding = styles.NewSideValues(units.Px(0), units.Px(0), units.Px(0), units.Px(0))
					s.Margin = styles.NewSideValues(units.Px(5), units.Px(5), units.Px(0), units.Px(5))
				})
				deleteBTN.SetIcon(icons.Delete)
				deleteBTN.SetType(core.ButtonAction)
				deleteBTN.OnClick(func(e events.Event) {
					if m.data.RemoveOrder(ItemID) {
						println("Order removed successfully")
					} else {
						println("Failed to remove order")
					}
					m.UpdateCart()
				})
			})

		}
	})
}
