package main

import (
	"fmt"

	"github.com/czas/gosip"
	strategy "github.com/czas/gosip/auth/ntlm"
)

func main() {
	configPath := "./config/private.onprem-ntlm.json"
	auth := &strategy.AuthCnfg{}

	err := auth.ReadConfig(configPath)
	if err != nil {
		fmt.Printf("unable to get config: %v\n", err)
		return
	}

	client := &gosip.SPClient{
		AuthCnfg: auth,
	}

	digest, err := gosip.GetDigest(client)
	if err != nil {
		fmt.Printf("unable to get digest: %v\n", err)
		return
	}

	fmt.Printf("Digest: %s\n", digest)

	cachedDigest, err := gosip.GetDigest(client)
	if err != nil {
		fmt.Printf("unable to get digest: %v\n", err)
		return
	}

	fmt.Printf("Cached digest: %s\n", cachedDigest)

}
