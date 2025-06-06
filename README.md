# DevTask

**DevTask** é um gerenciador de tarefas pessoais para desenvolvedores, totalmente em linha de comando (CLI), desenvolvido em Go. Ele permite adicionar, listar, concluir e excluir tarefas de forma rápida e organizada, usando armazenamento local com SQLite.

---

## Funcionalidades

- Adicionar tarefas com título e descrição  
- Listar tarefas pendentes, concluídas ou todas  
- Marcar tarefas como concluídas  
- Excluir tarefas  
- Gerar relatório diário de progresso  
- Armazenamento local SQLite com migração automática  
- Interface simples e direta no terminal  

---

## Tecnologias

- Linguagem: Go  
- CLI: [Cobra](https://github.com/spf13/cobra)  
- Banco de dados: SQLite via [modernc.org/sqlite]
- Arquitetura inspirada em Clean Architecture e CQRS  

---

## Instalação

### Requisitos

- Go 1.18+ instalado  
- SQLite instalado no sistema (geralmente já vem por padrão)  

### Passos

```bash
# Clonar o repositório
git clone https://github.com/gustavomelo20/devtask.git
cd devtask

# Baixar dependências
go mod download

# Buildar o binário
go build -o devtaskJ

# Rodar o programa
./devtask --help
