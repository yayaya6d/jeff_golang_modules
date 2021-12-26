package gin_helper

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type RestfulRouterSuite struct {
	suite.Suite
}

func (s *RestfulRouterSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)
}

func TestRestfulRouterSuite(t *testing.T) {
	suite.Run(t, new(RestfulRouterSuite))
}

func testRouter() Router {
	return NewRouter(
		"/test",
		[]gin.HandlerFunc{
			testMiddleware(),
		},
		[]Router{
			NewRouter(
				"/add_trace_id",
				[]gin.HandlerFunc{},
				[]Router{},
				Handler{
					GET: func(ctx *gin.Context) {
						ctx.String(http.StatusOK, ctx.GetHeader("trace_id"))
					},
				},
			),
		},
		Handler{},
	)
}

func testMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Header.Get("trace_id") == "" {
			traceID := "test_trace_id"
			ctx.Request.Header.Set("trace_id", traceID)
		}
		ctx.Next()
	}
}

func (s *RestfulRouterSuite) TestNewRouter_RunNewRouterAndSendRequest_ReturnExpectedResponse() {
	// arrange
	expectedResponse := "test_trace_id"
	request, _ := http.NewRequest(http.MethodGet, "/test/add_trace_id", nil)
	responseRecorder := httptest.NewRecorder()
	router := testRouter()
	ctx, engine := gin.CreateTestContext(responseRecorder)
	router.SetRouter(engine.Group("/"))

	// act
	ctx.Request = request
	engine.HandleContext(ctx)

	// assert
	s.Equal(http.StatusOK, responseRecorder.Code)
	s.Equal(expectedResponse, responseRecorder.Body.String())
}
