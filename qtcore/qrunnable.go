package qtcore

// /usr/include/qt/QtCore/qrunnable.h
// #include <qrunnable.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
type QRunnable struct {
	*qtrt.CObject
}
type QRunnable_ITF interface {
	QRunnable_PTR() *QRunnable
}

func (ptr *QRunnable) QRunnable_PTR() *QRunnable { return ptr }

func (this *QRunnable) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QRunnable) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQRunnableFromPointer(cthis unsafe.Pointer) *QRunnable {
	return &QRunnable{&qtrt.CObject{cthis}}
}
func (*QRunnable) NewFromPointer(cthis unsafe.Pointer) *QRunnable {
	return NewQRunnableFromPointer(cthis)
}

// /usr/include/qt/QtCore/qrunnable.h:58
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void run()

/*
Implement this pure virtual function in your subclass.
*/
func (this *QRunnable) Run() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QRunnable3runEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qrunnable.h:60
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QRunnable()

/*
Constructs a QRunnable. Auto-deletion is enabled by default.

See also autoDelete() and setAutoDelete().
*/
func (*QRunnable) NewForInherit() *QRunnable {
	return NewQRunnable()
}
func NewQRunnable() *QRunnable {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QRunnableC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRunnableFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQRunnable)
	return gothis
}

// /usr/include/qt/QtCore/qrunnable.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QRunnable()

/*

 */
func DeleteQRunnable(this *QRunnable) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QRunnableD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qrunnable.h:63
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool autoDelete() const

/*
Returns true is auto-deletion is enabled; false otherwise.

If auto-deletion is enabled, QThreadPool will automatically delete this runnable after calling run(); otherwise, ownership remains with the application programmer.

See also setAutoDelete() and QThreadPool.
*/
func (this *QRunnable) AutoDelete() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QRunnable10autoDeleteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qrunnable.h:64
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setAutoDelete(bool)

/*
Enables auto-deletion if autoDelete is true; otherwise auto-deletion is disabled.

If auto-deletion is enabled, QThreadPool will automatically delete this runnable after calling run(); otherwise, ownership remains with the application programmer.

Note that this flag must be set before calling QThreadPool::start(). Calling this function after QThreadPool::start() results in undefined behavior.

See also autoDelete() and QThreadPool.
*/
func (this *QRunnable) SetAutoDelete(_autoDelete bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QRunnable13setAutoDeleteEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), _autoDelete)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_10413() {
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
}

//  keep block end
