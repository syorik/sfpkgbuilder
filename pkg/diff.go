package pkg

import (
	"fmt"
	"os/exec"
	"strings"
)

const rootDir = "force-app/main/default"

func GetChangedFilesByDirectory(sourceCommit, targetCommit, repoPath string) (map[string][]string, error) {
	cmd := exec.Command("git", "diff", "--name-only", "--ignore-space-at-eol", "--ignore-space-change", "--ignore-all-space", sourceCommit, targetCommit)
	cmd.Dir = repoPath

	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error executing git diff: %w", err)
	}

	changedFiles := strings.Split(strings.TrimSpace(string(output)), "\n")
	result := make(map[string][]string)

	for _, file := range changedFiles {
		if !strings.HasPrefix(file, rootDir) {
			continue
		}

		relPath := strings.TrimPrefix(file, rootDir+"/")
		parts := strings.SplitN(relPath, "/", 2)

		if len(parts) < 2 {
			continue
		}

		dir, filePath := parts[0], parts[1]
		result[dir] = append(result[dir], filePath)
	}

	return result, nil
}
