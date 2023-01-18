package jsonnet

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	jsonnet "github.com/google/go-jsonnet"
)

// FileMemImporter is an importer for the jsonnet library allowing the usage of jsonnet code from disk or from memory
type FileMemImporter struct {
	Data    map[string]jsonnet.Contents
	JPaths  []string
	fsCache map[string]*fsCacheEntry
}

type fsCacheEntry struct {
	exists   bool
	contents jsonnet.Contents
}

func (importer *FileMemImporter) tryPath(dir, importedPath string) (found bool, contents jsonnet.Contents, foundHere string, err error) {
	if importer.fsCache == nil {
		importer.fsCache = make(map[string]*fsCacheEntry)
	}
	var absPath string
	if path.IsAbs(importedPath) {
		absPath = importedPath
	} else {
		absPath = path.Join(dir, importedPath)
	}
	var entry *fsCacheEntry
	if cacheEntry, isCached := importer.fsCache[absPath]; isCached {
		entry = cacheEntry
	} else {
		contentBytes, err := ioutil.ReadFile(absPath)
		if err != nil {
			if os.IsNotExist(err) {
				entry = &fsCacheEntry{
					exists: false,
				}
			} else {
				return false, jsonnet.Contents{}, "", err
			}
		} else {
			entry = &fsCacheEntry{
				exists:   true,
				contents: jsonnet.MakeContents(string(contentBytes)),
			}
		}
		importer.fsCache[absPath] = entry
	}
	return entry.exists, entry.contents, absPath, nil
}

// Import imports a map entry.
func (importer *FileMemImporter) Import(importedFrom, importedPath string) (contents jsonnet.Contents, foundAt string, err error) {
	// Memory importer
	if content, ok := importer.Data[importedPath]; ok {
		return content, importedPath, nil
	}

	// File importer
	dir, _ := path.Split(importedFrom)
	found, content, foundHere, err := importer.tryPath(dir, importedPath)
	if err != nil {
		return jsonnet.Contents{}, "", err
	}

	for i := len(importer.JPaths) - 1; !found && i >= 0; i-- {
		found, content, foundHere, err = importer.tryPath(importer.JPaths[i], importedPath)
		if err != nil {
			return jsonnet.Contents{}, "", err
		}
	}

	if !found {
		return jsonnet.Contents{}, "", fmt.Errorf("couldn't open import %#v: no match locally or in the Jsonnet library paths", importedPath)
	}
	return content, foundHere, nil
}
