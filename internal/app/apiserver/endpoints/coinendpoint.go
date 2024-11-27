package endpoints

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/VitalyCone/kuznecov_coins_api/internal/app"
	dto "github.com/VitalyCone/kuznecov_coins_api/internal/app/apiserver/dtos"
	"github.com/VitalyCone/kuznecov_coins_api/internal/app/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func DecodeJWT(tokenString string) (*Claims, error) {
    secretKey := []byte(app.CurrentToken) // Замените на ваш секретный ключ

    claims := &Claims{}
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        // Проверка алгоритма
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return secretKey, nil
    })

    if err != nil {
        return nil, err
    }

    if !token.Valid {
        return nil, fmt.Errorf("invalid token")
    }

    return claims, nil
}

// @Summary Get coins for user
// @Schemes
// @Description Get coins for user by token
// @Tags Coins
// @Accept json
// @Produce json
// @Router /user [GET]
func (ep *Endpoints) GetCoin(g *gin.Context) {
	token := strings.Split(g.Request.Header["Authorization"][0], " ")[1]

	claims, err := DecodeJWT(token)
    if err != nil {
        fmt.Println("Error decoding token:", err)
        return
    }

	model, err := ep.store.Coin().FindByUsername(claims.Subject)
	if err != nil{
		g.JSON(http.StatusNotFound, "Not found user: " + err.Error())
		return
	}

	returningModel := dto.CoinModelToCoinDetailsDto(model)
	g.JSON(200, returningModel)
}

// @Summary Get coins for user
// @Schemes
// @Description Get coins for user by token
// @Tags Coins
// @Accept json
// @Produce json
// @Router /user [GET]
func (ep *Endpoints) PostNewUser(g *gin.Context) {
	token := strings.Split(g.Request.Header["Authorization"][0], " ")[1]

	claims, err := DecodeJWT(token)
    if err != nil {
        g.JSON(http.StatusUnauthorized, "Error decoding token:" + err.Error())
        return
    }

	_, err = ep.store.Coin().FindByUsername(claims.Subject)
	if err == nil{
		g.JSON(http.StatusConflict, "User already exists")
		return
	}

	model := model.Coin{
		Username: claims.Subject,
		Coins: 0,
	}

	if err := ep.store.Coin().Create(&model); err != nil{
		g.JSON(http.StatusInternalServerError, "Faild to create user:" + err.Error())
		return
	}

	returningModel := dto.CoinModelToCoinDetailsDto(model)
	g.JSON(http.StatusCreated, returningModel)
}

// @Summary Get coins for user
// @Schemes
// @Description Get coins for user by token
// @Tags Coins
// @Accept json
// @Produce json
// @Param coins body int true "new user coins"
// @Router /user [PUT]
func (ep *Endpoints) PutCoin(g *gin.Context) {
	type Request struct{
		Coins int `json:"coins"`
	}
	req := Request{}

	err:= g.BindJSON(&req)
	if err != nil{
		g.JSON(http.StatusBadRequest, "Invalid request: " + err.Error())
		return
	}

	if req.Coins < 0{
		g.JSON(http.StatusBadRequest, "Coins can't be negative")
		return
	}

	token := strings.Split(g.Request.Header["Authorization"][0], " ")[1]

	claims, err := DecodeJWT(token)
    if err != nil {
        g.JSON(http.StatusUnauthorized, "Error decoding token: " + err.Error())
        return
    }

	model, err := ep.store.Coin().UpdateCoinsByUsername(claims.Subject, req.Coins)
	if err != nil{
		g.JSON(http.StatusNotFound, "Failed to put user: " + err.Error())
		return
	}

	returningModel := dto.CoinModelToCoinDetailsDto(model)
	g.JSON(http.StatusOK, returningModel)
}
