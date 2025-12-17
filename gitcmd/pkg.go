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
func New(workPathP *string) *GitCmd {
	return new(GitCmd).New(workPathP)
}

// Run "git clone <optionsP>".
//   - If <workPath> is empty/nil, current directory is used.
//   - Return a cmd.Cmd pointer for execution information.
func Clone(workPathP *string, optionsP *[]string) *cmd.Cmd {
	return git.New(workPathP).Clone(optionsP)
}

// Check git executable exist.
func ExecExist() bool { return git.ExecPath() != "" }

// Get git executable path.
//   - Return empty string if not found.
func ExecPath() string { return git.ExecPath() }

// Run "git init".
//   - If <workPath> is empty/nil, current directory is used.
//   - Return a cmd.Cmd pointer for execution information.
func Init(workPathP *string) *cmd.Cmd { return git.New(workPathP).Init() }

// Run "git branch --show-current".
//   - If <workPath> is empty/nil, current directory is used.
//   - Return a cmd.Cmd pointer for execution information.
func BranchCurrent(workPathP *string) *cmd.Cmd { return git.New(workPathP).BranchCurrent() }

// Run "git pull <optionsP>".
//   - If <workPath> is empty/nil, current directory is used.
//   - Return a cmd.Cmd pointer for execution information.
func Pull(workPathP *string, optionsP *[]string) *cmd.Cmd { return git.New(workPathP).Pull(optionsP) }

// Run "git push <optionsP>".
//   - If <workPath> is empty/nil, current directory is used.
//   - Return a cmd.Cmd pointer for execution information.
func Push(workPathP *string, optionsP *[]string) *cmd.Cmd {
	return git.New(workPathP).Push(optionsP)
}

// Run "git remote".
//   - If <workPath> is empty/nil, current directory is used.
//   - Return remotes in string array.
func Remote(workPathP *string, v bool) *[]string { return git.New(workPathP).Remote(v) }

// Run "git remote add <name> <gitUrl>".
//   - If <workPath> is empty/nil, current directory is used.
//   - Return a cmd.Cmd pointer for execution information.
func RemoteAdd(workPathP *string, name string, gitUrl string) *cmd.Cmd {
	return git.New(workPathP).RemoteAdd(name, gitUrl)
}

// Check if a git remote(by name) exist in workPath.
//   - If <workPath> is empty/nil, current directory is used.
func RemoteExist(workPathP *string, name string) bool { return git.New(workPathP).RemoteExist(name) }

// Run "git remote remove".
//   - If <workPath> is empty/nil, current directory is used.
//   - If remote exist Return a cmd.Cmd pointer for execution information.
//   - If remote does not exit, return nil.(Nothing to remove)
func RemoteRemove(workPathP *string, name string) *cmd.Cmd {
	return git.New(workPathP).RemoteRemove(name)
}

// Run "git remote remove" all git remotes
//   - If <workPath> is empty/nil, current directory is used.
func RemoteRemoveAll(workPathP *string) { git.New(workPathP).RemoteRemoveAll() }

// Get git root from current directory.
//   - If <workPath> is empty/nil, current directory is used.
//   - Return empty string if not a git dir.
func Root(workPathP *string) string { return git.New(workPathP).Root() }

// Get git submodule root from `workPath`.
//   - If <workPath> is empty/nil, current directory is used.
//   - Return empty string if not a submodule dir.
func RootSubmodule(workPathP *string) string { return git.New(workPathP).RootSubmodule() }

func Tag(workPathP *string) *[]string { return git.New(workPathP).Tag() }
