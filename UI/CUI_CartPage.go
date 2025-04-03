package UI

import (
	"cogentcore.org/core/colors"
	"cogentcore.org/core/core"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/styles/states"
	"cogentcore.org/core/styles/units"
)

func (u *UIManager) CartPageUI(parent *core.Frame) {
	CartPageFrame := u.NewFrameZero(parent, units.Pw(100), units.Ph(100))
	CartPageFrame.Styler(func(s *styles.Style) {
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
