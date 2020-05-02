package utility

import (
	"github.com/benammann/drinkspot-core/api/memory"
	"github.com/gin-gonic/gin"
)

// sets the memory_environment to the given gin context
func SetEnvironment(ctx *gin.Context, env *memory.Environment) {
	ctx.Set("memory_environment", env)
}

// extracts the current memory_environment from the given gin context
func GetEnvironment(ctx *gin.Context) *memory.Environment {

	env, _ := ctx.Get("memory_environment")

	return env.(*memory.Environment)

}
