package middleware

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"gitlab.id.vin/gami/ps2-gami-common/configs"
	"gitlab.id.vin/gami/ps2-gami-common/dtos"
	"gitlab.id.vin/gami/ps2-gami-common/errors"
	"gitlab.id.vin/gami/ps2-gami-common/logger"
	"gitlab.id.vin/gami/ps2-gami-common/utils"
	"gitlab.id.vin/gami/ps2-gami-common/utils/ctxutil"

	"github.com/gin-gonic/gin"
)

const (
	StatusOK               = 200000
	HeaderAuthorizationKey = "Authorization"
	LocalCacheAuth0JwksKey = "auth0-jwks"
)

var jwksMap = make(map[string]*dtos.Jwks)

func Auth0Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			bearerToken = c.GetHeader(HeaderAuthorizationKey)
			splits      = strings.Split(bearerToken, " ")
		)
		if len(splits) != 2 || splits[0] != "Bearer" {
			logger.Context(c.Request.Context()).Errorf("Authorization header is invalid format")
			utils.HandleError(c, errors.New(errors.UnauthorizedCodeError))
			c.Abort()
			return
		}
		jwtParser := &jwt.Parser{}
		jwtClaims := dtos.Auth0Claims{}
		token, _, err := jwtParser.ParseUnverified(splits[1], &jwtClaims)
		claim, ok := token.Claims.(*dtos.Auth0Claims)
		if !ok {
			logger.Context(c.Request.Context()).Errorf("Access token is invalid, err: %v", err)
			utils.HandleError(c, errors.New(errors.UnauthorizedCodeError))
			c.Abort()
			return
		}

		tmpCtx := ctxutil.SetAuth0Claim(c.Request.Context(), claim)
		tmpCtx = ctxutil.SetAuthClaims(tmpCtx, claim)
		tmpCtx = ctxutil.SetUserID(tmpCtx, claim.VinIdInfo.UserID)

		c.Request = c.Request.WithContext(tmpCtx)
	}
}

func parseClaims(tokenString string) (*dtos.Auth0Claims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &dtos.Auth0Claims{}, func(token *jwt.Token) (interface{}, error) {
		cert, err := getPemCert(token, configs.AppConfig.Auth0Service.JwksUri)
		if err != nil {
			logger.Errorf("[ParseWithClaims][getPemCert]Get Pem Cert error: %", err)
			return nil, err
		}
		result, err := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))

		if err != nil {
			logger.Errorf("[ParseWithClaims][ParseRSAPublicKeyFromPEM] error: %", err)
		}

		return result, err
	}, jwt.WithIssuer("https://"+configs.AppConfig.Auth0Service.Domain+"/"),
		jwt.WithAudience(configs.AppConfig.Auth0Service.Audience))

	if err != nil {
		logger.Errorf("[parseClaims][ParseWithClaims]Parse with claims error: %", err)
		return nil, err
	}

	claims, ok := token.Claims.(*dtos.Auth0Claims)

	if !ok || !token.Valid {
		return nil, errors.New(errors.UnauthorizedCodeError, "Invalid token")
	}

	return claims, nil

}

func getPemCert(token *jwt.Token, jwksUri string) (string, error) {
	cert := ""
	jwks, ok := jwksMap[LocalCacheAuth0JwksKey]

	if !ok {
		resp, err := http.Get(jwksUri)

		if err != nil {
			return cert, err
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				logger.Errorf("[getPemCert][Close]Close Body Resp error: %v", err)
			}
		}(resp.Body)

		err = json.NewDecoder(resp.Body).Decode(&jwks)

		if err != nil {
			logger.Errorf("[getPemCert][Decode] Decode jwks error: %v", err)
			return cert, err
		}

		jwksMap[LocalCacheAuth0JwksKey] = jwks
	}

	for k := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		logger.Errorf("[getPemCert][Cert]Unable to find appropriate key")
		return cert, errors.New(errors.UnauthorizedCodeError, "Unable to find appropriate key.")
	}

	return cert, nil
}
