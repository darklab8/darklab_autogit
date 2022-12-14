package git

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGitRepo(t *testing.T) {
	repo := (&Repository{}).TestNewRepo()
	repo.TestCommit("feat: test")
	repo.TestCommit("feat: test3")
	repo.TestCommit("feat: test5")
	repo.TestCreateTag("v0.0.1", repo.TestCommit("fix: thing"))
	repo.TestCommit("feat(api): test")
	repo.TestCreateTag("v0.0.2", repo.TestCommit("feat(api): test2"))
	repo.TestCommit("fix: test1")
	repo.TestCommit("fix: test2")
	repo.TestCommit("fix: test3")

	tags := repo.getUnorderedTags()
	fmt.Printf("tags=%v\n", tags)
	assert.Equal(t, 2, len(tags))
}
