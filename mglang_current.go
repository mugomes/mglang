// Copyright (C) 2025-2026 Murilo Gomes Julio
// SPDX-License-Identifier: MIT

// Site: https://mugomes.github.io

package mglang

type Provider interface {
	GetSystemLanguage() string
}

var Current Provider
