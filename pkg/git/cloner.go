package git

//nolint:typecheck

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"log"
	"os"
)

type ClonerReal struct{}

func (ClonerReal) Clone(path string, repoURL string, branch string) error {
	if pathExists(path) {
		return PullBranch(path, branch)
	}
	return CloneBranch(path, repoURL, branch)
}

func pathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func CloneBranch(path string, repoURL string, branch string) error {
	log.Printf("Cloning branch %s", branch)
	_, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:           repoURL,
		ReferenceName: plumbing.NewBranchReferenceName(branch),
		SingleBranch:  true})
	if err != nil {
		return nil
	}
	return nil
}

func PullBranch(path string, branch string) error {
	log.Printf("Pulling branch %s", branch)
	repo, err := git.PlainOpen(path)
	if err != nil {
		return err
	}

	workTree, err := repo.Worktree()
	if err != nil {
		return err
	}

	// TODO: We remove branch and re-fetch it because pulling with go-git doesn't work.
	branchRef := plumbing.NewBranchReferenceName(branch)
	_, err = repo.Reference(branchRef, true)
	if err == nil {
		branchRef := plumbing.NewBranchReferenceName(branch)
		err = repo.Storer.RemoveReference(branchRef)
	}

	log.Printf("Branch %s does not exist locally, creating it from remote.", branch)
	err = repo.Fetch(&git.FetchOptions{
		RefSpecs: []config.RefSpec{"refs/*:refs/*", "HEAD:refs/heads/HEAD"},
	})
	if err != nil {
		return err
	}

	err = workTree.Checkout(&git.CheckoutOptions{
		Branch: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", branch)),
		Force:  true,
	})
	if err != nil {
		return nil
	}

	return nil
}
