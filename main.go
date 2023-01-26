package main

import (
	"fmt"
	"github.com/sinlov/drone-file-browser-plugin/plugin"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/sinlov/drone-file-browser-plugin/drone_info"
	"github.com/urfave/cli/v2"
)

// Version of cli
var Version = "v0.1.2"

func action(c *cli.Context) error {

	isDebug := c.Bool("config.debug")

	drone := bindDroneInfo(c)

	if isDebug {
		log.Printf("load droneInfo finish at link: %v\n", drone.Build.Link)
	}

	config := plugin.Config{
		Debug:         c.Bool("config.debug"),
		TimeoutSecond: c.Int("config.timeout_second"),
		Webhook:       c.String("config.webhook"),
		Secret:        c.String("config.secret"),
		MsgType:       c.String("config.msg_type"),
	}

	p := plugin.Plugin{
		Drone:  *drone,
		Config: config,
	}
	err := p.Exec()

	if err != nil {
		log.Fatalf("err: %v", err)
		return err
	}

	return nil
}

func main() {
	app := cli.NewApp()
	app.Version = Version
	app.Name = "Drone Plugin"
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
	app.Flags = []cli.Flag{
		// plugin hidden start
		&cli.BoolFlag{
			Name:    "config.download_enable,download_enable",
			Usage:   "file_browser download mode, if use this mode only download. default: false",
			Value:   false,
			Hidden:  true,
			EnvVars: []string{"PLUGIN_DOWNLOAD_ENABLE"},
		},
		&cli.StringFlag{
			Name:    "config.file_browser_download_path,file_browser_download_path",
			Usage:   "file_browser download mode use remote path",
			Hidden:  true,
			EnvVars: []string{"PLUGIN_FILE_BROWSER_DOWNLOAD_PATH"},
		},
		&cli.StringFlag{
			Name:    "config.file_browser_download_local_path,file_browser_download_local_path",
			Usage:   "file_browser download mode local path",
			Hidden:  true,
			EnvVars: []string{"PLUGIN_FILE_BROWSER_DOWNLOAD_LOCAL_PATH"},
		},
		// plugin hidden end

		// plugin start
		&cli.BoolFlag{
			Name:    "config.debug,debug",
			Usage:   "debug mode",
			Value:   false,
			EnvVars: []string{"PLUGIN_DEBUG"},
		},
		&cli.UintFlag{
			Name:    "config.timeout_second,timeout_second",
			Usage:   "do request timeout setting second. default: 10",
			Value:   10,
			EnvVars: []string{"PLUGIN_TIMEOUT_SECOND"},
		},
		&cli.UintFlag{
			Name:    "config.file_browser_timeout_push_min,file_browser_timeout_push_minute",
			Usage:   "file_browser push each file timeout push minute default: 1",
			Value:   1,
			EnvVars: []string{"PLUGIN_FILE_BROWSER_TIMEOUT_PUSH_MINUTE"},
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
			Name:    "config.file_browser_dist_root,file_browser_dist_root",
			Usage:   "must set args, path of file_browser root like: dist/",
			EnvVars: []string{"PLUGIN_FILE_BROWSER_DIST_ROOT"},
		},
		&cli.StringFlag{
			Name:    "config.file_browser_dist_type,file_browser_dist_type",
			Usage:   "must set args, type of dist file graph only can use: git, custom",
			Value:   "git",
			EnvVars: []string{"PLUGIN_FILE_BROWSER_DIST_TYPE"},
		},
		&cli.StringFlag{
			Name:    "config.file_browser_dist_graph,file_browser_dist_graph",
			Usage:   "type of dist custom set as: {{ Commit.Branch }}/{{ Commit.Sha }}/{{ Build.FinishedAt }}/{{ Build.Tag }}",
			EnvVars: []string{"PLUGIN_FILE_BROWSER_DIST_GRAPH"},
		},
		&cli.StringFlag{
			Name:    "config.file_browser_target_file_regular,file_browser_target_file_regular",
			Usage:   "must set args, regular of send to file_browser",
			EnvVars: []string{"PLUGIN_FILE_BROWSER_TARGET_FILE_REGULAR"},
		},
		&cli.BoolFlag{
			Name:    "config.file_browser_share_link_enable,file_browser_share_link_enable",
			Usage:   "share dist dir as link, default: true",
			Value:   true,
			EnvVars: []string{"PLUGIN_FILE_BROWSER_SHARE_LINK_ENABLE"},
		},
		&cli.StringFlag{
			Name:    "config.file_browser_share_link_unit,file_browser_share_link_unit",
			Usage:   "take effect by open share_link, only can use as [ days hours minutes seconds ]",
			Value:   "days",
			EnvVars: []string{"PLUGIN_FILE_BROWSER_SHARE_LINK_UNIT"},
		},
		&cli.UintFlag{
			Name:    "config.file_browser_share_link_expires,file_browser_share_link_expires",
			Usage:   "if set 0, will allow share_link exist forever，default: 0",
			Value:   0,
			EnvVars: []string{"PLUGIN_FILE_BROWSER_SHARE_LINK_EXPIRES"},
		},
		&cli.BoolFlag{
			Name:    "config.file_browser_share_link_auto_password_enable,file_browser_share_link_auto_password_enable",
			Usage:   "password of share_link auto , if open this will cover settings.file_browser_share_link_password default: false",
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
		// plugin end
		// droneInfo start
		&cli.StringFlag{
			Name:    "commit.author.username",
			Usage:   "providers the author username for the current commit",
			EnvVars: []string{drone_info.EnvDroneCommitAuthor},
		},
		&cli.StringFlag{
			Name:    "commit.author.avatar",
			Usage:   "providers the author avatar url for the current commit",
			EnvVars: []string{drone_info.EnvDroneCommitAuthorAvatar},
		},
		&cli.StringFlag{
			Name:    "commit.author.email",
			Usage:   "providers the author email for the current commit",
			EnvVars: []string{drone_info.EnvDroneCommitAuthorEmail},
		},
		&cli.StringFlag{
			Name:    "commit.author.name",
			Usage:   "providers the author name for the current commit",
			EnvVars: []string{drone_info.EnvDroneCommitAuthor},
		},
		&cli.StringFlag{
			Name:    "commit.branch",
			Usage:   "providers the branch for the current build",
			EnvVars: []string{drone_info.EnvDroneCommitBranch},
			Value:   "master",
		},
		&cli.StringFlag{
			Name:    "commit.link",
			Usage:   "providers the http link to the current commit in the remote source code management system(e.g.GitHub)",
			EnvVars: []string{drone_info.EnvDroneCommitLink},
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
		&cli.StringFlag{
			Name:    "repo.full.name",
			Usage:   "providers the full name of the repository",
			EnvVars: []string{drone_info.EnvDroneRepo},
		},
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
			Name:    "repo.remote.url",
			Usage:   "providers the remote url of the repository",
			EnvVars: []string{drone_info.EnvDroneRemoteUrl},
		},
		&cli.StringFlag{
			Name:    "repo.owner",
			Usage:   "providers the owner of the repository",
			EnvVars: []string{drone_info.EnvDroneRepoOwner},
		},
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
			Name:    "build.tag",
			Usage:   "build tag",
			EnvVars: []string{drone_info.EnvDroneTag},
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
		// droneInfo end
	}

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

func bindDroneInfo(c *cli.Context) *drone_info.Drone {
	var drone = drone_info.Drone{
		//  repo info
		Repo: drone_info.Repo{
			ShortName: c.String("repo.name"),
			GroupName: c.String("repo.group"),
			OwnerName: c.String("repo.owner"),
			RemoteURL: c.String("repo.remote.url"),
			FullName:  c.String("repo.full.name"),
		},
		//  build info
		Build: drone_info.Build{
			Status:     c.String("build.status"),
			Number:     c.Uint64("build.number"),
			Tag:        c.String("build.tag"),
			Link:       c.String("build.link"),
			Event:      c.String("build.event"),
			StartAt:    c.Uint64("build.started"),
			FinishedAt: c.Uint64("build.finished"),
			PR:         c.String("pull.request"),
			DeployTo:   c.String("deploy.to"),
		},
		Commit: drone_info.Commit{
			Sha:     c.String("commit.sha"),
			Branch:  c.String("commit.branch"),
			Message: c.String("commit.message"),
			Link:    c.String("commit.link"),
			Author: drone_info.CommitAuthor{
				Avatar:   c.String("commit.author.avatar"),
				Email:    c.String("commit.author.email"),
				Name:     c.String("commit.author.name"),
				Username: c.String("commit.author.username"),
			},
		},
		Stage: drone_info.Stage{
			StartedAt:  c.Uint64("stage.started"),
			FinishedAt: c.Uint64("stage.finished"),
		},
	}
	return &drone
}
