package plugin_test

import (
	"github.com/sinlov/drone-file-browser-plugin/drone_info"
	"github.com/sinlov/drone-file-browser-plugin/plugin"
	"github.com/sinlov/filebrowser-client/web_api"
	"github.com/stretchr/testify/assert"
	"os"
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
	testDataFolderAbsPath, err := getOrCreateTestDataFolderFullPath()
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
	p.Config.FileBrowserSendModeConfig.FileBrowserTargetFileRegular = mockFileBrowserTargetFileRegularFail

	p.Config.TimeoutSecond = defTimeoutSecond
	p.Config.FileBrowserBaseConfig.FileBrowserTimeoutPushSecond = defTimeoutFileSecond

	p.Drone = *drone_info.MockDroneInfo("success")
	// verify Plugin
	assert.Equal(t, "sinlov", p.Drone.Repo.OwnerName)

	err = p.Exec()
	if err == nil {
		t.Error("args p.Drone.Build.WorkSpace should be catch!")
	}

	// change right workspace
	p.Drone.Build.WorkSpace = testDataFolderAbsPath
	if err == nil {
		t.Error("args p.Config.FileBrowserSendModeConfig.FileBrowserTargetFileRegular should be catch!")
	}

	p.Config.FileBrowserSendModeConfig.FileBrowserShareLinkEnable = true
	p.Config.FileBrowserSendModeConfig.FileBrowserShareLinkAutoPasswordEnable = true
	p.Config.FileBrowserSendModeConfig.FileBrowserShareLinkUnit = web_api.ShareUnitHours
	p.Config.FileBrowserSendModeConfig.FileBrowserShareLinkExpires = 4

	p.Config.FileBrowserSendModeConfig.FileBrowserDistType = plugin.DistTypeCustom
	p.Config.FileBrowserSendModeConfig.FileBrowserDistGraph = mockFileBrowserDistGraph

	// change
	p.Config.FileBrowserSendModeConfig.FileBrowserTargetFileRegular = mockFileBrowserTargetFileRegularOne

	err = p.Exec()
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, "", os.Getenv(plugin.EnvPluginDroneFileBrowserShareRemotePath))
	assert.NotEqual(t, "", os.Getenv(plugin.EnvPluginDroneFileBrowserSharePage))
	downloadUrl := os.Getenv(plugin.EnvPluginDroneFileBrowserShareDownloadUrl)
	assert.NotEqual(t, "", downloadUrl)
	t.Logf("download url: %s", downloadUrl)

	p.Config.FileBrowserSendModeConfig.FileBrowserDistType = plugin.DistTypeGit
	// change right file regular for one file
	p.Config.FileBrowserSendModeConfig.FileBrowserTargetFileRegular = mockFileBrowserTargetFileRegularOne

	err = p.Exec()
	if err != nil {
		t.Fatal(err)
	}
	downloadUrl = os.Getenv(plugin.EnvPluginDroneFileBrowserShareDownloadUrl)
	assert.NotEqual(t, "", downloadUrl)
	t.Logf("download url: %s", downloadUrl)

	// change right file regular for more than one
	p.Config.FileBrowserSendModeConfig.FileBrowserTargetFileRegular = mockFileBrowserTargetFileRegular

	err = p.Exec()
	if err != nil {
		t.Fatal(err)
	}
	downloadUrl = os.Getenv(plugin.EnvPluginDroneFileBrowserShareDownloadUrl)
	assert.NotEqual(t, "", downloadUrl)
	t.Logf("download url: %s", downloadUrl)
}
