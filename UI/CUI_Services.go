package UI

import (
	"cogentcore.org/core/colors"
	"cogentcore.org/core/core"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/styles/units"
)

func (u *UIManager) ServiceStatusUI(parent *core.Frame) {
	serviceStatusFrame := u.NewFrameZero(parent, units.Pw(100), units.Dp(80))
	serviceStatusFrame.Styler(func(s *styles.Style) {
		s.Display = styles.Flex
		//s.Background = colors.Scheme.Secondary.OnContainer
	})

	u.FiberStatusUI(serviceStatusFrame)
	u.DatabaseStatusUI(serviceStatusFrame)
	u.RFIDStatusUI(serviceStatusFrame)
	u.ScannerStatusUI(serviceStatusFrame)
}

func (u *UIManager) FiberStatusUI(parent *core.Frame) {
	fiberStatusFrame := u.NewFrameZero(parent, units.Pw(25), units.Dp(80))
	fiberStatusFrame.Styler(func(s *styles.Style) {
		s.Background = colors.Scheme.Secondary.Container
		s.Margin = styles.NewSideValues(units.Px(0), units.Px(2), units.Px(0), units.Px(0))
		s.Padding = styles.NewSideValues(units.Px(2))
		s.Display = styles.Flex
		s.Direction = styles.Column
		s.Border.Radius = styles.BorderRadiusExtraSmall
	})

	fiberStatusFrame_TextFrame := u.NewFrameZero(fiberStatusFrame, units.Pw(100), units.Ph(25))
	fiberStatusFrame_TextFrame.Styler(func(s *styles.Style) {
		s.Padding = styles.NewSideValues(units.Px(0), units.Px(3), units.Px(0), units.Px(0))
		s.Display = styles.Flex
		s.Direction = styles.Row
		s.Gap = units.XY{X: units.Px(3), Y: units.Px(0)}
	})

	fiberStatusFrame_TextFrame_text := core.NewText(fiberStatusFrame_TextFrame)
	fiberStatusFrame_TextFrame_text.SetText("API SERVER: ")

	fiberStatusFrame_TextFrame_textStatus := core.NewText(fiberStatusFrame_TextFrame)
	fiberStatusFrame_TextFrame_textStatus.SetText("OFFLINE")

	fiberStatusFrame_StatusBar := u.NewFrameZero(fiberStatusFrame, units.Pw(100), units.Ph(75))
	fiberStatusFrame_StatusBar.Styler(func(s *styles.Style) {
		s.Background = colors.Scheme.Primary.Base
		s.Margin = styles.NewSideValues(units.Px(5), units.Px(0), units.Px(5), units.Px(0))
	})
}

func (u *UIManager) DatabaseStatusUI(parent *core.Frame) {
	DataBaseStatusFrame := u.NewFrameZero(parent, units.Pw(25), units.Dp(80))
	DataBaseStatusFrame.Styler(func(s *styles.Style) {
		s.Background = colors.Scheme.Secondary.Container
		s.Margin = styles.NewSideValues(units.Px(0), units.Px(2), units.Px(0), units.Px(1))
		s.Padding = styles.NewSideValues(units.Px(2))
		s.Display = styles.Flex
		s.Direction = styles.Column
		s.Border.Radius = styles.BorderRadiusExtraSmall
	})

	DataBaseStatusFrame_TextFrame := u.NewFrameZero(DataBaseStatusFrame, units.Pw(100), units.Ph(25))
	DataBaseStatusFrame_TextFrame.Styler(func(s *styles.Style) {
		s.Padding = styles.NewSideValues(units.Px(0), units.Px(3), units.Px(0), units.Px(0))
		s.Display = styles.Flex
		s.Direction = styles.Row
		s.Gap = units.XY{X: units.Px(3), Y: units.Px(0)}
	})

	DataBaseStatusFrame_TextFrame_text := core.NewText(DataBaseStatusFrame_TextFrame)
	DataBaseStatusFrame_TextFrame_text.SetText("DATABASE: ")

	DataBaseStatusFrame_TextFrame_textStatus := core.NewText(DataBaseStatusFrame_TextFrame)
	DataBaseStatusFrame_TextFrame_textStatus.SetText("OFFLINE")

	DataBaseStatusFrame_StatusBar := u.NewFrameZero(DataBaseStatusFrame, units.Pw(100), units.Ph(75))
	DataBaseStatusFrame_StatusBar.Styler(func(s *styles.Style) {
		s.Background = colors.Scheme.Primary.Base
		s.Margin = styles.NewSideValues(units.Px(5), units.Px(0), units.Px(5), units.Px(0))
	})
}

