//  header block begin
// /usr/include/qt/QtCore/qeventloop.h
// #include <qeventloop.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

func init() {
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
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin
type QEventLoopLocker struct {
	*qtrt.CObject
}

func (this *QEventLoopLocker) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQEventLoopLockerFromPointer(cthis unsafe.Pointer) *QEventLoopLocker {
	return &QEventLoopLocker{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qeventloop.h:93
// index:0
// Public
// void QEventLoopLocker()
func NewQEventLoopLocker() *QEventLoopLocker {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QEventLoopLockerC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQEventLoopLockerFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qeventloop.h:94
// index:1
// Public
// void QEventLoopLocker(class QEventLoop *)
func NewQEventLoopLocker_1(loop unsafe.Pointer) *QEventLoopLocker {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QEventLoopLockerC2EP10QEventLoop", ffiqt.FFI_TYPE_VOID, cthis, loop)
	gopp.ErrPrint(err, rv)
	gothis := NewQEventLoopLockerFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qeventloop.h:95
// index:2
// Public
// void QEventLoopLocker(class QThread *)
func NewQEventLoopLocker_2(thread unsafe.Pointer) *QEventLoopLocker {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QEventLoopLockerC2EP7QThread", ffiqt.FFI_TYPE_VOID, cthis, thread)
	gopp.ErrPrint(err, rv)
	gothis := NewQEventLoopLockerFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qeventloop.h:96
// index:0
// Public
// void ~QEventLoopLocker()
func DeleteQEventLoopLocker(*QEventLoopLocker) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QEventLoopLockerD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

//  body block end