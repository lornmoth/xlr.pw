package db

import (
	"database/sql"
	"meguca/common"
	"meguca/config"
	"testing"
	"time"
)

func TestValidateOp(t *testing.T) {
	assertTableClear(t, "boards")
	writeSampleBoard(t)
	writeSampleThread(t)

	cases := [...]struct {
		id      uint64
		board   string
		isValid bool
	}{
		{1, "a", true},
		{15, "a", false},
	}

	for i := range cases {
		c := cases[i]
		t.Run("", func(t *testing.T) {
			t.Parallel()
			valid, err := ValidateOP(c.id, c.board)
			if err != nil {
				t.Fatal(err)
			}
			if valid != c.isValid {
				t.Fatal("unexpected result")
			}
		})
	}
}

func writeSampleBoard(t *testing.T) {
	t.Helper()
	b := BoardConfigs{
		BoardConfigs: config.BoardConfigs{
			ID:        "a",
			Eightball: []string{"yes"},
		},
	}
	err := InTransaction(false, func(tx *sql.Tx) error {
		return WriteBoard(tx, b)
	})
	if err != nil {
		t.Fatal(err)
	}
}

func writeSampleThread(t *testing.T) {
	t.Helper()
	thread := Thread{
		ID:    1,
		Board: "a",
	}
	op := Post{
		StandalonePost: common.StandalonePost{
			Post: common.Post{
				ID:   1,
				Time: time.Now().Unix(),
			},
			OP:    1,
			Board: "a",
		},
		IP: "::1",
	}
	if err := WriteThread(nil, thread, op); err != nil {
		t.Fatal(err)
	}
}
