package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/syorik/sfpkgbuilder/pkg"
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

	packageDefinition := pkg.NewPackage(pkg.WithVersion(args.APIVersion))

	// Add all metadata types with "*" as member
	packageDefinition.AddMember(pkg.ApexClassMdt, "*")
	packageDefinition.AddMember(pkg.ApexTriggerMdt, "*")
	packageDefinition.AddMember(pkg.CustomObjectMdt, "*")
	packageDefinition.AddMember(pkg.CustomLabelMdt, "*")
	packageDefinition.AddMember(pkg.CustomObjectTranslationMdt, "*")
	packageDefinition.AddMember(pkg.CustomPermissionMdt, "*")
	packageDefinition.AddMember(pkg.CustomTabMdt, "*")
	packageDefinition.AddMember(pkg.ExperienceBundleMdt, "*")
	packageDefinition.AddMember(pkg.FlexiPageMdt, "*")
	packageDefinition.AddMember(pkg.FlowMdt, "*")
	packageDefinition.AddMember(pkg.GlobalPicklistMdt, "*")
	packageDefinition.AddMember(pkg.GlobalValueSetMdt, "*")
	packageDefinition.AddMember(pkg.GlobalValueSetTranslationMdt, "*")
	packageDefinition.AddMember(pkg.LightningComponentBundleMdt, "*")
	packageDefinition.AddMember(pkg.LightningMessageChannelMdt, "*")
	packageDefinition.AddMember(pkg.MilestoneTypeMdt, "*")
	packageDefinition.AddMember(pkg.PermissionSetMdt, "*")
	packageDefinition.AddMember(pkg.PermissionSetGroupMdt, "*")
	packageDefinition.AddMember(pkg.PlatformEventChannelMdt, "*")
	packageDefinition.AddMember(pkg.ProfileMdt, "*")
	packageDefinition.AddMember(pkg.QueueMdt, "*")
	packageDefinition.AddMember(pkg.StandardValueSetMdt, "*")
	packageDefinition.AddMember(pkg.StandardValueSetTranslationMdt, "*")
	packageDefinition.AddMember(pkg.StaticResourceMdt, "*")

	xmlStr, err := packageDefinition.ToXMLString()
	if err != nil {
		fmt.Printf("Error generating package XML: %v\n", err)
		return
	}

	fmt.Println("Generated package XML:")
	fmt.Println(xmlStr)

	err = os.WriteFile("package.xml", []byte(xmlStr), 0644)
	if err != nil {
		fmt.Printf("Error saving package XML to file: %v\n", err)
		return
	}
	fmt.Println("Package XML saved to package.xml")
}

func handleDiffPackage(args DiffPackageArgs) {
	fmt.Println("Handling diff package generation")
	fmt.Printf("API Version: %s\n", args.APIVersion)
	fmt.Printf("Source Branch: %s\n", args.SourceBranch)
	fmt.Printf("Target Branch: %s\n", args.TargetBranch)
	fmt.Printf("Repository Path: %s\n", args.RepoPath)

	changedFiles, err := pkg.GetChangedFilesByDirectory(args.SourceBranch, args.TargetBranch, args.RepoPath)
	if err != nil {
		fmt.Printf("Error getting changed files: %v\n", err)
		return
	}

	fmt.Println("\n\nChanged files by directory:")
	for dir, files := range changedFiles {
		fmt.Printf("%s:\n", dir)
		for _, file := range files {
			fmt.Printf("  %s\n", file)
		}
	}
	fmt.Println()
	packageDefinition := pkg.NewPackage(pkg.WithVersion(args.APIVersion))
	for dir, files := range changedFiles {
		metadataType := pkg.MapDirectoryToMetadataType(dir)
		if metadataType == "" {
			fmt.Printf("Warning: Unknown metadata type for directory '%s'. Skipping.\n", dir)
			continue
		}

		if metadataType == pkg.ApexClassMdt ||
			metadataType == pkg.ProfileMdt ||
			metadataType == pkg.PermissionSetMdt ||
			metadataType == pkg.GlobalValueSetMdt ||
			metadataType == pkg.StandardValueSetMdt ||
			metadataType == pkg.FlowMdt ||
			metadataType == pkg.CustomPermissionMdt {
			for _, file := range files {
				parts := strings.Split(file, "/")
				fileName := parts[len(parts)-1]
				memberName := strings.Split(fileName, ".")[0]
				packageDefinition.AddMember(metadataType, memberName)
			}
		} else if metadataType == pkg.CustomFieldMdt {
			for _, file := range files {
				parts := strings.Split(file, "/")
				if len(parts) >= 4 {
					objectName := parts[1]
					fieldName := strings.Split(parts[3], ".")[0]
					memberName := fmt.Sprintf("%s.%s", objectName, fieldName)
					packageDefinition.AddMember(metadataType, memberName)
				}
			}
		} else if metadataType == pkg.LightningComponentBundleMdt {
			for _, file := range files {
				parts := strings.Split(file, "/")
				if len(parts) >= 2 {
					memberName := parts[0]
					packageDefinition.AddMember(metadataType, memberName)
				}
			}
		} else if metadataType == pkg.CustomObjectMdt {
			for _, file := range files {
				parts := strings.Split(file, "/")
				if len(parts) >= 1 {
					memberName := parts[0]
					packageDefinition.AddMember(metadataType, memberName)
				}
			}
		} else if metadataType == pkg.ListViewMdt {
			for _, file := range files {
				parts := strings.Split(file, "/")
				if len(parts) >= 4 {
					objectName := parts[1]
					listViewName := strings.Split(parts[3], ".")[0]
					memberName := fmt.Sprintf("%s.%s", objectName, listViewName)
					packageDefinition.AddMember(metadataType, memberName)
				}
			}
		}
	}

	xmlStr, err := packageDefinition.ToXMLString()
	if err != nil {
		fmt.Printf("Error generating package XML: %v\n", err)
		return
	}

	fmt.Println("Generated package XML:")
	fmt.Println(xmlStr)

	err = os.WriteFile("package.xml", []byte(xmlStr), 0644)
	if err != nil {
		fmt.Printf("Error saving package XML to file: %v\n", err)
		return
	}
	fmt.Println("Package XML saved to package.xml")
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
