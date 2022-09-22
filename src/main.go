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

	bobSecret, message := intercept(bobPk, messageCipher);
	fmt.Printf("Bob's secret is: %.0f, and the decrypted message is: %.0f", bobSecret, message);

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

func decrypt(commonKey, cipher float64) float64 {
	message := cipher / commonKey;
	return message;
}

func intercept(publicKey, cipher float64) (secret, message float64) {

	var testSecret float64;
	for testSecret = 1; testSecret < 1000; testSecret++ {
		
		if key := calculateKey(SHARED_BASE, SHARED_PRIME, testSecret); key == publicKey {

			message := decrypt(key, cipher);
			return testSecret, message;
		}
	}	
	return 0, 0;
}