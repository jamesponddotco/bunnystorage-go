package bunnystorage_test

import (
	"context"
	"net/http"
	"os"
	"reflect"
	"testing"

	"git.sr.ht/~jamesponddotco/bunnystorage-go"
	"git.sr.ht/~jamesponddotco/bunnystorage-go/internal/testutil"
)

const _testDataPath string = "tests/testdata"

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
	var (
		client        = testutil.SetupMockClient(t)
		mux, teardown = testutil.SetupMockServer(t)
		ctx           = context.Background()
	)

	defer t.Cleanup(func() {
		teardown()
	})

	tests := []struct {
		name    string
		handler http.HandlerFunc
		route   string
		path    string
		want    []*bunnystorage.Object
		wantErr bool
	}{
		{
			name: "valid_response",
			handler: func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodGet {
					t.Errorf("List() method = %v, want %v", r.Method, http.MethodGet)
				}

				w.Header().Set("Content-Type", "application/json")

				_, err := w.Write(testutil.ReadFile(t, _testDataPath+"/list-valid.json"))
				if err != nil {
					t.Fatalf("List() error = %v", err)
				}
			},
			route: "/mock/testdata/",
			path:  "/testdata",
			want: []*bunnystorage.Object{
				{
					UserID:          "737ccef2-ded9-4acf-9f2f-ce006c7a3bd1",
					Path:            "/bunnystorage-go/testdata/",
					ObjectName:      "5k6jh7.txt",
					LastChanged:     "2023-04-20T15:32:08.004",
					StorageZoneName: "bunnystorage-go",
					Checksum:        "3DE0A250A581BD2D18E0CEC41924AFAA691C633C09543AF456C6B7831E419A82",
					DateCreated:     "2023-04-20T15:32:08.004",
					GUID:            "fd3e5a08-1dd3-484a-839d-cdc66fd1fa6e",
					Length:          65,
					ServerID:        146,
					StorageZoneID:   996873,
				},
				{
					UserID:          "737ccef2-ded9-4acf-9f2f-ce006c7a3bd1",
					Path:            "/bunnystorage-go/testdata/",
					ObjectName:      "87dwda.txt",
					LastChanged:     "2023-04-20T15:32:08.535",
					StorageZoneName: "bunnystorage-go",
					Checksum:        "27307E0E4EBCA9D816582BBB6EEFB0FD6F892A4412398345408B7068AC236BDE",
					DateCreated:     "2023-04-20T15:32:08.535",
					GUID:            "6c5cf3ad-be04-4269-a0cb-43c05b155423",
					Length:          65,
					ServerID:        272,
					StorageZoneID:   996873,
				},
			},
			wantErr: false,
		},
		{
			name: "invalid_response",
			handler: func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodGet {
					t.Errorf("List() method = %v, want %v", r.Method, http.MethodGet)
				}

				w.Header().Set("Content-Type", "application/json")

				_, err := w.Write(testutil.ReadFile(t, _testDataPath+"/list-invalid.json"))
				if err != nil {
					t.Fatalf("List() error = %v", err)
				}
			},
			route:   "/mock/invalid/",
			path:    "/invalid",
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mux.HandleFunc(tt.route, tt.handler)

			got, _, err := client.List(ctx, tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Download(t *testing.T) {
	var (
		client        = testutil.SetupMockClient(t)
		mux, teardown = testutil.SetupMockServer(t)
		ctx           = context.Background()
	)

	defer t.Cleanup(func() {
		teardown()
	})

	tests := []struct {
		name     string
		handler  http.HandlerFunc
		route    string
		path     string
		filename string
		want     []byte
		wantCode int
		wantErr  bool
	}{
		{
			name: "valid_response",
			handler: func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodGet {
					t.Errorf("Download() method = %v, want %v", r.Method, http.MethodGet)
				}

				w.Header().Set("Content-Type", "image/jpeg")

				_, err := w.Write(testutil.ReadFile(t, _testDataPath+"/download-valid.jpg"))
				if err != nil {
					t.Fatalf("Download() error = %v", err)
				}
			},
			route:    "/mock/testdata/download-valid.jpg",
			path:     "/testdata",
			filename: "download-valid.jpg",
			want:     testutil.ReadFile(t, _testDataPath+"/download-valid.jpg"),
			wantCode: http.StatusOK,
			wantErr:  false,
		},
		{
			name: "not_found_response",
			handler: func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodGet {
					t.Errorf("Download() method = %v, want %v", r.Method, http.MethodGet)
				}

				w.WriteHeader(http.StatusNotFound)
			},
			route:    "/mock/testdata/download-not-found.json",
			path:     "/testdata",
			filename: "download-not-found.json",
			want:     nil,
			wantCode: http.StatusNotFound,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mux.HandleFunc(tt.route, tt.handler)

			got, resp, err := client.Download(ctx, tt.path, tt.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("Download() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(got) != len(tt.want) {
				t.Errorf("Download() got = %v, want %v", len(got), len(tt.want))
			}

			if resp.Status != tt.wantCode {
				t.Errorf("Download() got = %v, want %v", resp.Status, tt.wantCode)
			}
		})
	}
}

