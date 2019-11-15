package takeinlistobj
// takeinlistobj.go

import(
   "fmt"
   "strings"
   "strconv"
   "github.com/asccclass/sherrytime"
   "github.com/asccclass/sherrydb/mysql"
)

// Table Structure 
type Xtakeinlistobj struct {
ObjID	 string `json:"objID"`
Name	 string `json:"name"`
Ename	 string `json:"ename"`
Valuez	 string `json:"valuez"`
Lastupdate	 string `json:"lastupdate"`
}  

type Ntakeinlistobj struct {
   DBConfig SherryDB.DBConnect
}

// 新增takeinlistobj資料
func(c *Ntakeinlistobj) Create(m Xtakeinlistobj)(Xtakeinlistobj, error) {
   // 設定PrimaryKey為空值
   m.ObjID = ""

   // 自動產生日期
   st := sherrytime.NewSherryTime("Asia/Taipei", "-")
   
   m.Lastupdate = st.Today()

   var sql strings.Builder
   sql.WriteString("insert into Xtakeinlistobj(objID,name,ename,valuez,lastupdate) ")
   sql.WriteString("values(?,?,?,?,?)")
   conn, err := SherryDB.NewSherryDB(c.DBConfig)
   defer conn.Conn.Close()
   if err != nil {
      return m, err
   }
   ID, err := conn.Exec(sql.String(), m.ObjID,m.Name,m.Ename,m.Valuez,m.Lastupdate)
   if err != nil  {
      return m, fmt.Errorf("Insert meno information Error(%s)", err.Error())
   }
   m.ObjID = strconv.FormatInt(ID.(int64), 10)
   return m, nil
}

// 取得takeinlistobj資料
func(c *Ntakeinlistobj) Read(id string)([]Xtakeinlistobj, error) {
   var Tx []Xtakeinlistobj
   if id == "" {
     return Tx, fmt.Errorf("Primary Key ID can not empty.")   
   }
   conn, err := SherryDB.NewSherryDB(c.DBConfig)
   defer conn.Conn.Close()
   if err != nil {
      return Tx, err
   }
   var sql strings.Builder
   sql.WriteString("select objID,name,ename,valuez,lastupdate from Xtakeinlistobj ")
   if id != "" {
       sql.WriteString("where objID=?")
   }
   rows, err := conn.Conn.Query(sql.String(), id)
   if err != nil {
      if err.Error() == "sql: no rows in result set" {
         return Tx, nil
      } else {
         return Tx, err
      }
   }
   defer rows.Close()
   for rows.Next() {
      var m Xtakeinlistobj
      if err := rows.Scan(&m.ObjID,&m.Name,&m.Ename,&m.Valuez,&m.Lastupdate); err != nil {
         return Tx, err
      }
      Tx = append(Tx, m)
   }
   return Tx, nil
}

// 更新takeinlistobj資料
// Update 備註資料
func(c *Ntakeinlistobj) Update(m []Xtakeinlistobj) ([]Xtakeinlistobj, error) {
   var newT []Xtakeinlistobj
   conn, err := SherryDB.NewSherryDB(c.DBConfig)
   defer conn.Conn.Close()
   if err != nil {
      return newT, err
   }
   // 自動產生日期
   st := sherrytime.NewSherryTime("Asia/Taipei", "-")

   for _, elem := range m {
      if elem.ObjID == "" {   // 無Key資料,新增
         n, err := c.Create(elem)
         if err != nil {
            return newT, err
         }
         newT = append(newT, n)
      } else {		// 更新現有資料
         elem.Lastupdate = st.Today()
         sql := "update takeinlistobj set objID=?,name=?,ename=?,valuez=?,lastupdate=? where objID=?"
         _, err := conn.Exec(sql,elem.ObjID,elem.Name,elem.Ename,elem.Valuez,elem.Lastupdate, elem.ObjID)
         if err != nil {
            return newT, err
         }
         newT = append(newT, elem)
      }
   }
   return newT, nil
}

// Delete takeinlistobj
func(c *Ntakeinlistobj) Delete(ID string) (error) {
   conn, err := SherryDB.NewSherryDB(c.DBConfig)
   if err != nil {
      return fmt.Errorf("Can not connect databaser(%s).", err.Error())
   }
   _, err = conn.Exec("delete from takeinlistobj where objID=?", ID)
   if err != nil {
      fmt.Errorf("Delete meno Error(%s).", err.Error())
   }
   return nil
}
