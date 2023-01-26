package plugin

import (
	"github.com/sinlov/drone-file-browser-plugin/drone_info"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestPlugin(t *testing.T) {
	// mock Plugin
	t.Logf("~> mock Plugin")
	p := Plugin{}
	// do Plugin
	t.Logf("~> do Plugin")
	err := p.Exec()
	if nil == err {
		t.Error("file_browser_host empty error should be catch!")
	}

	envWebHook := os.Getenv("PLUGIN_FILE_BROWSER_HOST")
	if envWebHook == "" {
		t.Logf("please set env:PLUGIN_FILE_BROWSER_HOST to again")
		return
	}

	p.Config.Webhook = envWebHook

	err = p.Exec()
	if nil == err {
		t.Error("msg type empty error should be catch!")
	}

	p.Config.MsgType = "mock" // not support type
	err = p.Exec()
	if nil == err {
		t.Error("msg type not support error should be catch!")
	}

	envMsgType := os.Getenv("PLUGIN_MSG_TYPE")

	if envMsgType == "" {
		t.Error("please set env:PLUGIN_MSG_TYPE")
	}

	p.Config.MsgType = envMsgType

	p.Drone = *drone_info.MockDroneInfo("success")
	// verify Plugin

	assert.Equal(t, "sinlov", p.Drone.Repo.OwnerName)
}
