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
		Edition:             getInt(msg, "edition"),
		MasterTableNumber:   getInt(msg, "masterTableNumber"),
		DataCategory:        getInt(msg, "dataCategory"),
		DataSubCategory:     getInt(msg, "dataSubCategory"),
		TypicalDate:         getInt64(msg, "typicalDate"),
		TypicalTime:         getInt64(msg, "typicalTime"),
		BufrHeaderCenter:    getInt64(msg, "bufrHeaderCentre"),
		BufrHeaderSubCenter: getInt64(msg, "bufrHeaderSubCentre"),
		MasterTablesVersion: getInt(msg, "masterTablesVersionNumber"),
		LocalTablesVersion:  getInt(msg, "localTablesVersionNumber"),
		NumberOfSubsets:     getInt64(msg, "numberOfSubsets"),
		TotalLength:         getInt64(msg, "totalLength"),
	}
}
