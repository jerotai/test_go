package game

import (
        "github.com/labstack/echo"
        "net/http"

        "test_go/module/game"
)

type Game struct {
}

// new game list api object
func NewApiGame(c echo.Context) error {
        game_obj := game.NewGameList().List()

        return c.JSON(http.StatusOK, game_obj)
}
