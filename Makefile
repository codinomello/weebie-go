# Sistema operacional do binário
OS=LINUX

# Nome do binário gerado
BINARY_NAME=projeto

# Nome do arquivo
FILE_NAME=main.go

# Diretório do código fonte
SOURCE_DIRECTORY=cmd

# Comando para rodar o programa
run:
	templ generate && cd $(SOURCE_DIRECTORY) && go run $(FILE_NAME)

# Comando para compilar o programa
build:
	templ generate && cd $(SOURCE_DIRECTORY) && go build -o $(BINARY_NAME) $(FILE_NAME)

# Comando para limpar o binário gerado
clean:
	del $(BINARY_NAME)

# Comando para compilar os componentes
templ:
	templ generate

# Comando para compilar TypeScript para JavaScript
ts:
	tsc

# Comando padrão (executado ao rodar apenas 'make')
default: run