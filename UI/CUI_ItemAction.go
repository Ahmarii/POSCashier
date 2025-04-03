package UI

import (
	"cogentcore.org/core/colors"
	"cogentcore.org/core/core"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/styles/units"
)

func (u *UIManager) ItemActionUI(parent *core.Frame) {
	u.ItemActionFrame = u.NewFrameZero(parent, units.Pw(100), units.Dp(125.5))
	u.ItemActionFrame.Styler(func(s *styles.Style) {
		s.Display = styles.Flex
		s.Grow.Set(1, 1)
		s.Background = colors.Scheme.Secondary.On
		s.Align.Content = styles.Center
		s.Align.Items = styles.Center
		s.Justify.Content = styles.Center
		s.Justify.Items = styles.Center
		s.Gap = units.XY{X: units.Em(0.5), Y: units.Em(0.5)}
	})

	ProductByID_Field := core.NewTextField(u.ItemActionFrame).SetPlaceholder("Product ID")
	ProductByID_Field.Styler(func(s *styles.Style) {
		u.SetZero(s)
		u.SetSize(s, units.Dp(400), units.Dp(80))
		s.Background = colors.Scheme.Secondary.Container
		s.Padding = styles.NewSideValues(units.Px(5))
		s.Margin = styles.NewSideValues(units.Px(0), units.Px(0), units.Px(0), units.Px(5))
		s.Font.Size.Dp(30)
		s.Align.Content = styles.Center
		s.Align.Items = styles.Center
	})

	ProductByID_ButtonFrame := u.NewFrameZero(u.ItemActionFrame, units.Dp(95), units.Dp(95))
	ProductByID_ButtonFrame.Styler(func(s *styles.Style) {
		s.Align.Content = styles.Center
		s.Align.Items = styles.Center
		s.Justify.Content = styles.Center
		s.Justify.Items = styles.Center
	})

	ProductByID_Button := core.NewButton(ProductByID_ButtonFrame).SetType(core.ButtonAction).SetText("Add")
	ProductByID_Button.Styler(func(s *styles.Style) {
		u.SetZero(s)
		u.SetSize(s, units.Pw(100), units.Ph(100))

		s.Background = colors.Scheme.Error.Container
	})

	Refresh_Button := core.NewButton(u.ItemActionFrame).SetType(core.ButtonAction).SetText("Refresh")
	Refresh_Button.SetName("Refresh_Button")
	Refresh_Button.Styler(func(s *styles.Style) {
		u.SetZero(s)
		u.SetSize(s, units.Dp(95), units.Dp(95))
		s.Background = colors.Scheme.Error.Container
	})
}
