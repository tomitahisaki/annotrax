package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// annotationKeywords
// -----------------------------------------------------------------------------
// A list of annotation-style keywords that this tool will search for.
// If a line contains any of these strings, the line will be printed as a result.
// This list can later be made configurable.
var annotationKeywords = []string{
	"TODO",
	"FIXME",
	"NOTE",
	"HACK",
	"XXX",
}

func main() {
	//
	// Define the -dir option.
	// StringVar lets us bind a value directly to a named variable,
	// which is easier to understand than flag.String().
	//
	var directoryToScan string
	flag.StringVar(&directoryToScan, "dir", ".", "directory to scan")

	// Parse CLI args (e.g., `annotrax -dir src`)
	flag.Parse()

	//
	// Execute the main scanning process.
	//
	err := run(directoryToScan)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

// run
// -----------------------------------------------------------------------------
// Recursively walks through the directory tree starting at rootPath.
// For each file found, it performs annotation scanning.
// filepath.WalkDir helps traverse both files and directories.
func run(rootPath string) error {
	return filepath.WalkDir(rootPath, func(path string, entry fs.DirEntry, err error) error {
		// If there was an error accessing the file/directory, show a warning and continue.
		if err != nil {
			fmt.Fprintf(os.Stderr, "warn: cannot access %s: %v\n", path, err)
			return nil
		}

		// Handle directories
		if entry.IsDir() {
			directoryName := entry.Name()

			// Skip certain directories entirely, as they are usually irrelevant
			// or extremely large:
			//   - .git         : Git metadata
			//   - node_modules : npm packages
			//   - vendor       : dependency directories
			//
			if directoryName == ".git" || directoryName == "node_modules" || directoryName == "vendor" {
				return filepath.SkipDir
			}

			// No further processing needed for directories.
			return nil
		}

		// Skip binary-like files (images, archives, executables, etc.)
		if isNonTextFile(path) {
			return nil
		}

		// Scan the file for annotation-style keywords.
		scanErr := scanFile(path)
		if scanErr != nil {
			// Do not stop the entire process if one file cannot be read.
			fmt.Fprintf(os.Stderr, "warn: cannot read %s: %v\n", path, scanErr)
		}

		return nil
	})
}

// scanFile
// -----------------------------------------------------------------------------
// Reads a file line-by-line using bufio.Scanner.
// If any line contains a keyword such as TODO or FIXME,
// the line will be printed with the file path and line number.
//
// Note: bufio.Scanner is convenient for line-based reading,
// but it has a line length limit; this is usually fine for source code.
func scanFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		lineText := scanner.Text()

		// Check if this line contains any annotation keyword.
		found, keyword := containsAnnotation(lineText)
		if found {
			fmt.Printf(
				"%s:%d: [%s] %s\n",
				path,
				lineNumber,
				keyword,
				strings.TrimSpace(lineText),
			)
		}
	}

	// Return any scanning error (e.g., extremely long lines).
	return scanner.Err()
}

// containsAnnotation
// -----------------------------------------------------------------------------
// Checks whether a line contains any of the annotation keywords.
//
// Example:
//
//	"// TODO: fix this"  -> (true, "TODO")
//	"normal line"        -> (false, "")
func containsAnnotation(lineText string) (bool, string) {
	for _, keyword := range annotationKeywords {
		if strings.Contains(lineText, keyword) {
			return true, keyword
		}
	}
	return false, ""
}

// isProbablyBinary
// -----------------------------------------------------------------------------
// Determines whether a file is "probably" a binary file based on its extension.
// This is a lightweight heuristic, not a perfect check.
// Useful for skipping images, archives, and executables during scanning.
func isNonTextFile(path string) bool {
	binarySuffixes := []string{
		".png", ".jpg", ".jpeg", ".gif", ".ico",
		".pdf", ".zip", ".tar", ".gz",
		".exe", ".dll", ".so", ".webp", ".svg",
	}

	for _, suffix := range binarySuffixes {
		if strings.HasSuffix(path, suffix) {
			return true
		}
	}
	return false
}
