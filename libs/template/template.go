package sherryTemplate

import(
   "os"
   "fmt"
   "github.com/gorilla/mux"
   "github.com/asccclass/sherrydb/mysql"
   "github.com/asccclass/sherryschema"
)

type Template struct {
   DbConnect	*SherryDB.DBConnect
   Fields	[]DbSchema
}


func NewTemplate(dbName, tableName string)(*Template, error) {
   dbconnect := &SherryDB.DBConnect {
      DBMS: os.Getenv("DBMS"),
      DbServer: os.Getenv("DBSERVER"),
      DbPort: os.Getenv("DBPORT"),
      DbName: os.Getenv("DBNAME"),
      DbLogin: os.Getenv("DBLOGIN"),
      DbPasswd: os.Getenv("DBPASSWORD"),
   }

   schema, err := sherrySchema.NewSherrySchema(dbconnect)
   if err != nil {
      return nil, err
   } 
   schema.Database(dbName)
   fs, err := schema.GetColumns(tableName)
   if err != nil {
      return nil, err
   }
   schema.Conn.Disconnect()
   tpl := &Template{
      DbConnect: dbconnect,
      Fields: fs,
   }
   return tpl, nil
}

func(c *Template) AddTemplateRouter(router *mux.Router) {
   // router.HandleFunc("/template", c.Read).Methods("GET")   	//取得單筆資料內容
}
