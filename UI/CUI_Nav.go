package UI

import (
	"cogentcore.org/core/colors"
	"cogentcore.org/core/core"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/styles/units"
)

func (u *UIManager) NavUI(parent *core.Frame) {
	NavFrame := u.NewFrameZero(parent, units.Pw(100), units.Dp(80))
	NavFrame.Styler(func(s *styles.Style) {
		s.Display = styles.Flex
		s.Grow.Set(1, 1)
		s.Background = colors.Scheme.Primary.Base
		s.Align.Content = styles.Center
		s.Align.Items = styles.Center
	})
}
