package main   //sherryTemplate

// 轉換樣板用參數
type TemplateStructure struct {
   TableName		string		// 表格名稱
   PrimaryKey		string
   UPrimaryKey		string		// 第一個字母大寫的Primary Key
   TableStructure	string
   Fields		[]string	// 欄位名稱
   UFields		[]string	// 轉換後第一個字母為大寫的欄位名稱
   AutoDate		bool		// 是否有createDate, lastupdate
}
