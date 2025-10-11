# go-gitcmd

Provide git command using exec.Cmd wrapper [go-helper](https://github.com/J-Siu/go-helper)/v2/cmd.

<!-- TOC -->
- [Package Functions](#package-functions)
- [Doc](#doc)
- [Repository](#repository)
- [Change Log](#change-log)
- [License](#license)
<!-- /TOC -->
<!--more-->

### Package Functions

```go
func Git(workPathP *string, optionsP *[]string) *cmd.Cmd
func GitClone(workPathP *string, optionsP *[]string) *cmd.Cmd
func GitExecExist() bool
func GitExecPath() string
func GitInit(workPathP *string) *cmd.Cmd
func GitBranchCurrent(workPathP *string) *cmd.Cmd
func GitPull(workPathP *string, optionsP *[]string) *cmd.Cmd
func GitPush(workPathP *string, optionsP *[]string) *cmd.Cmd
func GitRemote(workPathP *string, v bool) *[]string
func GitRemoteAdd(workPathP *string, name string, git string) *cmd.Cmd
func GitRemoteExist(workPathP *string, name string) bool
func GitRemoteRemove(workPathP *string, name string) *cmd.Cmd
func GitRemoteRemoveAll(workPathP *string)
func GitRoot(workPathP *string) string
func GitRootSubmodule(workPathP *string) string
```

### Doc

- https://pkg.go.dev/github.com/J-Siu/go-gitcmd

### Repository

- [go-gitcmd](https://github.com/J-Siu/go-gitcmd)

### Change Log

- v0.0.1
  - Initial commit
- v0.0.2
  - Fix output
- v0.0.3
  - Update go-helper/v2
- v0.0.4
  - Update go-helper/v2
- v0.0.5
  - Update go-helper/v2
- v1.0.0
  - Update go-helper/v2

### License

The MIT License (MIT)

Copyright © 2025 John, Sing Dao, Siu <john.sd.siu@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
