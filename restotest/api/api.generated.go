package api

//
// ATTENTION: This file is generated automagically.
// Do not touch it. Do not pass go. Do not collect $200.
// Instead run 'go generate' or 'make gen' to build this file.
//

import (
	"github.com/buaazp/fasthttprouter"
	mw "github.com/kayteh/restokit/middleware"
)

// FetchAPIRoutes is a **generated** function that takes a router, and injects the api routes into it.
func FetchAPIRoutes(router *fasthttprouter.Router) {
	mw.Noop()

	router.Handle("GET", "/hello/:name", mw.NoLogging(hello))

	router.Handle("GET", "/json", mw.JSON(jsonTest))

	router.Handle("GET", "/localmw", mw.JSON(localMw(getlocalmw)),
	)

	router.Handle("GET", "/localmw2", mw.NoLogging(mw.JSON(getlocalmw2)),
	)

	router.Handle("GET", "/test",
		mw.VersionedRoute(mw.VersionedRouteMap{
			"default": testGet,

			"v1": testGetv1,

			"v2": testGet,
		}),
	)

}
