package utils

import (
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"log"
	"time"
)

// Git tracks the repository used through the application.
type Git struct {
	Repo *git.Repository

	// TODO: Consider whether it makes sense to allow path to be a value
	// that is not identical to Config.Workspace.
	Path string
}

// Open the repository.
func (g *Git) Open() error {
	var err error
	g.Repo, err = git.PlainOpen(g.Path)
	if err == git.ErrRepositoryNotExists {
		// Initialize it!
		g.Repo, err = git.PlainInit(g.Path, false)
		if err != nil {
			log.Fatalf("No repo, could not init: %v", err)
		}
	} else if err != nil {
		// Unexpected... panic provides trace.
		panic(err)
	}

	return nil
}

// GitDiffs string for pending changes.
func (g *Git) GitDiffs() {

}

// GitCommit any changes found.
func (g *Git) GitCommit() error {
	w, err := g.Repo.Worktree()
	if err != nil {
		return err
	}

	s, err := w.Status()
	if err != nil {
		return err
	}
	if s.IsClean() {
		// Nothing to do
		log.Println("No changes.")
		return nil
	}
	//log.Println(s.String())

	// Walk through the modified files.
	for e, ss := range s {
		switch ss.Worktree {
		case git.Untracked:
			log.Printf("New file: %v\n", e)
			w.Add(e)
		case git.Modified:
			log.Printf("Modified: %v\n", e)
			w.Add(e)
		default:
			log.Printf("Unhandled: [%v] %v", string(ss.Worktree), e)
		}
	}

	hash, err := w.Commit("Splendid commit.", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "Splendid",
			Email: "splendid@example.com",
			When:  time.Now(),
		},
	})
	if err != nil {
		return err
	}

	log.Printf("Committed: %v", hash)
	return nil
}

// GitHash for the filename.
func (g *Git) GitHash(filename string) string {

	l, err := g.Repo.Head()
	if err == plumbing.ErrReferenceNotFound {
		// TODO: A repo with no commits errors.
		log.Fatalf("HEAD reference not found.\nYou must make at least "+
			"one commit if you initialized the repository yourself.\n%v", err)
	} else if err != nil {
		// Unexpected... panic provides trace.
		panic(err)
	}
	log.Printf("Head hash: %v", l.Hash())

	cIter, err := g.Repo.Log(&git.LogOptions{})
	if err != nil {
		panic(err)
	}
	cIter.ForEach(func(commit *object.Commit) error {
		log.Println(commit.String())
		return nil
	})

	return l.String()
}
