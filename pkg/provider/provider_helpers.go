package provider

import (
	"crypto/rsa"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mitchellh/go-homedir"
	"github.com/youmark/pkcs8"
	"golang.org/x/crypto/ssh"
)

func mergeSchemas(schemaCollections ...map[string]*schema.Resource) map[string]*schema.Resource {
	out := map[string]*schema.Resource{}
	for _, schemaCollection := range schemaCollections {
		for name, s := range schemaCollection {
			out[name] = s
		}
	}
	return out
}

func getPrivateKey(privateKeyPath, privateKeyString, privateKeyPassphrase string) (*rsa.PrivateKey, error) {
	if privateKeyPath == "" && privateKeyString == "" {
		return nil, nil
	}
	privateKeyBytes := []byte(privateKeyString)
	var err error
	if len(privateKeyBytes) == 0 && privateKeyPath != "" {
		privateKeyBytes, err = readFile(privateKeyPath)
		if err != nil {
			return nil, fmt.Errorf("private Key file could not be read err = %w", err)
		}
	}
	return parsePrivateKey(privateKeyBytes, []byte(privateKeyPassphrase))
}

func readFile(privateKeyPath string) ([]byte, error) {
	expandedPrivateKeyPath, err := homedir.Expand(privateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("invalid Path to private key err = %w", err)
	}

	privateKeyBytes, err := os.ReadFile(expandedPrivateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("could not read private key err = %w", err)
	}

	if len(privateKeyBytes) == 0 {
		return nil, errors.New("private key is empty")
	}

	return privateKeyBytes, nil
}

func parsePrivateKey(privateKeyBytes []byte, passhrase []byte) (*rsa.PrivateKey, error) {
	privateKeyBlock, _ := pem.Decode(privateKeyBytes)
	if privateKeyBlock == nil {
		return nil, fmt.Errorf("could not parse private key, key is not in PEM format")
	}

	if privateKeyBlock.Type == "ENCRYPTED PRIVATE KEY" {
		if len(passhrase) == 0 {
			return nil, fmt.Errorf("private key requires a passphrase, but private_key_passphrase was not supplied")
		}
		privateKey, err := pkcs8.ParsePKCS8PrivateKeyRSA(privateKeyBlock.Bytes, passhrase)
		if err != nil {
			return nil, fmt.Errorf("could not parse encrypted private key with passphrase, only ciphers aes-128-cbc, aes-128-gcm, aes-192-cbc, aes-192-gcm, aes-256-cbc, aes-256-gcm, and des-ede3-cbc are supported err = %w", err)
		}
		return privateKey, nil
	}

	privateKey, err := ssh.ParseRawPrivateKey(privateKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("could not parse private key err = %w", err)
	}

	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("privateKey not of type RSA")
	}
	return rsaPrivateKey, nil
}

type GetRefreshTokenResponseBody struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}

func GetAccessTokenWithRefreshToken(
	tokenEndpoint string,
	clientID string,
	clientSecret string,
	refreshToken string,
) (string, error) {
	v := url.Values{}
	v.Set("grant_type", "refresh_token")
	v.Set("refresh_token", refreshToken)
	req, err := http.NewRequest("POST", tokenEndpoint, strings.NewReader(v.Encode()))
	if err != nil {
		return "", fmt.Errorf("new http request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(url.QueryEscape(clientID), url.QueryEscape(clientSecret))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("do http request: %w", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("read response body: %w", err)
	}
	if res.StatusCode != 200 {
		return "", fmt.Errorf("http response error: %s, %s", res.Status, body)
	}
	contentType, _, err := mime.ParseMediaType(res.Header.Get("Content-Type"))
	if err != nil {
		return "", fmt.Errorf("parse content type: %w", err)
	}
	var token GetRefreshTokenResponseBody
	switch contentType {
	case "application/x-www-form-urlencoded", "text/plain":
		vals, err := url.ParseQuery(string(body))
		if err != nil {
			return "", fmt.Errorf("parse query: %w", err)
		}
		token.AccessToken = vals.Get("access_token")
		token.TokenType = vals.Get("token_type")
		if e := vals.Get("expires_in"); e != "" {
			expires, _ := strconv.ParseInt(e, 10, 64)
			if expires != 0 {
				token.ExpiresIn = expires
			}
		}
	default:
		if err := json.Unmarshal(body, &token); err != nil {
			return "", fmt.Errorf("unmarshal json: %w", err)
		}
	}
	if token.AccessToken == "" {
		return "", fmt.Errorf("oauth2: server response missing access_token")
	}
	return token.AccessToken, nil
}
