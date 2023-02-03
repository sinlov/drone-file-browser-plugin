package main

import (
	_ "embed"
	"fmt"
	"github.com/sinlov/drone-file-browser-plugin/file_browser_plugin"
	"github.com/sinlov/drone-info-tools/pkgJson"
	"github.com/sinlov/drone-info-tools/template"
	"github.com/sinlov/filebrowser-client/web_api"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/sinlov/drone-info-tools/drone_info"
	"github.com/sinlov/drone-info-tools/drone_urfave_cli_v2"
	"github.com/urfave/cli/v2"
)

const (
	Name = "drone-file-browser-plugin"
)

//go:embed package.json
var packageJson string

func action(c *cli.Context) error {

	isDebug := c.Bool("config.debug")

	drone := drone_urfave_cli_v2.UrfaveCliBindDroneInfo(c)

	cliVersion := pkgJson.GetPackageJsonVersionGoStyle()
	if isDebug {
		log.Printf("debug: cli version is %s", cliVersion)
		log.Printf("debug: load droneInfo finish at link: %v\n", drone.Build.Link)
	}

	config := file_browser_plugin.Config{

		Debug:         c.Bool("config.debug"),
		TimeoutSecond: c.Uint("config.timeout_second"),

		FileBrowserBaseConfig: file_browser_plugin.FileBrowserBaseConfig{
			FileBrowserHost:              c.String("config.file_browser_host"),
			FileBrowserUsername:          c.String("config.file_browser_username"),
			FileBrowserUserPassword:      c.String("config.file_browser_user_password"),
			FileBrowserTimeoutPushSecond: c.Uint("config.file_browser_timeout_push_second"),
			FileBrowserWorkSpace:         c.String("config.file_browser_work_space"),
		},

		FileBrowserWorkMode: c.String("config.file_browser_work_mode"),

		FileBrowserSendModeConfig: file_browser_plugin.FileBrowserSendModeConfig{
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

		FileBrowserDownloadModeConfig: file_browser_plugin.FileBrowserDownloadModeConfig{
			FileBrowserDownloadEnable:    c.Bool("config.file_browser_download_enable"),
			FileBrowserDownloadPath:      c.String("config.file_browser_download_remote_path"),
			FileBrowserDownloadLocalPath: c.String("config.file_browser_download_local_path"),
		},
	}

	p := file_browser_plugin.FileBrowserPlugin{
		Name:    Name,
		Version: cliVersion,
		Drone:   drone,
		Config:  config,
	}
	err := p.Exec()

	if err != nil {
		log.Fatalf("err: %v", err)
		return err
	}

	return nil
}

// pluginFlag
// set file_browser_plugin flag at here
func pluginFlag() []cli.Flag {
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
			Value:   file_browser_plugin.WorkModeSend,
			EnvVars: []string{"PLUGIN_FILE_BROWSER_WORK_MODE"},
		},
		&cli.StringFlag{
			Name:    "config.file_browser_dist_type,file_browser_dist_type",
			Usage:   "must set args, type of dist file graph only can use: git, custom",
			Value:   file_browser_plugin.DistTypeGit,
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
		&cli.BoolFlag{
			Name:    "config.debug,debug",
			Usage:   "debug mode",
			Value:   false,
			EnvVars: []string{"PLUGIN_DEBUG"},
		},
		//&cli.StringFlag{
		//	Name:    "config.new_arg,new_arg",
		//	Usage:   "",
		//	EnvVars: []string{"PLUGIN_new_arg"},
		//},
		// file_browser_plugin end
	}
}

