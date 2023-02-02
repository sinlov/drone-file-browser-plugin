package file_browser_plugin

const (
	WorkModeSend     = "send"
	WorkModeDownload = "download"

	DistTypeGit    = "git"
	DistTypeCustom = "custom"

	EnvPluginFileBrowserResultShareHost        = "PLUGIN_FILE_BROWSER_RESULT_SHARE_HOST"
	EnvPluginFileBrowserResultSharePage        = "PLUGIN_FILE_BROWSER_RESULT_SHARE_PAGE"
	EnvPluginFileBrowserResultSharePasswd      = "PLUGIN_FILE_BROWSER_RESULT_SHARE_PASSWD"
	EnvPluginFileBrowserResultShareDownloadUrl = "PLUGIN_FILE_BROWSER_RESULT_SHARE_DOWNLOAD_URL"
	EnvPluginFileBrowserResultShareUser        = "PLUGIN_FILE_BROWSER_RESULT_SHARE_USER"
	EnvPluginFileBrowserResultShareRemotePath  = "PLUGIN_FILE_BROWSER_RESULT_SHARE_REMOTE_PATH"
)

var (
	// cleanResultEnvList
	// for clean result
	cleanResultEnvList = []string{
		EnvPluginFileBrowserResultShareHost,
		EnvPluginFileBrowserResultSharePage,
		EnvPluginFileBrowserResultSharePasswd,
		EnvPluginFileBrowserResultShareDownloadUrl,
		EnvPluginFileBrowserResultShareUser,
		EnvPluginFileBrowserResultShareRemotePath,
	}
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
