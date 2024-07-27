package util

import (
    "github.com/dgrijalva/jwt-go"
    "auth/lib/config"
    "time"
)

func GenerateJWT(userID int64) (string, error) {
    cfg := config.LoadConfig()
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "id":  userID,
        "exp": time.Now().Add(30 * 24 * time.Hour).Unix(),
    })
    signedToken, err := token.SignedString([]byte(cfg.JWT_SECRET))
    if err != nil {
        return "", err
    }
    return signedToken, nil
}

func ParseJWT(tokenStr string) (float64, error) {
    cfg := config.LoadConfig()
    token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
        return []byte(cfg.JWT_SECRET), nil
    })
    if err != nil || !token.Valid {
        return 0, err
    }
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return 0, err
    }
    return claims["id"].(float64), nil
}
