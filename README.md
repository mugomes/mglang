# MGLang

<!-- Badge opcional para deixar claro visualmente -->
![Status](https://img.shields.io/badge/status-arquivado--archived-red.svg)

> **Aviso importante:** Esta biblioteca foi **oficialmente arquivada** e não receberá mais atualizações, correções de bugs ou suporte para novas versões.

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
go get github.com/mugomes/mglang
```

---

## 🚀 Uso básico

### Definindo traduções

```go
import "github.com/mugomes/mglang"

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

## 👤 Autor

**Murilo Gomes Julio**

🔗 [https://mugomes.github.io](https://mugomes.github.io)

📺 [https://youtube.com/@mugomesoficial](https://youtube.com/@mugomesoficial)

---

## License

Copyright (c) 2025-2026 Murilo Gomes Julio

Licensed under the [MIT](https://github.com/mugomes/mglang/blob/main/LICENSE) license.

All contributions to the MGLang are subject to this license.
