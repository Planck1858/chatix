package logging

type Logger interface {
	Info(...interface{})
	Error(...interface{})
	Warn(...interface{})
	Fatal(...interface{})
	With(...interface{}) Logger
	Sync() error
}
