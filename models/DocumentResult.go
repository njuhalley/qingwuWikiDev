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

