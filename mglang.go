// Required Notice: Copyright (c) 2025-2026 Murilo Gomes Julio. All Rights Reserved. (https://profmugomes.com.br)

// Licensed under the PolyForm Perimeter License 1.0.1.
// See LICENSE.md for details.

package mglang

import (
	"fmt"
	"runtime"
)

type Translations map[string]string

var tr = make(Translations)

// Detecta o idioma do sistema e retorna apenas o código base (pt, en, es, etc.)
func GetLang() string {
	platform := runtime.GOOS
	var lang string
	if platform == "linux" {
		lang = Current.GetSystemLanguage()
	} else {
		lang = Current.GetSystemLanguage()
	}

	return lang
}

func Set(name string, value string) {
	tr[name] = value
}

// T retorna o texto traduzido com formatação opcional.
func T(key string, args ...any) string {
	msg, ok := tr[key]
	if !ok {
		msg = key // fallback se não achar
	}
	return fmt.Sprintf(msg, args...)
}
