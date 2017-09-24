package utils

import (
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"log"
	"os/exec"
)

// Shim using exec until we determine the best native Go implementation.

// GitExecDiff uses os/exec to pull a git diff.
func GitExecDiff(path, filename string) (string, error) {
	// Prep the command.
	cmd := exec.Command("git", "diff", "-U4", filename)
	// Switch to the workspace folder.
	cmd.Dir = path
	// Run!
	diff, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(diff), err
}

// Work In Progress -- attempts to use go-git

// GitDiff4 based on https://github.com/src-d/go-git/issues/468
func (g *Git) GitDiff4(filename string) string {
	cIter, err := g.Repo.CommitObjects()
	if err != nil {
		panic(err)
	}
	cIter.ForEach(func(commit *object.Commit) error {
		cTree, err := commit.Tree()
		if err != nil {
			panic(err)
		}
		commit.Parents().ForEach(func(parent *object.Commit) error {
			pTree, err := parent.Tree()
			if err != nil {
				panic(err)
			}
			changes, err := pTree.Diff(cTree)
			if err != nil {
				panic(err)
			}
			for _, change := range changes {
				if change.From.Name == filename ||
					change.To.Name == filename {
					log.Println("Commit affected file.")
					log.Println(commit.String())
					return nil
				}
			}
			return nil
		})

		return nil
	})
	return ""
}

// GitDiff just tests blame.
func (g *Git) GitDiff(filename string) string {
	h, err := g.Repo.Head()
	if err != nil {
		panic(err)
	}
	c, err := g.Repo.CommitObject(h.Hash())
	if err != nil {
		panic(err)
	}

	//log.Println(c)

	//f, err := c.File("test")
	//if err != nil {
	//	panic(err)
	//}

	b, err := git.Blame(c, "test")
	if err != nil {
		panic(err)
	}
	for _, line := range b.Lines {
		log.Printf("%v :  %v", line.Author, line.Text)
	}
	return ""
}

// GitDiff2 testing diff trees.
func (g *Git) GitDiff2(filename string) string {
	h, err := g.Repo.Head()
	if err != nil {
		panic(err)
	}
	c, err := g.Repo.CommitObject(h.Hash())
	if err != nil {
		panic(err)
	}
	log.Println(c)
	//f, err := c.File("test")
	//if err != nil {
	//	panic(err)
	//}
	t1, _ := c.Tree()
	object.DiffTree(t1, t1)
	return ""
}

// GitDiff3 testing file inspection.
func (g *Git) GitDiff3(filename string) string {
	w, err := g.Repo.Worktree()
	if err != nil {
		panic(err)
	}

	s, err := w.Status()
	if err != nil {
		panic(err)
	}
	fs := s.File("test")

	log.Println(s.String())
	log.Println(fs)
	log.Println(fs.Extra)
	return ""
}
