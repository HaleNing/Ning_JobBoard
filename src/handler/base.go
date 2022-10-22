package handler

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"github.com/HaleNing/Ning_JobBoard/src/database"
	"github.com/HaleNing/Ning_JobBoard/src/delegate"
	"github.com/HaleNing/Ning_JobBoard/src/param"
	"github.com/HaleNing/Ning_JobBoard/src/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

var baseCtx context.Context

func BaseAPi(app *fiber.App, ctx context.Context) {
	baseCtx = ctx
	app.Post("/login", loginHandler)
	app.Get("/set", setCookieHandler)
	app.Get("/logout", logoutHandler)
	app.Post("/createNewUser", createNewHandler)
}

// sign out
func logoutHandler(ctx *fiber.Ctx) error {
	cookieToken := ctx.Cookies("token")
	ctx.ClearCookie()

	result, err := database.RDB.Exists(baseCtx, cookieToken).Result()
	if err != nil {
		log.Println(err)
	}
	if result > 0 {
		log.Printf("del cookie:[%v]", cookieToken)
		database.RDB.Del(baseCtx, cookieToken)
	}

	return ctx.JSON("logout")
}

// success sign in to the system.
func setCookieHandler(ctx *fiber.Ctx) error {
	if utils.CheckLogin(ctx) {
		return fiber.NewError(fiber.StatusBadRequest, "you already  login")
	}
	token, _ := GenerateRandomString(32)
	////  store to redis
	result, err := database.RDB.Set(context.Background(), token, 1, 24*time.Hour).Result()
	log.Printf("store token:[%v] to redis", result)
	if err != nil {
		return err
	}
	ctx.Cookie(&fiber.Cookie{
		Secure:   true,
		HTTPOnly: true,
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
	})
	return ctx.JSON("set cookie~")
}

// in this handler, we will check user info and generate cookie
func loginHandler(ctx *fiber.Ctx) error {
	loginParam := new(param.UserLoginParam)
	// ensure that param field not empties
	if err := ctx.BodyParser(loginParam); err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, "login user msg is error")
	}
	//check username and password are existed or not.
	user, err := delegate.GetInfoByName(loginParam.UserName)
	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, "login fail,please check you account")
	}
	salt := user.Salt
	passwdHash := user.PasswdHash
	// remember, we can't equal two value to check password right or not.
	// the right way is to compare two value after hash.
	err = bcrypt.CompareHashAndPassword([]byte(passwdHash), []byte(loginParam.UserPass+salt))
	log.Printf("user:%v-salt:[%v]-compare:[%v]--passhash:[%v]", loginParam.UserName, salt, err, passwdHash)
	if err == nil {
		log.Println("login success")
		return ctx.JSON("login account success,we will set cookie later")
	} else {
		return ctx.JSON("fail login in, please check")
	}
}

// in this handler, we will create a new userinfo.
func createNewHandler(ctx *fiber.Ctx) error {
	loginParam := new(param.UserLoginParam)
	// ensure that param field not empties
	if err := ctx.BodyParser(loginParam); err != nil {
		log.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, "create user account is error")
	}
	exist := delegate.IsExistForUserName(loginParam.UserName)
	if exist {
		return ctx.JSON(param.Result{
			Code: fiber.StatusAccepted,
			Data: nil,
			Msg:  "user is already exist",
		})
	}
	salt, err := GenerateRandomString(len(loginParam.UserPass))
	if err != nil {
		log.Printf("generate salt fail %v", err)
		fiber.NewError(fiber.StatusOK, "you are not ready to go")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(loginParam.UserPass+salt), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("generate hashpasswd fail, reason:[%v]", err)
		fiber.NewError(fiber.StatusOK, "you are not ready to go")
	}
	info := delegate.CreateNewUserInfo(loginParam.UserName, salt, string(hash))
	if info != nil {
		return ctx.JSON("you are ready to go")
	}
	return fiber.NewError(fiber.StatusOK, "you are not ready to go")
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}