func (u *UIManager) RFIDStatusUI(parent *core.Frame) {
	u.RFID_StatusFrame = u.NewFrameZero(parent, units.Pw(25), units.Dp(80))
	u.RFID_StatusFrame.Styler(func(s *styles.Style) {
		s.Background = colors.Scheme.Secondary.Container
		s.Margin = styles.NewSideValues(units.Px(0), units.Px(1), units.Px(0), units.Px(1))
		s.Padding = styles.NewSideValues(units.Px(2))
		s.Display = styles.Flex
		s.Direction = styles.Column
		s.Border.Radius = styles.BorderRadiusExtraSmall
	})

	RFIDStatusFrame_TextFrame := u.NewFrameZero(u.RFID_StatusFrame, units.Pw(100), units.Ph(25))
	RFIDStatusFrame_TextFrame.SetName("TextFrame")
	RFIDStatusFrame_TextFrame.Styler(func(s *styles.Style) {
		s.Padding = styles.NewSideValues(units.Px(0), units.Px(3), units.Px(0), units.Px(0))
		s.Display = styles.Flex
		s.Direction = styles.Row
		s.Gap = units.XY{X: units.Px(3), Y: units.Px(0)}
	})

	RFIDStatusFrame_TextFrame_text := core.NewText(RFIDStatusFrame_TextFrame)
	RFIDStatusFrame_TextFrame_text.SetText("RFID: ")

	RFIDStatusFrame_TextFrame_textStatus := core.NewText(RFIDStatusFrame_TextFrame)
	RFIDStatusFrame_TextFrame_textStatus.SetName("TextStatus")
	RFIDStatusFrame_TextFrame_textStatus.SetText("OFFLINE")

	RFIDStatusFrame_StatusBarFrame := u.NewFrameZero(u.RFID_StatusFrame, units.Pw(100), units.Ph(75))
	RFIDStatusFrame_StatusBarFrame.SetName("StatusBarFrame")
	RFIDStatusFrame_StatusBarFrame.Styler(func(s *styles.Style) {
		//s.Background = colors.Scheme.Primary.Base
		s.Margin = styles.NewSideValues(units.Px(5), units.Px(0), units.Px(5), units.Px(0))
		s.Display = styles.Flex
		s.Direction = styles.Row
	})

	RFIDStatusFrame_StatusBarFrame_Status := u.NewFrameZero(RFIDStatusFrame_StatusBarFrame, units.Pw(50), units.Ph(100))
	RFIDStatusFrame_StatusBarFrame_Status.SetName("BarStatus")
	RFIDStatusFrame_StatusBarFrame_Status.Styler(func(s *styles.Style) {
		s.Background = colors.Scheme.Primary.Base
		s.Border.Radius = styles.BorderRadiusExtraSmall

	})
	u.RFID_Chooser = core.NewChooser(RFIDStatusFrame_StatusBarFrame)
	u.RFID_Chooser.SetName("Chooser")
	u.RFID_Chooser.Styler(func(s *styles.Style) {
		s.Background = colors.Scheme.Secondary.On
		u.SetZero(s)
		u.SetSize(s, units.Pw(50), units.Ph(100))
		s.Border.Radius = styles.BorderRadiusExtraSmall

	})
}

func (u *UIManager) ScannerStatusUI(parent *core.Frame) {
	ScannerStatusFrame := u.NewFrameZero(parent, units.Pw(25), units.Dp(80))
	ScannerStatusFrame.Styler(func(s *styles.Style) {
		s.Background = colors.Scheme.Secondary.Container
		s.Margin = styles.NewSideValues(units.Px(0), units.Px(0), units.Px(0), units.Px(1))
		s.Padding = styles.NewSideValues(units.Px(2))
		s.Display = styles.Flex
		s.Direction = styles.Column
		s.Border.Radius = styles.BorderRadiusExtraSmall
	})

	ScannerStatusFrame_TextFrame := u.NewFrameZero(ScannerStatusFrame, units.Pw(100), units.Ph(25))
	ScannerStatusFrame_TextFrame.Styler(func(s *styles.Style) {
		s.Padding = styles.NewSideValues(units.Px(0), units.Px(3), units.Px(0), units.Px(0))
		s.Display = styles.Flex
		s.Direction = styles.Row
		s.Gap = units.XY{X: units.Px(3), Y: units.Px(0)}
	})

	ScannerStatusFrame_TextFrame_text := core.NewText(ScannerStatusFrame_TextFrame)
	ScannerStatusFrame_TextFrame_text.SetText("Scanner: ")

	ScannerStatusFrame_TextFrame_textStatus := core.NewText(ScannerStatusFrame_TextFrame)
	ScannerStatusFrame_TextFrame_textStatus.SetText("OFFLINE")

	ScannerStatusFrame_StatusBar := u.NewFrameZero(ScannerStatusFrame, units.Pw(100), units.Ph(75))
	ScannerStatusFrame_StatusBar.Styler(func(s *styles.Style) {
		s.Background = colors.Scheme.Primary.Base
		s.Margin = styles.NewSideValues(units.Px(5), units.Px(0), units.Px(5), units.Px(0))
	})
}
