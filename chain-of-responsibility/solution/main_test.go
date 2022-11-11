package main

import (
	"testing"

	"github.com/arthurshafikov/patterns/chain-of-responsibility/solution/models"
	"github.com/arthurshafikov/patterns/chain-of-responsibility/solution/pages"
	"github.com/arthurshafikov/patterns/chain-of-responsibility/solution/router"
	"github.com/stretchr/testify/require"
)

var r *router.Router

func init() {
	r = router.NewRouter(router.Dependencies{
		AdminPage: pages.NewAdminPage(),
		IndexPage: pages.NewIndexPage(),
	})
}

// Admin Page

func TestShowAnAdminPageToAnAdmin(t *testing.T) {
	user := models.NewUser(true, false)

	response := r.ShowAdminPage(models.NewRequest(user))

	require.Equal(t, 200, response.Code)
}

func TestShowAnAdminPageToABannedUser(t *testing.T) {
	user := models.NewUser(false, true)

	response := r.ShowAdminPage(models.NewRequest(user))

	require.Equal(t, 401, response.Code)
}

func TestShowAnAdminPageToACommonUser(t *testing.T) {
	user := models.NewUser(false, false)

	response := r.ShowAdminPage(models.NewRequest(user))

	require.Equal(t, 403, response.Code)
}

func TestShowAnAdminPageToAnUnathorizedUser(t *testing.T) {
	response := r.ShowAdminPage(models.NewRequest(nil))

	require.Equal(t, 401, response.Code)
}

// Index Page

func TestShowAnIndexPageToAnAdmin(t *testing.T) {
	user := models.NewUser(true, false)

	response := r.ShowIndexPage(models.NewRequest(user))

	require.Equal(t, 200, response.Code)
}

func TestShowAnIndexPageToABannedUser(t *testing.T) {
	user := models.NewUser(false, true)

	response := r.ShowIndexPage(models.NewRequest(user))

	require.Equal(t, 401, response.Code)
}

func TestShowAnIndexPageToACommonUser(t *testing.T) {
	user := models.NewUser(false, false)

	response := r.ShowIndexPage(models.NewRequest(user))

	require.Equal(t, 200, response.Code)
}

func TestShowAnIndexPageToAnUnathorizedUser(t *testing.T) {
	response := r.ShowIndexPage(models.NewRequest(nil))

	require.Equal(t, 401, response.Code)
}
