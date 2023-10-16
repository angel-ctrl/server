package middlewares

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	user_domain "server/domain"
	"server/services"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(User_DB services.UserPort) gin.HandlerFunc {
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

		// Buscar al usuario
		user, err := User_DB.GetUser(User)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar el usuario"})
			c.Abort()
			return
		}

		// Verificar si esta baneado
		if user.UserState == "ban" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user baned"})
			c.Abort()
			return
		}

		// Verificar la propiedad Plan
		if user.Plan.IsZero() {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No tiene un plan activo"})
			c.Abort()
			return
		}

		if user.Plan.Before(time.Now()) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "El plan ha caducado"})
			c.Abort()
			return
		}

		c.Request.Body = io.NopCloser(bytes.NewReader(bodyData))

		c.Next()
	}
}
