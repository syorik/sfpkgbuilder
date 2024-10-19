package main

import (
	"flag"
	"fmt"
	"os"
)

type FullPackageArgs struct {
	APIVersion string
}

type DiffPackageArgs struct {
	APIVersion   string
	SourceBranch string
	TargetBranch string
	RepoPath     string
}

func main() {
	fullCmd := flag.NewFlagSet("full", flag.ExitOnError)
	fullAPIVersion := fullCmd.String("api-version", "", "API version for full package generation")

	diffCmd := flag.NewFlagSet("diff", flag.ExitOnError)
	diffAPIVersion := diffCmd.String("api-version", "", "API version for diff package generation")
	sourceBranch := diffCmd.String("source", "", "Source branch for diff")
	targetBranch := diffCmd.String("target", "", "Target branch for diff")
	repoPath := diffCmd.String("repo", "", "Path to git repository")

	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "full":
		fullCmd.Parse(os.Args[2:])
		args := FullPackageArgs{
			APIVersion: *fullAPIVersion,
		}
		handleFullPackage(args)
	case "diff":
		diffCmd.Parse(os.Args[2:])
		args := DiffPackageArgs{
			APIVersion:   *diffAPIVersion,
			SourceBranch: *sourceBranch,
			TargetBranch: *targetBranch,
			RepoPath:     *repoPath,
		}
		handleDiffPackage(args)
	case "help":
		printHelp()
	default:
		fmt.Println("Unknown command. Use 'help' for usage information.")
		os.Exit(1)
	}
}

func handleFullPackage(args FullPackageArgs) {
	fmt.Println("Handling full package generation")
	fmt.Printf("API Version: %s\n", args.APIVersion)
}

func handleDiffPackage(args DiffPackageArgs) {
	fmt.Println("Handling diff package generation")
	fmt.Printf("API Version: %s\n", args.APIVersion)
	fmt.Printf("Source Branch: %s\n", args.SourceBranch)
	fmt.Printf("Target Branch: %s\n", args.TargetBranch)
	fmt.Printf("Repository Path: %s\n", args.RepoPath)
}

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("  command [subcommand] [options]")
	fmt.Println("\nSubcommands:")
	fmt.Println("  full        Generate a full package")
	fmt.Println("  diff        Generate a diff package")
	fmt.Println("  help        Show this help message")
	fmt.Println("\nOptions for 'full':")
	fmt.Println("  -api-version string    API version for full package generation")
	fmt.Println("\nOptions for 'diff':")
	fmt.Println("  -api-version string    API version for diff package generation")
	fmt.Println("  -source string         Source branch for diff")
	fmt.Println("  -target string         Target branch for diff")
	fmt.Println("  -repo string           Path to git repository")
}
