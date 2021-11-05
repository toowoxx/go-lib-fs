package filepaths

import "testing"

func TestComponents(t *testing.T) {
	c := Components("/")
	if len(c) != 2 {
		t.Fatal("expected two empty strings, instead got", c)
	}
	if c[0] != "" {
		t.Fatal("expected first string to be empty, instead got", c[0])
	}
	if c[1] != "" {
		t.Fatal("expected second string to be empty, instead got", c[1])
	}

	c = Components("/1/2/3.txt")
	if len(c) != 4 {
		t.Fatal("expected 4 strings, instead got", c)
	}
	if c[0] != "" {
		t.Fatal("expected first string to be empty, instead got", c[0])
	}
	if c[1] != "1" {
		t.Fatal("expected first string to be '1', instead got", c[1])
	}
	if c[2] != "2" {
		t.Fatal("expected second string to be '2', instead got", c[2])
	}
	if c[3] != "3.txt" {
		t.Fatal("expected third string to be '3.txt', instead got", c[3])
	}

	c = Components("abc.txt")
	if len(c) != 1 {
		t.Fatal("expected one string, instead got", c)
	}
	if c[0] != "abc.txt" {
		t.Fatal("expected first string to be 'abc.txt', instead got", c[0])
	}

	c = Components("")
	if len(c) != 1 {
		t.Fatal("expected one string, instead got", c)
	}
	if c[0] != "" {
		t.Fatal("expected first string to be empty, instead got", c)
	}
}
