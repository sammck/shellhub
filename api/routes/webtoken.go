package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/shellhub-io/shellhub/api/apicontext"
	"github.com/shellhub-io/shellhub/api/nsadm"
	"golang.org/x/crypto/bcrypt"
)

const (
	GetTokenListURL = "/token"
)

func GenerateToken(c apicontext.Context) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(c.Tenant()), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hash to store:", string(hash))

	hasher := md5.New()
	hasher.Write(hash)
	return hex.EncodeToString(hasher.Sum(nil))
}

func GetTokenList(c apicontext.Context) error {
	svc := nsadm.NewService(c.Store())

	tokenList, count, err := svc.ListTokens()
	if err != nil {
		return err
	}

	c.Response().Header().Set("X-Total-Count", strconv.Itoa(count))

	return c.JSON(http.StatusOK, tokenList)
}
