package game

import (
        //"github.com/labstack/echo"
)

type GameList struct {
}

// new game list object
func NewGameList() *GameList {
        return &GameList{}
}

//game list object func
func (s *GameList) List() string {
        return "sa"
}
