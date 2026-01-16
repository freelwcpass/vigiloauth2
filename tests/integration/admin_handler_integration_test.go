package integration

import (
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/vigiloauth/vigilo/v2/idp/config"
	"github.com/vigiloauth/vigilo/v2/internal/constants"
	"github.com/vigiloauth/vigilo/v2/internal/web"
)

func TestAdminHandler_GetAuditEvents(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		testContext := NewVigiloTestContext(t)
		defer testContext.TearDown()

		testContext.WithAdminToken(testUserID, time.Duration(2*time.Hour))
		testContext.WithAuditEvents()
		headers := map[string]string{"Authorization": constants.BearerAuthHeader + testContext.JWTToken}

		from := time.Now().Add(-5 * time.Hour).UTC().Format(time.RFC3339)
		to := time.Now().UTC().Format(time.RFC3339)

		queryParams := url.Values{}
		queryParams.Add("from", from)
		queryParams.Add("to", to)
		queryParams.Add("UserID", testUserID)
		queryParams.Add("EventType", "login_attempt")
		queryParams.Add("Success", "false")
		queryParams.Add("IP", testIP)
		queryParams.Add("limit", "50")
		queryParams.Add("offset", "0")
		endpoint := web.AdminEndpoints.GetAuditEvents + "?" + queryParams.Encode()

		rr := testContext.SendHTTPRequest(http.MethodGet, endpoint, nil, headers)

		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("Error is returned when user does not have the required roles", func(t *testing.T) {
		testContext := NewVigiloTestContext(t)
		defer testContext.TearDown()

		testContext.WithUser([]string{constants.UserRole})
		testContext.WithJWTToken(testUserID, time.Duration(2*time.Hour))
		testContext.WithAuditEvents()
		headers := map[string]string{"Authorization": constants.BearerAuthHeader + testContext.JWTToken}

		from := time.Now().Add(-5 * time.Hour).UTC().Format(time.RFC3339)
		to := time.Now().UTC().Format(time.RFC3339)

		queryParams := url.Values{}
		queryParams.Add("from", from)
		queryParams.Add("to", to)
		queryParams.Add("UserID", testUserID)
		queryParams.Add("EventType", "login_attempt")
		queryParams.Add("Success", "false")
		queryParams.Add("IP", testIP)
		queryParams.Add("limit", "50")
		queryParams.Add("offset", "0")
		endpoint := web.AdminEndpoints.GetAuditEvents + "?" + queryParams.Encode()

		rr := testContext.SendHTTPRequest(http.MethodGet, endpoint, nil, headers)

		assert.Equal(t, http.StatusForbidden, rr.Code)
	})
}

func TestAdminHandler_GetPasswordPolicy(t *testing.T) {
	t.Run("Success with all requirements", func(t *testing.T) {
		testContext := NewVigiloTestContext(t)
		defer testContext.TearDown()

		testContext.WithAdminToken(testUserID, time.Duration(2*time.Hour))
		headers := map[string]string{"Authorization": constants.BearerAuthHeader + testContext.JWTToken}
		endpoint := web.AdminEndpoints.GetPasswordPolicy

		// set policy to known values using public configuration API
		testContext.WithCustomConfig(
			config.WithPasswordConfig(
				config.NewPasswordConfig(
					config.WithUppercase(),
					config.WithNumber(),
					config.WithMinLength(12),
				),
			),
		)

		rr := testContext.SendHTTPRequest(http.MethodGet, endpoint, nil, headers)

		// Assert policy is returned
		assert.Contains(t, rr.Body.String(), `"require_upper":true`)
		assert.Contains(t, rr.Body.String(), `"require_number":true`)
		assert.Contains(t, rr.Body.String(), `"require_symbol":false`)
		assert.Contains(t, rr.Body.String(), `"min_length":12`)

		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("Success with some requirements", func(t *testing.T) {
		testContext := NewVigiloTestContext(t)
		defer testContext.TearDown()

		testContext.WithAdminToken(testUserID, time.Duration(2*time.Hour))
		headers := map[string]string{"Authorization": constants.BearerAuthHeader + testContext.JWTToken}
		endpoint := web.AdminEndpoints.GetPasswordPolicy

		// set policy to known values using public configuration API
		testContext.WithCustomConfig(
			config.WithPasswordConfig(
				config.NewPasswordConfig(
					config.WithUppercase(),
					config.WithSymbol(),
					config.WithMinLength(8),
				),
			),
		)

		rr := testContext.SendHTTPRequest(http.MethodGet, endpoint, nil, headers)

		// Assert policy is returned
		assert.Contains(t, rr.Body.String(), `"require_upper":true`)
		assert.Contains(t, rr.Body.String(), `"require_number":false`)
		assert.Contains(t, rr.Body.String(), `"require_symbol":true`)
		assert.Contains(t, rr.Body.String(), `"min_length":8`)

		assert.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("Error is returned when user does not have the required roles", func(t *testing.T) {
		testContext := NewVigiloTestContext(t)
		defer testContext.TearDown()

		testContext.WithUser([]string{constants.UserRole})
		testContext.WithJWTToken(testUserID, time.Duration(2*time.Hour))

		headers := map[string]string{"Authorization": constants.BearerAuthHeader + testContext.JWTToken}
		endpoint := web.AdminEndpoints.GetPasswordPolicy

		rr := testContext.SendHTTPRequest(http.MethodGet, endpoint, nil, headers)
		
		assert.Equal(t, http.StatusForbidden, rr.Code)
	})
}
