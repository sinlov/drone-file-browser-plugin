package plugin

const (
	WorkModeSend     = "send"
	WorkModeDownload = "download"

	DistTypeGit    = "git"
	DistTypeCustom = "custom"

	EnvPluginDroneFileBrowserShareRemotePath  = "PLUGIN_DRONE_FILE_BROWSER_SHARE_REMOTE_PATH"
	EnvPluginDroneFileBrowserSharePage        = "PLUGIN_DRONE_FILE_BROWSER_SHARE_PAGE"
	EnvPluginDroneFileBrowserSharePasswd      = "PLUGIN_DRONE_FILE_BROWSER_SHARE_PASSWD"
	EnvPluginDroneFileBrowserShareDownloadUrl = "PLUGIN_DRONE_FILE_BROWSER_SHARE_DOWNLOAD_URL"
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
	}

	FileBrowserSendModeConfig struct {
		FileBrowserDistType           string
		FileBrowserDistGraph          string
		FileBrowserRemoteRootPath     string
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

	// Config plugin private config
	Config struct {
		Debug bool

		TimeoutSecond uint

		FileBrowserBaseConfig FileBrowserBaseConfig

		FileBrowserWorkMode string

		FileBrowserSendModeConfig FileBrowserSendModeConfig

		FileBrowserDownloadModeConfig FileBrowserDownloadModeConfig
	}
)
