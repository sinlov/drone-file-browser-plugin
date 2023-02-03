[![go-ubuntu](https://github.com/sinlov/drone-file-browser-plugin/workflows/go-ubuntu/badge.svg?branch=main)](https://github.com/sinlov/drone-file-browser-plugin/actions)
[![GoDoc](https://godoc.org/github.com/sinlov/drone-file-browser-plugin?status.png)](https://godoc.org/github.com/sinlov/drone-file-browser-plugin/)
[![GoReportCard](https://goreportcard.com/badge/github.com/sinlov/drone-file-browser-plugin)](https://goreportcard.com/report/github.com/sinlov/drone-file-browser-plugin)
[![codecov](https://codecov.io/gh/sinlov/drone-file-browser-plugin/branch/main/graph/badge.svg)](https://codecov.io/gh/sinlov/drone-file-browser-plugin)
[![docker version semver](https://img.shields.io/docker/v/sinlov/drone-file-browser-plugin?sort=semver)](https://hub.docker.com/r/sinlov/drone-file-browser-plugin/tags?page=1&ordering=last_updated)
[![docker image size](https://img.shields.io/docker/image-size/sinlov/drone-file-browser-plugin)](https://hub.docker.com/r/sinlov/drone-file-browser-plugin)
[![docker pulls](https://img.shields.io/docker/pulls/sinlov/drone-file-browser-plugin)](https://hub.docker.com/r/sinlov/drone-file-browser-plugin/tags?page=1&ordering=last_updated)

## for what

- this project used to drone CI to support [file browser](https://github.com/filebrowser/filebrowser)
- get file browser `host` like `https://filebrowser.xxx.com`
- this plugin need file browser `username`
- and need file browser user need `password`

## Pipeline Settings (.drone.yml)

- sample

```yaml
steps:
  - name: drone-file-browser-send
    image: sinlov/drone-file-browser-plugin:latest
    # pull: if-not-exists
    settings:
      # debug: false # plugin debug switch
      file_browser_host: "http://127.0.0.1:80" # must set args, file_browser host like http://127.0.0.1:80
      file_browser_username: # must set args, file_browser username
        # https://docs.drone.io/pipeline/environment/syntax/#from-secrets
        from_secret: file_browser_user_name
      file_browser_user_password: # must set args, file_browser user password
        from_secret: file_browser_user_password
      file_browser_remote_root_path: dist/ # must set args, send to file_browser base path
      file_browser_target_file_globs: # must set args, globs list of send to file_browser under file_browser_target_dist_root_path
        - "**/*.tar.gz"
        - "**/*.sha256"
      file_browser_share_link_expires: 0 # if set 0, will allow share_link exist forever，default: 0
      file_browser_share_link_unit: days # take effect by open share_link, only can use as [ days hours minutes seconds ]
      file_browser_share_link_auto_password_enable: true # password of share_link auto , if open this will cover settings.file_browser_share_link_password. default: false
```

- full config

```yaml
steps:
  - name: drone-file-browser-send
    image: sinlov/drone-file-browser-plugin:latest
    # pull: if-not-exists
    settings:
      debug: false # plugin debug switch
      timeout_second: 10 # api timeout default: 10
      file_browser_timeout_push_second: 60 # push each file timeout push second, must gather than 60.default: 60
      file_browser_host: # must set args, file_browser host like http://127.0.0.1:80
          from_secret: file_browser_host
      file_browser_username: # must set args, file_browser username
          # https://docs.drone.io/pipeline/environment/syntax/#from-secrets
          from_secret: file_browser_user_name
      file_browser_user_password: # must set args, file_browser user password
          from_secret: file_browser_user_password
      file_browser_work_space: "" # file_browser work space. default "" will use env:DRONE_WORKSPACE
      file_browser_work_mode: send # 1.0 only support send
      file_browser_remote_root_path: dist/ # must set args, send to file_browser base path
      file_browser_dist_type: custom # must set args, type of dist file graph only can use: git, custom
      file_browser_dist_graph: "{{ Repo.HostName }}/{{ Repo.GroupName }}/{{ Repo.ShortName }}/s/{{ Build.Number }}/{{ Stage.Name }}-{{ Build.Number }}-{{ Stage.FinishedTime }}" # type of dist custom
      file_browser_target_dist_root_path: dist/ # path of file_browser work on root, can set "". default: ""
      file_browser_target_file_globs: # must set args, globs list of send to file_browser under file_browser_target_dist_root_path
        - "**/*.tar.gz"
        - "**/*.sha256"
      file_browser_target_file_regular: .*.tar.gz # must set args, regular of send to file_browser under file_browser_target_dist_root_path
      file_browser_share_link_enable: true # share dist dir as link, default: true
      file_browser_share_link_expires: 0 # if set 0, will allow share_link exist forever，default: 0
      file_browser_share_link_unit: days # take effect by open share_link, only can use as [ days hours minutes seconds ]
      file_browser_share_link_password: "" # password of share_link, if not set will not use password, default: ""
      file_browser_share_link_auto_password_enable: false # password of share_link auto , if open this will cover settings.file_browser_share_link_password. default: false
    when:
      event: # https://docs.drone.io/pipeline/exec/syntax/conditions/#by-event
        - promote
        - rollback
        - push
        - pull_request
        - tag
      status: # only support failure/success,  both open will send anything
        - failure
        # - success
```

- `1.x` drone-exec only support env

- download by [https://github.com/sinlov/drone-file-browser-plugin/releases](https://github.com/sinlov/drone-file-browser-plugin/releases) to get platform binary, then has local path
- binary path like `C:\Drone\drone-runner-exec\plugins\drone-file-browser-plugin.exe` can be drone run env like `EXEC_DRONE_FILE_BROWSER_PLUGIN_FULL_PATH`
- env:EXEC_DRONE_FILE_BROWSER_PLUGIN_FULL_PATH can set at file which define as [DRONE_RUNNER_ENVFILE](https://docs.drone.io/runner/exec/configuration/reference/drone-runner-envfile/) to support each platform

```yaml
steps:
  - name: drone-file-browser-send-exec # must has env EXEC_DRONE_FILE_BROWSER_PLUGIN_FULL_PATH and exec tools
    environment:
      PLUGIN_DEBUG: false
      # PLUGIN_NTP_TARGET: "pool.ntp.org" # if not set will not sync
      PLUGIN_TIMEOUT_SECOND: 10 # default 10
      PLUGIN_FILE_BROWSER_TIMEOUT_PUSH_SECOND: 60 # push each file timeout push second, must gather than 60.default: 60
      PLUGIN_FILE_BROWSER_HOST: "http://127.0.0.1:80" # must set args, file_browser host like http://127.0.0.1:80
      PLUGIN_FILE_BROWSER_USERNAME: # must set args, file_browser username
        # https://docs.drone.io/pipeline/environment/syntax/#from-secrets
        from_secret: file_browser_user_name
      PLUGIN_FILE_BROWSER_USER_PASSWORD: # must set args, file_browser user password
        from_secret: file_browser_user_password
      PLUGIN_FILE_BROWSER_REMOTE_ROOT_PATH: dist/ # must set args, send to file_browser base path
      PLUGIN_FILE_BROWSER_TARGET_DIST_ROOT_PATH: dist/ # path of file_browser work on root, can set "". default: ""
      PLUGIN_FILE_BROWSER_TARGET_FILE_GLOBS: # must set args, globs list of send to file_browser under file_browser_target_dist_root_path
        - "**/*.tar.gz"
        - "**/*.sha256"
      PLUGIN_FILE_BROWSER_TARGET_FILE_REGULAR: .*.tar.gz # must set args, regular of send to file_browser under file_browser_target_dist_root_path
      PLUGIN_FILE_BROWSER_SHARE_LINK_EXPIRES: 0 # if set 0, will allow share_link exist forever，default: 0
      PLUGIN_FILE_BROWSER_SHARE_LINK_UNIT: days # take effect by open share_link, only can use as [ days hours minutes seconds ]
      PLUGIN_FILE_BROWSER_SHARE_LINK_AUTO_PASSWORD_ENABLE: true # password of share_link auto , if open this will cover settings.file_browser_share_link_password. default: false
    commands:
      - ${EXEC_DRONE_FILE_BROWSER_PLUGIN_FULL_PATH} `
        ""
    when:
      event: # https://docs.drone.io/pipeline/exec/syntax/conditions/#by-event
        - promote
        - rollback
        - push
        - pull_request
        - tag
      status: # only support failure/success,  both open will send anything
        - failure
        - success
```

### settings.debug

- if open `settings.debug` will try send file browser use `override` for debug.
- please close `settings.debug` in production models

### file_browser_dist_type

- if use file_browser_dist_type = `git`, send to filebrowser file tree like

```
# not use tag
${file_browser_remote_root_path}/
	{{Repo.HostName}}/
		{{Repo.GroupName}}/
			{{Repo.ShortName}}/
				b/
					{{Build.Number}}/
						{{Commit.Branch}}/
							{{Commit.Sha[0:8]}}

# use tag
${file_browser_remote_root_path}/
	{{Repo.HostName}}/
		{{Repo.GroupName}}/
			{{Repo.ShortName}}/
				tag/
					{{Build.Tag}}/
						{{Build.Number}}/
							{{Commit.Sha[0:8]}}
```

- you can use file_browser_dist_type = `custom`, like

```
{{ Repo.HostName }}/{{ Repo.GroupName }}/{{ Repo.ShortName }}/s/{{ Build.Number }}/{{ Stage.Name }}-{{ Build.Number }}-{{ Stage.FinishedTime }}
```

template use struct [Drone](https://github.com/sinlov/drone-info-tools/blob/main/drone_info/droneInfo.go#L358)

# Features

- more see [features/README.md](features/README.md)

# dev

## depends

in go mod project

```bash
# warning use privte git host must set
# global set for once
# add private git host like github.com to evn GOPRIVATE
$ go env -w GOPRIVATE='github.com'
# use ssh proxy
# set ssh-key to use ssh as http
$ git config --global url."git@github.com:".insteadOf "http://github.com/"
# or use PRIVATE-TOKEN
# set PRIVATE-TOKEN as gitlab or gitea
$ git config --global http.extraheader "PRIVATE-TOKEN: {PRIVATE-TOKEN}"
# set this rep to download ssh as https use PRIVATE-TOKEN
$ git config --global url."ssh://github.com/".insteadOf "https://github.com/"

# before above global settings
# test version info
$ git ls-remote -q https://github.com/sinlov/drone-file-browser-plugin.git

# test depends see full version
$ go list -mod=readonly -v -m -versions github.com/sinlov/drone-file-browser-plugin
# or use last version add go.mod by script
$ echo "go mod edit -require=$(go list -mod=readonly -m -versions github.com/sinlov/drone-file-browser-plugin | awk '{print $1 "@" $NF}')"
$ echo "go mod vendor"
```

```

```bash
make test
```

- see help

```bash
make dev
```

update main.go file set env then and run

```bash
export PLUGIN_MSG_TYPE= \
  export PLUGIN_WEBHOOK= \
  export DRONE_REPO=sinlov/drone-file-browser-plugin \
  export DRONE_REPO_NAME=drone-file-browser-plugin \
  export DRONE_REPO_NAMESPACE=sinlov \
  export DRONE_REMOTE_URL=https://github.com/sinlov/drone-file-browser-plugin \
  export DRONE_REPO_OWNER=sinlov \
  export DRONE_COMMIT_AUTHOR=sinlov \
  export DRONE_COMMIT_AUTHOR_AVATAR=  \
  export DRONE_COMMIT_AUTHOR_EMAIL=sinlovgmppt@gmail.com \
  export DRONE_COMMIT_BRANCH=main \
  export DRONE_COMMIT_LINK=https://github.com/sinlov/drone-file-browser-plugin/commit/68e3d62dd69f06077a243a1db1460109377add64 \
  export DRONE_COMMIT_SHA=68e3d62dd69f06077a243a1db1460109377add64 \
  export DRONE_COMMIT_REF=refs/heads/main \
  export DRONE_COMMIT_MESSAGE="mock message commit" \
  export DRONE_STAGE_STARTED=1674531206 \
  export DRONE_STAGE_FINISHED=1674532106 \
  export DRONE_BUILD_STATUS=success \
  export DRONE_BUILD_NUMBER=1 \
  export DRONE_BUILD_LINK=https://drone.xxx.com/sinlov/drone-file-browser-plugin/1 \
  export DRONE_BUILD_EVENT=push \
  export DRONE_BUILD_STARTED=1674531206 \
  export DRONE_BUILD_FINISHED=1674532206
```

- then run

```bash
make run
```

## docker

```bash
# then test build as test/Dockerfile
$ make dockerTestRestartLatest
# if run error
# like this error
# err: missing webhook, please set webhook
#  fix env settings then test

# see run docker fast
$ make dockerTestRunLatest

# clean test build
$ make dockerTestPruneLatest

# see how to use
$ docker run --rm sinlov/drone-file-browser-plugin:latest -h
```
