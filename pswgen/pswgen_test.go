package pswgen

import "testing"
import "time"

func TestGenPassword(t *testing.T) {
	key := "wjei@g8r4"

	serials := make([]string, 0, 25)
	scap := cap(serials)
	for i := 0; i < scap; i++ {
		serials = append(serials, GenSerial())
		time.Sleep(time.Millisecond * 150)
	}

	passwords := make([]string, 0, scap)
	for i := 0; i < scap; i++ {
		passwords = append(passwords, GenPassword(serials[i], key))
	}

	for i := range serials {
		if !VerifyPassword(serials[i], key, passwords[i]) {
			t.Error("Serial ", serials[i], " Password ", passwords[i], " invalid")
			t.FailNow()
		}
		t.Log("Serial ", serials[i], " Password ", passwords[i], " pass")
	}
}
