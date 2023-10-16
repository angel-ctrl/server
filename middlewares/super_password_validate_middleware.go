package middlewares

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	configs "server/config"
	user_domain "server/domain"

	"github.com/gin-gonic/gin"
)

func SuperPassValidateMiddleware(Env *configs.Env) gin.HandlerFunc {
	return func(c *gin.Context) {

		bodyCopy := new(bytes.Buffer)
		// Read the whole body
		_, err := io.Copy(bodyCopy, c.Request.Body)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "error en los datos recibidos parte 2"})
			return
		}
		bodyData := bodyCopy.Bytes()

		User := &user_domain.Users{}
		err = json.Unmarshal(bodyData, User)
		if err != nil {
			// Devuelve una respuesta de error si hay un error al deserializar los datos
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "error en los datos recibidos parte 1"})
			return
		}

		if User.Secret != Env.Secret {
			c.JSON(http.StatusUnauthorized, gin.H{"error: ": "codigo de autorizacion invalido"})
			return
		}

		c.Request.Body = io.NopCloser(bytes.NewReader(bodyData))

		c.Next()
	}
}
