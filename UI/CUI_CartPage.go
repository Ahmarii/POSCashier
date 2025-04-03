package UI

import (
	"cogentcore.org/core/colors"
	"cogentcore.org/core/core"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/styles/states"
	"cogentcore.org/core/styles/units"
)

func (u *UIManager) CartPageUI(parent *core.Frame) {

	u.CartFrame = u.NewFrameZero(parent, units.Pw(100), units.Ph(100))
	u.CartFrame.Styler(func(s *styles.Style) {
		s.Grow.Set(1, 1)
		s.Overflow.Y = styles.OverflowAuto
		s.Justify.Content = styles.Center
		s.Justify.Items = styles.Center
		s.Border.Radius = styles.BorderRadiusExtraSmall
	})

	CartPage := u.NewFrameZero(u.CartFrame, units.Pw(100), units.Ph(100))
	CartPage.SetName("CartPage")
	CartPage.Styler(func(s *styles.Style) {
		s.Direction = styles.Column
		s.Grow.Set(1, 1)
		s.Background = colors.Scheme.Secondary.On
		s.Overflow.Y = styles.OverflowAuto
		s.ScrollbarWidth = units.Px(1)
		s.SetState(false, states.Invisible)
		s.Align.Content = styles.Center
		s.Align.Items = styles.Center
	})
}
