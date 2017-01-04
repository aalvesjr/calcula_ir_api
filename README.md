# API to Calculate IR (Imposto de Renda - Brazilian Tax)
[![CircleCI](https://circleci.com/gh/aalvesjr/calcula_ir_api.svg?style=svg)](https://circleci.com/gh/aalvesjr/calcula_ir_api)

Urls available to see working: 
[heroku/?salary=5122.32](https://calcula-ir.herokuapp.com/api/calculate-ir?salary=5122.32)
[heroku/?salary=5122.32&discounts=156.72](https://calcula-ir.herokuapp.com/api/calculate-ir?salary=5122.32&discounts=156.72)
[heroku/?discounts=156.72](https://calcula-ir.herokuapp.com/api/calculate-ir?discounts=156.72)
[heroku/salary/5122.32](https://calcula-ir.herokuapp.com/api/salary/5122.32)
[heroku/salary/5122.32/discounts/156.72](https://calcula-ir.herokuapp.com/api/salary/5122.32/discounts/156.72)
[heroku/discounts/156.72](https://calcula-ir.herokuapp.com/api/discounts/156.72)

Project to test:
  - [Gorilla Mux](https://github.com/gorilla/mux)
  - [Salary](https://github.com/aalvesjr/salary)

## Setup

```
mkdir -p $GOPATH/src/github.com/aalvesjr
cd $GOPATH/src/github.com/aalvesjr
git clone https://github.com/aalvesjr/calcula_ir_api.git
cd calcula_ir_api

glide install
go build
```

## Testing

To avoid run test from vendor folder, run locally:

```
go test -v $(go list ./... | grep -v '/vendor/')
```

## Running

```
export PORT=3000

./calcula_ir_api

```

API endpoints available:

```
curl -X GET -H "Content-Type: application/json" "http://localhost:3000/api/calculate-ir?salary=5138.79"
curl -X GET -H "Content-Type: application/json" "http://localhost:3000/api/calculate-ir?salary=5138.79&discounts=243.12"
curl -X GET -H "Content-Type: application/json" "http://localhost:3000/api/calculate-ir?discounts=243.12"

curl -X GET -H "Content-Type: application/json" "http://localhost:3000/api/salary/5138.95"
curl -X GET -H "Content-Type: application/json" "http://localhost:3000/api/salary/5138.95/discounts/243.12"
curl -X GET -H "Content-Type: application/json" "http://localhost:3000/api/discounts/243.12"
```

## Contributing
- Fork it
- Create your feature branch (`git checkout -b my-new-feature`)
- Commit your changes (`git commit -am 'Add some feature'`)
- Push to the branch (`git push origin my-new-feature`)
- Create new Pull Request
