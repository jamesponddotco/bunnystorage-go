package bunnystorage_test

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"testing"

	"git.sr.ht/~jamesponddotco/bunnystorage-go"
	"git.sr.ht/~jamesponddotco/bunnystorage-go/internal/testutil"
)

const _testPath string = "/testdata"

var (
	client *bunnystorage.Client
	err    error
)

func TestMain(m *testing.M) {
	client, err = testutil.SetupClient()
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}

func TestClient_NewClient(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		giveConfig *bunnystorage.Config
		wantErr    bool
	}{
		{
			name: "success with default Application",
			giveConfig: &bunnystorage.Config{
				Application: bunnystorage.DefaultApplication(),
				StorageZone: "my-storage-zone",
				Key:         "my-key",
				ReadOnlyKey: "my-read-only-key",
				Endpoint:    bunnystorage.EndpointFalkenstein,
				MaxRetries:  bunnystorage.DefaultMaxRetries,
				Timeout:     bunnystorage.DefaultTimeout,
			},
			wantErr: false,
		},
		{
			name: "success with custom Application and Config",
			giveConfig: &bunnystorage.Config{
				Application: nil,
				StorageZone: "my-storage-zone",
				Key:         "my-key",
				ReadOnlyKey: "my-read-only-key",
				Endpoint:    bunnystorage.EndpointFalkenstein,
			},
			wantErr: false,
		},
		{
			name: "error with default Application",
			giveConfig: &bunnystorage.Config{
				Application: bunnystorage.DefaultApplication(),
			},
			wantErr: true,
		},
		{
			name:       "error with nil Config",
			giveConfig: nil,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			_, err := bunnystorage.NewClient(tt.giveConfig)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClient_List(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	files, resp, err := client.List(ctx, _testPath)
	if err != nil {
		t.Fatal(err)
	}

	if len(files) == 0 {
		t.Errorf("expected file list to be non-zero, got %d", len(files))
	}

	if resp.Status != 200 {
		t.Errorf("expected status code to be 200, got %d", resp.Status)
	}
}

func TestClient_UploadDownloadDelete(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	testFile, testFileSize, err := testutil.SetupFile(t)
	if err != nil {
		t.Fatal(err)
	}

	defer t.Cleanup(func() {
		if err = os.Remove(testFile); err != nil {
			t.Fatal(err)
		}
	})

	file, err := os.Open(testFile)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	filename := filepath.Base(testFile)

	checksum, err := bunnystorage.ComputeSHA256(file)
	if err != nil {
		t.Fatal(err)
	}

	if _, err = file.Seek(0, 0); err != nil {
		t.Fatal(err)
	}

	resp, err := client.Upload(ctx, _testPath, filename, checksum, file)
	if err != nil {
		t.Fatalf("upload error: %v", err)
	}

	if resp.Status != 201 {
		t.Errorf("expected status code to be 201, got %d", resp.Status)
	}

	body, resp, err := client.Download(ctx, _testPath, testFile)
	if err != nil {
		t.Fatalf("download error: %v", err)
	}

	if resp.Status < 200 || resp.Status >= 300 {
		t.Errorf("expected file to be downloaded, got status %d", resp.Status)
	}

	if len(body) != int(testFileSize) {
		t.Errorf("expected file size to be %d, got %d", testFileSize, len(body))
	}

	resp, err = client.Delete(ctx, _testPath, testFile)
	if err != nil {
		t.Fatalf("delete error: %v", err)
	}

	if resp.Status != 200 {
		t.Errorf("expected status code to be 204, got %d", resp.Status)
	}

	_, resp, err = client.Download(ctx, _testPath, testFile)
	if err != nil {
		t.Fatalf("download error: %v", err)
	}

	if resp.Status != 404 {
		t.Errorf("expected status code to be 404, got %d", resp.Status)
	}
}
