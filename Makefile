# Sistema operacional do binário
OS=WINDOWS

# Nome do binário gerado
BINARY_NAME=projeto.exe

# Nome do arquivo
FILE_NAME=main.go

# Diretório do código fonte
SOURCE_DIRECTORY=cmd

# Comando para compilar o programa
build:
	templ generate && cd $(SOURCE_DIRECTORY) && go build -o $(BINARY_NAME) $(FILE_NAME)

# Comando para rodar o programa
run:
	templ generate && cd $(SOURCE_DIRECTORY) && go run $(FILE_NAME)

# Comando para limpar o binário gerado
clean:
	del $(BINARY_NAME)

templ:
	templ generate

# Comando padrão (executado ao rodar apenas 'make')
default: build