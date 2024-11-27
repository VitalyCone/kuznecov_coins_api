package endpoints

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/VitalyCone/account/internal/app"
	"github.com/VitalyCone/account/internal/app/apiserver/dtos"
	"github.com/VitalyCone/account/internal/app/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var(
    bcryptSalt = 10
)

func createToken(user model.User)(string, jwt.Claims, error){

    claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Username,                    // Subject (user identifier)
        "firstname": user.FirstName,
        "secondname": user.SecondName,
        "created_at": user.CreatedAt.Unix(),
        "updated_at": user.UpdatedAt.Unix(),
		//"aud": getRole(username),           // Audience (user role)
		//"exp": time.Now().Add(time.Hour).Unix(), // Expiration time
		"iat": time.Now().Unix(),                 // Issued at
	})

    tokenString, err := claims.SignedString([]byte(app.CurrentToken))
    if err!= nil{
        return "", nil, err
    }

    return tokenString, claims.Claims, nil
}

func verifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(app.CurrentToken), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}

func verifyPassword(hashPass, password string) error{
    return bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(password))
}


// @Summary Register for user
// @Schemes
// @Description Register in api
// @Tags Authorization
// @Accept json
// @Produce json
// @Param createUserDto body dtos.CreateUserDto true "Create user dto for register in"
// @Router /account/register [POST]
func (ep *Endpoints) RegisterUser(g *gin.Context){
    var CreateUserDto dtos.CreateUserDto

    if err := g.BindJSON(&CreateUserDto); err != nil {
        g.JSON(http.StatusBadRequest, gin.H{"Invalid form data": err.Error()})
        return
    }

    validate := validator.New()
    if err := validate.Struct(CreateUserDto); err != nil{
        g.JSON(http.StatusBadRequest, gin.H{"Failed validation": err.Error()})
        return
    }

    _, exists := ep.store.User().FindUserByUsername(strings.ToLower(CreateUserDto.Username))
    if exists == nil{
        g.JSON(http.StatusBadRequest, "Username already exists")
        return
    }

    passHash, err := bcrypt.GenerateFromPassword([]byte(CreateUserDto.Password), bcryptSalt)
    if err != nil{
        g.JSON(http.StatusBadRequest, gin.H{"Failed to generate password hash": err.Error()})
        return
    }

    user:= CreateUserDto.ToModel(string(passHash))

    if err := ep.store.User().CreateUser(&user); err != nil{
        g.JSON(http.StatusInternalServerError, gin.H{"Failed to create user on db": err.Error()})
        return
    }
    
    tokenString, _, err := createToken(user)
    if err != nil {
        ep.store.User().DeleteUserByUsername(user.Username)
        g.JSON(http.StatusInternalServerError, gin.H{"Error creating token" : err.Error()})
        return
    }

    //g.SetCookie("token", tokenString, 3600, "/", "localhost", false, true)

    g.JSON(http.StatusCreated, tokenString)
}

// @Summary Login for user
// @Schemes
// @Description login in api
// @Tags Authorization
// @Accept json
// @Produce json
// @Param userDto body dtos.UserDto true "User  dto for login in"
// @Router /account/login [post]
func (ep *Endpoints) LoginUser(g *gin.Context){
    var userDto dtos.UserDto

    if err := g.BindJSON(&userDto); err != nil {
        g.JSON(http.StatusBadRequest, gin.H{"Invalid form data": err.Error()})
        return
    }

    user, err := ep.store.User().FindUserByUsername(strings.ToLower(userDto.Username))
    if err != nil{
        g.JSON(http.StatusBadRequest, gin.H{"User doesn't exist": err.Error()})
        return
    }

    verify:= verifyPassword(user.PasswordHash, userDto.Password)

    if verify != nil {
        g.JSON(http.StatusBadRequest, "Invalid password")
        return
    }
    
    tokenString, _, err := createToken(user)
    if err != nil {
        g.JSON(http.StatusInternalServerError, "Error creating token")
        return
    }
    
    //g.SetCookie("token", tokenString, 3600, "/", "localhost", false, true)

    g.JSON(http.StatusOK, tokenString)
}


