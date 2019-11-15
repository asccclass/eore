package customer 
/*
   客戶資料管理＿備註
*/

import(
   "fmt"
   "net/http"
   "io/ioutil"
   "strconv"
   "encoding/json"
   // "database/sql"
   "github.com/gorilla/mux"
   "github.com/asccclass/sherrytime"
   "github.com/asccclass/sherrydb/mysql"
   // "github.com/asccclass/sherryschema"
)

// 客戶備註
type CustomerMeno struct {
   MenoID	string	`json:"menoID"`
   CustomerID	string	`json:"customerID"`
   Menoz	string	`json:"menoz"`
   UsrNo	string	`json:"usrNo"`
   Lastupdate	string	`json:"lastupdate"`
}

// 新增備註資料
func(c *Customer) CreateMeno(m CustomerMeno)(CustomerMeno, error) {
   st := sherrytime.NewSherryTime("Asia/Taipei", "-")
   m.Lastupdate = st.Today()
   sql := "insert into customermeno(customerID,menoz,usrNo,lastupdate) value(?,?,?,?)"
   conn, err := SherryDB.NewSherryDB(c.DBConfig)
   defer conn.Conn.Close()
   if err != nil {
      return m, err
   }
   ID, err := conn.Exec(sql,m.CustomerID,m.Menoz,0,m.Lastupdate)
   if err != nil  {
      return m, fmt.Errorf("Insert meno information Error(%s)", err.Error())
   }
   m.MenoID = strconv.FormatInt(ID.(int64), 10)
   return m, nil
}

// 取得備註資料
func(c *Customer) ReadMeno(id string)([]CustomerMeno, error) {
   var meno []CustomerMeno
   if id == "" {
     return meno, fmt.Errorf("Primary Key ID can not empty.")   
   }
   conn, err := SherryDB.NewSherryDB(c.DBConfig)
   defer conn.Conn.Close()
   if err != nil {
      return meno, err
   }
   rows, err := conn.Conn.Query("select * from customermeno where customerID=?", id)
   if err != nil {
      if err.Error() == "sql: no rows in result set" {
         return meno, nil
      } else {
         return meno, err
      }
   }
   defer rows.Close()
   for rows.Next() {
      var m CustomerMeno
      if err := rows.Scan(&m.MenoID,&m.CustomerID,&m.Menoz,&m.UsrNo,&m.Lastupdate); err != nil {
         return meno, err
      }
      meno = append(meno, m)
   }
   return meno, nil
}

// Update 備註資料
func(c *Customer) UpdateMeno(parentID string, m []CustomerMeno) ([]CustomerMeno, error) {
   var newMeno []CustomerMeno
   conn, err := SherryDB.NewSherryDB(c.DBConfig)
   defer conn.Conn.Close()
   if err != nil {
      return newMeno, err
   }
   for _, elem := range m {
      if elem.MenoID == "" {   // 無資料,新增
         elem.CustomerID = parentID
         n, err := c.CreateMeno(elem)
         if err != nil {
            return newMeno, err
         }
         newMeno = append(newMeno, n)
      } else {		// 更新現有資料
         st := sherrytime.NewSherryTime("Asia/Taipei", "-")
         sql := "update customermeno set customerID=?, menoz=?,lastupdate=? where menoID=?"
         _, err := conn.Exec(sql, parentID, elem.Menoz, st.Today(), elem.MenoID)
         if err != nil {
            return newMeno, err
         }
         newMeno = append(newMeno, elem)
      }
   }
   return newMeno, nil
}

// Delete 備註資料
func(c *Customer) DeleteMeno(menoID string) (error) {
   conn, err := SherryDB.NewSherryDB(c.DBConfig)
   if err != nil {
      return fmt.Errorf("Can not connect databaser(%s).", err.Error())
   }
   _, err = conn.Exec("delete from customermeno where menoID=?", menoID)
   if err != nil {
      fmt.Errorf("Delete meno Error(%s).", err.Error())
   }
   return nil
}

// DELETE 備註（） 
func(c *Customer) DeleteMenoFromWeb(w http.ResponseWriter, r *http.Request) {
   w.Header().Set("Content-Type", "application/json;charset=UTF-8")
   w.WriteHeader(http.StatusOK)
   b, err := ioutil.ReadAll(r.Body)
   if err != nil {
      c.Error2Web(w, err)
      return
   }
   meno := CustomerMeno{}
   if !json.Valid(b)  {
      c.Error2Web(w, fmt.Errorf("JSON data is invalid.") )
      return
   }
   defer r.Body.Close()
   if err := json.Unmarshal(b, &meno); err != nil {
      c.Error2Web(w, err)
      return
   }
   urlParams := mux.Vars(r)
   if urlParams["menoID"] != meno.MenoID {
      c.Error2Web(w, fmt.Errorf("Params Error."))
      return
   }
   if err = c.DeleteMeno(meno.MenoID); err != nil {
      c.Error2Web(w, err)
      return
   }
   fmt.Fprintf(w, "{\"Status\": \"200\"}")
}
