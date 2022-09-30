echo "Finally something worthy to download"
cd ../ && go mod tidy && go mod vendor &&  gofmt ./ && golangci-lint run && cd deploy ;

echo "Checking dependencies and formatting..."
go build -o bin/ ./../cmd/datastation/;

echo "Default Config Copying to build directory..."
cp env bin/.env
echo "Build complete!";