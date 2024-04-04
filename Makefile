

# .PHONE: release
# release:
# 	GOOS=linux GOARCH=amd64 go build -o dist/klc-linux-amd64
# 	GOOS=linux GOARCH=arm go build -o dist/klc-linux-arm
# 	GOOS=darwin GOARCH=amd64 go build -o dist/klc-darwin-amd64
# 	GOOS=freebsd GOARCH=amd64 go build -o dist/klc-freebsd-amd64
# 	GOOS=windows GOARCH=amd64 go build -o dist/klc-windows-amd64.exe

# .PHONY: release
# release: release-window release-linux release-drawin release-freebsd

# .PHONY: release-window
# release-window:
# 	GOOS=windows GOARCH=amd64 go build -o dist/klc-windows-amd64.exe
# 	GOOS=windows GOARCH=amd64 go build -o dist/klcc-windows-amd64.exe