# Geacon



## log


20210402:

* 优化代码，统一配置至config.go


20210401:

* 信任https,简单支持c2profile,添加异常容错处理,解决cmd小黑窗,支持CDN上线


## use


### macos
```
export GOOS="darwin" && export GOARCH="amd64" && rm -rf ./main &&go build -ldflags="-s -w" cmd/main.go &&./main
```
### linux
```
export GOOS="linux" && export GOARCH="amd64" && rm -rf ./main &&go build -ldflags="-s -w" cmd/main.go &&./main
```
### windows

WINDOWS环境需要**去除`/src/geacon/cmd/packet/commands.go`中第13和64行的注释，否则执行命令会有小黑窗**！！！

```
export GOOS="windows" && export GOARCH="amd64"&&go build -ldflags="-s -w -H windowsgui" cmd/main.go
```




**Implement CobaltStrike's Beacon in Go**

----

*This project is only used for learning protocol analysis and reverse engineering. If someone's rights are violated, contact me to delete the project, and the last DO NOT USE IT ILLEGALLY*



## How to play

1. Setup the teamserver and start a http lisenter, the teamserver will generate the file `.cobaltstrike.beacon_keys`.
2. Compile beacontoo with Jetbrains Idea, use command `java -jar BeaconTool.jar ` to convert java keystore to PEM format.
3. Replace the RSA key pair in the file `cmd/config/config.go` (the RSA private key is not required, I wrote it in the code just for the record)
4. Compile geacon for what platform you want to run, use command `export GOOS="darwin" && export GOARCH="amd64" && go build cmd/main.go`
5. Having fun ! PR and issue is welcome ;)
6. Geacon has just been tested on CobaltStrike 3.14 and only support default c2profile, so many hardcode in the project and I will not try to implement more C2profile support at this moment.
7. Thanks for **[@xxxxxyyyy](https://github.com/xxxxxyyyy)**'s PR, and now Geacon support **CobaltStrike 4.x** 

## Screenshot

Grab Geacon's command execution results in Linux
![login](https://github.com/darkr4y/geacon/raw/master/screenshots/sc.png)



## Protocol analysis

To be continued, I will update as soon as I have time ...



## Todo

~~1. Support CobaltStrike 4.0~~

2. Fix the OS icon issue in session table

*_DarkRay@RedCore*

