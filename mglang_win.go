// Required Notice: Copyright (c) 2025-2026 Murilo Gomes Julio. All Rights Reserved. (https://profmugomes.com.br)

// Licensed under the PolyForm Perimeter License 1.0.1.
// See LICENSE.md for details.

//go:build windows

package mglang

import (
	"golang.org/x/sys/windows/registry"
	"strings"
)

type winProvider struct{}

func init() {
	Current = winProvider{}
}

func (winProvider) GetSystemLanguage() string {
	k, _, err := registry.CreateKey(
		registry.CURRENT_USER,
		`Control Panel\International`,
		registry.QUERY_VALUE,
	)
	if err != nil {
		return "en"
	}
	defer k.Close()

	s, _, err := k.GetStringValue("LocaleName")
	if err != nil || len(s) < 2 {
		return "en"
	}

	return strings.ToLower(s[:2])
}
