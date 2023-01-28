package plugin

const (
	msgTypeText        = "text"
	msgTypePost        = "post"
	msgTypeInteractive = "interactive"
)

var (
	// supportMsgType
	supportMsgType = []string{
		msgTypeText,
		msgTypePost,
		msgTypeInteractive,
	}
)

type (
	// Config plugin private config
	Config struct {
		Webhook string
		Secret  string
		MsgType string

		FileBrowserTimeoutPushMin              uint
		FileBrowserUsername                    string
		FileBrowserUserPassword                string
		FileBrowserDistRoot                    string
		FileBrowserDistType                    string
		FileBrowserDistGraph                   string
		FileBrowserTargetFileRegular           string
		FileBrowserShareLinkEnable             bool
		FileBrowserShareLinkUnit               string
		FileBrowserShareLinkExpires            uint
		FileBrowserShareLinkAutoPasswordEnable bool
		FileBrowserShareLinkPassword           string

		Debug bool

		TimeoutSecond                uint
		DownloadEnable               bool
		FileBrowserDownloadPath      string
		FileBrowserDownloadLocalPath string
	}
)
