package bunnystorage

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"git.sr.ht/~jamesponddotco/httpx-go"
	"git.sr.ht/~jamesponddotco/xstd-go/xerrors"
	"git.sr.ht/~jamesponddotco/xstd-go/xstrings"
	"golang.org/x/time/rate"
)

const (
	// ErrConfigRequired is returned when a Client is created without a Config.
	ErrConfigRequired xerrors.Error = "config is required"
)

type (
	// Client is the LanguageTool API client.
	Client struct {
		// httpc is the underlying HTTP client used by the API client.
		httpc *httpx.Client

		// cfg specifies the configuration used by the API client.
		cfg *Config
	}
)

// NewClient returns a new bunny.net Edge Storage API client.
func NewClient(cfg *Config) (*Client, error) {
	if cfg == nil {
		return nil, ErrConfigRequired
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	cfg.init()

	return &Client{
		httpc: &httpx.Client{
			RateLimiter: rate.NewLimiter(rate.Limit(2), 1),
			RetryPolicy: httpx.DefaultRetryPolicy(),
			UserAgent:   cfg.Application.UserAgent(),
		},
		cfg: cfg,
	}, nil
}

// do performs an HTTP request using the underlying HTTP client.
func (c *Client) do(ctx context.Context, req *httpx.Request) (*Response, error) {
	ret, err := c.httpc.Do(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	defer func() {
		if err = httpx.DrainResponseBody(ret); err != nil {
			log.Fatal(err)
		}
	}()

	body, err := io.ReadAll(ret.Body)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	response := &Response{
		Header: ret.Header.Clone(),
		Body:   body,
		Status: ret.StatusCode,
	}

	return response, nil
}

// List lists the files in the storage zone.
func (c *Client) List(ctx context.Context, path string) ([]*Object, *Response, error) {
	path = strings.TrimPrefix(path, "/")

	uri := xstrings.JoinWithSeparator("/", c.cfg.Endpoint.String(), c.cfg.StorageZone, path+"/")

	headers := map[string]string{
		"Accept":    "application/json",
		"AccessKey": c.cfg.AccessKey(OperationRead),
	}

	req, err := httpx.NewRequest(ctx, http.MethodGet, uri, headers, http.NoBody)
	if err != nil {
		return nil, nil, fmt.Errorf("%w", err)
	}

	resp, err := c.do(ctx, req)
	if err != nil {
		return nil, nil, fmt.Errorf("%w", err)
	}

	var files []*Object
	if err := json.Unmarshal(resp.Body, &files); err != nil {
		return nil, nil, fmt.Errorf("%w", err)
	}

	return files, resp, nil
}

// Download downloads a file from the storage zone.
func (c *Client) Download(ctx context.Context, path, filename string) ([]byte, *Response, error) {
	path = strings.TrimPrefix(path, "/")

	uri := xstrings.JoinWithSeparator("/", c.cfg.Endpoint.String(), c.cfg.StorageZone, path, filename)

	headers := map[string]string{
		"Accept":    "*/*",
		"AccessKey": c.cfg.AccessKey(OperationRead),
	}

	req, err := httpx.NewRequest(ctx, http.MethodGet, uri, headers, http.NoBody)
	if err != nil {
		return nil, nil, fmt.Errorf("%w", err)
	}

	resp, err := c.do(ctx, req)
	if err != nil {
		return nil, nil, fmt.Errorf("%w", err)
	}

	return resp.Body, resp, nil
}

// Upload uploads a file to the storage zone.
func (c *Client) Upload(ctx context.Context, path, file string) (*Response, error) {
	path = strings.TrimPrefix(path, "/")

	fileData, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	var (
		filename = filepath.Base(file)
		uri      = xstrings.JoinWithSeparator("/", c.cfg.Endpoint.String(), c.cfg.StorageZone, path, filename)
	)

	headers := map[string]string{
		"Content-Type": "application/octet-stream",
		"AccessKey":    c.cfg.AccessKey(OperationWrite),
	}

	req, err := httpx.NewRequest(ctx, http.MethodPut, uri, headers, bytes.NewReader(fileData))
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	req.Req.ContentLength = int64(len(fileData))

	resp, err := c.do(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return resp, nil
}

// Delete deletes a file from the storage zone.
func (c *Client) Delete(ctx context.Context, path, filename string) (*Response, error) {
	path = strings.TrimPrefix(path, "/")

	uri := xstrings.JoinWithSeparator("/", c.cfg.Endpoint.String(), c.cfg.StorageZone, path, filename)

	headers := map[string]string{
		"AccessKey": c.cfg.AccessKey(OperationWrite),
	}

	req, err := httpx.NewRequest(ctx, http.MethodDelete, uri, headers, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	resp, err := c.do(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return resp, nil
}
