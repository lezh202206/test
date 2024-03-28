package auth

import (
	"fmt"
	"testing"
)

func TestJwtConf_Get(t *testing.T) {
	claim := CustomerClaim{
		UserId:     1,
		MerchantId: 2,
		Corpid:     "",
		TokenType:  1,
	}
	genJwtToken, err := GenJwtToken(claim)
	if err != nil {
		return
	}
	fmt.Printf("%+v", genJwtToken)

}
func TestJwtConf_Parse(t *testing.T) {
	var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VyX25hbWUiOiIiLCJtZXJjaGFudF9pZCI6MiwiY29ycGlkIjoiIiwidG9rZW5fdHlwZSI6MSwic291cmNlIjowLCJzdXBlcl9hZG1pbiI6MCwiZXhwIjoxNzE0MTE0MjIyLCJpYXQiOjE3MTE1MjIyMjIsImlzcyI6InpieSIsIm5iZiI6MTcxMTUyMjIyMiwic3ViIjoiemJ5In0.BZkb38ZxsO0G1g11aK7Pw9swcRDgI44VlVthXk_j_XQ"
	jwtToken, err := ParseJwtToken(token)
	if err != nil {
		return
	}
	fmt.Printf("%+v", jwtToken)
}
