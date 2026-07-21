package main

import (
	"fmt"
	"syscall"
)

type StorageEntry struct {
	Name string
	Size string
}

type StorageInfo struct {
	Entries   []StorageEntry
	TotalUsed float64
	TotalSize float64
}

func getStorageInfo() StorageInfo {
	rootUsed, rootTotal := getDiskStats("/")
	homeUsed, homeTotal := getDiskStats("/home")

	return StorageInfo{
		Entries: []StorageEntry{
			{Name: "/  (root)", Size: fmt.Sprintf("%.0f / %.0f GB", rootUsed, rootTotal)},
			{Name: "/home", Size: fmt.Sprintf("%.0f / %.0f GB", homeUsed, homeTotal)},
		},
		TotalUsed: rootUsed + homeUsed,
		TotalSize: rootTotal + homeTotal,
	}
}

func formatSize(bytes int64) string {
	gb := float64(bytes) / 1e9
	if gb >= 1 {
		return fmt.Sprintf("%.1f GB", gb)
	}
	mb := float64(bytes) / 1e6
	if mb >= 1 {
		return fmt.Sprintf("%.0f MB", mb)
	}
	return fmt.Sprintf("%.0f KB", float64(bytes)/1e3)
}

// getDiskStats returns used and total GB for a given mount point
func getDiskStats(path string) (used, total float64) {
	var stat syscall.Statfs_t
	if err := syscall.Statfs(path, &stat); err != nil {
		return 0, 0
	}
	totalBytes := stat.Blocks * uint64(stat.Bsize)
	freeBytes := stat.Bfree * uint64(stat.Bsize)
	usedBytes := totalBytes - freeBytes
	return float64(usedBytes) / 1e9, float64(totalBytes) / 1e9
}