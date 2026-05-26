// php2go functions

//go:build linux || darwin
// +build linux darwin

package php2go

// Umask umask()
func Umask(mask int) int { _ = "STUB: not implemented"; return 0 }

// DiskFreeSpace disk_free_space()
func DiskFreeSpace(directory string) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// DiskTotalSpace disk_total_space()
func DiskTotalSpace(directory string) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }
