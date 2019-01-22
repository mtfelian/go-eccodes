package bufr

import codes "github.com/mtfelian/go-eccodes"

// Header ...
type Header struct {
	msg                 codes.Message
	Edition             int
	MasterTableNumber   int
	DataCategory        int
	DataSubCategory     int
	TypicalDate         int64
	TypicalTime         int64
	BufrHeaderCenter    int64
	BufrHeaderSubCenter int64
	MasterTablesVersion int
	LocalTablesVersion  int
	NumberOfSubsets     int64
	TotalLength         int64
}

// NewHeader ...
func NewHeader(msg codes.Message) *Header {
	return &Header{
		msg:                 msg,
		Edition:             GetInt(msg, "edition"),
		MasterTableNumber:   GetInt(msg, "masterTableNumber"),
		DataCategory:        GetInt(msg, "dataCategory"),
		DataSubCategory:     GetInt(msg, "dataSubCategory"),
		TypicalDate:         GetInt64(msg, "typicalDate"),
		TypicalTime:         GetInt64(msg, "typicalTime"),
		BufrHeaderCenter:    GetInt64(msg, "bufrHeaderCentre"),
		BufrHeaderSubCenter: GetInt64(msg, "bufrHeaderSubCentre"),
		MasterTablesVersion: GetInt(msg, "masterTablesVersionNumber"),
		LocalTablesVersion:  GetInt(msg, "localTablesVersionNumber"),
		NumberOfSubsets:     GetInt64(msg, "numberOfSubsets"),
		TotalLength:         GetInt64(msg, "totalLength"),
	}
}
