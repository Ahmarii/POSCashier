package UI

import (
	"cogentcore.org/core/core"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/styles/units"
	"cogentcore.org/core/tree"
)

// set Max and Min of Element Styles
func (u *UIManager) SetSize(s *styles.Style, x units.Value, y units.Value) {
	s.Min.Set(x, y)
	s.Max.Set(x, y)
}

func (u *UIManager) SetZero(s *styles.Style) {
	s.Gap = units.XY{X: units.Px(0), Y: units.Px(0)}
	s.Border.Radius.Zero()
	s.Border.Width.Set(units.Dp(0))
	s.MaxBorder = styles.Border{}
	s.Border = styles.Border{}
	s.Margin = styles.NewSideValues(units.Px(0), units.Px(0), units.Px(0), units.Px(0))
	s.Padding = styles.NewSideValues(units.Px(0), units.Px(0), units.Px(0), units.Px(0))
	s.MaxBoxShadow = []styles.Shadow{styles.Shadow{}}
}

func (u *UIManager) NewFrameZero(parent tree.Node, x units.Value, y units.Value) *core.Frame {
	newFrame := core.NewFrame(parent)
	newFrame.Styler(func(s *styles.Style) {
		u.SetZero(s)
		u.SetSize(s, x, y)
	})
	return newFrame
}

func (u *UIManager) AsyncUpdateElement(w *core.WidgetBase) {
	w.AsyncLock()
	w.Update()
	w.AsyncUnlock()
}