func TestClient_Upload(t *testing.T) {
	var (
		client        = testutil.SetupMockClient(t)
		mux, teardown = testutil.SetupMockServer(t)
		ctx           = context.Background()
	)

	defer t.Cleanup(func() {
		teardown()
	})

	tests := []struct {
		name     string
		handler  http.HandlerFunc
		route    string
		path     string
		filename string
		wantCode int
		wantErr  bool
	}{
		{
			name: "valid_response",
			handler: func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodPut {
					t.Errorf("Upload() method = %v, want %v", r.Method, http.MethodPut)
				}

				w.WriteHeader(http.StatusCreated)
			},
			route:    "/mock/testdata/upload-valid.txt",
			path:     "/testdata",
			filename: "upload-valid.txt",
			wantCode: http.StatusCreated,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mux.HandleFunc(tt.route, tt.handler)

			file, err := os.Open(_testDataPath + "/" + tt.filename)
			if err != nil {
				t.Fatalf("Upload() error = %v", err)
			}
			defer file.Close()

			checksum, err := bunnystorage.ComputeSHA256(file)
			if err != nil {
				t.Fatalf("ComputeSHA256() error = %v", err)
			}

			resp, err := client.Upload(ctx, tt.path, tt.filename, checksum, file)
			if (err != nil) != tt.wantErr {
				t.Errorf("Upload() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if resp.Status != tt.wantCode {
				t.Errorf("Upload() got = %v, want %v", resp.Status, tt.wantCode)
			}
		})
	}
}

func TestClient_Delete(t *testing.T) {
	var (
		client        = testutil.SetupMockClient(t)
		mux, teardown = testutil.SetupMockServer(t)
		ctx           = context.Background()
	)

	defer t.Cleanup(func() {
		teardown()
	})

	tests := []struct {
		name     string
		handler  http.HandlerFunc
		route    string
		path     string
		filename string
		wantCode int
		wantErr  bool
	}{
		{
			name: "valid_response",
			handler: func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodDelete {
					t.Errorf("Delete() method = %v, want %v", r.Method, http.MethodDelete)
				}

				w.WriteHeader(http.StatusNoContent)
			},
			route:    "/mock/testdata/delete-valid.txt",
			path:     "/testdata",
			filename: "delete-valid.txt",
			wantCode: http.StatusNoContent,
			wantErr:  false,
		},
		{
			name: "not_found_response",
			handler: func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodDelete {
					t.Errorf("Delete() method = %v, want %v", r.Method, http.MethodDelete)
				}

				w.WriteHeader(http.StatusNotFound)
			},
			route:    "/mock/testdata/delete-not-found.json",
			path:     "/testdata",
			filename: "delete-not-found.json",
			wantCode: http.StatusNotFound,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mux.HandleFunc(tt.route, tt.handler)

			resp, err := client.Delete(ctx, tt.path, tt.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if resp.Status != tt.wantCode {
				t.Errorf("Delete() got = %v, want %v", resp.Status, tt.wantCode)
			}
		})
	}
}
