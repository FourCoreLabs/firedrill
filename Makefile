all: ransomware discovery uac_bypass

mockransomware:
	go build github.com/FourCoreLabs/firedrill/cmd/mockransomware

ransomware:
	go build github.com/FourCoreLabs/firedrill/cmd/ransomware

discovery:
	go build github.com/FourCoreLabs/firedrill/cmd/discovery

uac_bypass:
	go build github.com/FourCoreLabs/firedrill/cmd/uac_bypass

registry_run:
	go build github.com/FourCoreLabs/firedrill/cmd/runkeyregistry
	
gorelease:
	goreleaser release --rm-dist --snapshot 
