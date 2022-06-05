# Experiment: Connect on Lambda

> A simple echo server using [Connect](https://connect.build/) running on [AWS Lambda](https://aws.amazon.com/lambda/).

## Run the Server Locally

```
go run ./cmd/connect-server
```

## Testing Locally

```
go run ./cmd/echo-client "Hello Lambda"
go run ./cmd/echo-client --grpc "Hello Lambda"
go run ./cmd/echo-client --grpc-web "Hello Lambda"
go run ./cmd/echo-client --protojson "Hello Lambda"

curl \
  --header "Content-Type: application/json" \
  --data '{"message": "Hello"}' \
  http://localhost:3000/htdvisser.echo.v1.EchoService/Echo
```

## Build and Deploy the Lambda

```
make deps build deploy
```

## Testing the Lambda

```
base_url=$(terraform -chdir=deploy output -raw lambda_function_url)

go run ./cmd/echo-client --base-url $base_url "Hello Lambda"
go run ./cmd/echo-client --base-url $base_url --grpc "Hello Lambda" # This doesn't work.
go run ./cmd/echo-client --base-url $base_url --grpc-web "Hello Lambda"
go run ./cmd/echo-client --base-url $base_url --protojson "Hello Lambda"

curl \
  --header "Content-Type: application/json" \
  --data '{"message": "Hello"}' \
  ${base_url%"/"}/htdvisser.echo.v1.EchoService/Echo
```
