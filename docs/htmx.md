# HTMX ⚡

## Introdução
HTMX é uma biblioteca JavaScript que permite adicionar interatividade às aplicações web utilizando atributos HTML, eliminando a necessidade de escrever muito código JavaScript.

## Instalação
- Visite o [site oficial do HTMX](https://htmx.org/) para mais informações.
- Inclua a biblioteca em seu projeto adicionando o seguinte script ao seu HTML:
```html
<script src="https://unpkg.com/htmx.org@1.8.4"></script>

<!-- Botão que carrega conteúdo dinamicamente -->
<button hx-get="/rota" hx-target="#conteudo">
  Carregar Conteúdo
</button>
<div id="conteudo"></div>