//  header block begin

// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qstyleoption.h
// #include <qstyleoption.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// /usr/include/qt/QtWidgets/qstyleoption.h:442
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStyleOptionViewItem()

/*

 */
func (*QStyleOptionViewItem) NewForInherit() *QStyleOptionViewItem {
	return NewQStyleOptionViewItem()
}
func NewQStyleOptionViewItem() *QStyleOptionViewItem {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QStyleOptionViewItemC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionViewItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionViewItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qstyleoption.h:446
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QStyleOptionViewItem(int)

/*

 */
func (*QStyleOptionViewItem) NewForInherit1(version int) *QStyleOptionViewItem {
	return NewQStyleOptionViewItem1(version)
}
func NewQStyleOptionViewItem1(version int) *QStyleOptionViewItem {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QStyleOptionViewItemC2Ei", qtrt.FFI_TYPE_POINTER, version)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStyleOptionViewItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQStyleOptionViewItem)
	return gothis
}

func DeleteQStyleOptionViewItem(this *QStyleOptionViewItem) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QStyleOptionViewItemD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_11028() {
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
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
