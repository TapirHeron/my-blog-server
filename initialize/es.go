package initialize

import (
	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8"
	"go.uber.org/zap"
	"os"
	"server/global"
)

func ConnectES() *elasticsearch.TypedClient {
	esConfig := global.Config.ES
	config := elasticsearch.Config{
		Addresses: []string{esConfig.URL},
		Username:  esConfig.Username,
		Password:  esConfig.Password,
	}
	if esConfig.IsConsolePrint {
		config.Logger = &elastictransport.ColorLogger{
			Output:             os.Stdout,
			EnableRequestBody:  true,
			EnableResponseBody: true,
		}
	}

	client, err := elasticsearch.NewTypedClient(config)
	if err != nil {
		global.Log.Error("Failed to connect to Elasticsearch", zap.Error(err))
		os.Exit(1)
	}
	global.Log.Info("Connected to Elasticsearch")
	return client
}
