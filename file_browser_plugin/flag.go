package file_browser_plugin

import (
	"fmt"
	"github.com/sinlov/drone-info-tools/drone_info"
	"github.com/sinlov/filebrowser-client/web_api"
	"github.com/urfave/cli/v2"
)

func BindFlag(c *cli.Context, cliVersion, cliName string, drone drone_info.Drone) FileBrowserPlugin {
	config := Config{

		Debug:         c.Bool("config.debug"),
		TimeoutSecond: c.Uint("config.timeout_second"),

		FileBrowserBaseConfig: FileBrowserBaseConfig{
			FileBrowserHost:              c.String("config.file_browser_host"),
			FileBrowserUsername:          c.String("config.file_browser_username"),
			FileBrowserUserPassword:      c.String("config.file_browser_user_password"),
			FileBrowserTimeoutPushSecond: c.Uint("config.file_browser_timeout_push_second"),
			FileBrowserWorkSpace:         c.String("config.file_browser_work_space"),
		},

		FileBrowserWorkMode: c.String("config.file_browser_work_mode"),

		FileBrowserSendModeConfig: FileBrowserSendModeConfig{
			FileBrowserDistType:           c.String("config.file_browser_dist_type"),
			FileBrowserDistGraph:          c.String("config.file_browser_dist_graph"),
			FileBrowserRemoteRootPath:     c.String("config.file_browser_remote_root_path"),
			FileBrowserTargetDistRootPath: c.String("config.file_browser_target_dist_root_path"),
			FileBrowserTargetFileGlob:     c.StringSlice("config.file_browser_target_file_globs"),
			FileBrowserTargetFileRegular:  c.String("config.file_browser_target_file_regular"),

			FileBrowserShareLinkEnable:             c.Bool("config.file_browser_share_link_enable"),
			FileBrowserShareLinkUnit:               c.String("config.file_browser_share_link_unit"),
			FileBrowserShareLinkExpires:            c.Uint("config.file_browser_share_link_expires"),
			FileBrowserShareLinkAutoPasswordEnable: c.Bool("config.file_browser_share_link_auto_password_enable"),
			FileBrowserShareLinkPassword:           c.String("config.file_browser_share_link_password"),
		},

		FileBrowserDownloadModeConfig: FileBrowserDownloadModeConfig{
			FileBrowserDownloadEnable:    c.Bool("config.file_browser_download_enable"),
			FileBrowserDownloadPath:      c.String("config.file_browser_download_remote_path"),
			FileBrowserDownloadLocalPath: c.String("config.file_browser_download_local_path"),
		},
	}

	p := FileBrowserPlugin{
		Name:    cliName,
		Version: cliVersion,
		Drone:   drone,
		Config:  config,
	}
	return p
}

