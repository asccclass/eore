package main   //sherryTemplate

import(
   "os"
   "fmt"
   "strings"
   "github.com/gorilla/mux"
   "github.com/asccclass/sherrydb/mysql"
   "github.com/asccclass/sherryschema"
)

// 樣板資料
type Template struct {
   DbConnect	*SherryDB.DBConnect
   Fields	[]sherrySchema.DbSchema
   TmplStructure	*TemplateStructure
}

// 將字串第一個字母轉換為大寫
func(tpl *Template) TransFirstChar2Upper(s string)(string, error) {
   if len(s) == 0 {
      return "", fmt.Errorf("empty string")
   }
   var str strings.Builder
   str.WriteString(strings.ToUpper(s[0:1])) 
   str.WriteString(s[1:])
   return str.String(), nil
}

// ProcessSchema() 輸出資料庫結構
func(tpl *Template) ProcessSchema()(*TemplateStructure) {
   var str strings.Builder
   for _, elem := range tpl.Fields {
      name, _ := tpl.TransFirstChar2Upper(elem.Field)
      str.WriteString(name)
      str.WriteString("\t string `json:\"")
      str.WriteString(elem.Field)
      str.WriteString("\"`\n")
      // 判斷是否為PK
      if elem.Key == "PRI" {
         tpl.TmplStructure.PrimaryKey = elem.Field
         tpl.TmplStructure.UPrimaryKey, _ = tpl.TransFirstChar2Upper(elem.Field)
      }
      tpl.TmplStructure.Fields = append(tpl.TmplStructure.Fields, elem.Field)
      tpl.TmplStructure.UFields = append(tpl.TmplStructure.UFields, name)
      if elem.Field == "createDate" || elem.Field == "lastupdate" {
         tpl.TmplStructure.AutoDate = true
      }
   }
   tpl.TmplStructure.TableStructure = str.String()
   return tpl.TmplStructure
}

// 初始化
func NewTemplate(dbName, tableName string)(*Template, error) {
   dbconnect := SherryDB.DBConnect {
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
   tmplstructure := &TemplateStructure {
      TableName: tableName,
   }
   tpl := &Template{
      DbConnect: &dbconnect,
      Fields: fs,
      TmplStructure: tmplstructure,
   }
   return tpl, nil
}

func(c *Template) AddTemplateRouter(router *mux.Router) {
   // router.HandleFunc("/template", c.Read).Methods("GET")   	//取得單筆資料內容
}
