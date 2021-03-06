package qtqml

// /usr/include/qt/QtQml/qqmlprivate.h
// #include <qqmlprivate.h>
// #include <QtQml>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  body block begin

/*

 */
type CachedQmlUnit struct {
	*qtrt.CObject
}
type CachedQmlUnit_ITF interface {
	CachedQmlUnit_PTR() *CachedQmlUnit
}

func (ptr *CachedQmlUnit) CachedQmlUnit_PTR() *CachedQmlUnit { return ptr }

func (this *CachedQmlUnit) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *CachedQmlUnit) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewCachedQmlUnitFromPointer(cthis unsafe.Pointer) *CachedQmlUnit {
	return &CachedQmlUnit{&qtrt.CObject{cthis}}
}
func (*CachedQmlUnit) NewFromPointer(cthis unsafe.Pointer) *CachedQmlUnit {
	return NewCachedQmlUnitFromPointer(cthis)
}

func DeleteCachedQmlUnit(this *CachedQmlUnit) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13CachedQmlUnitD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_11495() {
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
		qtnetwork.KeepMe()
	}
}

//  keep block end
