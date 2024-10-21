package pkg

import (
	"fmt"
	"os/exec"
	"strings"
)

const rootDir = "force-app/main/default"

func GetChangedFilesByDirectory(sourceCommit, targetCommit, repoPath string) (map[string][]string, error) {
	cmd := exec.Command("git", "diff", "--name-status", "--ignore-space-at-eol", "--ignore-space-change", "--ignore-all-space", sourceCommit, targetCommit)
	cmd.Dir = repoPath

	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error executing git diff: %w", err)
	}

	changedFiles := strings.Split(strings.TrimSpace(string(output)), "\n")
	result := make(map[string][]string)

	for _, line := range changedFiles {
		parts := strings.SplitN(line, "\t", 2)
		if len(parts) != 2 {
			continue
		}

		status, file := parts[0], parts[1]
		if status == "D" || !strings.HasPrefix(file, rootDir) {
			continue
		}

		relPath := strings.TrimPrefix(file, rootDir+"/")
		parts = strings.SplitN(relPath, "/", 3)

		if len(parts) < 2 {
			continue
		}

		dir := parts[0]
		filePath := strings.Join(parts[1:], "/")

		if dir == "objects" {
			if strings.Contains(filePath, "/fields/") {
				dir = "fields"
				filePath = relPath
			} else if strings.Contains(filePath, "/listViews/") {
				dir = "listViews"
				filePath = relPath
			}
		}

		result[dir] = append(result[dir], filePath)
	}

	return result, nil
}
