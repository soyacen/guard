package log

import (
	"github.com/sirupsen/logrus"
)

/*func init() {
	filenameHook := NewHook()
	filenameHook.Field = "line"
	logrus.AddHook(filenameHook)
}
*/
func Fatal(v ...interface{}) {
	logrus.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	logrus.Fatalf(format, v...)
}

func Fatalln(v ...interface{}) {
	logrus.Fatalln(v...)
}

func Println(v ...interface{}) {
	logrus.Println(v...)
}

func Print(v ...interface{}) {
	logrus.Print(v...)
}

func Printf(format string, v ...interface{}) {
	logrus.Printf(format, v...)
}

func Errorln(v ...interface{}) {
	logrus.Errorln(v...)
}

func Error(v ...interface{}) {
	logrus.Error(v...)
}

func Errorf(format string, v ...interface{}) {
	logrus.Errorf(format, v...)
}
