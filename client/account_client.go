package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"io"
	"net/http"
	"net/url"
)

// CreateAccountPayload is the account create action payload.
type CreateAccountPayload struct {
	// Name of account
	Name string `json:"name" xml:"name"`
}

// CreateAccountPath computes a request path to the create action of account.
func CreateAccountPath() string {
	return fmt.Sprintf("/cellar/accounts")
}

// Create new account
func (c *Client) CreateAccount(ctx context.Context, path string, payload *CreateAccountPayload) (*http.Response, error) {
	var body io.Reader
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize body: %s", err)
	}
	body = bytes.NewBuffer(b)
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	c.SignerAdminPass.Sign(ctx, req)
	return c.Client.Do(ctx, req)
}

// DeleteAccountPath computes a request path to the delete action of account.
func DeleteAccountPath(accountID int) string {
	return fmt.Sprintf("/cellar/accounts/%v", accountID)
}

// DeleteAccount makes a request to the delete action endpoint of the account resource
func (c *Client) DeleteAccount(ctx context.Context, path string) (*http.Response, error) {
	var body io.Reader
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return c.Client.Do(ctx, req)
}

// ShowAccountPath computes a request path to the show action of account.
func ShowAccountPath(accountID int) string {
	return fmt.Sprintf("/cellar/accounts/%v", accountID)
}

// Retrieve account with given id. IDs 1 and 2 pre-exist in the system.
func (c *Client) ShowAccount(ctx context.Context, path string) (*http.Response, error) {
	var body io.Reader
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return c.Client.Do(ctx, req)
}

// UpdateAccountPayload is the account update action payload.
type UpdateAccountPayload struct {
	// Name of account
	Name string `json:"name" xml:"name"`
}

// UpdateAccountPath computes a request path to the update action of account.
func UpdateAccountPath(accountID int) string {
	return fmt.Sprintf("/cellar/accounts/%v", accountID)
}

// Change account name
func (c *Client) UpdateAccount(ctx context.Context, path string, payload *UpdateAccountPayload) (*http.Response, error) {
	var body io.Reader
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize body: %s", err)
	}
	body = bytes.NewBuffer(b)
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PUT", u.String(), body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return c.Client.Do(ctx, req)
}