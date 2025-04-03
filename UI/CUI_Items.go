package UI

import (
	"cogentcore.org/core/colors"
	"cogentcore.org/core/core"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/styles/units"
)

func (u *UIManager) ItemUI(parent *core.Frame) {
	u.ItemsFrame = u.NewFrameZero(parent, units.Pw(100), units.Dp(758))
	u.ItemsFrame.Styler(func(s *styles.Style) {
		s.Grow.Set(1, 1)
		s.Overflow.Y = styles.OverflowAuto
		s.Justify.Content = styles.Center
		s.Justify.Items = styles.Center
		s.Border.Radius = styles.BorderRadiusExtraSmall
	})

	ItemPage := u.NewFrameZero(u.ItemsFrame, units.Pw(100), units.Ph(100))
	ItemPage.Styler(func(s *styles.Style) {
		s.Display = styles.Grid
		s.Columns = 4
		s.Grow.Set(1, 1)
		s.Background = colors.Scheme.Secondary.On
		s.Overflow.Y = styles.OverflowAuto
		s.ScrollbarWidth = units.Px(1)
		s.Justify.Content = styles.Center
		s.Justify.Items = styles.Center
		s.Border.Radius = styles.BorderRadiusExtraSmall
	})

}
