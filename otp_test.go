package otp

import (
	"fmt"
	"testing"
)

func Test_Gen(t *testing.T) {
	fmt.Println("test gen")
	secret, err := GenSecretKey("sha1")
	if err != nil {
		t.Error("secret", err)
		return
	}
	fmt.Println("secret", secret)
	cfg := InitOTPConfig(secret)
	fmt.Println("url", cfg.ProvisionURIWithIssuer("cjinle", "GitHub"))
}

func Test_Auth(t *testing.T) {
	fmt.Println("test auth")

	secret := "GY4GMN3DHBSTIMTCGFSTKMRUHE4WCMLGGFRWEY3CG5TDSYZTMI4TSNRZMJRGMNRZ"
	// fmt.Println(ComputeCode(secret, 123123))
	cfg := InitOTPConfig(secret)
	fmt.Println(cfg.Authenticate("221597"))
}
