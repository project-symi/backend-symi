package controllers

import (
	"fmt"
	"os"
	"time"

	"project-symi-backend/app/interfaces/database"
	"project-symi-backend/app/usecase"

	jwt "github.com/dgrijalva/jwt-go"
)

type UserAuthController struct {
	Interactor usecase.UserAuthInteractor
}

func NewUserAuthController(sqlHandler database.SqlHandler) *UserAuthController {
	return &UserAuthController{
		Interactor: usecase.UserAuthInteractor{
			UserAuthRepository: &database.UserAuthRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserAuthController) LoginUser(c Context) {
	//READ HEADER INTO A CREDENTIALS STRUCT
	type UserCredentials struct {
		EmployeeId string `json:"userId"`
		Pass       string `json:"password"`
	}
	var user UserCredentials
	if err := c.BindJSON(&user); err != nil {
		panic(err)
	}

	//CHECK IF PASSED CREDENTIALS MATCH
	tokenId, permissionLevel, err := controller.Interactor.CheckUserPass(user.EmployeeId, user.Pass)
	if err != nil {
		c.JSON(401, NewError(err))
		return
	}

	//GET SECRET SIGNING PASSWORD
	signingKey := []byte(os.Getenv("SIGNING_KEY"))

	//GENERATE JWT FROM TOKEN ID
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["jti"] = tokenId
	claims["permissionLevel"] = permissionLevel
	claims["tokenValid"] = true
	claims["exp"] = time.Now().Add(time.Hour * 5).Unix()

	tokenString, err := token.SignedString(signingKey)

	//CREATE A RESPONSE OBJ TO BE SENT AS JSON
	type TokenResponse struct {
		Token           string `json:"token"`
		PermissionLevel string `json:"permission"`
	}

	response := TokenResponse{Token: tokenString, PermissionLevel: permissionLevel}

	c.JSON(200, response)
}

func (controller *UserAuthController) LogoutUser(c Context) {
	token := c.GetHeader("token")
	//PARSE JWT TOKEN
	tokenId, err := getTokenId(token)
	if err != nil && tokenId != "" {
		c.JSON(500, NewError(err))
		return
	}

	//REMOVE THE TOKEN FROM ACTIVE LIST
	amountOfDeleted, err := controller.Interactor.EndUserSession(tokenId)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	if amountOfDeleted == 0 {
		c.JSON(400, NewError(err)) //TODO: create another error
		return
	}

	//RETURN SUCCESS IF NO ERRORS
	c.JSON(200, amountOfDeleted)
}

func (controller *UserAuthController) Authenticate(c Context) {
	token := c.GetHeader("token")

	//PARSE JWT TOKEN
	tokenId, err := getTokenId(token)
	if err != nil && tokenId != "" {
		c.JSON(500, NewError(err))
		return
	}

	tokenValid, err := controller.Interactor.CheckSessionValidity(tokenId)
	if !tokenValid || err != nil {
		c.AbortWithStatus(401)
		return
	}
	return
}

//**************************
//****** HELPER FUNCS ******
//**************************
func getTokenId(tokenString string) (tokenId string, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate that the signing methids is the same
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SIGNING_KEY")), nil
	})

	if err != nil {
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return
	}
	fmt.Println(claims["jti"], claims["permissionLevel"])
	return claims["jti"].(string), nil
}
