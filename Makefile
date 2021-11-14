ransomware:
	go build github.com/FourCoreLabs/firedrill/cmd/ransomware

gorelease:
	goreleaser release --rm-dist --snapshot 
