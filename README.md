## Cash

### Introduction

Dead simple CLI for currency conversion. Because Googling takes too long.

Rates come from [currency-api](https://github.com/fawazahmed0/exchange-api). A list of supported currencies can be seen here: [supported currencies](https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies.json).

### How to use

`$ cash [amount] [from] [to]`

Example:

`$ cash 10 usd brl`

`$ BRL $57.89`

### Installing

Download the pre-built binary. Boom.

### Build from source

To build this from source you will need Go 1.24 or newer.

- Clone this repository
- `cd cash`
- `go build`
- Done

### License

[The Unlicense](https://unlicense.org/)
