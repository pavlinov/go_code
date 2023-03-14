package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

func TestGetIntEnv(t *testing.T) {
	// Test case 1: Environment variable is set to valid integer
	os.Setenv("MY_INT_VAR", "123")
	defer os.Unsetenv("MY_INT_VAR")
	assert.Equal(t, 123, getIntEnv("MY_INT_VAR", 456), "getIntEnv should return the value of the environment variable")

	// Test case 2: Environment variable is not set, so should return the default value
	os.Unsetenv("MY_INT_VAR")
	assert.Equal(t, 456, getIntEnv("MY_INT_VAR", 456), "getIntEnv should return the default value when the environment variable is not set")

	// Test case 3: Environment variable is set to a non-integer value, so should return the default value
	os.Setenv("MY_INT_VAR", "not an integer")
	assert.Equal(t, 456, getIntEnv("MY_INT_VAR", 456), "getIntEnv should return the default value when the environment variable is set to a non-integer value")
}

func TestRemoveFolder(t *testing.T) {
	testDir := "test_dir"
	err := os.MkdirAll(testDir+"/subdir1/subdir2", 0755)
	assert.NoError(t, err, "Error creating test directories")

	_, err = os.Create(testDir + "/file1")
	assert.NoError(t, err, "Error creating test file")

	err = removeFolder(testDir)
	assert.NoError(t, err, "Error removing test directory")

	_, err = os.Stat(testDir)
	assert.Error(t, err, "Expected test directory to be removed, but it still exists")

	_, err = os.Stat(testDir + "/subdir1/subdir2")
	assert.Error(t, err, "Expected test subdirectory to be removed, but it still exists")

	_, err = os.Stat(testDir + "/file1")
	assert.Error(t, err, "Expected test file to be removed, but it still exists")

	// Test case: Try to remove a file instead of a directory
	_, err = os.Create("test_file")
	assert.NoError(t, err, "Failed to create test file")
	defer os.Remove("test_file")
	err = removeFolder("test_file")
	assert.Error(t, err, "removeFolder should return an error when trying to remove a file instead of a directory")

}

func TestCreateFileStructureParallel(t *testing.T) {
	tests := []struct {
		name       string
		folderSrc  string
		foldersCnt int
		filesCnt   int
		wantErr    bool
	}{
		{
			name:       "Create 2 folders with 2 files each",
			folderSrc:  "",
			foldersCnt: 2,
			filesCnt:   2,
			wantErr:    false,
		},
		{
			name:       "Create 1 folder with 1 file",
			folderSrc:  "",
			foldersCnt: 1,
			filesCnt:   1,
			wantErr:    false,
		},
		{
			name:       "Create 2 folders with 1 file each",
			folderSrc:  "",
			foldersCnt: 2,
			filesCnt:   1,
			wantErr:    false,
		},
		{
			name:       "Create 0 folders with 0 files",
			folderSrc:  "",
			foldersCnt: 0,
			filesCnt:   0,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir, err := ioutil.TempDir("", "test")
			require.NoError(t, err)
			defer os.RemoveAll(tmpDir)

			if tt.folderSrc == "" {
				tt.folderSrc = tmpDir
			} else {
				tt.folderSrc = filepath.Join(tmpDir, tt.folderSrc)
			}

			err = createFileStructure(tt.folderSrc, tt.foldersCnt, tt.filesCnt)

			if tt.wantErr {
				require.Error(t, err, "createFileStructure() should have returned an error")
				return
			}
			require.NoError(t, err, "createFileStructure() returned an error")

			for i := 1; i <= tt.foldersCnt; i++ {
				folderName := fmt.Sprintf("folder_%d", i)
				folderPath := filepath.Join(tt.folderSrc, folderName)

				require.DirExists(t, folderPath, "%s folder should have been created but was not", folderPath)

				for j := 1; j <= tt.filesCnt; j++ {
					fileName := "file_" + strconv.Itoa(i) + "-" + strconv.Itoa(j)
					filePath := filepath.Join(folderPath, fileName)

					require.FileExists(t, filePath, "%s file should have been created but was not", filePath)
				}
			}
		})
	}
}
