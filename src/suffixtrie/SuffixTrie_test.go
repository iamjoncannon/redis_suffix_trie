package suffixtrie

import (
	"reflect"
	"testing"
)

func TestInsertIntoTrie(t *testing.T) {

	firstTrie := NewTrie()

	firstTrie.InsertIntoTrie("a", "123")

	// test one

	if _, ok := (*firstTrie)["a"]; ok != true {

		t.Errorf("the trie would store one string as a key")
	}

	// test two

	if (*firstTrie)["a"][0] != "123" {

		t.Errorf("value is a slice containing the payload as the first")
	}

	secondTrie := NewTrie()

	secondTrie.InsertIntoTrie("ab", "123")

	_, threeOk1 := (*secondTrie)["a"]
	_, threeOk2 := (*secondTrie)["b"]
	_, threeOk3 := (*secondTrie)["ab"]

	// test three

	if (threeOk1 && threeOk2 && threeOk3) != true {

		t.Errorf("given a two char key, it would store separate")
		t.Errorf("keys for each char, and one key for both chars")
	}

	thirdTrie := NewTrie()

	thirdTrie.InsertIntoTrie("ab", "123")
	thirdTrie.InsertIntoTrie("ab", "456")

	if len((*thirdTrie)["a"]) != 2 {

		t.Errorf("given two entries same key with two different payloads")
		t.Errorf("stores both payloads in []string, in the order of insertion")
	}

	testThreeExpectOne := (*thirdTrie)["a"][0] == "123"
	testThreeExpectTwo := (*thirdTrie)["a"][1] == "456"

	if (testThreeExpectOne && testThreeExpectTwo) != true {

		t.Errorf("given two entries with same key and two different payloads")
		t.Errorf("stores both payloads on []string")
	}

	// test four

	fourthTrie := NewTrie()

	fourthTrie.InsertIntoTrie("aaa", "123")

	if len((*fourthTrie)["a"]) != 1 {

		t.Errorf("if payload already exists in key, doesn't append again")
	}

}

func TestContains(t *testing.T) {

	firstTrie := NewTrie()

	firstTrie.InsertIntoTrie("ab", "123")

	testOne := firstTrie.Contains("a")

	if reflect.TypeOf(testOne).String() != "[]string" {

		t.Errorf("given an existing key, trie.Contains returns the associated value as []string")
	}

	if testOne[0] != "123" {

		t.Errorf("given an existing key, trie.Contains returns the associated value as []string")
	}

	testTwo := firstTrie.Contains("b")

	if testTwo[0] != "123" {

		t.Errorf("given an existing key, trie.Contains returns the associated value as []string")
	}

	testThree := firstTrie.Contains("ab")

	if testThree[0] != "123" {

		t.Errorf("given an existing key, trie.Contains returns the associated value as []string")
	}

	testFour := firstTrie.Contains("c")

	if testFour[0] != "<NOT FOUND>" {

		t.Errorf("given a non-existent key, Contains returns '<NOT FOUND>'")
	}
}

func TestInsertIntoRedisTrie(t *testing.T) {

	/*
		TODO- insert a mock into this function
		instead of testing a connection to a real Redis instance

		e.g. https://github.com/alicebob/miniredis

		I couldn't figure out how to do this in a type safe way-
		in a dynamically typed language you could just have the
		function switch out the real and mocked connection, but
		golang requires the types to be identical-- and you can't
		simply use interface{} as a placeholder between the mock
		and actual client
	*/

	pool := rediscache.NewPool()
	conn := pool.Get()
	defer conn.Close()

	// we expect it to behave exactly like the in application
	// function above, but using get and set to redis connection
	// using lists

	

}
