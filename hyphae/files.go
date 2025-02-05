package hyphae

import (
	"log"
	"os"
	"path/filepath"

	"github.com/bouncepaw/mycorrhiza/mimetype"
)

// Index finds all hypha files in the full `path` and saves them to the hypha storage.
func Index(path string) {
	byNames = make(map[string]*Hypha)
	ch := make(chan *Hypha, 5)

	go func(ch chan *Hypha) {
		indexHelper(path, 0, ch)
		close(ch)
	}(ch)

	for h := range ch {
		// It's safe to ignore the mutex because there is a single worker right now.
		if oh := ByName(h.Name); oh.Exists {
			oh.mergeIn(h)
		} else {
			h.insert()
		}
	}
	log.Println("Indexed", Count(), "hyphae")
}

// indexHelper finds all hypha files in the full `path` and sends them to the
// channel. Handling of duplicate entries and attachment and counting them is
// up to the caller.
func indexHelper(path string, nestLevel uint, ch chan *Hypha) {
	nodes, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, node := range nodes {
		// If this hypha looks like it can be a hypha path, go deeper. Do not
		// touch the .git and static folders for they have an administrative
		// importance!
		if node.IsDir() && IsValidName(node.Name()) && node.Name() != ".git" &&
			!(nestLevel == 0 && node.Name() == "static") {
			indexHelper(filepath.Join(path, node.Name()), nestLevel+1, ch)
			continue
		}

		var (
			hyphaPartPath           = filepath.Join(path, node.Name())
			hyphaName, isText, skip = mimetype.DataFromFilename(hyphaPartPath)
			hypha                   = &Hypha{Name: hyphaName, Exists: true}
		)
		if !skip {
			if isText {
				hypha.TextPath = hyphaPartPath
			} else {
				hypha.BinaryPath = hyphaPartPath
			}
			ch <- hypha
		}
	}
}
