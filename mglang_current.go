// Required Notice: Copyright (c) 2025-2026 Murilo Gomes Julio. All Rights Reserved. (https://profmugomes.com.br)

// Licensed under the PolyForm Perimeter License 1.0.1.
// See LICENSE.md for details.

package mglang

type Provider interface {
	GetSystemLanguage() string
}

var Current Provider
