# MGLang

MGLang é um componente que detecta e retorna o código base do idioma do sistema operacional.

## Recursos

- GetLang: Retorna pt, en, es e assim por diante, de acordo com o idioma do sistema operacional.
- Set: Adicione o texto original e a tradução.
- T: Retorna a tradução

## Instalação

`go get github.com/mugomes/mglang`

## Exemplo

```
import "github.com/mugomes/mglang"

func main() {
	if mglang.GetLang() == "pt" {
		mglang.Set("File", "Arquivo")
	}
	
	print(mglang.T("File"))
}
```

## Information

 - [Page MGLang](https://github.com/mugomes/mglang)

## Requirement

 - Go 1.25.3

## Support

- GitHub: https://github.com/sponsors/mugomes
- More: https://www.mugomes.com.br/apoie.html

## License

Copyright (c) 2025 Murilo Gomes Julio

Licensed under the [MIT](https://github.com/mugomes/mglang/blob/main/LICENSE) license.

All contributions to the MGLang are subject to this license.
