package qtgui

// /usr/include/qt/QtGui/qbrush.h
// #include <qbrush.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 19
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
type QConicalGradient struct {
	*QGradient
}
type QConicalGradient_ITF interface {
	QGradient_ITF
	QConicalGradient_PTR() *QConicalGradient
}

func (ptr *QConicalGradient) QConicalGradient_PTR() *QConicalGradient { return ptr }

func (this *QConicalGradient) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGradient.GetCthis()
	}
}
func (this *QConicalGradient) SetCthis(cthis unsafe.Pointer) {
	this.QGradient = NewQGradientFromPointer(cthis)
}
func NewQConicalGradientFromPointer(cthis unsafe.Pointer) *QConicalGradient {
	bcthis0 := NewQGradientFromPointer(cthis)
	return &QConicalGradient{bcthis0}
}
func (*QConicalGradient) NewFromPointer(cthis unsafe.Pointer) *QConicalGradient {
	return NewQConicalGradientFromPointer(cthis)
}

// /usr/include/qt/QtGui/qbrush.h:480
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QConicalGradient()

/*

 */
func (*QConicalGradient) NewForInherit() *QConicalGradient {
	return NewQConicalGradient()
}
func NewQConicalGradient() *QConicalGradient {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QConicalGradientC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQConicalGradientFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQConicalGradient)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:481
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QConicalGradient(const QPointF &, qreal)

/*

 */
func (*QConicalGradient) NewForInherit1(center qtcore.QPointF_ITF, startAngle float64) *QConicalGradient {
	return NewQConicalGradient1(center, startAngle)
}
func NewQConicalGradient1(center qtcore.QPointF_ITF, startAngle float64) *QConicalGradient {
	var convArg0 unsafe.Pointer
	if center != nil && center.QPointF_PTR() != nil {
		convArg0 = center.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QConicalGradientC2ERK7QPointFd", qtrt.FFI_TYPE_POINTER, convArg0, startAngle)
	qtrt.ErrPrint(err, rv)
	gothis := NewQConicalGradientFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQConicalGradient)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:482
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QConicalGradient(qreal, qreal, qreal)

/*

 */
func (*QConicalGradient) NewForInherit2(cx float64, cy float64, startAngle float64) *QConicalGradient {
	return NewQConicalGradient2(cx, cy, startAngle)
}
func NewQConicalGradient2(cx float64, cy float64, startAngle float64) *QConicalGradient {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QConicalGradientC2Eddd", qtrt.FFI_TYPE_POINTER, cx, cy, startAngle)
	qtrt.ErrPrint(err, rv)
	gothis := NewQConicalGradientFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQConicalGradient)
	return gothis
}

// /usr/include/qt/QtGui/qbrush.h:484
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF center() const

/*

 */
func (this *QConicalGradient) Center() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QConicalGradient6centerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtGui/qbrush.h:485
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCenter(const QPointF &)

/*

 */
func (this *QConicalGradient) SetCenter(center qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if center != nil && center.QPointF_PTR() != nil {
		convArg0 = center.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QConicalGradient9setCenterERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:486
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setCenter(qreal, qreal)

/*

 */
func (this *QConicalGradient) SetCenter1(x float64, y float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QConicalGradient9setCenterEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:488
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal angle() const

/*

 */
func (this *QConicalGradient) Angle() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QConicalGradient5angleEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qbrush.h:489
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAngle(qreal)

/*

 */
func (this *QConicalGradient) SetAngle(angle float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QConicalGradient8setAngleEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), angle)
	qtrt.ErrPrint(err, rv)
}

func DeleteQConicalGradient(this *QConicalGradient) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QConicalGradientD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10753() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
