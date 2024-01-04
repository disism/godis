package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// https://datatracker.ietf.org/doc/html/rfc7519
// JWT https://datatracker.ietf.org/doc/html/rfc7519#section-4

type SigningMethod int

const (
	SigningMethodHS256 SigningMethod = iota
)

var signingMethods = map[SigningMethod]jwt.SigningMethod{
	SigningMethodHS256: jwt.SigningMethodHS256,
}

type JWT struct {
	*jwt.RegisteredClaims
	signingMethod jwt.SigningMethod
}

func New(method SigningMethod, opts ...Option) *JWT {
	now := time.Now()
	j := &JWT{
		RegisteredClaims: &jwt.RegisteredClaims{
			IssuedAt: jwt.NewNumericDate(now),
		},
		signingMethod: signingMethods[method],
	}
	for _, opt := range opts {
		opt(j)
	}
	return j
}

func (j *JWT) Generate(secret string) (string, error) {
	token := jwt.NewWithClaims(j.signingMethod, j.RegisteredClaims)
	return token.SignedString([]byte(secret))
}

func (j *JWT) Validate(token, secret string) (*jwt.RegisteredClaims, error) {
	claims := &jwt.RegisteredClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if token.Method != j.signingMethod {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}

type Option func(*JWT)

func WithIssuer(issuer string) Option {
	return func(j *JWT) {
		j.Issuer = issuer
	}
}

func WithSubject(subject string) Option {
	return func(j *JWT) {
		j.Subject = subject
	}
}

func WithAudience(audience []string) Option {
	return func(j *JWT) {
		j.Audience = audience
	}
}

func WithExpiresAt(expiresAt time.Time) Option {
	return func(j *JWT) {
		j.ExpiresAt = jwt.NewNumericDate(expiresAt)
	}
}

func WithNotBefore(notBefore time.Time) Option {
	return func(j *JWT) {
		j.NotBefore = jwt.NewNumericDate(notBefore)
	}
}

func WithIssuedAt(issuedAt time.Time) Option {
	return func(j *JWT) {
		j.IssuedAt = jwt.NewNumericDate(issuedAt)
	}
}

func WithID(id string) Option {
	return func(j *JWT) {
		j.ID = id
	}
}
