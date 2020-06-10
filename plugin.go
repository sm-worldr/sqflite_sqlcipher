package sqflite_sqlcipher

import (
	//"github.com/pkg/errors"

	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
)

const channelName = "com.davidmartos96.sqflite_sqlcipher"

type SqlCipherPlugin struct{}

var _ flutter.Plugin = &SqlCipherPlugin{}

// plugin init
func (p *SqlCipherPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.handleFunc("test", p.test)
	return nil
}

// plugin methods
func (p *SqlCipherPlugin) test(arguments interface{}) (reply string, err error) {
	return "", nil
}
