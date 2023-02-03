package file_browser_plugin

import (
	"fmt"
	"github.com/sinlov/drone-info-tools/drone_info"
	"github.com/sinlov/drone-info-tools/template"
	tools "github.com/sinlov/drone-info-tools/tools/str_tools"
	"github.com/sinlov/filebrowser-client/file_browser_client"
	"github.com/sinlov/filebrowser-client/tools/folder"
	"github.com/sinlov/filebrowser-client/web_api"
	"log"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type (
	// FileBrowserPlugin file_browser_plugin all config
	FileBrowserPlugin struct {
		Name              string
		Version           string
		Drone             drone_info.Drone
		Config            Config
		fileBrowserClient file_browser_client.FileBrowserClient
	}
)

func (p *FileBrowserPlugin) CleanResultEnv() error {
	for _, envItem := range cleanResultEnvList {
		err := os.Unsetenv(envItem)
		if err != nil {
			return fmt.Errorf("at FileBrowserPlugin.CleanResultEnv [ %s ], err: %v", envItem, err)
		}
	}
	return nil
}

func (p *FileBrowserPlugin) Exec() error {
	if p.Config.Debug {
		for _, e := range os.Environ() {
			log.Println(e)
		}
	}

	var err error

	if p.Config.FileBrowserBaseConfig.FileBrowserHost == "" {
		msg := "missing file_browser_host, please check"
		return fmt.Errorf(msg)
	}

	if p.Config.FileBrowserBaseConfig.FileBrowserUsername == "" {
		msg := "missing file_browser_username, please check"
		return fmt.Errorf(msg)
	}

	if !(tools.StrInArr(p.Config.FileBrowserWorkMode, pluginWorkModeSupport)) {
		return fmt.Errorf("file_browser_plugin file_browser_work_mode type only support: %v", pluginWorkModeSupport)
	}

	// check default TimeoutSecond
	if p.Config.TimeoutSecond < 10 {
		p.Config.TimeoutSecond = 10
	}
	// check default FileBrowserTimeoutPushSecond
	if p.Config.FileBrowserBaseConfig.FileBrowserTimeoutPushSecond < 60 {
		p.Config.FileBrowserBaseConfig.FileBrowserTimeoutPushSecond = 60
	}
	// check default p.Config.FileBrowserBaseConfig.FileBrowserWorkSpace
	if p.Config.FileBrowserBaseConfig.FileBrowserWorkSpace == "" {
		p.Config.FileBrowserBaseConfig.FileBrowserWorkSpace = p.Drone.Build.WorkSpace
	}

	fileBrowserClient, err := file_browser_client.NewClient(
		p.Config.FileBrowserBaseConfig.FileBrowserUsername,
		p.Config.FileBrowserBaseConfig.FileBrowserUserPassword,
		p.Config.FileBrowserBaseConfig.FileBrowserHost,
		p.Config.TimeoutSecond,
		p.Config.FileBrowserBaseConfig.FileBrowserTimeoutPushSecond,
	)
	if err != nil {
		return err
	}

	fileBrowserClient.Debug(p.Config.Debug)

	p.fileBrowserClient = fileBrowserClient

	switch p.Config.FileBrowserWorkMode {
	default:
		return fmt.Errorf("file_browser_plugin file_browser_work_mode not support: %v", p.Config.FileBrowserWorkMode)
	case WorkModeSend:
		err = workOnSend(p)
	}

	log.Printf("=> plugin %s version %s", p.Name, p.Version)
	return err
}

func workOnSend(p *FileBrowserPlugin) error {
	sendModeConfig := p.Config.FileBrowserSendModeConfig
	if !(tools.StrInArr(sendModeConfig.FileBrowserDistType, pluginDistTypeSupport)) {
		return fmt.Errorf("file_browser_plugin file_browser_dist_type dist type only support: %v", pluginDistTypeSupport)
	}

	if sendModeConfig.FileBrowserRemoteRootPath == "" {
		return fmt.Errorf("file_browser_plugin file_browser_remote_root_path not be empty")
	}

	//if sendModeConfig.FileBrowserTargetDistRootPath == "" {
	//	return fmt.Errorf("file_browser_plugin file_browser_target_dist_root_path not be empty")
	//}

	var remoteRealRootPath = strings.TrimRight(sendModeConfig.FileBrowserRemoteRootPath, "/")

	switch sendModeConfig.FileBrowserDistType {
	default:
		return fmt.Errorf("send dist type not support %s", sendModeConfig.FileBrowserDistType)
	case DistTypeGit:
		if p.Drone.Build.Tag == "" {
			remoteRealRootPath = fmt.Sprintf("%s/%s/%s/%s/%s/%d/%s/%s",
				remoteRealRootPath,
				p.Drone.Repo.HostName, p.Drone.Repo.GroupName, p.Drone.Repo.ShortName,
				"b",
				p.Drone.Build.Number,
				p.Drone.Commit.Branch,
				string([]rune(p.Drone.Commit.Sha))[:8],
			)
		} else {
			remoteRealRootPath = fmt.Sprintf("%s/%s/%s/%s/%s/%s/%d/%s",
				remoteRealRootPath,
				p.Drone.Repo.HostName, p.Drone.Repo.GroupName, p.Drone.Repo.ShortName,
				"tag",
				p.Drone.Build.Tag,
				p.Drone.Build.Number,
				string([]rune(p.Drone.Commit.Sha))[:8],
			)
		}
	case DistTypeCustom:
		renderPath, err := template.RenderTrim(sendModeConfig.FileBrowserDistGraph, p.Drone)
		if err != nil {
			return fmt.Errorf("setting file_browser_dist_graph as %s \nerr: %v", sendModeConfig.FileBrowserDistGraph, err)
		}
		remoteRealRootPath = fmt.Sprintf("%s/%s",
			remoteRealRootPath,
			renderPath,
		)
	}
	targetRootPath := filepath.Join(p.Config.FileBrowserBaseConfig.FileBrowserWorkSpace, sendModeConfig.FileBrowserTargetDistRootPath)
	if p.Config.Debug {
		log.Printf("debug: workOnSend fileBrowserDistType: %s", sendModeConfig.FileBrowserDistType)
		if sendModeConfig.FileBrowserDistType == DistTypeCustom {
			log.Printf("debug: workOnSend fileBrowserDistGraph: %s", sendModeConfig.FileBrowserDistGraph)
		}
		log.Printf("debug: workOnSend remoteRealRootPath: %s", remoteRealRootPath)
		log.Printf("debug: workOnSend targetRootPath: %s", targetRootPath)
		log.Printf("debug: workOnSend fileBrowserWorkSpace: %s", p.Config.FileBrowserBaseConfig.FileBrowserWorkSpace)
		log.Printf("debug: workOnSend targetDistRootPath: %s", sendModeConfig.FileBrowserTargetDistRootPath)
		log.Printf("debug: workOnSend targetFileRegular: %s", sendModeConfig.FileBrowserTargetFileRegular)
	}

	if !(folder.PathExistsFast(targetRootPath)) {
		return fmt.Errorf("file browser want send file local path not exists at: %s", targetRootPath)
	}

	var fileSendPathList []string
	if folder.PathIsFile(targetRootPath) {
		fileSendPathList = append(fileSendPathList, targetRootPath)
	} else {
		if sendModeConfig.FileBrowserTargetFileGlob != nil && len(sendModeConfig.FileBrowserTargetFileGlob) > 0 {
			for _, glob := range sendModeConfig.FileBrowserTargetFileGlob {
				walkByGlob, errWalkAllByGlob := folder.WalkAllByGlob(targetRootPath, glob, true)
				if errWalkAllByGlob != nil {
					return fmt.Errorf("file browser want send file local path with glob %s be err: %v", targetRootPath, errWalkAllByGlob)
				}
				fileSendPathList = append(fileSendPathList, walkByGlob...)
			}
		}
		if sendModeConfig.FileBrowserTargetFileRegular != "" {
			matchPath, err := folder.WalkAllByMatchPath(targetRootPath, sendModeConfig.FileBrowserTargetFileRegular, true)
			if err != nil {
				return fmt.Errorf("file browser want send file local path with file regular %s be err: %v", targetRootPath, err)
			}
			fileSendPathList = append(fileSendPathList, matchPath...)
		}
	}

	if len(fileSendPathList) == 0 {
		return fmt.Errorf("file browser want send file local path not find any file at: %s", targetRootPath)
	}

	if p.Config.Debug {
		log.Printf("debug: now send path len %d", len(fileSendPathList))
	}

	fileSendPathList = tools.StrArrRemoveDuplicates(fileSendPathList)

	if p.Config.Debug {
		log.Printf("debug: send path remove duplicates len %d", len(fileSendPathList))
	}

	err := p.fileBrowserClient.Login()
	if err != nil {
		return err
	}

	if len(fileSendPathList) == 1 {
		localFileAbsPath := fileSendPathList[0]
		remotePath := fetchRemotePathByLocalRoot(localFileAbsPath, targetRootPath, remoteRealRootPath)
		var resourcePostOne = file_browser_client.ResourcePostFile{
			LocalFilePath:  localFileAbsPath,
			RemoteFilePath: remotePath,
		}
		errSendOneFile := p.fileBrowserClient.ResourcesPostFile(resourcePostOne, p.Config.Debug)
		if err != nil {
			return errSendOneFile
		}
		if sendModeConfig.FileBrowserShareLinkEnable {
			errSendFileShare := shareBySendConfig(*p, remotePath, false)
			if errSendFileShare != nil {
				return errSendFileShare
			}
		}

	} else {
		for _, item := range fileSendPathList {
			var resourcePost = file_browser_client.ResourcePostFile{
				LocalFilePath:  item,
				RemoteFilePath: fetchRemotePathByLocalRoot(item, targetRootPath, remoteRealRootPath),
			}
			errSendOneFile := p.fileBrowserClient.ResourcesPostFile(resourcePost, p.Config.Debug)
			if err != nil {
				return errSendOneFile
			}
		}
		if sendModeConfig.FileBrowserShareLinkEnable {
			errSendFileShare := shareBySendConfig(*p, remoteRealRootPath, true)
			if errSendFileShare != nil {
				return errSendFileShare
			}
		}
	}

	return nil
}

func shareBySendConfig(p FileBrowserPlugin, remotePath string, isDir bool) error {
	expires := strconv.Itoa(int(p.Config.FileBrowserSendModeConfig.FileBrowserShareLinkExpires))
	passWord := p.Config.FileBrowserSendModeConfig.FileBrowserShareLinkPassword
	if p.Config.FileBrowserSendModeConfig.FileBrowserShareLinkAutoPasswordEnable {
		passWord = genPwd(8)
	}
	if isDir {
		remotePath = fmt.Sprintf("%s/", remotePath)
	}
	shareResource := file_browser_client.ShareResource{
		RemotePath: remotePath,
		ShareConfig: web_api.ShareConfig{
			Password: passWord,
			Expires:  expires,
			Unit:     p.Config.FileBrowserSendModeConfig.FileBrowserShareLinkUnit,
		},
	}
	sharePost, errSendShareFile := p.fileBrowserClient.SharePost(shareResource)
	if errSendShareFile != nil {
		return errSendShareFile
	}
	log.Printf("=> share page: %s", sharePost.DownloadPage)
	if passWord != "" {
		log.Printf("=> share pwd: %s", sharePost.DownloadPasswd)
	}
	log.Printf("=> share user name: %s", p.Config.FileBrowserBaseConfig.FileBrowserUsername)
	log.Printf("=> share remote path: %s", sharePost.RemotePath)
	setEnvFromStr(EnvPluginFileBrowserResultShareHost, p.Config.FileBrowserBaseConfig.FileBrowserHost)
	setEnvFromStr(EnvPluginFileBrowserResultSharePage, sharePost.DownloadPage)
	setEnvFromStr(EnvPluginFileBrowserResultSharePasswd, sharePost.DownloadPasswd)
	setEnvFromStr(EnvPluginFileBrowserResultShareDownloadUrl, sharePost.DownloadUrl)
	setEnvFromStr(EnvPluginFileBrowserResultShareUser, p.Config.FileBrowserBaseConfig.FileBrowserUsername)
	setEnvFromStr(EnvPluginFileBrowserResultShareRemotePath, remotePath)
	return nil
}

func fetchRemotePathByLocalRoot(localAbsPath, localRootPath, remoteRootPath string) string {
	remotePath := strings.Replace(localAbsPath, localRootPath, "", -1)
	remotePath = strings.TrimPrefix(remotePath, "/")
	return path.Join(remoteRootPath, remotePath)
}

func genPwd(cnt uint) string {
	if cnt == 0 {
		return ""
	}

	return randomStrBySed(cnt, "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz_-!")
}

// randomStr
// new random string by cnt
func randomStr(cnt uint) string {
	var letters = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	result := make([]byte, cnt)
	keyL := len(letters)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(keyL)]
	}
	return string(result)
}

// randomStr
// new random string by cnt
func randomStrBySed(cnt uint, sed string) string {
	var letters = []byte(sed)
	result := make([]byte, cnt)
	keyL := len(letters)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(keyL)]
	}
	return string(result)
}

func setEnvFromStr(key string, val string) {
	err := os.Setenv(key, val)
	if err != nil {
		log.Fatalf("set env key [%v] string err: %v", key, err)
	}
}
