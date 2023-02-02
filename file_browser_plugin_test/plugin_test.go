package plugin_test

import (
	"github.com/sinlov/drone-file-browser-plugin/file_browser_plugin"
	"github.com/sinlov/drone-info-tools/drone_info"
	"github.com/sinlov/filebrowser-client/web_api"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestPluginSendMode(t *testing.T) {
	// mock FileBrowserPlugin
	t.Logf("~> mock FileBrowserPlugin")
	p := file_browser_plugin.FileBrowserPlugin{
		Version: mockVersion,
	}
	// do FileBrowserPlugin
	t.Logf("~> do FileBrowserPlugin")

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

	t.Logf("-> now start test FileBrowserWorkMode %s", file_browser_plugin.WorkModeSend)
	p.Config.FileBrowserWorkMode = file_browser_plugin.WorkModeSend

	// start test mode
	p.Config.FileBrowserSendModeConfig.FileBrowserDistType = "other"
	err = p.Exec()
	if nil == err {
		t.Error("args FileBrowserDistType should be catch!")
	}

	p.Config.FileBrowserSendModeConfig.FileBrowserDistType = file_browser_plugin.DistTypeGit

	err = p.Exec()
	if nil == err {
		t.Error("args file_browser_remote_root_path should be catch!")
	}

	p.Config.FileBrowserSendModeConfig.FileBrowserRemoteRootPath = mockFileBrowserRemoteRootPath

	//err = p.Exec()
	//if nil == err {
	//	t.Error("args file_browser_target_dist_root_path should be catch!")
	//}

	//p.Config.FileBrowserSendModeConfig.FileBrowserTargetDistRootPath = mockFileBrowserTargetDistRootPath
	p.Config.FileBrowserSendModeConfig.FileBrowserTargetFileRegular = mockFileBrowserTargetFileRegularFail

	p.Config.TimeoutSecond = defTimeoutSecond
	p.Config.FileBrowserBaseConfig.FileBrowserTimeoutPushSecond = defTimeoutFileSecond

	p.Drone = *drone_info.MockDroneInfo("success")
	// verify FileBrowserPlugin
	assert.Equal(t, "sinlov", p.Drone.Repo.OwnerName)

	err = p.Exec()
	if err == nil {
		t.Fatal("args file browser want send file local path not find any file should be catch!")
	}

	// change right workspace
	p.Drone.Build.WorkSpace = testDataFolderAbsPath
	p.Config.FileBrowserBaseConfig.FileBrowserWorkSpace = p.Drone.Build.WorkSpace

	p.Config.FileBrowserSendModeConfig.FileBrowserShareLinkEnable = true
	p.Config.FileBrowserSendModeConfig.FileBrowserShareLinkAutoPasswordEnable = true
	p.Config.FileBrowserSendModeConfig.FileBrowserShareLinkUnit = web_api.ShareUnitHours
	p.Config.FileBrowserSendModeConfig.FileBrowserShareLinkExpires = 4

	p.Config.FileBrowserSendModeConfig.FileBrowserDistType = file_browser_plugin.DistTypeCustom
	p.Config.FileBrowserSendModeConfig.FileBrowserDistGraph = mockFileBrowserDistGraph

	// change FileRegular
	p.Config.FileBrowserSendModeConfig.FileBrowserTargetFileRegular = mockFileBrowserTargetFileRegularOne
	t.Logf("-> share by file browser work space: %s", p.Config.FileBrowserBaseConfig.FileBrowserWorkSpace)
	t.Logf("-> share by dist type: %s", p.Config.FileBrowserSendModeConfig.FileBrowserDistType)
	t.Logf("-> share by dist remote path: %s", p.Config.FileBrowserSendModeConfig.FileBrowserRemoteRootPath)
	t.Logf("-> share by file target path: %s", p.Config.FileBrowserSendModeConfig.FileBrowserTargetDistRootPath)
	t.Logf("-> share by file target regular: %s", p.Config.FileBrowserSendModeConfig.FileBrowserTargetFileRegular)

	err = p.Exec()
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, "", os.Getenv(file_browser_plugin.EnvPluginFileBrowserResultShareHost))
	assert.NotEqual(t, "", os.Getenv(file_browser_plugin.EnvPluginFileBrowserResultShareRemotePath))
	assert.NotEqual(t, "", os.Getenv(file_browser_plugin.EnvPluginFileBrowserResultSharePage))
	t.Logf("share host: %s", os.Getenv(file_browser_plugin.EnvPluginFileBrowserResultShareHost))
	t.Logf("share page: %s", os.Getenv(file_browser_plugin.EnvPluginFileBrowserResultSharePage))
	t.Logf("share password: %s", os.Getenv(file_browser_plugin.EnvPluginFileBrowserResultSharePasswd))
	t.Logf("share user: %s", os.Getenv(file_browser_plugin.EnvPluginFileBrowserResultShareUser))
	t.Logf("share path: %s", os.Getenv(file_browser_plugin.EnvPluginFileBrowserResultShareRemotePath))
	t.Logf("")

	p.Config.FileBrowserSendModeConfig.FileBrowserDistType = file_browser_plugin.DistTypeGit
	// change right file regular for one file
	p.Config.FileBrowserSendModeConfig.FileBrowserTargetFileRegular = mockFileBrowserTargetFileRegularOne

	t.Logf("-> share by file browser work space: %s", p.Config.FileBrowserBaseConfig.FileBrowserWorkSpace)
	t.Logf("-> share by dist type: %s", p.Config.FileBrowserSendModeConfig.FileBrowserDistType)
	t.Logf("-> share by dist remote path: %s", p.Config.FileBrowserSendModeConfig.FileBrowserRemoteRootPath)
	t.Logf("-> share by file target path: %s", p.Config.FileBrowserSendModeConfig.FileBrowserTargetDistRootPath)
	t.Logf("-> share by file target regular: %s", p.Config.FileBrowserSendModeConfig.FileBrowserTargetFileRegular)

	err = p.Exec()
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, "", os.Getenv(file_browser_plugin.EnvPluginFileBrowserResultShareHost))
	assert.NotEqual(t, "", os.Getenv(file_browser_plugin.EnvPluginFileBrowserResultShareRemotePath))
	assert.NotEqual(t, "", os.Getenv(file_browser_plugin.EnvPluginFileBrowserResultSharePage))
	t.Logf("share host: %s", os.Getenv(file_browser_plugin.EnvPluginFileBrowserResultShareHost))
	t.Logf("share page: %s", os.Getenv(file_browser_plugin.EnvPluginFileBrowserResultSharePage))
	t.Logf("share password: %s", os.Getenv(file_browser_plugin.EnvPluginFileBrowserResultSharePasswd))
	t.Logf("share user: %s", os.Getenv(file_browser_plugin.EnvPluginFileBrowserResultShareUser))
	t.Logf("share path: %s", os.Getenv(file_browser_plugin.EnvPluginFileBrowserResultShareRemotePath))
	t.Logf("")

	// change right file regular for more than one
	p.Config.FileBrowserSendModeConfig.FileBrowserTargetFileRegular = mockFileBrowserTargetFileRegular

	t.Logf("-> share by file browser work space: %s", p.Config.FileBrowserBaseConfig.FileBrowserWorkSpace)
	t.Logf("-> share by dist type: %s", p.Config.FileBrowserSendModeConfig.FileBrowserDistType)
	t.Logf("-> share by dist remote path: %s", p.Config.FileBrowserSendModeConfig.FileBrowserRemoteRootPath)
	t.Logf("-> share by file target path: %s", p.Config.FileBrowserSendModeConfig.FileBrowserTargetDistRootPath)
	t.Logf("-> share by file target regular: %s", p.Config.FileBrowserSendModeConfig.FileBrowserTargetFileRegular)

	err = p.Exec()
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, "", os.Getenv(file_browser_plugin.EnvPluginFileBrowserResultShareHost))
	assert.NotEqual(t, "", os.Getenv(file_browser_plugin.EnvPluginFileBrowserResultShareRemotePath))
	assert.NotEqual(t, "", os.Getenv(file_browser_plugin.EnvPluginFileBrowserResultSharePage))
	t.Logf("share host: %s", os.Getenv(file_browser_plugin.EnvPluginFileBrowserResultShareHost))
	t.Logf("share page: %s", os.Getenv(file_browser_plugin.EnvPluginFileBrowserResultSharePage))
	t.Logf("share password: %s", os.Getenv(file_browser_plugin.EnvPluginFileBrowserResultSharePasswd))
	t.Logf("share user: %s", os.Getenv(file_browser_plugin.EnvPluginFileBrowserResultShareUser))
	t.Logf("share path: %s", os.Getenv(file_browser_plugin.EnvPluginFileBrowserResultShareRemotePath))
	t.Logf("")

	// test Result
	err = p.CleanResultEnv()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "", os.Getenv(file_browser_plugin.EnvPluginFileBrowserResultShareHost))
	assert.Equal(t, "", os.Getenv(file_browser_plugin.EnvPluginFileBrowserResultShareRemotePath))
	assert.Equal(t, "", os.Getenv(file_browser_plugin.EnvPluginFileBrowserResultSharePage))
}
