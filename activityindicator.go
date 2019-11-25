package qtwidgets

import (
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"math"
)

type ActivityIndicator struct {
	widgets.QWidget

	interval int
	index int

	dotCount int
	dotColor *gui.QColor
	maxDiameter float64
	minDiameter float64

	timer *core.QTimer

	radiusList []int
	locationList [][2]float64
}

func CreateActivityIndicator(parent widgets.QWidget_ITF) *ActivityIndicator {
	this := NewActivityIndicator(parent, core.Qt__Widget)
	this.dotCount = 10
	this.dotColor = gui.NewQColor3(49, 177, 190, 255)
	this.maxDiameter = 12
	this.minDiameter = 5
	this.interval = 100

	this.SetFixedSize2(50, 50)

	this.timer = core.NewQTimer(this)
	this.timer.ConnectTimeout(this.Repaint)

	this.ConnectResizeEvent(this.ResizeEvent)
	this.ConnectPaintEvent(this.PaintEvent)

	return this
}

func(this *ActivityIndicator) Start() {
	this.timer.SetInterval(this.interval)
	this.timer.Start2()
}

func (this *ActivityIndicator) Stop() {
	this.timer.Stop()
}

func (this *ActivityIndicator) calculate() {
	var squareWidth = this.Width()
	if this.Height() < squareWidth {
		squareWidth = this.Height()
	}

	half := float64(squareWidth)/2
	centerDistance := half- this.maxDiameter / 2 - 1

	gap := (this.maxDiameter-this.minDiameter)/(float64(this.dotCount)-1)/2
	angleGap := math.Pi * 2 / float64(this.dotCount)

	this.locationList = nil
	this.radiusList = nil

	for i := 0; i < this.dotCount; i++ {
		this.radiusList = append(this.radiusList, int(this.maxDiameter/2-float64(i)*gap))
		radian := -angleGap*float64(i)
		x := half + centerDistance * math.Cos(radian)
		y := half- centerDistance * math.Sin(radian)
		this.locationList = append(this.locationList, [2]float64{x, y})
	}
}

func (this *ActivityIndicator) ResizeEvent(*gui.QResizeEvent) {
	this.calculate()
}

func (this *ActivityIndicator) PaintEvent(event *gui.QPaintEvent) {
	var painter = gui.NewQPainter2(this)
	defer painter.DestroyQPainter()

	painter.SetRenderHint(gui.QPainter__Antialiasing, true)

	painter.SetPen2(this.dotColor)
	painter.SetBrush(gui.NewQBrush3(this.dotColor, core.Qt__SolidPattern))

	for i := 0; i < this.dotCount; i++ {
		radius := this.radiusList[(this.index + this.dotCount - i)%this.dotCount]
		x := int(this.locationList[i][0])
		y := int(this.locationList[i][1])
		painter.DrawEllipse3(x, y, radius, radius)
	}

	this.index++
}