// @Summary Get user data by token
// @Schemes
// @Description Get user data by token
// @Security ApiKeyAuth
// @Tags Users
// @Accept json
// @Produce json
// @Router /account/info [GET]
func (ep *Endpoints) GetUserInfo(g *gin.Context){
    tokenString := g.GetHeader("token")
    if tokenString == ""{
        g.JSON(500, "token nil")
        return
    }

    token, err := verifyToken(tokenString)
    if err != nil{
        g.JSON(500, "token not verifed or nil")
        return
    }

    username, err := token.Claims.GetSubject()
    if err != nil{
        g.JSON(500, "Failed to get subject from token")
        return
    }

    user, err := ep.store.User().FindUserByUsername(username)
    if err != nil{
        g.JSON(500, "User not found")
        return
    }
    g.JSON(200, user)
}

// @Summary Get user token data
// @Schemes
// @Description Get user token data
// @Security ApiKeyAuth
// @Tags Users
// @Accept json
// @Produce json
// @Param ModifyUserDto body dtos.ModifyUserDto true "User dto for modify user"
// @Router /account/info [PUT]
func (ep *Endpoints) PutUserInfo(g *gin.Context){
    var userDto dtos.ModifyUserDto

    tokenString := g.GetHeader("token")
    if tokenString == ""{
        g.JSON(http.StatusUnauthorized, "token nil")
        return
    }

    token, err := verifyToken(tokenString)
    if err != nil{
        g.JSON(http.StatusUnauthorized, "token not verifed or nil")
        return
    }


    username, err := token.Claims.GetSubject()
    if err != nil{
        g.JSON(http.StatusBadRequest, gin.H{"Username nil": err.Error()})
    }

    
    if err := g.BindJSON(&userDto); err != nil {
        g.JSON(http.StatusBadRequest, gin.H{"Invalid json data": err.Error()})
        return
    }

    validate := validator.New()
    if err := validate.Struct(userDto); err != nil{
        g.JSON(http.StatusBadRequest, gin.H{"Failed validation": err.Error()})
        return
    }

    userDto.Username = strings.ToLower(userDto.Username)

    userModel, err := ep.store.User().FindUserByUsername(username)
    if err != nil{
        g.JSON(http.StatusInternalServerError, gin.H{"User not found": err.Error()})
    }
    
    if userDto.NewPassword != ""{
        if userDto.OldPassword == ""{
            g.JSON(http.StatusBadRequest, "Old password is nil")
            return
        }else{
            if err := verifyPassword(userModel.PasswordHash, userDto.OldPassword); err != nil{
                g.JSON(http.StatusBadRequest, gin.H{"Old password is incorrect": err.Error()})
                return
            }

            passHash, err := bcrypt.GenerateFromPassword([]byte(userDto.NewPassword), bcryptSalt)
            if err != nil{
                g.JSON(http.StatusBadRequest, gin.H{"Failed to generate password hash": err.Error()})
                return
            }

            userModel.PasswordHash = string(passHash)
        }
    }

    if userDto.FirstName != ""{
        userModel.FirstName = userDto.FirstName
    }

    if userDto.SecondName != ""{
        userModel.SecondName = userDto.SecondName
    }

    if userDto.Avatar != ""{
        // Тут бы добавить проверку на валидность изображения
        userModel.Avatar = []byte(userDto.Avatar)
    }

    if userDto.Username != ""{
        _, exist := ep.store.User().FindUserByUsername(userDto.Username)
        if exist == nil{
            g.JSON(http.StatusBadRequest, "Username already exists")
            return
        }

        userModel.Username = userDto.Username
    }

    err = ep.store.User().ModifyUser(username, &userModel)
    if err != nil{
        g.JSON(http.StatusBadRequest, gin.H{"Failed to save user": err.Error()})
        return 
    }

    tokenString, _, err = createToken(userModel)
    if err != nil {
        g.JSON(http.StatusInternalServerError, gin.H{"Error creating token" : err.Error()})
        return
    }

    g.JSON(http.StatusCreated, tokenString)
}

// @Summary Delete user by token data
// @Schemes
// @Description Delete user by token data
// @Security ApiKeyAuth
// @Tags Users
// @Accept json
// @Produce json
// @Router /account/delete [DELETE]
func (ep *Endpoints) DeleteUserInfo(g *gin.Context){
    tokenString := g.GetHeader("token")
    if tokenString == ""{
        g.JSON(http.StatusUnauthorized, "token nil")
        return
    }

    token, err := verifyToken(tokenString)
    if err != nil{
        g.JSON(http.StatusUnauthorized, "token not verifed or nil")
        return
    }

    username, err := token.Claims.GetSubject()
    if err != nil{
        g.JSON(http.StatusBadRequest, gin.H{"Username nil": err.Error()})
    }

    ep.store.User().DeleteUserByUsername(username)

    tokenString = ""

    g.JSON(http.StatusOK, tokenString)
}