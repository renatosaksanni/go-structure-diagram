package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"go-structure-diagram/internal/analyzer"
	"go-structure-diagram/internal/generator"
	"go-structure-diagram/internal/renderer"
)

func main() {
	var outputFormat string
	var outputPath string
	var analyzeStructs bool
	var analyzePackages bool
	var analyzeGlobals bool
	var analyzeMethods bool

	var rootCmd = &cobra.Command{
		Use:   "gostructure",
		Short: "Generate diagrams from your Go project",
		Run: func(cmd *cobra.Command, args []string) {
			projectPath := "."
			fmt.Println("Starting analysis...")
			config := analyzer.Config{
				AnalyzeStructs:  analyzeStructs,
				AnalyzePackages: analyzePackages,
				AnalyzeGlobals:  analyzeGlobals,
				AnalyzeMethods:  analyzeMethods,
			}
			interactions, err := analyzer.Analyze(projectPath, config)
			if err != nil {
				fmt.Println("Error during analysis:", err)
				os.Exit(1)
			}

			fmt.Println("Generating diagrams...")
			diagrams := generator.Generate(interactions)

			fmt.Println("Rendering diagrams...")
			err = renderer.Render(diagrams, outputFormat, outputPath)
			if err != nil {
				fmt.Println("Error during rendering:", err)
				os.Exit(1)
			}

			fmt.Println("Diagrams generated successfully!")
		},
	}

	rootCmd.Flags().StringVarP(&outputFormat, "format", "f", "mermaid", "Output format (mermaid/png)")
	rootCmd.Flags().StringVarP(&outputPath, "output", "o", "diagram.mmd", "Output file path")
	rootCmd.Flags().BoolVar(&analyzeStructs, "structs", true, "Analyze structs and interfaces")
	rootCmd.Flags().BoolVar(&analyzePackages, "packages", true, "Analyze package dependencies")
	rootCmd.Flags().BoolVar(&analyzeGlobals, "globals", true, "Analyze global variables and constants")
	rootCmd.Flags().BoolVar(&analyzeMethods, "methods", true, "Analyze method and function calls")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
