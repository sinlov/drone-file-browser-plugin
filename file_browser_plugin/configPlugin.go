package file_browser_plugin

const (
	WorkModeSend     = "send"
	WorkModeDownload = "download"

	DistTypeGit    = "git"
	DistTypeCustom = "custom"

	EnvPluginDroneFileBrowserSharePage        = "PLUGIN_DRONE_FILE_BROWSER_SHARE_PAGE"
	EnvPluginDroneFileBrowserSharePasswd      = "PLUGIN_DRONE_FILE_BROWSER_SHARE_PASSWD"
	EnvPluginDroneFileBrowserShareDownloadUrl = "PLUGIN_DRONE_FILE_BROWSER_SHARE_DOWNLOAD_URL"
	EnvPluginDroneFileBrowserShareUser        = "PLUGIN_DRONE_FILE_BROWSER_SHARE_USER"
	EnvPluginDroneFileBrowserShareRemotePath  = "PLUGIN_DRONE_FILE_BROWSER_SHARE_REMOTE_PATH"
)

var (
	// pluginWorkModeSupport
	pluginWorkModeSupport = []string{
		WorkModeSend,
		WorkModeDownload,
	}

	// pluginDistTypeSupport
	pluginDistTypeSupport = []string{
		DistTypeGit,
		DistTypeCustom,
	}
)

type (
	FileBrowserBaseConfig struct {
		FileBrowserHost              string
		FileBrowserUsername          string
		FileBrowserUserPassword      string
		FileBrowserTimeoutPushSecond uint
		FileBrowserWorkSpace         string
	}

	FileBrowserSendModeConfig struct {
		FileBrowserRemoteRootPath     string
		FileBrowserDistType           string
		FileBrowserDistGraph          string
		FileBrowserTargetDistRootPath string
		FileBrowserTargetFileRegular  string
		FileBrowserShareLinkEnable    bool
		// FileBrowserShareLinkUnit
		// use [ web_api.ShareUnitDays web_api.ShareUnitHours
		// web_api.ShareUnitMinutes
		// web_api.ShareUnitSeconds ]
		FileBrowserShareLinkUnit               string
		FileBrowserShareLinkExpires            uint
		FileBrowserShareLinkAutoPasswordEnable bool
		FileBrowserShareLinkPassword           string
	}

	FileBrowserDownloadModeConfig struct {
		FileBrowserDownloadEnable    bool
		FileBrowserDownloadPath      string
		FileBrowserDownloadLocalPath string
	}

	// Config file_browser_plugin private config
	Config struct {
		Debug bool

		TimeoutSecond uint

		FileBrowserBaseConfig FileBrowserBaseConfig

		FileBrowserWorkMode string

		FileBrowserSendModeConfig FileBrowserSendModeConfig

		FileBrowserDownloadModeConfig FileBrowserDownloadModeConfig
	}
)