// pluginHideFlag
// set file_browser_plugin hide flag at here
func pluginHideFlag() []cli.Flag {
	return []cli.Flag{
		// file_browser_plugin hidden start
		&cli.UintFlag{
			Name:    "config.timeout_second,timeout_second",
			Usage:   "do request timeout setting second. gather than 10",
			Hidden:  true,
			Value:   10,
			EnvVars: []string{"PLUGIN_TIMEOUT_SECOND"},
		},
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

// droneInfoFlag
// Please do not edit unless you understand drone's environment variables.
func droneInfoFlag() []cli.Flag {
	return []cli.Flag{
		// droneInfo start
		&cli.StringFlag{
			Name:    "repo.name",
			Usage:   "providers the name of the repository",
			EnvVars: []string{drone_info.EnvDroneRepoName},
		},
		&cli.StringFlag{
			Name:    "repo.group",
			Usage:   "providers the group of the repository",
			EnvVars: []string{drone_info.EnvDroneRepoNamespace},
		},
		&cli.StringFlag{
			Name:    "repo.full.name",
			Usage:   "providers the full name of the repository",
			EnvVars: []string{drone_info.EnvDroneRepo},
		},
		&cli.StringFlag{
			Name:    "repo.owner",
			Usage:   "providers the owner of the repository",
			EnvVars: []string{drone_info.EnvDroneRepoOwner},
		},
		&cli.StringFlag{
			Name:    "repo.scm",
			Usage:   "Provides the repository type for the current running build",
			EnvVars: []string{drone_info.EnvDroneRepoScm},
		},
		&cli.StringFlag{
			Name:    "repo.remote.url",
			Usage:   "Provides the git+ssh url that should be used to clone the repository.",
			EnvVars: []string{drone_info.EnvDroneRemoteUrl},
		},
		&cli.StringFlag{
			Name:    "repo.http.url",
			Usage:   "Provides the http url that should be used to clone the repository",
			EnvVars: []string{drone_info.EnvDroneGitHttpUrl},
		},
		&cli.StringFlag{
			Name:    "repo.ssh.url",
			Usage:   "Provides the ssh url that should be used to clone the repository",
			EnvVars: []string{drone_info.EnvDroneGitSshUrl},
		},

		// drone_info.Build
		&cli.StringFlag{
			Name:    "build.workspace",
			Usage:   "drone’s working directory for a pipeline",
			EnvVars: []string{drone_info.EnvDroneBuildWorkSpace},
		},
		&cli.StringFlag{
			Name:    "build.status",
			Usage:   "build status",
			Value:   "success",
			EnvVars: []string{drone_info.EnvDroneBuildStatus},
		},
		&cli.Uint64Flag{
			Name:    "build.number",
			Usage:   "providers the current build number",
			EnvVars: []string{drone_info.EnvDroneBuildNumber},
		},
		&cli.StringFlag{
			Name:    "build.tag",
			Usage:   "build tag",
			EnvVars: []string{drone_info.EnvDroneTag},
		},
		&cli.StringFlag{
			Name:    "build.target_branch",
			Usage:   "This environment variable can be used in conjunction with the source branch variable to get the pull request base and head branch.",
			EnvVars: []string{drone_info.EnvDroneTargetBranch},
		},
		&cli.StringFlag{
			Name:    "build.link",
			Usage:   "build link",
			EnvVars: []string{drone_info.EnvDroneBuildLink},
		},
		&cli.StringFlag{
			Name:    "build.event",
			Usage:   "build event",
			EnvVars: []string{drone_info.EnvDroneBuildEvent},
		},
		&cli.Uint64Flag{
			Name:    "build.started",
			Usage:   "build started",
			EnvVars: []string{drone_info.EnvDroneBuildStarted},
		},
		&cli.Uint64Flag{
			Name:    "build.finished",
			Usage:   "build finished",
			EnvVars: []string{drone_info.EnvDroneBuildFinished},
		},
		&cli.StringFlag{
			Name:    "pull.request",
			Usage:   "pull request",
			EnvVars: []string{drone_info.EnvDronePR},
		},
		&cli.StringFlag{
			Name:    "deploy.to",
			Usage:   "provides the target deployment environment for the running build. This value is only available to promotion and rollback pipelines.",
			EnvVars: []string{drone_info.EnvDroneDeployTo},
		},
		&cli.StringFlag{
			Name:    "failed.stages",
			Usage:   "Provides a comma-separate list of failed pipeline stages for the current running build.",
			EnvVars: []string{drone_info.EnvDroneFailedStages},
		},
		&cli.StringFlag{
			Name:    "failed.steps",
			Usage:   "Provides a comma-separate list of failed pipeline steps",
			EnvVars: []string{drone_info.EnvDroneFailedSteps},
		},

		&cli.StringFlag{
			Name:    "commit.author.username",
			Usage:   "Provides the commit author name for the current running build. Note this is a user-defined value and may be empty or inaccurate",
			EnvVars: []string{drone_info.EnvDroneCommitAuthorName},
		},
		&cli.StringFlag{
			Name:    "commit.author.email",
			Usage:   "Provides the commit email address for the current running build. Note this is a user-defined value and may be empty or inaccurate",
			EnvVars: []string{drone_info.EnvDroneCommitAuthorEmail},
		},
		&cli.StringFlag{
			Name:    "commit.author.name",
			Usage:   "Provides the commit author username for the current running build. This is the username from source control management system (e.g. GitHub username)",
			EnvVars: []string{drone_info.EnvDroneCommitAuthor},
		},
		&cli.StringFlag{
			Name:    "commit.author.avatar",
			Usage:   "Provides the commit author avatar for the current running build. This is the avatar from source control management system (e.g. GitHub)",
			EnvVars: []string{drone_info.EnvDroneCommitAuthorAvatar},
		},
		&cli.StringFlag{
			Name:    "commit.link",
			Usage:   "providers the http link to the current commit in the remote source code management system(e.g.GitHub)",
			EnvVars: []string{drone_info.EnvDroneCommitLink},
		},
		&cli.StringFlag{
			Name:    "commit.branch",
			Usage:   "providers the branch for the current build",
			EnvVars: []string{drone_info.EnvDroneCommitBranch},
			Value:   "master",
		},
		&cli.StringFlag{
			Name:    "commit.message",
			Usage:   "providers the commit message for the current build",
			EnvVars: []string{drone_info.EnvDroneCommitMessage},
		},
		&cli.StringFlag{
			Name:    "commit.sha",
			Usage:   "providers the commit sha for the current build",
			EnvVars: []string{drone_info.EnvDroneCommitSha},
		},
		&cli.StringFlag{
			Name:    "commit.ref",
			Usage:   "providers the commit ref for the current build",
			EnvVars: []string{drone_info.EnvDroneCommitRef},
		},

		// drone_info.Stage
		&cli.Uint64Flag{
			Name:    "stage.started",
			Usage:   "stage started ",
			EnvVars: []string{drone_info.EnvDroneStageStarted},
		},
		&cli.Uint64Flag{
			Name:    "stage.finished",
			Usage:   "stage finished",
			EnvVars: []string{drone_info.EnvDroneStageFinished},
		},
		&cli.StringFlag{
			Name:    "stage.machine",
			Usage:   "stage machine",
			EnvVars: []string{drone_info.EnvDroneStageMachine},
		},
		&cli.StringFlag{
			Name:    "stage.os",
			Usage:   "stage OS",
			EnvVars: []string{drone_info.EnvDroneStageOs},
		},
		&cli.StringFlag{
			Name:    "stage.arch",
			Usage:   "stage arch",
			EnvVars: []string{drone_info.EnvDroneStageArch},
		},
		&cli.StringFlag{
			Name:    "stage.variant",
			Usage:   "stage variant",
			EnvVars: []string{drone_info.EnvDroneStageVariant},
		},
		&cli.StringFlag{
			Name:    "stage.type",
			Usage:   "stage type",
			EnvVars: []string{drone_info.EnvDroneStageType},
		},
		&cli.StringFlag{
			Name:    "stage.kind",
			Usage:   "stage kind",
			EnvVars: []string{drone_info.EnvDroneStageKind},
		},
		&cli.StringFlag{
			Name:    "stage.name",
			Usage:   "stage name",
			EnvVars: []string{drone_info.EnvDroneStageName},
		},

		// drone_info.DroneSystem
		&cli.StringFlag{
			Name:    "drone.system.version",
			Usage:   "Provides the version of the Drone server.",
			EnvVars: []string{drone_info.EnvDroneSystemVersion},
		},
		&cli.StringFlag{
			Name:    "drone.system.host",
			Usage:   "Provides the host used by the Drone server. This can be combined with the protocol to construct to the server url.",
			EnvVars: []string{drone_info.EnvDroneSystemHost},
		},
		&cli.StringFlag{
			Name:    "drone.system.hostname",
			Usage:   "Provides the hostname used by the Drone server. This can be combined with the protocol to construct to the server url.",
			EnvVars: []string{drone_info.EnvDroneSystemHostName},
		},
		&cli.StringFlag{
			Name:    "drone.system.proto",
			Usage:   "Provides the protocol used by the Drone server. This can be combined with the hostname to construct to the server url.",
			EnvVars: []string{drone_info.EnvDroneSystemProto},
		},
		// droneInfo end
	}
}

func main() {
	pkgJson.InitPkgJsonContent(packageJson)
	template.RegisterSettings(template.DefaultFunctions)
	app := cli.NewApp()
	app.Version = pkgJson.GetPackageJsonVersionGoStyle()
	app.Name = "Drone FileBrowserPlugin"
	app.Usage = ""
	year := time.Now().Year()
	app.Copyright = fmt.Sprintf("© 2022-%d sinlov", year)
	author := &cli.Author{
		Name:  "sinlov",
		Email: "sinlovgmppt@gmail.com",
	}
	app.Authors = []*cli.Author{
		author,
	}

	app.Action = action
	flags := drone_urfave_cli_v2.UrfaveCliAppendCliFlag(drone_urfave_cli_v2.DroneInfoUrfaveCliFlag(), pluginFlag())
	flags = drone_urfave_cli_v2.UrfaveCliAppendCliFlag(flags, pluginHideFlag())
	app.Flags = flags

	// kubernetes runner patch
	if _, err := os.Stat("/run/drone/env"); err == nil {
		errDotEnv := godotenv.Overload("/run/drone/env")
		if errDotEnv != nil {
			log.Fatalf("load /run/drone/env err: %v", errDotEnv)
		}
	}

	// app run as urfave
	if err := app.Run(os.Args); nil != err {
		log.Println(err)
	}
}
