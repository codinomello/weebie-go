
---

**Arquivo: `templ.md`**
```md
# Documentação do Front-end com Templ 🖼️

## Introdução
Templ é uma ferramenta para criação de templates que facilita a renderização de interfaces de usuário. Sua integração com Go torna o desenvolvimento front-end mais simples e organizado.

## Instalação
- Consulte o [repositório oficial](https://github.com/arschles/templ) para instruções de instalação e configuração.
- Pode ser utilizado com gerenciadores de pacotes como Go Modules.

## Conceitos Básicos
- **Templates Reutilizáveis:** Permite a criação de componentes e layouts consistentes.
- **Integração com Go:** Facilita a incorporação de lógica de servidor com a renderização de templates.

## Exemplo de Uso
```go
// Exemplo básico de template utilizando Templ
package main

import (
    "github.com/arschles/templ"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    templ.Render(w, "index.html", map[string]string{
        "Titulo": "Olá, Templ!",
    })
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
