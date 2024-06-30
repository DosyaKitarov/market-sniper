# market-sniper
Golang part of backend for Market Sniper

## Installation
1. Clone the repository
2. Run `go get` to install dependencies
3. Generate a self-signed certificate and key and place them in the `tls` directory by `go run path/to/generate_cert.go --host=localhost`
4. Run `go run ./cmd` to start the server
5. The server will be running on `https://localhost:4000`


## Endpoints

### GET /api/v1/getInfo/{ASIN}
Returns the info about product

curl --request GET \
  --url http://{host}:{gateway_port}/v1/getInfo/{ASIN} \
  --header 'Content-Type: application/json' \
  --header 'X-Api-Key: your_api_key' \
  --data '{
    "country": "US",
    "tld": "com"
  }'

RESPONSE
```json
{
    "ASIN": "B07VGRJDFY",
    "Title": "Apple iPhone 11 Pro Max, 256GB, Midnight Green, Fully Unlocked (Renewed)",
    "Brand": "Apple",
    "Price": "999.99 $",
    "Previous_price": "1099.99 $",
    "Change_date": "2021-07-01T00:00:00Z"
}
```

