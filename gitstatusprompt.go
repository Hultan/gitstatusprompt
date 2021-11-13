package gitStatusPrompt

import (
	"fmt"
	"strconv"

	gitStatus "github.com/hultan/gitstatus"
)

const (
	version            = "1.0.0"
)

// GitStatusPrompt : Creates a short string showing status of a git repository
type GitStatusPrompt struct {
}

// GetAdvancedPrompt : Creates a short string showing status of a git repository
func (g *GitStatusPrompt) GetAdvancedPrompt(path, format string) string {
	status := &gitStatus.GitStatus{}
	info, err := status.GetStatus(path)
	if err != nil {
		return err.Error()
	}

	// If it is not a git repository, just leave
	if !info.IsGit() {
		return ""
	}

	advanced := &advancedPrompt{info, format}
	return advanced.getAdvancedPrompt()
}

func (g *GitStatusPrompt) GetVerbosePrompt(path string) string {
	status := &gitStatus.GitStatus{}
	info, err := status.GetStatus(path)
	if err != nil {
		return err.Error()
	}

	// If it is not a git repository, just leave
	if !info.IsGit() {
		return ""
	}

	var result string

	result += fmt.Sprintf("Branch    : %s (%d ahead, %d behind)\n", info.Branch(), info.Ahead(), info.Behind())
	result += fmt.Sprintf("Staged    : %d\n", info.Staged())
	result += fmt.Sprintf("Modified  : %d\n", info.Modified())
	result += fmt.Sprintf("Deleted   : %d\n", info.Deleted())
	result += fmt.Sprintf("Unmerged  : %d\n", info.Unmerged())
	result += fmt.Sprintf("Untracked : %d\n", info.Untracked())

	return result
}

func (g *GitStatusPrompt) GetPrompt(path string) string {
	status := &gitStatus.GitStatus{}
	info, err := status.GetStatus(path)
	if err != nil {
		return err.Error()
	}

	// If it is not a git repository, just leave
	if !info.IsGit() {
		return ""
	}

	// Create and return the normal git prompt
	var result string

	if info.Branch() != "" {
		// Branch
		result = info.Branch()
	} else {
		// Detached head
		result = ":HEAD"
	}
	if info.Ahead() > 0 {
		result += "↑" + strconv.Itoa(info.Ahead())
	}
	if info.Behind() > 0 {
		result += "↓" + strconv.Itoa(info.Behind())
	}
	if info.Untracked()+info.Modified()+info.Deleted()+info.Unmerged()+info.Staged() > 0 {
		result += "|"
	}
	if info.Untracked() > 0 {
		result += "+" + strconv.Itoa(info.Untracked())
	}
	if info.Modified() > 0 {
		result += "~" + strconv.Itoa(info.Modified())
	}
	if info.Deleted() > 0 {
		result += "-" + strconv.Itoa(info.Deleted())
	}
	if info.Unmerged() > 0 {
		result += "x" + strconv.Itoa(info.Unmerged())
	}
	if info.Staged() > 0 {
		result += "•" + strconv.Itoa(info.Staged())
	}

	return result
}