// Flag
// set flag at here
func Flag() []cli.Flag {
	return []cli.Flag{
		// file_browser_plugin start
		&cli.StringFlag{
			Name:    "config.file_browser_host,file_browser_host",
			Usage:   "must set args, file_browser host",
			EnvVars: []string{"PLUGIN_FILE_BROWSER_HOST"},
		},
		&cli.StringFlag{
			Name:    "config.file_browser_username,file_browser_username",
			Usage:   "must set args, file_browser username",
			EnvVars: []string{"PLUGIN_FILE_BROWSER_USERNAME"},
		},
		&cli.StringFlag{
			Name:    "config.file_browser_user_password,file_browser_user_password",
			Usage:   "must set args, file_browser user password",
			EnvVars: []string{"PLUGIN_FILE_BROWSER_USER_PASSWORD"},
		},
		&cli.StringFlag{
			Name:    "config.file_browser_work_space,file_browser_work_space",
			Usage:   fmt.Sprintf("file_browser work space. default will use env:%s", drone_info.EnvDroneBuildWorkSpace),
			EnvVars: []string{"PLUGIN_FILE_BROWSER_WORK_SPACE"},
		},
		&cli.UintFlag{
			Name:    "config.file_browser_timeout_push_second,file_browser_timeout_push_second",
			Usage:   "file_browser push each file timeout push second, must gather than 60",
			Value:   60,
			EnvVars: []string{"PLUGIN_FILE_BROWSER_TIMEOUT_PUSH_SECOND"},
		},
		&cli.StringFlag{
			Name:    "config.file_browser_work_mode,file_browser_work_mode",
			Usage:   "must set args, work mode only can use: send, download",
			Value:   WorkModeSend,
			EnvVars: []string{"PLUGIN_FILE_BROWSER_WORK_MODE"},
		},
		&cli.StringFlag{
			Name:    "config.file_browser_dist_type,file_browser_dist_type",
			Usage:   "must set args, type of dist file graph only can use: git, custom",
			Value:   DistTypeGit,
			EnvVars: []string{"PLUGIN_FILE_BROWSER_DIST_TYPE"},
		},
		&cli.StringFlag{
			Name:    "config.file_browser_dist_graph,file_browser_dist_graph",
			Usage:   "type of dist custom set as struct [ drone_info.Drone ]",
			Value:   "{{ Repo.HostName }}/{{ Repo.GroupName }}/{{ Repo.ShortName }}/s/{{ Build.Number }}/{{ Stage.Name }}-{{ Build.Number }}-{{ Stage.FinishedTime }}",
			EnvVars: []string{"PLUGIN_FILE_BROWSER_DIST_GRAPH"},
		},
		&cli.StringFlag{
			Name:    "config.file_browser_remote_root_path,file_browser_remote_root_path",
			Usage:   "must set args, this will append by file_browser_dist_type at remote",
			EnvVars: []string{"PLUGIN_FILE_BROWSER_REMOTE_ROOT_PATH"},
		},
		&cli.StringFlag{
			Name:    "config.file_browser_target_dist_root_path,file_browser_target_dist_root_path",
			Usage:   "path of file_browser local work on root, can set \"\"",
			Value:   "",
			EnvVars: []string{"PLUGIN_FILE_BROWSER_TARGET_DIST_ROOT_PATH"},
		},
		&cli.StringSliceFlag{
			Name:    "config.file_browser_target_file_globs,file_browser_target_file_globs",
			Usage:   "must set args, globs list of send to file_browser under file_browser_target_dist_root_path",
			EnvVars: []string{"PLUGIN_FILE_BROWSER_TARGET_FILE_GLOBS"},
		},
		&cli.StringFlag{
			Name:    "config.file_browser_target_file_regular,file_browser_target_file_regular",
			Usage:   "must set args, regular of send to file_browser under file_browser_target_dist_root_path",
			EnvVars: []string{"PLUGIN_FILE_BROWSER_TARGET_FILE_REGULAR"},
		},
		&cli.BoolFlag{
			Name:    "config.file_browser_share_link_enable,file_browser_share_link_enable",
			Usage:   "share dist dir as link",
			Value:   true,
			EnvVars: []string{"PLUGIN_FILE_BROWSER_SHARE_LINK_ENABLE"},
		},
		&cli.StringFlag{
			Name:    "config.file_browser_share_link_unit,file_browser_share_link_unit",
			Usage:   "take effect by open share_link, only can use as [ days hours minutes seconds ]",
			Value:   web_api.ShareUnitDays,
			EnvVars: []string{"PLUGIN_FILE_BROWSER_SHARE_LINK_UNIT"},
		},
		&cli.UintFlag{
			Name:    "config.file_browser_share_link_expires,file_browser_share_link_expires",
			Usage:   "if set 0, will allow share_link exist forever, default: 0",
			Value:   0,
			EnvVars: []string{"PLUGIN_FILE_BROWSER_SHARE_LINK_EXPIRES"},
		},
		&cli.BoolFlag{
			Name:    "config.file_browser_share_link_auto_password_enable,file_browser_share_link_auto_password_enable",
			Usage:   "password of share_link auto , if open this will cover settings.file_browser_share_link_password",
			Value:   false,
			EnvVars: []string{"PLUGIN_FILE_BROWSER_SHARE_LINK_AUTO_PASSWORD_ENABLE"},
		},
		&cli.StringFlag{
			Name:    "config.file_browser_share_link_password,file_browser_share_link_password",
			Usage:   "password of share_link, if not set will not use password, default: \"\"",
			Value:   "",
			EnvVars: []string{"PLUGIN_FILE_BROWSER_SHARE_LINK_PASSWORD"},
		},
		//&cli.StringFlag{
		//	Name:    "config.new_arg,new_arg",
		//	Usage:   "",
		//	EnvVars: []string{"PLUGIN_new_arg"},
		//},
		// file_browser_plugin end
	}
}

// CommonFlag
// Other modules also have flags
func CommonFlag() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:    "config.debug,debug",
			Usage:   "debug mode",
			Value:   false,
			EnvVars: []string{"PLUGIN_DEBUG"},
		},
		&cli.UintFlag{
			Name:    "config.timeout_second,timeout_second",
			Usage:   "do request timeout setting second. gather than 10",
			Hidden:  true,
			Value:   10,
			EnvVars: []string{"PLUGIN_TIMEOUT_SECOND"},
		},
	}
}

// HideFlag
// set file_browser_plugin hide flag at here
func HideFlag() []cli.Flag {
	return []cli.Flag{
		// file_browser_plugin hidden start
		&cli.BoolFlag{
			Name:    "config.file_browser_download_enable,file_browser_download_enable",
			Usage:   "file_browser download mode, if use this mode only can check or download. default: false",
			Value:   false,
			Hidden:  true,
			EnvVars: []string{"PLUGIN_FILE_BROWSER_DOWNLOAD_ENABLE"},
		},
		&cli.StringFlag{
			Name:    "config.file_browser_download_remote_path,file_browser_download_remote_path",
			Usage:   "file_browser download mode use remote path",
			Hidden:  true,
			EnvVars: []string{"PLUGIN_FILE_BROWSER_DOWNLOAD_REMOTE_PATH"},
		},
		&cli.StringFlag{
			Name:    "config.file_browser_download_local_path,file_browser_download_local_path",
			Usage:   "file_browser download mode local path",
			Hidden:  true,
			EnvVars: []string{"PLUGIN_FILE_BROWSER_DOWNLOAD_LOCAL_PATH"},
		},
		// file_browser_plugin hidden end
	}
}
