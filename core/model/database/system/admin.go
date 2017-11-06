package system

import (
        "test_go/core/model/database/dto/site"
        "test_go/core/model/database"
        "fmt"
)

func AdminList() (ret []*site.BackCommisionSet, err error) {
        //var params []interface{}
        //params = append(params, GameID)

        test_conn := database.GetConn("c4_site")
        rows, err := test_conn.Query("SELECT `id`,`name`,`code` FROM back_commision_set")


        defer rows.Close()
        //var name []byte
        for rows.Next() {
                data := &site.BackCommisionSet{}
                err := rows.Scan(&data.Id, &data.Name, &data.Code)
                if err != nil {
                        fmt.Print(err)
                }
                //fmt.Print(data)
                ret = append(ret, data)
        }
        fmt.Print(ret[0])
        return
}