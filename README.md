# 🌍 Weebie

Bem-vindo ao Weebie - Sistema de Iniciativas Comunitárias 🏘️

Este sistema web permite a criação, gestão e compartilhamento de iniciativas de projetos com foco no desenvolvimento de comunidades locais. Nosso objetivo é promover práticas sustentáveis e facilitar a colaboração entre indivíduos e organizações comprometidas com um futuro mais colaborativo.

[![Go Version](https://img.shields.io/badge/Go-1.24%2B-blue?logo=go)](https://golang.org/)
[![Firebase](https://img.shields.io/badge/Firebase-Auth-orange?logo=firebase)](https://firebase.google.com/)
[![MongoDB](https://img.shields.io/badge/MongoDB-8.0%2B-green?logo=mongodb)](https://www.mongodb.com/)
[![License](https://img.shields.io/badge/License-MIT-yellow)](https://github.com/codinomello/weebie-go/blob/main/LICENSE)

## 🎨 Visual da Página Web

O Weebie foi projetado com um design moderno e responsivo, garantindo uma experiência de usuário agradável tanto em dispositivos desktop quanto móveis. Abaixo, você pode conferir o logo do projeto e algumas capturas de tela que ilustram o visual da página web.

### Logo do Weebie
<img src="https://raw.githubusercontent.com/codinomello/weebie-go/main/images/icons/house_with_garden.png" width="50" height="50"/>

### Capturas de Tela

<!--#### Visual em Desktop
<img src="https://raw.githubusercontent.com/codinomello/weebie-go/main/repo/desktop.png" width="800" height="450"/>-->

#### Visual em Celular
<img src="https://raw.githubusercontent.com/codinomello/weebie-go/main/images/repo/phone.png" width="435" height="800"/>


## 🛠️ Tecnologias Utilizadas

- **Back-end**: [Go](https://golang.org/) 🐹
- **Front-end**: [Templ](https://github.com/arschles/templ) 🖼️
- **Banco de Dados**: [MongoDB](https://www.mongodb.com/) 🍃
- **Autenticação**: [Firebase](https://firebase.google.com/products/auth) 🔐
- **Interatividade**: [HTMX](https://htmx.org/) ⚡
- **UI e UX**: [Flowbite](https://flowbite.com/) 🎨 + [Tailwind CSS](https://tailwindcss.com/) 🖌

## 🚀 Funcionalidades Principais

- **🌱 Criação de Iniciativas**:  
  Cadastre novas iniciativas sustentáveis com detalhes como título, descrição, objetivos e impacto esperado.

- **📊 Gestão de Projetos**:  
  Acompanhe o progresso das iniciativas, atribua tarefas e defina prazos de forma eficiente.

- **🤝 Colaboração**:  
  Conecte-se com outros usuários e organizações para colaborar em projetos sustentáveis.

- **📈 Relatórios e Análises**:  
  Gere relatórios sobre o impacto das iniciativas e visualize métricas de sucesso.

## 🛠️ Como Executar o Projeto

### 📋 Pré-requisitos

- **Go** 1.24 ou superior 🐹
- **MongoDB (Atlas)** 8.0 ou superior 🍃

## 📩 Instalação

1. **Clone o repositório**:
   ```bash
   git clone https://github.com/codinomello/weebie-go.git
   cd weebie-go

2. **Adicione o .env + firebase.json**:
   ```bash
   touch .env

3. **Configure as variáveis de ambiente**:
   ```bash
   # Porta do servidor HTTP
   PORT=...

   # URI do MongoDB
   MONGODB_URI=...

   # Base de dados MongoDB
   MONGODB_DATABASE=...

   # Configuração do Firebase
   FIREBASE_CREDENTIALS=...

4. **Instale o templ**:
   ```go
   go install github.com/a-h/templ/cmd/templ@latest

5. **Execute o projeto**:
   ```bash
   make
