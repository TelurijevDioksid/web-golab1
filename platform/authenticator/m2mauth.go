package authenticator

import (
	"context"
	"errors"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
)

type CustomClaims struct {
	Scope    string `json:"scope"`
	Audience string `json:"aud"`
	Issuer   string `json:"iss"`
	Subject  string `json:"sub"`
}

func (c CustomClaims) Validate(ctx context.Context) error {
	return nil
}

type M2MAuthenticator struct {
	JwtValidator *validator.Validator
}

func NewM2M() (*M2MAuthenticator, error) {
	issUrl, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/")
	if err != nil {
		log.Fatalf("Failed to parse issuer URL: %v", err)
	}

	provider := jwks.NewCachingProvider(issUrl, 5*time.Minute)

	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issUrl.String(),
		[]string{os.Getenv("AUTH0_AUDIENCE")},
		validator.WithCustomClaims(
			func() validator.CustomClaims {
				return &CustomClaims{}
			},
		),
		validator.WithAllowedClockSkew(time.Minute),
	)
	if err != nil {
		log.Fatalf("Failed to create JWT validator: %v", err)
	}

	return &M2MAuthenticator{
		JwtValidator: jwtValidator,
	}, nil
}

func (a *M2MAuthenticator) VerifyM2MToken(ctx context.Context, authHead string) error {
	parts := strings.Split(authHead, " ")

	if len(parts) != 2 {
		return errors.New("Invalid token")
	}

	if parts[0] != "Bearer" {
		return errors.New("Invalid token")
	}

	_, err := a.JwtValidator.ValidateToken(ctx, parts[1])
	if err != nil {
		return errors.New("Invalid token")
	}

	return nil
}
