package ch13

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type remoteSolver struct {
	MathServerURL string
	Client        *http.Client
}

func (rs remoteSolver) Resolve(ctx context.Context, exp string) (float64, error) {
	//req, err := http.NewRequestWithContext(ctx, http.MethodGet, rs.MathServerURL+"?expression="+url.QueryEscape(expression), nil)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, rs.MathServerURL+"?exp="+url.QueryEscape(exp), nil)

	if err != nil {
		return 0, err
	}
	resp, err := rs.Client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	contents, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, errors.New(string(contents))
	}
	result, err := strconv.ParseFloat(string(contents), 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}
