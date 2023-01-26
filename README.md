[![go-ubuntu](https://github.com/sinlov/drone-file-browser-plugin/workflows/go-ubuntu/badge.svg?branch=main)](https://github.com/sinlov/drone-file-browser-plugin/actions)
[![GoDoc](https://godoc.org/github.com/sinlov/drone-file-browser-plugin?status.png)](https://godoc.org/github.com/sinlov/drone-file-browser-plugin/)
[![GoReportCard](https://goreportcard.com/badge/github.com/sinlov/drone-file-browser-plugin)](https://goreportcard.com/report/github.com/sinlov/drone-file-browser-plugin)
[![codecov](https://codecov.io/gh/sinlov/drone-file-browser-plugin/branch/main/graph/badge.svg)](https://codecov.io/gh/sinlov/drone-file-browser-plugin)

## for what

- this project used to drone CI to support [file browser](https://github.com/filebrowser/filebrowser)

## Pipeline Settings (.drone.yml)

- sample

```yaml
steps:
  - name: drone-plugin-template
    image: sinlov/drone-file-browser-plugin:latest
    pull: if-not-exists
    settings:
      debug: false # plugin debug switch
      file_browser_host: http://127.0.0.1 # host of file_browser # must set args, file_browser host
      file_browser_username: # must set args, file_browser username
        # https://docs.drone.io/pipeline/environment/syntax/#from-secrets
        from_secret: file_browser_user_name
      file_browser_user_password: # must set args, file_browser user password
        from_secret: file_browser_user_password
      file_browser_dist_root: dist/ # must set args, path of file_browser root default: dist/
      file_browser_dist_type: git # must set args, type of dist file graph only can use: git, custom
      file_browser_target_file_regular: dist/*.tar.gz # must set args, regular of send to file_browser
```

- full

```yaml
steps:
  - name: drone-plugin-template
    image: sinlov/drone-file-browser-plugin:latest
    pull: if-not-exists
    settings:
      debug: false # plugin debug switch
      timeout_second: 10 # api timeout default: 10
      file_browser_timeout_push_minute: 1 # file_browser push each file timeout push minute default: 1
      file_browser_host: http://127.0.0.1 # host of file_browser # must set args, file_browser host
      file_browser_username: # must set args, file_browser username
        # https://docs.drone.io/pipeline/environment/syntax/#from-secrets
        from_secret: file_browser_user_name
      file_browser_user_password: # must set args, file_browser user password
        from_secret: file_browser_user_password
      file_browser_dist_root: dist/ # must set args, path of file_browser root default: dist/
      file_browser_dist_type: git # must set args, type of dist file graph only can use: git, custom
      file_browser_dist_graph: "{{ Commit.Branch }}/{{ Commit.Sha }}/{{ Build.FinishedAt }}/{{ Build.Tag }}" # type of dist custom set as this
      file_browser_target_file_regular: dist/*.tar.gz # must set args, regular of send to file_browser
      file_browser_share_link_enable: true # share dist dir as link, default: true
      file_browser_share_link_unit: days # take effect by open share_link, only can use as [ days hours minutes seconds ]
      file_browser_share_link_expires: 0 # if set 0, will allow share_link exist forever，default: 0
      file_browser_share_link_password: "" # password of share_link, if not set will not use password, default: ""
      file_browser_share_link_auto_password_enable: false # password of share_link auto , if open this will cover settings.file_browser_share_link_password default: false
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
$ go list -v -m -versions github.com/sinlov/drone-file-browser-plugin
# or use last version add go.mod by script
$ echo "go mod edit -require=$(go list -m -versions github.com/sinlov/drone-file-browser-plugin | awk '{print $1 "@" $NF}')"
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
