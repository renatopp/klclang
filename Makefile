default:
	clear && go build -o dist/main.exe && ./dist/main.exe

repr:
	clear && go build -o dist/repr.exe ./cmd/klcr/klcr.go && ./dist/repr.exe
	
build:
	clear && go build -o dist/klc.exe ./cmd/klc/klc.go
	clear && go build -o dist/repr.exe ./cmd/klcr/klcr.go
	