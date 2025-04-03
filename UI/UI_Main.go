package UI

import (
	"cogentcore.org/core/colors"
	"cogentcore.org/core/core"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/styles/units"
)

type UIManager struct {
	Body *core.Body

	RFID_StatusFrame *core.Frame
	RFID_Chooser     *core.Chooser

	ItemsFrame *core.Frame

	ItemActionFrame *core.Frame
}

func (u *UIManager) InitUI() {
	u.Body = core.NewBody()

	splitFrame := u.NewFrameZero(u.Body, units.Pw(100), units.Pw(100))
	splitFrame.Styler(func(s *styles.Style) {
		s.Display = styles.Flex
		s.Min.Set(units.Pw(0), units.Pw(0))
		s.Gap = units.XY{X: units.Em(0.5), Y: units.Em(0.5)}
		s.Grow.Set(1, 1)
	})

	u.LeftUI(splitFrame)
	u.RightUI(splitFrame)
}

func (u *UIManager) RunUI() {
	u.Body.RunMainWindow()
}

func (u *UIManager) LeftUI(splitFrame *core.Frame) {
	leftFrame := u.NewFrameZero(splitFrame, units.Pw(70), units.Ph(100))
	leftFrame.Styler(func(s *styles.Style) {
		//s.Background = colors.Scheme.Primary.Base
		s.Direction = styles.Column
		s.Gap = units.XY{X: units.Em(0.5), Y: units.Em(0.5)}
	})

	u.ServiceStatusUI(leftFrame)
	u.ItemUI(leftFrame)
	u.ItemActionUI(leftFrame)
}

func (u *UIManager) RightUI(splitFrame *core.Frame) {
	rightFrame := u.NewFrameZero(splitFrame, units.Pw(30), units.Ph(100))
	rightFrame.Styler(func(s *styles.Style) {
		//s.Background = colors.Scheme.Primary.Base
		s.Direction = styles.Column
		s.Gap = units.XY{X: units.Em(0.5), Y: units.Em(0.5)}
	})

	NavFrame := u.NewFrameZero(rightFrame, units.Pw(100), units.Dp(80))
	NavFrame.Styler(func(s *styles.Style) {
		s.Display = styles.Flex
		s.Grow.Set(1, 1)
		s.Background = colors.Scheme.Primary.Base
		s.Align.Content = styles.Center
		s.Align.Items = styles.Center
	})

	CartFrame := u.NewFrameZero(rightFrame, units.Pw(100), units.Dp(550))
	CartFrame.Styler(func(s *styles.Style) {
		s.Display = styles.Custom
		s.Grow.Set(1, 1)
		s.Background = colors.Scheme.Primary.Base
		s.Overflow.Y = styles.OverflowAuto
		s.ScrollbarWidth = units.Px(1)
	})

	InfoFrame := u.NewFrameZero(rightFrame, units.Pw(100), units.Dp(200))
	InfoFrame.Styler(func(s *styles.Style) {
		s.Display = styles.Flex
		s.Grow.Set(1, 1)
		s.Background = colors.Scheme.Secondary.On
	})

	CartActionFrame := u.NewFrameZero(rightFrame, units.Pw(100), units.Dp(125.5))
	CartActionFrame.Styler(func(s *styles.Style) {
		s.Display = styles.Flex
		s.Background = colors.Scheme.Primary.Base
		s.Align.Content = styles.Center
		s.Align.Items = styles.Center
		s.Justify.Content = styles.Center
		s.Justify.Items = styles.Center
	})

}
