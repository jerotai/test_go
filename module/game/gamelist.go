package game

import (
        //"github.com/labstack/echo"
        "test_go/core/model/database/system"
        "test_go/core/model/database/dto/site"
)

type GameList struct {
}

// new game list object
func NewGameList() *GameList {
        return &GameList{}
}

//game list object func
func (s *GameList) List() ([]*site.GameList) {
        return system.GameList()
}
