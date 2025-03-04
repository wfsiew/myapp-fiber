package utils

import (
    "time"
    "github.com/golang-jwt/jwt/v5"
    "app/config"
)

func GenerateToken(id uint, email string) (string, error) {
    claims := jwt.MapClaims{
        "user_id": id,
        "email": email,
        "exp": time.Now().Add(time.Hour * 24 * 365).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    /* token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": id,
        "email": email,
        "exp": time.Now().Add(time.Minute * 1).Unix(),
        //"exp": time.Now().Add(time.Hour * 24 * 365).Unix(),
    }) */
   
    t, err := token.SignedString([]byte(config.Config("JWT_SECRET")))
    if err != nil {
        return "", err
    }
   
    return t, nil
}

func VerifyToken(tokenString string) (bool, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte(config.Config("JWT_SECRET")), nil
    })
    if err != nil {
        return false, err
    }
   
    return token.Valid, nil
}