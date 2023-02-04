# Changelog

All notable changes to this project will be documented in this file. See [standard-version](https://github.com/conventional-changelog/standard-version) for commit guidelines.

## [1.5.0](https://github.com/sinlov/drone-file-browser-plugin/compare/v1.4.0...v1.5.0) (2023-02-04)


### Features

* change cli flag to package file_browser_plugin ([f3f974a](https://github.com/sinlov/drone-file-browser-plugin/commit/f3f974ada8d9a84ac0fac3a8f50c7b50d3f867f7))

## [1.4.0](https://github.com/sinlov/drone-file-browser-plugin/compare/v1.3.0...v1.4.0) (2023-02-03)


### Features

* embed package.json to config this cli ([0164279](https://github.com/sinlov/drone-file-browser-plugin/commit/016427917a3ba15c68cc995627e83fb9cdecd74a))

## [1.3.0](https://github.com/sinlov/drone-file-browser-plugin/compare/v1.2.1...v1.3.0) (2023-02-03)


### Features

* update to github.com/sinlov/drone-info-tools v1.3.0 ([29a2f5d](https://github.com/sinlov/drone-file-browser-plugin/commit/29a2f5d9b1de95e0e5bfd285df7fefd7b166df5f))

### [1.2.1](https://github.com/sinlov/drone-file-browser-plugin/compare/v1.2.0...v1.2.1) (2023-02-03)


### Bug Fixes

* fix github.com/sinlov/drone-info-tools/template init error ([708ef95](https://github.com/sinlov/drone-file-browser-plugin/commit/708ef955be44d4110d467ca6dee2143ef8105df9))

## [1.2.0](https://github.com/sinlov/drone-file-browser-plugin/compare/v1.1.0...v1.2.0) (2023-02-03)

## [1.1.0](https://github.com/sinlov/drone-file-browser-plugin/compare/v1.0.0...v1.1.0) (2023-02-03)


### Features

* add file_browser_target_file_globs for support glob list to send ([031fe05](https://github.com/sinlov/drone-file-browser-plugin/commit/031fe05a09be1181f9f56f4339e29f27003c22fb))

## 1.0.0 (2023-02-02)


### Features

* add base print at send finish ([5d2e976](https://github.com/sinlov/drone-file-browser-plugin/commit/5d2e9766255749ac5661fc62defa19af2fe85a23))
* add CleanResultEnv() and ResultEnv at plugin run ok ([e4c6ac7](https://github.com/sinlov/drone-file-browser-plugin/commit/e4c6ac73bc986ba6bf00170601c1ad210c599dde))
* add depend github.com/sinlov/drone-info-tools v1.0.1 ([d30bf4f](https://github.com/sinlov/drone-file-browser-plugin/commit/d30bf4fd1bc767783e593add361410afb9db3666))
* add file_browser_work_space for suport different work space to use ([b856c56](https://github.com/sinlov/drone-file-browser-plugin/commit/b856c56d4a1742a943db0377440b095f0e196590))
* add more debug info at send mode ([554bfb5](https://github.com/sinlov/drone-file-browser-plugin/commit/554bfb5fc0a59d573c630edd472714dae446bfad))
* add more drone env and change plugin.Config more clear ([c3c643a](https://github.com/sinlov/drone-file-browser-plugin/commit/c3c643acd283ea955dec90b65f4ea29e9c2e60e7))
* add plugin show name of cli ([c21dbf1](https://github.com/sinlov/drone-file-browser-plugin/commit/c21dbf1f5d71085e6a0fe253f343fdf457818154))
* add plugin.Config bind by cli ([4b60cb8](https://github.com/sinlov/drone-file-browser-plugin/commit/4b60cb8ba42ca439ae97bd2c4070601b87ecc745))
* back to env way to let plugin load ([d5a5fca](https://github.com/sinlov/drone-file-browser-plugin/commit/d5a5fca88fbef8d8edcff09e95dcaf76f6d4f864))
* change file_browser_target_dist_root_path default value ([2818187](https://github.com/sinlov/drone-file-browser-plugin/commit/2818187601a1ad674d3a7fc33ddcb251cf22d401))
* file_browser_target_dist_root_path can set "" ([187b1b7](https://github.com/sinlov/drone-file-browser-plugin/commit/187b1b775eff0339f54490fd9acd430931833de4))
* let share can used ([3778da0](https://github.com/sinlov/drone-file-browser-plugin/commit/3778da0f0d497df3c58825053eb180cc0bf39fed))
* mark final verison v1.0.0 ([ec04e0c](https://github.com/sinlov/drone-file-browser-plugin/commit/ec04e0cb0a6f7ffc85bf304ff36378a587f5e85c))
* mark version v1.0.0 ([9ce6179](https://github.com/sinlov/drone-file-browser-plugin/commit/9ce6179663b3f1edf7f7e9f7fcd1b3089eb84dba))
* rename plugin package for other plugin develop ([e220943](https://github.com/sinlov/drone-file-browser-plugin/commit/e220943f2b36daab9bd09d6c9a73a55d49d51958))
* support file_browser_dist_graph and add ci to build docker image ([b661e5a](https://github.com/sinlov/drone-file-browser-plugin/commit/b661e5a15aec013f4a3758e75b9b778af31fad3e))
* try file_browser.host parse by yaml ([4f4fe7c](https://github.com/sinlov/drone-file-browser-plugin/commit/4f4fe7c253e23763fedc8f3d27f907742578acec))
* update config setting by yaml ([0190973](https://github.com/sinlov/drone-file-browser-plugin/commit/01909730b33ab896013886d095690d40aed92f7e))
* update github.com/sinlov/filebrowser-client v0.1.3 and debug to show cli version ([9c18cc3](https://github.com/sinlov/drone-file-browser-plugin/commit/9c18cc31a14eb5b3fdc83909f8c692a2662b8e8a))


### Bug Fixes

* fix cli load file_browser_remote_root_path will be cover by env:PLUGIN_FILE_BROWSER_DIST_TYPE ([6e993b4](https://github.com/sinlov/drone-file-browser-plugin/commit/6e993b44d03d99ad26713e14e9507a774fe0b387))
