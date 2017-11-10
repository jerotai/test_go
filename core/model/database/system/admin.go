package system

import (
        "test_go/core/model/database/dto/site"
        "test_go/core/model/database"
        "fmt"
)

func AdminList() (ret []*site.AdminList, err error) {
        //var params []interface{}
        //params = append(params, GameID)

        test_conn := database.GetConn("c4_site")
        rows, err := test_conn.Query("SELECT `id`,`name`, `office_id` FROM admin")


        defer rows.Close()
        for rows.Next() {
                data := &site.AdminList{}
                err := rows.Scan(&data.Id, &data.Name, &data.OfficeId)
                if err != nil {
                        fmt.Print(err)
                }
                //fmt.Print(data)
                ret = append(ret, data)
        }
        fmt.Print(ret[0])
        return
}