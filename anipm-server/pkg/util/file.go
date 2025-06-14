package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func WriteFileFrom(path string, src io.Reader) error {
	dir, filename := filepath.Dir(path), filepath.Base(path)

	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	tmpFilePath := filepath.Join(dir, fmt.Sprintf(".%s.tmp", filename))

	if err := os.Remove(tmpFilePath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("failed to remove file: %w", err)
	}
	dst, err := os.Create(tmpFilePath)
	if err != nil {
		return fmt.Errorf("failed to create temporary file: %w", err)
	}
	defer func() {
		dst.Close()
		os.Remove(tmpFilePath)
	}()

	_, err = io.Copy(dst, src)
	if err != nil {
		return fmt.Errorf("failed to copy file: %w", err)
	}
	dst.Close()

	backupPath := path + ".backup"
	if _, err := os.Stat(path); err == nil {
		if err := os.Remove(backupPath); err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("failed to remove file: %w", err)
		}
		if err := os.Rename(path, backupPath); err != nil {
			return fmt.Errorf("failed to backup old file: %w", err)
		}
	}

	if err := os.Rename(tmpFilePath, path); err != nil {
		return fmt.Errorf("failed to move temp file to final destination: %w", err)
	}
	return nil
}

func WriteFile(path string, data []byte) error {
	return WriteFileFrom(path, bytes.NewReader(data))
}

func WriteFileWithJson(path string, data interface{}) error {
	j, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal json: %w", err)
	}
	return WriteFile(path, j)
}

func RemoveEmptyDirs(root string) error {
	dirs := []string{}

	queue := []string{root}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		dirs = append(dirs, curr)

		entries, err := os.ReadDir(curr)
		if err != nil {
			return err
		}
		for _, entry := range entries {
			if entry.IsDir() {
				queue = append(queue, filepath.Join(curr, entry.Name()))
			}
		}
	}

	for i := len(dirs) - 1; i >= 0; i-- {
		dir := dirs[i]
		entries, err := os.ReadDir(dir)
		if err != nil {
			return err
		}
		if len(entries) == 0 && dir != root {
			if err := os.Remove(dir); err != nil {
				return err
			}
		}
	}
	return nil
}
