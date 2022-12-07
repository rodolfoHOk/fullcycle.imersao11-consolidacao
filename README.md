# Imersão Fullstack & FullCycle 11 - Cartola FullCycle

> Microsserviço de Consolidação

## Tecnologias

- Go
- MySQL
- Kafka

## Guia

- Instalação do Go:

        curl -OL https://go.dev/dl/go1.19.3.linux-amd64.tar.gz
        sudo tar -C /usr/local -xvf go1.19.3.linux-amd64.tar.gz
        sudo nano ~/.profile
          export PATH=$PATH:/usr/local/go/bin
          CTRL+X / Y / ENTER
        source ~/.profile
        go version

- Extensão vs-code:

        instalar: golang.go
        CTRL+SHIFT+P
        Go: Install/Update Tools
        Selecionar todas e OK

- Instalação sqlc:

        curl -OL https://github.com/kyleconroy/sqlc/releases/download/v1.16.0/sqlc_1.16.0_linux_amd64.tar.gz
        sudo tar -C /usr/local/sqlc -xvf sqlc_1.16.0_linux_amd64.tar.gz
        sudo nano ~/.profile
          export PATH=$PATH:/usr/local/sqlc
          CTRL+X / Y / ENTER
        source ~/.profile
        sqlc version

- Geração das queries com sqlc:

        criar o sqlc.yaml
        sqlc generate

- Iniciando o projeto Go:

        go mod init github.com/rodolfoHOk/fullcycle.imersao11-consolidacao
        go mod tidy

- Baixar as dependências:

        go mod tidy

## Links dos repositórios da Imersão

- https://github.com/rodolfoHOk/fullcycle.imersao11-consolidacao
- https://github.com/rodolfoHOk/fullcycle.imersao11-django
