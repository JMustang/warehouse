# Define variáveis para caminhos e nomes de arquivos
BINARY_NAME=warehouse
BINARY_PATH=bin/$(BINARY_NAME)
MAIN_FILE=cmd/main.go

# Verifica e cria o diretório bin se necessário, então constrói o binário
build:
	@mkdir -p bin
	@go build -o $(BINARY_PATH) $(MAIN_FILE)

# Executa testes com verbose
test:
	@go test -v ./...

# Remove binário antigo e executa a construção e execução
run: clean build
	@$(BINARY_PATH)

# Remove o binário criado
clean:
	@rm -f $(BINARY_PATH)
