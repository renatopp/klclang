default:
	clear && go build -o dist/main.exe && ./dist/main.exe

repr:
	clear && go build -o dist/repr.exe ./cmd/klcr/klcr.go && ./dist/repr.exe