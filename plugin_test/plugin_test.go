package plugin_test

import (
	"github.com/sinlov/drone-file-browser-plugin/drone_info"
	"github.com/sinlov/drone-file-browser-plugin/plugin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPluginSendMode(t *testing.T) {
	// mock Plugin
	t.Logf("~> mock Plugin")
	p := plugin.Plugin{}
	// do Plugin
	t.Logf("~> do Plugin")

	if envCheck(t) {
		return
	}

	// use env:ENV_DEBUG
	p.Config.Debug = envDebug
	testDataDistFolderPath, err := initTestDataPostFileDir()
	if err != nil {
		t.Error(err)
	}

	assert.NotEqual(t, "", testDataDistFolderPath)

	p.Config.FileBrowserBaseConfig.FileBrowserHost = envFileBrowserHost
	p.Config.FileBrowserBaseConfig.FileBrowserUsername = envUserName
	p.Config.FileBrowserBaseConfig.FileBrowserUserPassword = envPassword

	err = p.Exec()
	if nil == err {
		t.Error("args FileBrowserWorkMode error should be catch!")
	}

	t.Logf("-> now start test FileBrowserWorkMode %s", plugin.WorkModeSend)
	p.Config.FileBrowserWorkMode = plugin.WorkModeSend

	// start test mode
	p.Config.FileBrowserSendModeConfig.FileBrowserDistType = "other"
	err = p.Exec()
	if nil == err {
		t.Error("args FileBrowserDistType should be catch!")
	}

	p.Config.FileBrowserSendModeConfig.FileBrowserDistType = plugin.DistTypeGit

	err = p.Exec()
	if nil == err {
		t.Error("args file_browser_remote_root_path should be catch!")
	}

	p.Config.FileBrowserSendModeConfig.FileBrowserRemoteRootPath = mockFileBrowserRemoteRootPath

	err = p.Exec()
	if nil == err {
		t.Error("args file_browser_target_dist_root_path should be catch!")
	}

	p.Config.FileBrowserSendModeConfig.FileBrowserTargetDistRootPath = mockFileBrowserTargetDistRootPath
	p.Config.FileBrowserSendModeConfig.FileBrowserTargetFileRegular = mockFileBrowserTargetFileRegular

	p.Config.TimeoutSecond = defTimeoutSecond
	p.Config.FileBrowserBaseConfig.FileBrowserTimeoutPushSecond = defTimeoutFileSecond

	p.Drone = *drone_info.MockDroneInfo("success")
	// verify Plugin
	assert.Equal(t, "sinlov", p.Drone.Repo.OwnerName)

	err = p.Exec()
	if err != nil {
		t.Fatal(err)
	}

}
