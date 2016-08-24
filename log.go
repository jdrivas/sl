package sl
  
import (
  "github.com/Sirupsen/logrus"
)

type SL struct {
  Logger *logrus.Logger
}

func New() (*SL) {
// func New(l *logrus.Logger) (*SL) {
  sl := new(SL)
  sl.Logger = logrus.New()
  return sl
}

func (l *SL) CheckFatalError(fields logrus.Fields, msg string, err error) {
  if err != nil { 
    l.Fatal(fields, msg, err)
  }
}

func (l *SL) Panic(fields logrus.Fields, msg string, err error) {
  log := l.Logger
  if err == nil {
    log.WithFields(fields).Panic(msg)
  } else {
    log.WithFields(fields).WithError(err).Panic(msg)
  }
}

func (l *SL) Fatal(fields logrus.Fields, msg string, err error) {
  log := l.Logger
  if err == nil {
    log.WithFields(fields).Fatal(msg)
  } else {
    log.WithFields(fields).WithError(err).Fatal(msg)
  }
}

func (l *SL) Error(fields logrus.Fields, msg string, err error) {
  log := l.Logger
  if err == nil {
    log.WithFields(fields).Error(msg)
  } else {
    log.WithFields(fields).WithError(err).Error(msg)
  }
}

func (l *SL) Warn(fields logrus.Fields, msg string) {
  l.Logger.WithFields(fields).Warn(msg)
}

func (l *SL)Info(fields logrus.Fields, msg string) {
  l.Logger.WithFields(fields).Info(msg)
}

func (l *SL) Debug(fields logrus.Fields, msg string) {
  l.Logger.WithFields(fields).Debug(msg)
}

func (l *SL) SetFormatter(f logrus.Formatter) {
  l.Logger.Formatter = f
}

func (l *SL) SetLevel(f logrus.Level) {
  l.Logger.Level = f
}
