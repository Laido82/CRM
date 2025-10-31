run:
	cd ../../cmd && air
build:
	cd ../../cmd && go build -o app app.go
build-windows64:
	cd ../../cmd && GOOS=windows GOARCH=amd64 go build -o app.exe app.go
build-windows32:
	cd ../../cmd && GOOS=windows GOARCH=386 go build -o app.exe app.go
