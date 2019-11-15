package takeinlistobj

import(
   "os"
   "fmt"
   "net/http"
   "encoding/json"
   "io/ioutil"
   "github.com/gorilla/mux"
   "github.com/asccclass/sherrydb/mysql"
)

// Error2Web 輸出錯誤訊息至web client
func(c *Ntakeinlistobj) Error2Web(w http.ResponseWriter, err error) {
   w.Header().Set("Content-Type", "application/json;charset=UTF-8")
   w.WriteHeader(http.StatusOK)
   fmt.Fprintf(w, "{\"errMsg\": \"%s(server)\"}", err.Error())
}

// GetDataFromWeb 從web取得資料
func(c *Ntakeinlistobj) ParseDataFromWeb(w http.ResponseWriter, r *http.Request)([]byte) {
   b, err := ioutil.ReadAll(r.Body)
   if err != nil {
      c.Error2Web(w, err)
      return b
   }
   if !json.Valid(b)  {
      c.Error2Web(w, fmt.Errorf("JSON data is invalid.") )
      return b
   }
   defer r.Body.Close()
   return b
}

// CreateFromWeb 新增資料
func(c *Ntakeinlistobj)CreateFromWeb(w http.ResponseWriter, r *http.Request) {
   // 取得web資料
   b := c.ParseDataFromWeb(w, r)
   tbl := Xtakeinlistobj{}
   if err := json.Unmarshal(b, &tbl); err != nil {
      c.Error2Web(w, err)
      return
   }
   tbl, err := c.Create(tbl)
   if err != nil {
      c.Error2Web(w, err)
      return
   }
   jstr, err := json.Marshal(tbl)
   if err != nil {
      c.Error2Web(w, err)
      return
   }
   w.Header().Set("Content-Type", "application/json;charset=UTF-8")
   w.WriteHeader(http.StatusOK)
   fmt.Fprintf(w, string(jstr))
}

// ReadFromWeb 讀取資料
func(c *Ntakeinlistobj) ReadFromWeb(w http.ResponseWriter, r *http.Request) {
   urlParams := mux.Vars(r)
   var id string = ""
   if urlParams["objID"] != "" {
      id = urlParams["objID"]
   }
   tbls, err := c.Read(id)
   jstr, err := json.Marshal(tbls)
   if err != nil {
      c.Error2Web(w, err)
      return
   }

   w.Header().Set("Content-Type", "application/json;charset=UTF-8")
   w.WriteHeader(http.StatusOK)
   if string(jstr) == "" {
      fmt.Fprintf(w, "[{}]")
      return
   }
   w.Write(jstr)
}

// UpdateFromWeb 修改資料
func(c *Ntakeinlistobj) UpdateFromWeb(w http.ResponseWriter, r *http.Request) {
   // 取得web資料
   b := c.ParseDataFromWeb(w, r)
   tbl := Xtakeinlistobj{}
   if err := json.Unmarshal(b, &tbl); err != nil {
      c.Error2Web(w, err)
      return
   }
   urlParams := mux.Vars(r)
   if urlParams["objID"] != tbl.ObjID {
      c.Error2Web(w, fmt.Errorf("params error"))
      return
   }
   tbls := []Xtakeinlistobj{}
   tbls = append(tbls, tbl)
   tbls, err := c.Update(tbls)
   if err != nil {
      c.Error2Web(w, err)
      return
   }

   // 輸出 json
   response, err := json.Marshal(tbls[0])
   if err != nil {
      c.Error2Web(w, err)
      return
   }
   w.Header().Set("Content-Type", "application/json;charset=UTF-8")
   w.WriteHeader(http.StatusOK)
   fmt.Fprintf(w, string(response))
}


// DELETE takeinlistobj from web
func(c *Ntakeinlistobj) DeleteFromWeb(w http.ResponseWriter, r *http.Request) {
   // 取得web資料
   b := c.ParseDataFromWeb(w, r)
   tbl := Xtakeinlistobj{}
   if err := json.Unmarshal(b, &tbl); err != nil {
      c.Error2Web(w, err)
      return
   }

   urlParams := mux.Vars(r)
   if urlParams["objID"] != tbl.ObjID {
      c.Error2Web(w, fmt.Errorf("Params Error."))
      return
   }
   if err := c.Delete(tbl.ObjID); err != nil {
      c.Error2Web(w, err)
      return
   }
   w.Header().Set("Content-Type", "application/json;charset=UTF-8")
   w.WriteHeader(http.StatusOK)
   fmt.Fprintf(w, "{\"Status\": \"200\"}")
}

func(c *Ntakeinlistobj) AddtakeinlistobjRouter(router *mux.Router) {
   router.HandleFunc("/takeinlistobj", c.CreateFromWeb).Methods("POST")       	            //C:新增資料
   router.HandleFunc("/takeinlistobj", c.ReadFromWeb).Methods("GET")    	                  //R:取得收件單s
   router.HandleFunc("/takeinlistobj/{ objID }", c.ReadFromWeb).Methods("GET")           //R:取得收件單
   router.HandleFunc("/takeinlistobj/{ objID }", c.UpdateFromWeb).Methods("PUT")         //U:修改單筆資料
   router.HandleFunc("/takeinlistobj/{ objID }", c.DeleteFromWeb).Methods("DELETE")      //D:刪除單筆資料
}

func NewNtakeinlistobj()(*Ntakeinlistobj, error) {
   dbconnect := SherryDB.DBConnect {
      DBMS: os.Getenv("DBMS"),
      DbServer: os.Getenv("DBSERVER"),
      DbPort: os.Getenv("DBPORT"),
      DbName: os.Getenv("DBNAME"),
      DbLogin: os.Getenv("DBLOGIN"),
      DbPasswd: os.Getenv("DBPASSWORD"),
   }

   return &Ntakeinlistobj {
      DBConfig: dbconnect,
   }, nil
}
