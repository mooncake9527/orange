package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/mooncake9527/orange-core/config"
)

func InitMiddleware(r *gin.Engine, cfg *config.AppCfg) {
	r.Use(ReqId)
	r.Use(AppKey)
	r.Use(CustomError)
	r.Use(LoggerToFile())
	if cfg.AccessLimit.Enable {
		r.Use(AccessLimit())
	}
	if cfg.Cors.Enable {
		r.Use(CorsByRules(&cfg.Cors))
	}
}
