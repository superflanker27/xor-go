package main

import (
	"fmt"
	"os"
)

// call when we need to encrypt/decrypt
func pseudoProtect(dhana, cilc []byte) []byte {
	nxjerr := make([]byte, len(dhana))
	for i := 0; i < len(dhana); i++ {
		nxjerr[i] = dhana[i] ^ cilc[i%len(cilc)]
	}
	return nxjerr
}

func main() {
	secret := []byte("i hate niggers, not just niggers. i also hate faggots !")
	cilc := []byte("topsecretciasponsoredkey")

	//mbylleMorTiMbylle(secret, cilc)
	hapeMorTiHape(secret, cilc)
}

func mbylleMorTiMbylle(secret, cilc []byte) {
	// enc
	skap := pseudoProtect(secret, cilc)

	// save
	err := os.WriteFile("kryvi.txt", skap, 0644)
	if err != nil {
		fmt.Println("save deshtoj...")
	}
}

func hapeMorTiHape(secret, cilc []byte) {
	// dec
	tdhanatEmbyllura, err := os.ReadFile("kryvi.txt")
	if err != nil {
		fmt.Println("leximi deshtoj...")
	}

	hapeTdhanat := pseudoProtect(tdhanatEmbyllura, cilc)

	// write decrypted data to file

	err = os.WriteFile("kryvi_dec.txt", hapeTdhanat, 0644)
	if err != nil {
		fmt.Println("save deshtoj...")
	}

}
