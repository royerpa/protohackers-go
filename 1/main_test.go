package main

import (
	"testing"
)

func TestIsPrime(t *testing.T) {
	// Test prime numbers
	if !isPrime(2) {
		t.Error("Expected 2 to be prime")
	}
	if !isPrime(3) {
		t.Error("Expected 3 to be prime")
	}
	if !isPrime(5) {
		t.Error("Expected 5 to be prime")
	}
	if !isPrime(7) {
		t.Error("Expected 7 to be prime")
	}
	if !isPrime(11) {
		t.Error("Expected 11 to be prime")
	}

	// Test non-prime numbers
	if isPrime(1) {
		t.Error("Expected 1 to not be prime")
	}
	if isPrime(4) {
		t.Error("Expected 4 to not be prime")
	}
	if isPrime(6) {
		t.Error("Expected 6 to not be prime")
	}
	if isPrime(8) {
		t.Error("Expected 8 to not be prime")
	}
	if isPrime(9) {
		t.Error("Expected 9 to not be prime")
	}
	if isPrime(882779287409319459912806758585208738419463647745336263878) {
		t.Error("Expected 882779287409319459912806758585208738419463647745336263878 to not be prime")
	}
}
