/*
The MIT License

Copyright © 2025 John, Sing Dao, Siu <john.sd.siu@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

// Provide git command using exec.Cmd wrapper go-helper/v2/cmd
package gitcmd

import (
	"github.com/J-Siu/go-helper/v2/cmd"
)

var (
	git = new(GitCmd)
)

// New
//   - If <workPath> is empty/nil, current directory is used.
//   - Create and Return a new GitCmd pointer
func New(workPath string) *GitCmd {
	return new(GitCmd).New(workPath)
}

// Run "git clone <optionsP>".
//   - If <workPath> is empty/nil, current directory is used.
//   - Return a cmd.Cmd pointer for execution information.
func Clone(workPath string, options []string) *cmd.Cmd {
	return git.New(workPath).Clone(options).Run()
}

// Check git executable exist.
func ExecExist() bool { return git.ExecPath() != "" }

// Get git executable path.
//   - Return empty string if not found.
func ExecPath() string { return git.ExecPath() }

// Run "git init".
//   - If <workPath> is empty/nil, current directory is used.
//   - Return a cmd.Cmd pointer for execution information.
func Init(workPath string) *cmd.Cmd { return git.New(workPath).Init().Run() }

// Run "git branch --show-current".
//   - If <workPath> is empty/nil, current directory is used.
//   - Return a cmd.Cmd pointer for execution information.
func BranchCurrent(workPath string) *cmd.Cmd { return git.New(workPath).BranchCurrent().Run() }

// Run "git pull <optionsP>".
//   - If <workPath> is empty/nil, current directory is used.
//   - Return a cmd.Cmd pointer for execution information.
func Pull(workPath string, options []string) *cmd.Cmd {
	return git.New(workPath).Pull(options).Run()
}

// Run "git push <optionsP>".
//   - If <workPath> is empty/nil, current directory is used.
//   - Return a cmd.Cmd pointer for execution information.
func Push(workPath string, options []string) *cmd.Cmd {
	return git.New(workPath).Push(options).Run()
}

// Run "git remote".
//   - If <workPath> is empty/nil, current directory is used.
//   - Return remotes in string array.
func Remote(workPath string, v bool) *[]string { return git.New(workPath).Remote(v) }

// Run "git remote add <name> <gitUrl>".
//   - If <workPath> is empty/nil, current directory is used.
//   - Return a cmd.Cmd pointer for execution information.
func RemoteAdd(workPath string, name string, gitUrl string) *cmd.Cmd {
	return git.New(workPath).RemoteAdd(name, gitUrl).Run()
}

// Check if a git remote(by name) exist in workPath.
//   - If <workPath> is empty/nil, current directory is used.
func RemoteExist(workPath string, name string) bool { return git.New(workPath).RemoteExist(name) }

// Run "git remote remove".
//   - If <workPath> is empty/nil, current directory is used.
//   - If remote exist Return a cmd.Cmd pointer for execution information.
//   - If remote does not exit, return nil.(Nothing to remove)
func RemoteRemove(workPath string, name string) *cmd.Cmd {
	return git.New(workPath).RemoteRemove(name).Run()
}

// Run "git remote remove" all git remotes
//   - If <workPath> is empty/nil, current directory is used.
func RemoteRemoveAll(workPath string) { git.New(workPath).RemoteRemoveAll() }

// Get git root from current directory.
//   - If <workPath> is empty/nil, current directory is used.
//   - Return empty string if not a git dir.
func Root(workPath string) string { return git.New(workPath).Root() }

// Get git submodule root from `workPath`.
//   - If <workPath> is empty/nil, current directory is used.
//   - Return empty string if not a submodule dir.
func RootSubmodule(workPath string) string { return git.New(workPath).RootSubmodule() }

func TagList(workPath string) *[]string { return git.New(workPath).TagList() }
