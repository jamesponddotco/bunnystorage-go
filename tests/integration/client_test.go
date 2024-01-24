package integration_test

import (
	"context"
	"flag"
	"log"
	"net/http"
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
	// Call flag.Parse explicitly to prevent testing.Short() from panicking.
	flag.Parse()

	if testing.Short() {
		os.Exit(0)
	}

	client, err = testutil.SetupClient()
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
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

	if resp.Status != http.StatusOK {
		t.Errorf("expected status code to be %d, got %d", http.StatusOK, resp.Status)
	}
}

func TestClient_Download(t *testing.T) {
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

		_, err = client.Delete(ctx, _testPath, testFile)
		if err != nil {
			t.Fatalf("delete error: %v", err)
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

	_, err = client.Upload(ctx, _testPath, filename, checksum, file)
	if err != nil {
		t.Fatalf("upload error: %v", err)
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
}

func TestClient_Upload(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	testFile, _, err := testutil.SetupFile(t)
	if err != nil {
		t.Fatal(err)
	}

	defer t.Cleanup(func() {
		if err = os.Remove(testFile); err != nil {
			t.Fatal(err)
		}

		_, err = client.Delete(ctx, _testPath, testFile)
		if err != nil {
			t.Fatalf("delete error: %v", err)
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
}

func TestClient_Delete(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	testFile, _, err := testutil.SetupFile(t)
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

	_, err = client.Upload(ctx, _testPath, filename, checksum, file)
	if err != nil {
		t.Fatalf("upload error: %v", err)
	}

	resp, err := client.Delete(ctx, _testPath, testFile)
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
