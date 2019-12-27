package main

import (
	"log"
	"os"
	"testing"
)

func TestLevelToFloat(t *testing.T) {
	given := "level 4   -  20%"
	wanted := 4.20
	res := LevelToFloat(given)
	if res != wanted {
		t.Errorf("\nGOT:[%.2f]\nORG:[%.2f]\nGIVEN:[%s]\n", wanted, res, given)
	}

	given = "level 5   -  2%"
	wanted = 5.02
	res = LevelToFloat(given)
	if res != wanted {
		t.Errorf("\nGOT:[%.2f]\nORG:[%.2f]\nGIVEN:[%s]\n", wanted, res, given)
	}

	given = "level 4   -  0%"
	wanted = 4.00
	res = LevelToFloat(given)
	if res != wanted {
		t.Errorf("\nGOT:[%.2f]\nORG:[%.2f]\nGIVEN:[%s]\n", wanted, res, given)
	}

	given = "level 4 -  "
	wanted = 4.00
	res = LevelToFloat(given)
	if res != wanted {
		t.Errorf("\nGOT:[%.2f]\nORG:[%.2f]\nGIVEN:[%s]\n", wanted, res, given)
	}

	given = "level 4"
	wanted = 0.00
	res = LevelToFloat(given)
	if res != wanted {
		t.Errorf("\nGOT:[%.2f]\nORG:[%.2f]\nGIVEN:[%s]\n", wanted, res, given)
	}
}

func TestGetLevelString(t *testing.T) {
	r, err := os.Open("../../data/mtuomine")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	wanted := "level 5 - 20%"
	actual := GetLevelString(r)
	if wanted != actual {
		t.Errorf("\nGOT:[%s]\nORG:[%s]\n", actual, wanted)
	}
}
