mocks:
	cd ./store/mocks/; go generate;
	cd ./service/mocks/; go generate;

build:
	go build -o ./build/doxod-api cmd/api/main.go

run:
	bash ./rundev.sh