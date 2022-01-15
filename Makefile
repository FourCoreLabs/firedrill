all: ransomware discovery uac_bypass

ransomware:
	go build github.com/FourCoreLabs/firedrill/cmd/ransomware

discovery:
	go build github.com/FourCoreLabs/firedrill/cmd/discovery

uac_bypass:
	go build github.com/FourCoreLabs/firedrill/cmd/uac_bypass

gorelease:
	goreleaser release --rm-dist --snapshot 
