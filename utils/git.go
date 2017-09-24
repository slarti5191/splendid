package utils

import (
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"log"
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
	fs := s.File("test")

	log.Println(s.String())
	log.Println(fs)
	log.Println(fs.Extra)
	return ""
}
