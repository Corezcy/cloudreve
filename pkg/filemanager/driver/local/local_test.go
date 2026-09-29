package local

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/cloudreve/Cloudreve/v4/pkg/filemanager/fs"
)

func TestPutRelocatesCompletedDownloadWithoutCopying(t *testing.T) {
	tempDir := t.TempDir()
	sourcePath := filepath.Join(tempDir, "download.tmp")
	destinationPath := filepath.Join(tempDir, "storage", "file.bin")
	content := []byte("downloaded content")
	if err := os.WriteFile(sourcePath, content, 0600); err != nil {
		t.Fatal(err)
	}

	source, err := os.Open(sourcePath)
	if err != nil {
		t.Fatal(err)
	}
	defer source.Close()

	var progress int64
	req := &fs.UploadRequest{
		Props:      &fs.UploadProps{SavePath: destinationPath, Size: int64(len(content))},
		File:       source,
		MoveSource: true,
		ProgressFunc: func(_, diff, _ int64) {
			progress += diff
		},
	}

	if err := (&Driver{}).Put(context.Background(), req); err != nil {
		t.Fatal(err)
	}
	if !req.SourceRelocated {
		t.Fatal("expected same-file-system source relocation")
	}
	if progress != int64(len(content)) {
		t.Fatalf("progress = %d, want %d", progress, len(content))
	}

	sourceInfo, err := os.Stat(sourcePath)
	if err != nil {
		t.Fatal(err)
	}
	destinationInfo, err := os.Stat(destinationPath)
	if err != nil {
		t.Fatal(err)
	}
	if !os.SameFile(sourceInfo, destinationInfo) {
		t.Fatal("source and destination should be hard links to the same file")
	}

	// This is the post-CompleteUpload cleanup performed by manager.Update.
	if err := os.Remove(sourcePath); err != nil {
		t.Fatal(err)
	}
	if actual, err := os.ReadFile(destinationPath); err != nil || string(actual) != string(content) {
		t.Fatalf("destination content = %q, err = %v", actual, err)
	}
}
