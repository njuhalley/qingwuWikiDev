package models

//var (
//	exportLimitWorkerChannel = gopool.NewChannelPool(conf.GetExportLimitNum(), conf.GetExportQueueLimitNum())
//)

type DocumentResult struct {
	DocumentId         int       `json:"document_id"`
	DocumentName       string    `json:"document_name"`
	Identify		   string    `json:"identify"`
	BookId			   int 		 `json:"book_id"`
}

func NewDocumentResult() *DocumentResult {
	return &DocumentResult{}
}


type DocumentStarResult struct {
	DocumentId         int       `json:"document_id"`
	DocumentName       string    `json:"document_name"`
	Identify		   string    `json:"identify"`
	BookId			   int 		 `json:"book_id"`
	ParentId		   int 	 	 `json:"parent_id"`
}

func NewDocumentStarResult() *DocumentStarResult {
	return &DocumentStarResult{}
}


type DocumentParentResult struct {
	DocumentId                  int          `json:"document_id"`
	DocumentName                string       `json:"document_name"`
	Identify		            string       `json:"identify"`
	BookId			            int 		 `json:"book_id"`
	FirstParentDocId            int 	 	 `json:"first_parent_doc_id"`
	FirstParentDocIdentify      string 	 	 `json:"first_parent_doc_identify"`
	FirstParentDocName	        string 	     `json:"first_parent_doc_name"`
	SecondParentDocId           int 	 	 `json:"second_parent_doc_id"`
	SecondParentDocIdentify     string 	 	 `json:"second_parent_doc_identify"`
	SecondParentDocName	        string 	     `json:"second_parent_doc_name"`
	ThirdParentDocId            int 	 	 `json:"third_parent_doc_id"`
	ThirdParentDocIdentify      string 	 	 `json:"third_parent_doc_identify"`
	ThirdParentDocName	        string 	     `json:"third_parent_doc_name"`
	ThirdParentPId              int 	 	 `json:"third_parent_p_id"`
}

func NewDocumentParentResult() *DocumentParentResult {
	return &DocumentParentResult{}
}