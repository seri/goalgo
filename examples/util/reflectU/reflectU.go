package reflectU

import (
    "strings"
    "reflect"
    "runtime"
)

func BaseName(s string) string {
    i := strings.LastIndex(s, ".")
    if i == -1 {
        return s
    }
    return s[(i + 1):]
}

func TypeName(x interface{}) string {
    t := reflect.TypeOf(x)
    switch t.Kind() {
    case reflect.Ptr:
        return t.Elem().Name()
    case reflect.Array, reflect.Slice:
        return "[]" + t.Elem().Name()
    case reflect.Func:
        return BaseName(runtime.FuncForPC(reflect.ValueOf(x).Pointer()).Name())
    }
    return t.Name()
}
