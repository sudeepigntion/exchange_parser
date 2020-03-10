package main

import(
	"fmt"
    "byte_parser"
)	

func main(){

    var jsonMap = make(map[string]interface{})

    jsonMap["BodyLen"] = 176
    jsonMap["TemplateID"] = 28500
    jsonMap["NetworkMsgID"] = ""
    jsonMap["Pad2"] = ""
    jsonMap["MsgSeqNum"] = 1
    jsonMap["SenderSubID"] = 123456
    jsonMap["Price"] = 850000000
    jsonMap["StopPx"] = 850000000
    jsonMap["MaxPricePercentage"] = 850000000
    jsonMap["SenderLocationID"] = 21212121
    jsonMap["ClOrdID"] = 43566763434
    jsonMap["MessageTag"] = 2222
    jsonMap["OrderQty"] = 2
    jsonMap["MaxShow"] = 5
    jsonMap["ExpireDate"] = 1583693033
    jsonMap["MarketSegmentID"] = 33232
    jsonMap["SimpleSecurityID"] = 545454
    jsonMap["PartyIDTakeUpTradingFirm"] = "TEST1"
    jsonMap["PartyIDOrderOriginationFirm"] = "TET"
    jsonMap["PartyIDBeneficiary"] = "TTETST"
    jsonMap["AccountType"] = 1
    jsonMap["ApplSeqIndicator"] = 2
    jsonMap["Side"] = 3
    jsonMap["OrdType"] = 2.0
    jsonMap["PriceValidityCheckType"] = 3
    jsonMap["TimeInForce"] = 3
    jsonMap["ExecInst"] = 1
    jsonMap["TradingSessionSubID"] = 2
    jsonMap["TradingCapacity"] = 4
    jsonMap["Account"] = "IN"
    jsonMap["PositionEffect"] = "Y"
    jsonMap["RegulatoryText"] = "79LABS"
    jsonMap["FreeText1"] = "79LABS"
    jsonMap["FreeText2"] = "79LABS"
    jsonMap["FreeText3"] = "79LABS"
    jsonMap["Pad3"] = "LAB"

    fmt.Println(jsonMap)

    var byteParser = &byte_parser.ByteParser{
    	ParserType:"json",
    	Path:"D:\\pounze_go_project\\crypto_backend_engine\\src\\templates\\request_templates\\json",
    	EncodingType:"bigendian",
    }

    byteParser.ParseJson()

    var byteArray = byteParser.ParseJsonToByte(jsonMap, "28500")

    fmt.Println(byteArray)

    var obj =  byteParser.ParseToObject(byteArray, "28500")

    fmt.Println(obj)

}