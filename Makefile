ransomware:
	go build github.com/FourCoreLabs/firedrill/cmd/ransomware

discovery:
	go build github.com/FourCoreLabs/firedrill/cmd/discovery

gorelease:
	goreleaser release --rm-dist --snapshot 
