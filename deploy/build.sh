echo "analizing this golang project"
cd ../ && go mod tidy && go mod vendor && gofmt ./ && golangci-lint run && cd deploy;
