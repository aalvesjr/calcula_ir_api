# API para Calculo do Imposto de Renda (IR)

Url disponível para testar: [calcula-ir](https://calcula-ir.herokuapp.com/api/calcula-ir?salario=5122.32&descontos=156.72)

Projeto para testar:
  - Gorilla [Mux](https://github.com/gorilla/mux)
  - struct [Salario](https://github.com/aalvesjr/salario)

## Setup

```
mkdir -p $GOPATH/src/github.com/aalvesjr
cd $GOPATH/src/github.com/aalvesjr
git clone https://github.com/aalvesjr/calcula_ir_api.git
cd calcula_ir_api

glide install
go build
```

## Running

```
export PORT=3000

./calcula_ir_api

```

endpoints disponiveis para api:

```
curl -X GET -H "Content-Type: application/json" "http://localhost:3000/api/calcula-ir/salario/5138.79/descontos/243.12"
curl -X GET -H "Content-Type: application/json" "http://localhost:3000/api/salario/5138.95"
curl -X GET -H "Content-Type: application/json" "http://localhost:3000/api/salario/5138.95/descontos/243.12"
```
