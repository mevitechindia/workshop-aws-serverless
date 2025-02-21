# workshop-aws-serverless
AWS serverless workshop

`serverless` folder contains the AWS SAM (serverless) deployment for an API that shows the source IP address

`beforeserverless` folder containers the source IP address API which is to be deployed in a traditional server setup. It contains go and python.

## SAM
To initialise a new projects
```
sam init
```

To build the application
```
sam build
```

To deploy for the first time (AWS profile named playground)
```
sam deploy --guided --profile playground
```

To deploy
```
sam deploy --config-env prod
```

To delete
```
sam delete --config-env prod
```

## Go
To initialise module
```
go mod init ipaddress-api
```

To set up dependencies
```
go mod tidy
```

To run directly
```
go run main.py
```

To build a binary
```
go build
```

## Python
To install deps
```
pip install flask
```

To run
```
python server.py
```