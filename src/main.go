package main

import (
	"fmt"
	"math/big"
)

const (
	SHARED_BASE  int64 = 666
	SHARED_PRIME int64 = 6661
)

func main() {

	bobPk := big.NewInt(2227)
	// Assignment 1
	alicePk, messageCipher := encrypt(*big.NewInt(6), *bobPk, *big.NewInt(2000))
	fmt.Printf("Alice's public key is: %s, and the encrypted message is: %s\n", alicePk.Text(10), messageCipher.Text(10))

	// Assignment 2
	bobSecret, message := intercept(*bobPk, *alicePk, *messageCipher)
	fmt.Printf("Bob's secret is: %s, and the decrypted message is: %s\n", bobSecret.Text(10), message.Text(10))

	// Assignment 3
	modifiedCipher := messageCipher.Mul(big.NewInt(3), messageCipher);
	bobMessage := decrypt(*alicePk, bobSecret, *modifiedCipher)
	fmt.Printf("Message after Mallory has tampered with the ciphertext: %s\n", bobMessage.Text(10))
}

func calculateKey(base, prime, secret big.Int) *big.Int {

	result := big.NewInt(0);
	result.Exp(&base, &secret, nil);
	result.Mod(result, &prime);
	return result;
}

func encrypt(secretKey, publicKey, text big.Int) (key, cipher *big.Int) {
	fmt.Printf("Encrypting...\n")
	sharedBase := big.NewInt(SHARED_BASE)
	sharedPrime := big.NewInt(SHARED_PRIME)

	secret := calculateKey(*sharedBase, *sharedPrime, secretKey)
	commonKey := calculateKey(publicKey, *sharedPrime, secretKey)
	cipher = big.NewInt(0);
	cipher.Mul(commonKey, &text)

	return secret, cipher
}

func decrypt(pk, secret, cipher big.Int) big.Int {
	sharedKey := calculateKey(pk, *big.NewInt(SHARED_PRIME), secret);

	result := big.NewInt(0);
	return *result.Div(&cipher, sharedKey);
}

func intercept(targetPk, pk, cipher big.Int) (secret, message big.Int) {
	var limit big.Int = *big.NewInt(1000)
	incrementer := big.NewInt(1)
	var testSecret big.Int
	for testSecret = *big.NewInt(1); testSecret.Cmp(&limit) < 0; testSecret.Add(&testSecret, incrementer) {
		key := calculateKey(*big.NewInt(SHARED_BASE), *big.NewInt(SHARED_PRIME), testSecret)
		if key.Cmp(&targetPk) == 0 {
			
			message := decrypt(pk, testSecret, cipher)
			return testSecret, message
		}
	}
	return *big.NewInt(0), *big.NewInt(0)
}
