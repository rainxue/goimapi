package core

type Conversation struct {
	Id uint32	`json:"id"`
	ConvType string `json:"conv_type"`
	Title string `json:"title"`
	// members []uint32
	// ext map
	CreateBy uint32 `json:"create_by"`
	// create_date 
	
}

