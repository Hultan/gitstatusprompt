package gitStatusPrompt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGitPrompt_NotAGitRepository(t *testing.T) {
	prompt := &GitStatusPrompt{}
	assert.NotNil(t, prompt, "should not be nil")

	p := prompt.GetAdvancedPrompt("/home/per", " $(BRANCH)$(AHEAD)$(BEHIND) | $(UNTRACKED)$(MODIFIED)$(DELETED)$(UNMERGED)$(STAGED) ")
	assert.Equal(t, "", p, "should not be a git repository")
}

func TestGitPrompt_GitRepository(t *testing.T) {
	prompt := &GitStatusPrompt{}
	assert.NotNil(t, prompt, "should not be nil")

	p := prompt.GetAdvancedPrompt("/home/per/code/gitprompt-go-test", " $(BRANCH)$(AHEAD)$(BEHIND) | $(UNTRACKED)$(MODIFIED)$(DELETED)$(UNMERGED)$(STAGED) ")
	assert.Equal(t, " ⎇ testBranch | +1~1-1•1 ", p, "invalid prompt")
}

func TestGitPrompt_EmptyPath(t *testing.T) {
	prompt := &GitStatusPrompt{}
	assert.NotNil(t, prompt, "should not be nil")

	p := prompt.GetAdvancedPrompt("", " $(BRANCH)$(AHEAD)$(BEHIND) | $(UNTRACKED)$(MODIFIED)$(DELETED)$(UNMERGED)$(STAGED) ")
	assert.NotEqual(t, p, "invalid prompt")
}

func TestGitPrompt_Verbose(t *testing.T) {
	prompt := &GitStatusPrompt{}
	assert.NotNil(t, prompt, "should not be nil")

	p := prompt.GetVerbosePrompt("/home/per/code/gitprompt-go-test")
	assert.Equal(t, `Branch    : testBranch (0 ahead, 0 behind)
Staged    : 1
Modified  : 1
Deleted   : 1
Unmerged  : 0
Untracked : 1
`, p, "invalid prompt")
}
