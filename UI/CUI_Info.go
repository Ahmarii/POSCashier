package UI

import (
	"cogentcore.org/core/colors"
	"cogentcore.org/core/core"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/styles/units"
)

func (u *UIManager) InfoUI(parent *core.Frame) {
	InfoFrame := u.NewFrameZero(parent, units.Pw(100), units.Dp(200))
	InfoFrame.Styler(func(s *styles.Style) {
		s.Display = styles.Flex
		s.Grow.Set(1, 1)
		s.Background = colors.Scheme.Secondary.On
	})
}
