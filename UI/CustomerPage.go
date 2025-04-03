package UI

import (
	"cogentcore.org/core/colors"
	"cogentcore.org/core/core"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/styles/states"
	"cogentcore.org/core/styles/units"
)

func (u *UIManager) CustomerPage(parent *core.Frame) {
	CustomerPageFrame := u.NewFrameZero(parent, units.Pw(100), units.Ph(100))
	CustomerPageFrame.Styler(func(s *styles.Style) {
		s.Direction = styles.Column
		s.Grow.Set(1, 1)
		s.Background = colors.Scheme.Secondary.On
		s.Overflow.Y = styles.OverflowAuto
		s.ScrollbarWidth = units.Px(1)
		s.SetState(true, states.Invisible)
		s.Align.Content = styles.Center
		s.Align.Items = styles.Center
	})

}
