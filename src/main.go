package main

import (
	"flag"
	"fmt"
	"math/big"
)

const (
	SHARED_BASE  int64 = 666
	SHARED_PRIME int64 = 6661
)

func main() {
	secret := flag.Int64("n", 99, "The secret to be used in exercise. It only accepts integers.")
	flag.Parse()

	bobPk := big.NewInt(2227)

	// ----------------------------------------
	// Assignment 1
	// ----------------------------------------
	fmt.Printf("Assignment 1:\n")
	alicePk, messageCipher := encrypt(*big.NewInt(*secret), *bobPk, *big.NewInt(2000))
	fmt.Printf("Alice's public key is: %s, and the encrypted message is: %s\n", alicePk.Text(10), messageCipher.Text(10))
	fmt.Println()



	// ----------------------------------------
	// Assignment 2
	// ----------------------------------------
	fmt.Printf("Assignment 2:\n")
	fmt.Println("Eve has seen the encrypted message from alice and tries to intercept it...")
	bobSecret, message := intercept(*bobPk, *alicePk, *messageCipher)
	fmt.Printf("Bob's secret is: %s, and the decrypted message is: %s\n", bobSecret.Text(10), message.Text(10))
	fmt.Println()


	// ----------------------------------------
	// Assignment 3
	// ----------------------------------------
	fmt.Printf("Assignment 3:\n")
	fmt.Println("Mallory also sees the message from Alice but do not have the resources to calculate Bob's secret in order to see the message.")
	fmt.Println("Instead Mallory tries to tamper with the message...")
	modifiedCipher := messageCipher.Mul(big.NewInt(3), messageCipher)
	bobMessage := decrypt(bobSecret, *alicePk, *modifiedCipher)
	fmt.Printf("Message after Mallory has tampered with the ciphertext: %s\n", bobMessage.Text(10))
	fmt.Println()
}

/*
Calculates the key.
This can both be the public key or the shared key based on the input
*/
func calculateKey(base, prime, secret big.Int) *big.Int {

	result := big.NewInt(0)
	result.Exp(&base, &secret, nil)
	result.Mod(result, &prime)
	return result
}

/*
Encrypts message by El Gamal's method with a secret and public key
*/
func encrypt(secretKey, publicKey, message big.Int) (key, cipher *big.Int) {
	fmt.Printf("Encrypting...\n")
	sharedBase := big.NewInt(SHARED_BASE)
	sharedPrime := big.NewInt(SHARED_PRIME)

	secret := calculateKey(*sharedBase, *sharedPrime, secretKey)
	commonKey := calculateKey(publicKey, *sharedPrime, secretKey)
	cipher = big.NewInt(0)
	cipher.Mul(commonKey, &message)

	return secret, cipher
}

/*
Decrypts message based on El Gamal's method using a public key and a secret
*/
func decrypt(secret, pk, cipher big.Int) big.Int {

	fmt.Println("Decrypting message...")
	sharedKey := calculateKey(pk, *big.NewInt(SHARED_PRIME), secret)

	result := big.NewInt(0)
	return *result.Div(&cipher, sharedKey)
}

/*
Intercepts message and decrypts it based on the both public keys and the cipher text
*/
func intercept(targetPk, pk, cipher big.Int) (secret, message big.Int) {

	fmt.Println("Intercepting message...")

	base := big.NewInt(SHARED_BASE)
	prime := big.NewInt(SHARED_PRIME)
	var limit big.Int = *big.NewInt(1000)
	incrementer := big.NewInt(1)
	
	for s := *big.NewInt(1); s.Cmp(&limit) < 0; s.Add(&s, incrementer) {

		key := calculateKey(*base, *prime, s)

		if key.Cmp(&targetPk) == 0 {

			fmt.Println("Secret has been found!")
			message := decrypt(s, pk, cipher)
			return s, message
		}
	}

	return *big.NewInt(0), *big.NewInt(0)
}
