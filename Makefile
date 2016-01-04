
.clean:
	rm -f coverage.out

.test: .clean
	go test -coverprofile=coverage.out ./
