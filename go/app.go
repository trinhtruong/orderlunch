// AppObj.go

package swagger

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"

	// Some comment here
	_ "github.com/lib/pq"
)

// App app
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// TokenDetails store value token to authenticate
type TokenDetails struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	AtExpires    int64
	RtExpires    int64
}

// AppObj AppObj
var AppObj App

// Initialize Initialize
func Initialize(user, password, dbname string) {
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s host=10.1.110.64 sslmode=disable", user, password, dbname)

	var err error
	AppObj.DB, err = sql.Open("postgres", connectionString) //"user=postgres password=1 dbname=orderlunch sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	AppObj.Router = NewRouter()
}

// Run Run
func Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, AppObj.Router))
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	w.Write(response)
}

// CreateToken create a token for user
func CreateToken(userID int64) (*TokenDetails, error) {
	td := &TokenDetails{}
	// Set time live for access token is 15 minutes
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	// Set time live for refresh token is 7 days
	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()

	var err error
	// Creating Access Token
	os.Setenv("ACCESS_SECRET", "goodorbad") // Todo: set to an .env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userID
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	// Creating Refresh token
	os.Setenv("REFRESH_SECRET", "badorsad") // Todo: set to an .env file
	rtClaims := jwt.MapClaims{}
	rtClaims["user_id"] = userID
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	return td, nil
}

func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func verifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func tokenValid(r *http.Request) error {
	token, err := verifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

// TokenAuthMiddleware Check token valid
func TokenAuthMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := tokenValid(r)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}
		h.ServeHTTP(w, r)
	})
}
