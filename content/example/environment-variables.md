+++
title = "Environment Variables"
date = 2023-08-07T13:57:13+08:00
type = "docs"
description = ""
isCJKLanguage = true
draft = false

+++

# Environment Variables

> 原文：https://gobyexample.com/environment-variables

```go
// Note:
// This code is from https://gobyexample.com.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO")) // FOO: 1
	fmt.Println("BAR:", os.Getenv("BAR")) // BAR: 

	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}

```



```bash
PS D:\Dev\Go\byExample\environment_variables> go run main.go  
FOO: 1
BAR:

FOO
TERM_SESSION_ID
ZES_ENABLE_SYSMAN
ProgramW6432
CommonProgramW6432
FIG_JETBRAINS_SHELL_INTEGRATION
GOPATH
USERNAME
USERPROFILE
ALLUSERSPROFILE
_INTELLIJ_FORCE_SET_GOROOT
PROCESSOR_REVISION
POWERSHELL_DISTRIBUTION_CHANNEL
IDEA_INITIAL_DIRECTORY
FPS_BROWSER_APP_PROFILE_STRING
PUBLIC
Path
GOROOT
DriverData
HOMEDRIVE
SESSIONNAME
TERMINAL_EMULATOR
LOGONSERVER
SystemRoot
HOMEPATH
_INTELLIJ_FORCE_PREPEND_PATH
NVM_HOME
LOCALAPPDATA
APPDATA
GoLand
PROCESSOR_IDENTIFIER
PATHEXT
PSModulePath
ProgramFiles(x86)
_INTELLIJ_FORCE_SET_GO111MODULE
OS
PROCESSOR_ARCHITECTURE
NVM_SYMLINK
NUMBER_OF_PROCESSORS
ComSpec
PROCESSOR_LEVEL
windir
USERDOMAIN_ROAMINGPROFILE
ProgramFiles
_INTELLIJ_FORCE_SET_GOPATH
TMP
TEMP
CommonProgramFiles(x86)
OneDrive
USERDOMAIN
SystemDrive
GO111MODULE
COMPUTERNAME
ProgramData
EFC_10276
FPS_BROWSER_USER_PROFILE_STRING

CommonProgramFiles

```

