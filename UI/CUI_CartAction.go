package UI

import (
	"cogentcore.org/core/colors"
	"cogentcore.org/core/core"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/styles/units"
)

func (u *UIManager) CartActionUI(parent *core.Frame) {
	CartActionFrame := u.NewFrameZero(parent, units.Pw(100), units.Dp(125.5))
	CartActionFrame.Styler(func(s *styles.Style) {
		s.Display = styles.Flex
		s.Background = colors.Scheme.Primary.Base
		s.Align.Content = styles.Center
		s.Align.Items = styles.Center
		s.Justify.Content = styles.Center
		s.Justify.Items = styles.Center
	})
}
