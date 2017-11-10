package system

import (
        "test_go/core/model/database/dto/site"
        "test_go/core/model/database"
        "fmt"
)

func GameList() (ret []*site.GameList) {
        test_conn := database.GetConn("sys_backend")
        rows, qu_err := test_conn.Query("SELECT `id`,`game_item` FROM game_list")

        if qu_err != nil {
                fmt.Print(qu_err)
        }

        defer rows.Close()
        for rows.Next() {
                data := &site.GameList{}
                err := rows.Scan(&data.Id, &data.GameItem)
                if err != nil {
                        fmt.Print(err)
                }
                ret = append(ret, data)
        }

        return ret
}
