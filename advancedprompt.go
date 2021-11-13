package gitStatusPrompt

import (
	"strconv"
	"strings"

	gitStatus "github.com/hultan/gitstatus"
)

// advancedPrompt : Handles the advanced git prompt based off of the config file
type advancedPrompt struct {
	gitStatusInfo *gitStatus.GitStatusInfo
	format        string
}

func (a *advancedPrompt) getAdvancedPrompt() string {
	parts := strings.Split(a.format, "$(SEPARATOR)")

	result := a.getPromptPart(parts[0])
	if len(parts) > 1 {
		var part string
		for i := 1; i < len(parts); i++ {
			part = a.getPromptPart(parts[i])
			if len(result) > 0 && len(strings.Trim(part, " ")) > 0 {
				result = strings.Join([]string{result, part}, "|")
			}
		}
	}

	return a.escape(result)
}

func (a *advancedPrompt) escape(text string) string {
	return strings.ReplaceAll(text, "$(ESC)", "\x1b")
}

func (a *advancedPrompt) getPromptPart(part string) string {
	result := part

	branch := a.getBranch(a.gitStatusInfo)
	result = strings.ReplaceAll(result, "$(BRANCH)", branch)
	if a.gitStatusInfo.Ahead() > 0 {
		ahead := a.getAhead(a.gitStatusInfo)
		result = strings.ReplaceAll(result, "$(AHEAD)", ahead)
	} else {
		result = strings.ReplaceAll(result, "$(AHEAD)", "")
	}
	if a.gitStatusInfo.Behind() > 0 {
		behind := a.getBehind(a.gitStatusInfo)
		result = strings.ReplaceAll(result, "$(BEHIND)", behind)
	} else {
		result = strings.ReplaceAll(result, "$(BEHIND)", "")
	}
	if a.gitStatusInfo.Untracked() > 0 {
		untracked := a.getUntracked(a.gitStatusInfo)
		result = strings.ReplaceAll(result, "$(UNTRACKED)", untracked)
	} else {
		result = strings.ReplaceAll(result, "$(UNTRACKED)", "")
	}
	if a.gitStatusInfo.Modified() > 0 {
		modified := a.getModified(a.gitStatusInfo)
		result = strings.ReplaceAll(result, "$(MODIFIED)", modified)
	} else {
		result = strings.ReplaceAll(result, "$(MODIFIED)", "")
	}
	if a.gitStatusInfo.Deleted() > 0 {
		deleted := a.getDeleted(a.gitStatusInfo)
		result = strings.ReplaceAll(result, "$(DELETED)", deleted)
	} else {
		result = strings.ReplaceAll(result, "$(DELETED)", "")
	}
	if a.gitStatusInfo.Unmerged() > 0 {
		unmerged := a.getUnmerged(a.gitStatusInfo)
		result = strings.ReplaceAll(result, "$(UNMERGED)", unmerged)
	} else {
		result = strings.ReplaceAll(result, "$(UNMERGED)", "")
	}
	if a.gitStatusInfo.Staged() > 0 {
		staged := a.getStaged(a.gitStatusInfo)
		result = strings.ReplaceAll(result, "$(STAGED)", staged)
	} else {
		result = strings.ReplaceAll(result, "$(STAGED)", "")
	}

	return result
}

func (a *advancedPrompt) getBranch(info *gitStatus.GitStatusInfo) string {
	return "⎇ " + info.Branch()
}

func (a *advancedPrompt) getAhead(info *gitStatus.GitStatusInfo) string {
	return "↑" + strconv.Itoa(info.Ahead())
}

func (a *advancedPrompt) getBehind(info *gitStatus.GitStatusInfo) string {
	return "↓" + strconv.Itoa(info.Behind())
}

func (a *advancedPrompt) getUntracked(info *gitStatus.GitStatusInfo) string {
	return "+" + strconv.Itoa(info.Untracked())
}

func (a *advancedPrompt) getModified(info *gitStatus.GitStatusInfo) string {
	return "~" + strconv.Itoa(info.Modified())
}
func (a *advancedPrompt) getDeleted(info *gitStatus.GitStatusInfo) string {
	return "-" + strconv.Itoa(info.Deleted())
}
func (a *advancedPrompt) getUnmerged(info *gitStatus.GitStatusInfo) string {
	return "x" + strconv.Itoa(info.Unmerged())
}
func (a *advancedPrompt) getStaged(info *gitStatus.GitStatusInfo) string {
	return "•" + strconv.Itoa(info.Staged())
}
