package plugin

const (
	WorkModeSend     = "send"
	WorkModeDownload = "download"

	DistTypeGit    = "git"
	DistTypeCustom = "custom"
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
		FileBrowserDistType                    string
		FileBrowserDistGraph                   string
		FileBrowserRemoteRootPath              string
		FileBrowserTargetDistRootPath          string
		FileBrowserTargetFileRegular           string
		FileBrowserShareLinkEnable             bool
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
