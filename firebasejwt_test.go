package firebasejwt

import (
	"strings"
	"testing"
)

// Test negative scenarios
func TestExpiredToken(t *testing.T) {
	str := "eyJhbGciOiJSUzI1NiIsImtpZCI6IjAyMTAwNzE2ZmRkOTA0ZTViNGQ0OTExNmZmNWRiZGZjOTg5OTk0MDEiLCJ0eXAiOiJKV1QifQ.eyJuYW1lIjoiZXVsemJheSBldWx6YmF5IiwicGljdHVyZSI6Imh0dHBzOi8vbGgzLmdvb2dsZXVzZXJjb250ZW50LmNvbS9hL0FDZzhvY0xMd01KX0RYcFRQa0theU9sZ3o2NldacmxPMzQwbXhEMDRpNTVqU25pX241MjJsQT1zOTYtYyIsImlzcyI6Imh0dHBzOi8vc2VjdXJldG9rZW4uZ29vZ2xlLmNvbS9iaXN1cmlpIiwiYXVkIjoiYmlzdXJpaSIsImF1dGhfdGltZSI6MTcyNTY5MzYwOCwidXNlcl9pZCI6IkZabnJpbVpyS01ieHV5a1F5V1R1MENvdEZQWTIiLCJzdWIiOiJGWm5yaW1acktNYnh1eWtReVdUdTBDb3RGUFkyIiwiaWF0IjoxNzI1NjkzNjA4LCJleHAiOjE3MjU2OTcyMDgsImVtYWlsIjoiZXVsemJheUBnbWFpbC5jb20iLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiZmlyZWJhc2UiOnsiaWRlbnRpdGllcyI6eyJnb29nbGUuY29tIjpbIjEwMDk4NTAwMDMzOTk5MTE0ODk1NiJdLCJlbWFpbCI6WyJldWx6YmF5QGdtYWlsLmNvbSJdfSwic2lnbl9pbl9wcm92aWRlciI6Imdvb2dsZS5jb20ifX0.MwYXKk3I_vPAH2lDLceGFFBFQTfKFYgcU2pPPoBU1aYif0B-7ipMW95LBHsGv0yJs-TLMGSXHHr9_0qsJQLIGyh48MVDstvX4JyYPhvbggNgjz8mXCC5kolvrDhZ2To3Ajsmz-rFlpGyRbQDNIOZtBh9dioY5O3f2nayBRYUEFgPrXpAbdvDe_tzw7lqteOzGHz8Wgz5o5nT-qudk6Pna4lmNoxq_uwDIkQGCQOx-ECek85Ai6ETH6r13qegfp31lYBoVSn5drP7H85sJh_AFgsWNnUWiSu1a2C5HlxGOYUC6QWN1DwsCbFnV6aJ6yzKAk5UWno3TPmg_ig8TrTrnA"
	actual, err := ParseFirebaseJWT(str)

	if err == nil {
		t.Fatalf(`Expected error, but got no error`)
	}
	if actual == nil {
		t.Fatalf(`Expected token, but got nil`)
	}
	if !strings.Contains(err.Error(), "token is expired") {
		t.Fatalf(`Expected token expired eror, but got '%s'`, err)
	}
}

// Test empty token
func TestEmptyString(t *testing.T) {
	actual, err := ParseFirebaseJWT("")

	if err == nil {
		t.Fatalf(`Expected error, but got no error`)
	}

	if !strings.Contains(err.Error(), "token string is empty") {
		t.Fatalf(`Expected emptry token eror, but got '%s'`, err)
	}

	if actual != nil {
		t.Fatalf(`Expected no token, but got '%s'`, actual)
	}
}

// Test malformed token
func TestMalformedToken(t *testing.T) {
	tokens := []string{"aa.dddd-d", "asdbc", "....", ".", "-sdaiks ."}

	for _, token := range tokens {
		actual, err := ParseFirebaseJWT(token)
		if err == nil {
			t.Fatalf(`Expected error when parsing '%s', but got no error`, token)
		}
		if actual != nil {
			t.Fatalf(`Expected no token, but got %s`, actual)
		}
	}

}
