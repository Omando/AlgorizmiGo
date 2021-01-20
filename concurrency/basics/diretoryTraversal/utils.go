package diretoryTraversal

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// walkDir recursively walks the directory tree rooted at dir.
// It sends the size of each processed file into fileSizes channel
func walkDir(dir string, fileSizes chan <- int64) {
	// Get dirEntries present under dir. Log error and return if unable to process files
	// for the given dir
	dirEntries, err := ioutil.ReadDir(dir) // ioutil.ReadDir returns a slice of os.FileInfo
	if err != nil {
		fmt.Printf("Error reading %s", dir)
		return
	}

	// We have dirEntries that exist under dir; if entry is a file, report its size,
	// otherwise, keep on walking the directory
	for _, entry := range dirEntries {
		if entry.IsDir() {
			// Get full path for this subdirectory
			fullDir := filepath.Join(dir, entry.Name())
			fmt.Printf("Processing %s\n", fullDir)

			// Walk this directory
			walkDir(fullDir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func printDiskUsage(fileCount int32, totalSize int64) {
	fmt.Printf("%d files: %.2f GB\n", fileCount,  (float64(totalSize) / 1e9))
}

