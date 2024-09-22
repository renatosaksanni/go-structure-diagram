package analyzer

import (
	"go-structure-diagram/internal/analyzer/globals"
	"go-structure-diagram/internal/analyzer/methods"
	"go-structure-diagram/internal/analyzer/packages"
	"go-structure-diagram/internal/analyzer/structs"
	"go-structure-diagram/pkg/diagram"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// Config holds the configuration for what to analyze
type Config struct {
	AnalyzeStructs  bool
	AnalyzePackages bool
	AnalyzeGlobals  bool
	AnalyzeMethods  bool
}

// Analyze performs static analysis on the Go project and returns the interactions found.
// It uses concurrency to speed up the analysis process.
func Analyze(projectPath string, config Config) ([]diagram.Interaction, error) {
	var interactions []diagram.Interaction
	fset := token.NewFileSet()

	fileCh := make(chan string, 100)                  // Buffered channel for file paths
	resultCh := make(chan []diagram.Interaction, 100) // Buffered channel for results
	errCh := make(chan error, 1)                      // Channel for errors

	var wg sync.WaitGroup
	numWorkers := 8 // Adjust based on CPU or needs

	// Worker function
	worker := func() {
		defer wg.Done()
		for path := range fileCh {
			file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
			if err != nil {
				errCh <- err
				return
			}
			var fileInteractions []diagram.Interaction

			if config.AnalyzeStructs {
				fileInteractions = append(fileInteractions, structs.Analyze(file, projectPath)...)
			}
			if config.AnalyzePackages {
				fileInteractions = append(fileInteractions, packages.Analyze(file, projectPath)...)
			}
			if config.AnalyzeGlobals {
				fileInteractions = append(fileInteractions, globals.Analyze(file, projectPath)...)
			}
			if config.AnalyzeMethods {
				fileInteractions = append(fileInteractions, methods.Analyze(file, projectPath)...)
			}

			resultCh <- fileInteractions
		}
	}

	// Start workers
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go worker()
	}

	// Walk the file tree and send file paths to fileCh
	go func() {
		err := filepath.Walk(projectPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") && !strings.HasSuffix(info.Name(), "_test.go") {
				fileCh <- path
			}
			return nil
		})
		if err != nil {
			errCh <- err
		}
		close(fileCh)
	}()

	// Collect results
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// Monitor errors and collect interactions
	for {
		select {
		case err := <-errCh:
			return interactions, err
		case fileInteractions, ok := <-resultCh:
			if !ok {
				return interactions, nil
			}
			interactions = append(interactions, fileInteractions...)
		}
	}
}
