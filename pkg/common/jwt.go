package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// StandardClaim defines the structure of JWT claims with custom fields
type StandardClaim struct {
	UserId int      `json:"id,omitempty"`
	RoleId int      `json:"roleId,omitempty"`
	Email  string   `json:"email,omitempty"`
	Roles  []string `json:"roles"`
	jwt.RegisteredClaims
}

// DecodeJwtToken decodes and validates a JWT token string and returns the claims
func DecodeJwtToken(tString string) (*StandardClaim, error) {
	mapClaim := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tString, mapClaim, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if token == nil || err != nil {
		switch {
		case err.Error() == "Token is expired":
			return nil, errors.New("token is expired")
		case err.Error() == "signature is invalid":
			return nil, errors.New("invalid token signature")
		default:
			return nil, fmt.Errorf("failed to decode token: %v", err)
		}
	}

	claimResult := &StandardClaim{}
	if err := mapClaimsToStandardClaim(mapClaim, claimResult); err != nil {
		return nil, err
	}

	return claimResult, nil
}

// EncodeJwtToken encodes the given claims into a JWT token string
func EncodeJwtToken(customClaim *StandardClaim) (string, error) {
	claim := customClaim
	if claim == nil {
		duration, _ := strconv.Atoi(os.Getenv("JWT_EXPIRY"))
		claim = &StandardClaim{
			UserId: 0,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "example@domain.com",
				Subject:   "users-authentication",
				Audience:  []string{"xavier"},
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(duration) * time.Minute)),
				NotBefore: jwt.NewNumericDate(time.Now()),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
		}
	}

	sign := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := sign.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return token, nil
}

// DecodeTokenWithoutValidate decodes a JWT token string without validation and returns the claims
func DecodeTokenWithoutValidate(token string) (*StandardClaim, error) {
	mapClaim := jwt.MapClaims{}
	jwt.ParseWithClaims(token, mapClaim, nil)

	claimResult := &StandardClaim{}
	if err := mapClaimsToStandardClaim(mapClaim, claimResult); err != nil {
		return nil, err
	}

	return claimResult, nil
}

// mapClaimsToStandardClaim maps JWT claims from mapClaim to StandardClaim struct
func mapClaimsToStandardClaim(mapClaim jwt.MapClaims, claimResult *StandardClaim) error {
	if iss, ok := mapClaim["iss"].(string); ok {
		claimResult.Issuer = iss
	}

	if sub, ok := mapClaim["sub"].(string); ok {
		claimResult.Subject = sub
	}

	if aud, ok := mapClaim["aud"].([]string); ok {
		claimResult.Audience = aud
	}

	if jti, ok := mapClaim["jti"].(string); ok {
		claimResult.ID = jti
	}

	if roles, ok := mapClaim["roles"].([]string); ok {
		claimResult.Roles = roles
	} else {
		claimResult.Roles = []string{}
	}

	if userId, ok := mapClaim["id"].(float64); ok {
		claimResult.UserId = int(userId)
	}

	if roleId, ok := mapClaim["roleId"].(float64); ok {
		claimResult.RoleId = int(roleId)
	} else if rid, ok := mapClaim["rid"].(float64); ok {
		fmt.Println("roleId is nil, using rid")
		claimResult.RoleId = int(rid)
	}

	if exp, ok := parseTimeClaim(mapClaim["exp"]); ok {
		claimResult.ExpiresAt = jwt.NewNumericDate(exp)
	}

	if iat, ok := parseTimeClaim(mapClaim["iat"]); ok {
		claimResult.IssuedAt = jwt.NewNumericDate(iat)
	}

	return nil
}

// parseTimeClaim parses a time claim from a JWT token claim
func parseTimeClaim(claim interface{}) (time.Time, bool) {
	switch v := claim.(type) {
	case float64:
		return time.Unix(int64(v), 0), true
	case json.Number:
		if val, err := v.Int64(); err == nil {
			return time.Unix(val, 0), true
		}
	}
	return time.Time{}, false
}
