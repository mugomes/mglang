// Copyright (C) 2025 Murilo Gomes Julio
// SPDX-License-Identifier: MIT

// Site: https://www.mugomes.com.br

package mglang

type Provider interface {
	GetSystemLanguage() string
}

var Current Provider
