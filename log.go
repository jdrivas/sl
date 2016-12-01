package sl
  
import (
  "fmt"
  "github.com/Sirupsen/logrus"
)

type SL struct {
  Logger *logrus.Logger
  DefFields logrus.Fields
}

func New() (*SL) {
// func New(l *logrus.Logger) (*SL) {
  sl := new(SL)
  sl.Logger = logrus.New()
  sl.DefFields = make(logrus.Fields,0)
  return sl
}



func (l *SL) CheckFatalError(fields logrus.Fields, msg string, err error) {
  if err != nil { 
    l.Fatal(fields, msg, err)
  }
}


func (l *SL) Panic(fields logrus.Fields, msg string, err error) {
  e := l.prepare(fields)
  if err == nil {
    e.Panic(msg)
  } else {
    e.WithError(err).Panic(msg)
  }
}

func (l *SL) Fatal(fields logrus.Fields, msg string, err error) {
  e := l.prepare(fields)
  if err == nil {
    e.Fatal(msg)
  } else {
    e.WithError(err).Fatal(msg)
  }
}

func (l *SL) Error(fields logrus.Fields, msg string, err error) {
  e := l.prepare(fields)
  if err == nil {
    e.Error(msg)
  } else {
    e.Error(msg)
  }
}

func (l *SL) Warn(fields logrus.Fields, msg string) {
  l.prepare(fields).Warn(msg)
}

func (l *SL)Info(fields logrus.Fields, msg string) {
  l.prepare(fields).Info(msg)
}

func (l *SL) Debug(fields logrus.Fields, msg string) {

  l.prepare(fields).Debug(msg)
}

func (l *SL) SetFormatter(f logrus.Formatter) {
  l.Logger.Formatter = f
}

func (l *SL) SetLevel(f logrus.Level) {
  l.Logger.Level = f
}

// Copies the fields in f into the default fields 
// into the log. Existing default fields are lost.
func (l *SL) SetDefaultFields(f logrus.Fields) {
  l.DefFields = make(logrus.Fields, 0)
  l.AddDefaultFields(f)
}

// Adds fields to the default fields.
func (l *SL) AddDefaultFields(f logrus.Fields) {
  for k, v := range f {
    l.DefFields[k] = v
  }
}

func (l *SL) prepare(fields logrus.Fields) (*logrus.Entry) {
  fmt.Printf("Preparing ... with %#v\n", fields)
  fmt.Printf("Default: ... with %#v\n", l.DefFields)
  f := l.mergeDefault(fields)
  fmt.Printf("Merged is: %#v\n", f)
  return l.Logger.WithFields(f)
}

func (l *SL) mergeDefault(f logrus.Fields) (logrus.Fields) {
  n := make(logrus.Fields, 0)
  for k, v := range l.DefFields {
    n[k] = v
  }
  for k, v := range f {
    n[k] = v
  }
  return n
}

