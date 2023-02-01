package plugin

import (
	"fmt"
	"github.com/sinlov/drone-file-browser-plugin/drone_info"
	"github.com/sinlov/drone-file-browser-plugin/tools"
	"github.com/sinlov/filebrowser-client/file_browser_client"
	"log"
	"os"
	"strings"
)

type (
	// Plugin plugin all config
	Plugin struct {
		Drone             drone_info.Drone
		Config            Config
		fileBrowserClient file_browser_client.FileBrowserClient
	}
)

func (p *Plugin) Exec() error {
	if p.Config.Debug {
		for _, e := range os.Environ() {
			log.Println(e)
		}
	}

	var err error

	if p.Config.FileBrowserBaseConfig.FileBrowserHost == "" {
		msg := "missing file_browser_host, please check"
		return fmt.Errorf(msg)
	}

	if p.Config.FileBrowserBaseConfig.FileBrowserUsername == "" {
		msg := "missing file_browser_username, please check"
		return fmt.Errorf(msg)
	}

	if !(tools.StrInArr(p.Config.FileBrowserWorkMode, pluginWorkModeSupport)) {
		return fmt.Errorf("plugin file_browser_work_mode type only support: %v", pluginWorkModeSupport)
	}

	// check default TimeoutSecond
	if p.Config.TimeoutSecond < 10 {
		p.Config.TimeoutSecond = 10
	}
	// check default FileBrowserTimeoutPushSecond
	if p.Config.FileBrowserBaseConfig.FileBrowserTimeoutPushSecond < 30 {
		p.Config.FileBrowserBaseConfig.FileBrowserTimeoutPushSecond = 30
	}

	fileBrowserClient, err := file_browser_client.NewClient(
		p.Config.FileBrowserBaseConfig.FileBrowserUsername,
		p.Config.FileBrowserBaseConfig.FileBrowserUserPassword,
		p.Config.FileBrowserBaseConfig.FileBrowserHost,
		p.Config.TimeoutSecond,
		p.Config.FileBrowserBaseConfig.FileBrowserTimeoutPushSecond,
	)
	if err != nil {
		return err
	}

	fileBrowserClient.Debug(p.Config.Debug)

	p.fileBrowserClient = fileBrowserClient

	switch p.Config.FileBrowserWorkMode {
	default:
		return fmt.Errorf("plugin file_browser_work_mode not support: %v", p.Config.FileBrowserWorkMode)
	case WorkModeSend:
		err = workOnSend(p)
	}

	return err
}

func workOnSend(p *Plugin) error {
	sendModeConfig := p.Config.FileBrowserSendModeConfig
	if !(tools.StrInArr(sendModeConfig.FileBrowserDistType, pluginDistTypeSupport)) {
		return fmt.Errorf("plugin file_browser_dist_type dist type only support: %v", pluginDistTypeSupport)
	}

	if sendModeConfig.FileBrowserRemoteRootPath == "" {
		return fmt.Errorf("plugin file_browser_remote_root_path not be empty")
	}

	if sendModeConfig.FileBrowserTargetDistRootPath == "" {
		return fmt.Errorf("plugin file_browser_target_dist_root_path not be empty")
	}

	var remoteRealPath = strings.TrimRight(sendModeConfig.FileBrowserRemoteRootPath, "/")

	switch sendModeConfig.FileBrowserDistType {
	default:
		return fmt.Errorf("send dist type not support %s", sendModeConfig.FileBrowserDistType)
	case DistTypeGit:
		if p.Drone.Build.Tag == "" {
			remoteRealPath = fmt.Sprintf("%s/%s/%s/%s/%d/%s/%s",
				remoteRealPath, p.Drone.Repo.GroupName, p.Drone.Repo.ShortName,
				"b",
				p.Drone.Build.Number,
				p.Drone.Commit.Branch, string([]rune(p.Drone.Commit.Sha))[:8],
			)
		} else {
			remoteRealPath = fmt.Sprintf("%s/%s/%s/%s/%s/%d/%s",
				remoteRealPath, p.Drone.Repo.GroupName, p.Drone.Repo.ShortName,
				"tag",
				p.Drone.Build.Tag,
				p.Drone.Build.Number,
				string([]rune(p.Drone.Commit.Sha))[:8],
			)
		}

	}
	if p.Config.Debug {
		log.Printf("debug: workOnSend remoteRealPath: %s", remoteRealPath)
	}

	return nil
}
