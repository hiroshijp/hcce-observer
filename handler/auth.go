package handler

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"io"
	"log"
	"net/http"

	"github.com/coreos/go-oidc"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

type AuthHandler struct {
	ClientID     string
	ClientSecret string
	OidcConfig   *oidc.Config
	Config       oauth2.Config
	Verifier     *oidc.IDTokenVerifier
}

func NewAuthHandler(e *echo.Echo, clientID, clientSecret string) {
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, "https://accounts.google.com")
	if err != nil {
		log.Fatal(err)
		return
	}
	oidcConfig := &oidc.Config{
		ClientID: clientID,
	}
	verifier := provider.Verifier(oidcConfig)

	config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  "http://127.0.0.1:8080/auth/google/callback",
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}

	handler := &AuthHandler{
		OidcConfig: oidcConfig,
		Config:     config,
		Verifier:   verifier,
	}

	e.GET("/auth", handler.login)
	e.GET("/auth/google/callback", handler.callback)
}

func (ah *AuthHandler) login(c echo.Context) error {

	state, _ := randString(16)

	return c.Redirect(302, ah.Config.AuthCodeURL(state))
}

func (ah *AuthHandler) callback(c echo.Context) error {
	ctx := c.Request().Context()
	oauth2Token, err := ah.Config.Exchange(ctx, c.QueryParam("code"))
	if err != nil {
		return err
	}
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		return err
	}
	idToken, err := ah.Verifier.Verify(ctx, rawIDToken)
	if err != nil {
		return err
	}

	var myClaim struct {
		Email         string `json:"email"`
		EmailVerified bool   `json:"email_verified"`
	}

	if err := idToken.Claims(&myClaim); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, myClaim)
}

func randString(nByte int) (string, error) {
	b := make([]byte, nByte)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}
