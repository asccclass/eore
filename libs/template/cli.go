package main

import(
   "os"
   "fmt"
   "text/template"
)

func main() {
   dbName := "laitaian"
   table := "takeinlistobj"
   tpl, err := NewTemplate(dbName, table)
   if err != nil {
      fmt.Println(err)
      return 
   }
   // 判斷目錄是否存在
   dir := "./" + dbName + "/" + table
   if _, err := os.Stat(dir); os.IsNotExist(err) {
      if err = os.MkdirAll(dir, 0755); err != nil {
         panic(err)
      }
   }
   f, err := os.Create(dir + "/" + table + ".go")
   if err != nil {
      fmt.Errorf("Create %s error:%s\n", dir, err.Error())
      return
   }
   tmpl := template.Must(template.ParseFiles("./crud.tmpl"))
   params := tpl.ProcessSchema()
   tmpl.Execute(f, params)
   // tmpl.Execute(os.Stdout, params)	// write to command line
   // tmpl.Execute(w, nil)		// write to web
   f.Close()

   f, err = os.Create(dir + "/" + table + "web.go")
   if err != nil {
      fmt.Errorf("Create %s error:%s\n", dir, err.Error())
      return
   }
   tmpl = template.Must(template.ParseFiles("./webinterface.tmpl"))
   params = tpl.ProcessSchema()
   tmpl.Execute(f, params)
}
