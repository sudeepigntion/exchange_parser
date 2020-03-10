# exchange_parser
Exchange Parser

This repository is all about parsing byte array to hashmap and vice versa.
Why one should use byte array rather than json. There are many reasons behind that

1) Small Packet
2) Hackers cannot easily decode unless they have the schema
3) Used mainly by exchanges.

There are lots of exchanges who uses FIX/FAST protocols for encoding and decoding the packets but there are also exchanges who uses byte array as they schema for transmission of data. Which exchanges uses byte array well they are India's biggest stock exchange like NSE (National Stock Exchange) and BSE (Bombay Stock Exchange).

    BSE : https://www.bseindia.com/downloads1/ETI_API_Manual_1.4.8.pdf
    NSE/CM: https://www1.nseindia.com/technology/content/nnf/TP_CM_Trimmed_NNF_PROTOCOL_4.1.pdf
    NSE/FO: https://www1.nseindia.com/technology/content/nnf/TP_FO_Trimmed_NNF_PROTOCOL_9.25.pdf
    NSE/CD: https://www1.nseindia.com/technology/content/nnf/TP_COM_Trimmed_NNF_PROTOCOL_1.5.pdf

Using this parser you can easily convert byte array to hashmap and hashmap to byte array.

    package main

    import(
      "fmt"
        "byte_parser"
    )	

    func main(){
    
      // creating hashmap
      
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
    
      // reading the schema of the byte array
      
      var byteParser = &byte_parser.ByteParser{
        ParserType:"json",
        Path:"D:\\pounze_go_project\\crypto_backend_engine\\src\\templates\\request_templates\\json",
        EncodingType:"bigendian",
      }
      
      // loading the schema
      
      byteParser.ParseJson()
    
      // converting hashmap to byte array
      
      // "28500" is the file name unique id to know the schema
      
      var byteArray = byteParser.ParseJsonToByte(jsonMap, "28500")

      fmt.Println(byteArray)
    
      // converting byte array to hashmap
      
      var obj =  byteParser.ParseToObject(byteArray, "28500")

      fmt.Println(obj)

    }
