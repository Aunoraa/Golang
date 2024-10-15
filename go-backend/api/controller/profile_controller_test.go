package controller_test

import (
	"encoding/json"
	"errors"
	domain2 "github.com/amitshekhariitbhu/go-backend-clean-architecture/internal/domain"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/internal/domain/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/amitshekhariitbhu/go-backend-clean-architecture/api/controller"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func setUserID(userID string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("x-user-id", userID)
		c.Next()
	}
}

func TestFetch(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		mockProfile := &domain2.Profile{
			Name:  "Test Name",
			Email: "test@gmail.com",
		}

		userObjectID := primitive.NewObjectID()
		userID := userObjectID.Hex()

		mockProfileUsecase := new(mocks.ProfileUsecase)

		mockProfileUsecase.On("GetProfileByID", mock.Anything, userID).Return(mockProfile, nil)

		gin := gin.Default()

		rec := httptest.NewRecorder()

		pc := &controller.ProfileController{
			ProfileUsecase: mockProfileUsecase,
		}

		gin.Use(setUserID(userID))
		gin.GET("/profile", pc.Fetch)

		body, err := json.Marshal(mockProfile)
		assert.NoError(t, err)

		bodyString := string(body)

		req := httptest.NewRequest(http.MethodGet, "/profile", nil)
		gin.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		assert.Equal(t, bodyString, rec.Body.String())

		mockProfileUsecase.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		userObjectID := primitive.NewObjectID()
		userID := userObjectID.Hex()

		mockProfileUsecase := new(mocks.ProfileUsecase)

		customErr := errors.New("Unexpected")

		mockProfileUsecase.On("GetProfileByID", mock.Anything, userID).Return(nil, customErr)

		gin := gin.Default()

		rec := httptest.NewRecorder()

		pc := &controller.ProfileController{
			ProfileUsecase: mockProfileUsecase,
		}

		gin.Use(setUserID(userID))
		gin.GET("/profile", pc.Fetch)

		body, err := json.Marshal(domain2.ErrorResponse{Message: customErr.Error()})
		assert.NoError(t, err)

		bodyString := string(body)

		req := httptest.NewRequest(http.MethodGet, "/profile", nil)
		gin.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		assert.Equal(t, bodyString, rec.Body.String())

		mockProfileUsecase.AssertExpectations(t)
	})

}
