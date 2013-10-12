package gotcp

import (
	"github.com/xuyu/logging"
	"strconv"
)

type EntryName interface {
	GetId() int64
	GetName() string
	GetEntryName() string
}
type Entry struct {
	Id           int64
	Name         string
	GetEntryName func() string
}

func (self *Entry) formatHead(format string) string {
	if self.GetEntryName != nil {
		return self.GetEntryName() + "[" + strconv.FormatInt(self.Id, 10) + "," + self.Name + "]" + format
	}
	return format
}
func (self *Entry) Debug(format string, v ...interface{}) {
	logging.Debug(self.formatHead(format), v...)
}
func (self *Entry) Info(format string, v ...interface{}) {
	logging.Info(self.formatHead(format), v...)
}
func (self *Entry) Error(format string, v ...interface{}) {
	logging.Error(self.formatHead(format), v...)
}
