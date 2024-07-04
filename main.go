package teeamai

import (
	"github.com/tee-am-ai/backend/routes"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("WebHook", routes.URL)
}