package main

import (
	"fmt"
	"math"
)
const (
	SHARED_BASE float64 = 666;
	SHARED_PRIME float64 = 6661;
)

func main() {

	var bobPk float64 = 2227;

	aliceKey, messageCipher := encrypt(8, bobPk, 2000);
	fmt.Printf("Alice's key is: %.0f, and the encrypted message is: %.0f\n", aliceKey, messageCipher);

}

func calculateKey(base, prime, secret float64) float64 {

	return math.Mod(math.Pow(base, secret), prime);
}

func encrypt(secretKey, publicKey, text float64) (key, cipher float64) {
	fmt.Printf("Encrypting...\n");

	secret := calculateKey(SHARED_BASE, SHARED_PRIME, secretKey);
	commonKey := calculateKey(publicKey, SHARED_PRIME, secretKey); 
	cipherText := text * commonKey;

	return secret, cipherText;
}

func intercept(publicKey, cipher float64) (secret, message float64) {

	var testSecret float64;
	for testSecret = 1; testSecret < 1000; testSecret++ {
		key := calculateKey(SHARED_BASE, SHARED_PRIME, testSecret);		
		if key == publicKey {
			return testSecret, 0;
		}
	}	
}