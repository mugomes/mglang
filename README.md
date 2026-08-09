# MGLang

[![License](https://img.shields.io/badge/license-PolyForm%20Perimeter%201.0.1-5351FB)](LICENSE.md)

Uma biblioteca simples e leve em Go para **gerenciamento de traduções (i18n)**, com suporte a **detecção automática do idioma do sistema** e **formatação dinâmica de strings**.

Ideal para aplicações desktop, CLI ou bibliotecas que precisam de internacionalização sem dependências pesadas.

---

## ✨ Recursos

* 🌍 Detecção automática do idioma do sistema
* 🗂️ Registro simples de traduções por chave
* 🔁 Fallback automático quando a tradução não existe
* 🧩 Suporte a para textos dinâmicos
* 🚀 API minimalista e fácil de integrar

---

## 📦 Instalação

```bash
go get github.com/profmugomes/mglang
```

---

## 🚀 Uso básico

### Definindo traduções

```go
import "github.com/profmugomes/mglang"

func main() {
	mglang.Set("hello", "Olá")
	mglang.Set("Welcome %s", "Bem-vindo, %s!")
}
```

---

### Obtendo tradução simples

```go
text := mglang.T("hello")
// Resultado: Olá
```

---

### Tradução com parâmetros

```go
text := mglang.T("Welcome %s", "Murilo")
// Resultado: Bem-vindo, Murilo!
```

---

### Fallback automático

```go
text := mglang.T("not_exists")
// Resultado: not_exists
```

---

## 🌐 Detecção de idioma

```go
lang := mglang.GetLang()
fmt.Println(lang) // ex: pt, en, es
```

A função detecta o idioma do sistema operacional e retorna apenas o **código base** (`pt`, `en`, `es`, etc).

---

## 🧠 Estrutura interna

* `Set(key, value)` → registra tradução
* `T(key, args...)` → retorna texto traduzido
* `GetLang()` → detecta idioma do sistema

---

## 🧩 Compatibilidade

* Go 1.26.5+

---

## 👤 Autor

**Murilo Gomes Julio**

🔗 [https://www.profmugomes.com.br](https://www.profmugomes.com.br)

📺 [https://youtube.com/@profmugomes](https://youtube.com/@profmugomes)

---

## License

Copyright (c) 2025-2026 Murilo Gomes Julio. All Rights Reserved.

This project is licensed under the PolyForm Perimeter License 1.0.1.

### Summary

This software is available for commercial and noncommercial use, subject to the terms of the PolyForm Perimeter License 1.0.1.

You may:

* ✔ Use the software for commercial and noncommercial purposes.
* ✔ Inspect and study the source code.
* ✔ Modify the software.
* ✔ Create derivative works based on the software.
* ✔ Redistribute the software and permitted modifications.

You may not:

* ✖ Provide a product that competes with the software.

See the full license terms at LICENSE.md.

This summary is provided for convenience only and does not replace or modify the full license terms.
