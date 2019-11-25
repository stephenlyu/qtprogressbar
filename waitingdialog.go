package qtwidgets

import (
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"github.com/therecipe/qt/core"
)

type WaitingDialog struct {
	widgets.QDialog
	indicator *ActivityIndicator
	VerticalLayout *widgets.QVBoxLayout
	parent widgets.QWidget_ITF
}

func CreateWaitingDialog(parent widgets.QWidget_ITF) *WaitingDialog {
	this := NewWaitingDialog(parent, core.Qt__Dialog)
	this.SetWindowFlags(core.Qt__MSWindowsFixedSizeDialogHint | core.Qt__FramelessWindowHint | core.Qt__Dialog)

	this.parent = parent
	this.indicator = CreateActivityIndicator(this)

	this.SetObjectName("Dialog")
	this.SetFixedSize2(50, 50)
	this.VerticalLayout = widgets.NewQVBoxLayout2(this)
	this.VerticalLayout.SetObjectName("verticalLayout")
	this.VerticalLayout.SetContentsMargins(0, 0, 0, 0)
	this.VerticalLayout.SetSpacing(0)

	this.VerticalLayout.AddWidget(this.indicator, 0, 0)

	return this
}

func (this *WaitingDialog) updatePosition() {
	rect := this.parent.QWidget_PTR().Geometry()
	rcDialog := this.Rect()
	this.Move2(rect.X() + int(rect.Width() - rcDialog.Width()) / 2,
		rect.Y() + int(rect.Height() - rcDialog.Height()) / 2)
}

func (this *WaitingDialog) Start() {
	if this.parent != nil {
		this.parent.QWidget_PTR().ConnectResizeEvent(func(*gui.QResizeEvent) {
			this.updatePosition()
		})
		this.parent.QWidget_PTR().ConnectMoveEvent(func(*gui.QMoveEvent) {
			this.updatePosition()
		})
	}

	this.indicator.Start()
	this.Exec()
}

func (this *WaitingDialog) Stop() {
	this.indicator.Stop()
	this.Accept()
	if this.parent != nil {
		this.parent.QWidget_PTR().DisconnectResizeEvent()
		this.parent.QWidget_PTR().DisconnectMoveEvent()
	}
}
