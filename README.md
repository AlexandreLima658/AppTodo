# Aplicação Todo List

## Descrição
Este é um projeto de aplicativo Todo List desenvolvido em Go usando o framework GoFiber e PostgreSQL como banco de dados.

## Requisitos

Certifique-se de ter os seguintes requisitos instalados em seu sistema antes de executar o aplicativo:

    Go (versão 1.16 ou superior)
    PostgreSQL (versão 9.6 ou superior)

## Configuração
    git clone https://github.com/seu-usuario/todo-list.git
    cd todo-list
    go mod download

### Crie um banco de dados PostgreSQL em sua máquina.

    No diretório raiz do projeto, crie um arquivo chamado .env
    e defina as seguintes variáveis de ambiente:
    
    DB_HOST=localhost
    DB_PORT=5432
    DB_NAME=todo
    DB_USER=seu-usuario
    DB_PASSWORD=sua-senha

### Criar tabela no banco de dados

    CREATE TABLE todo (
        id SERIAL PRIMARY KEY,
        item VARCHAR(255) NOT NULL,
    );
    
### Executando o aplicativo

    go run main.go