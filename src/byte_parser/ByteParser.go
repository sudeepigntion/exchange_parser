package byte_parser

import(
	"fmt"
    "io/ioutil"
    "path/filepath"
    "os"
    "encoding/json"
    "ByteBuffer"
    "reflect"
    "strings"
    "encoding/xml"
)

type ByteParser struct{
	ParserType string
	Path string
	EncodingType string
	Template map[string]interface{}
}

func (bb *ByteParser) ParseJson(){

	files, err := ioutil.ReadDir(bb.Path)

    if err != nil {

        panic(err)

        return
    }

    bb.Template = make(map[string]interface{})

    for _, f := range files {

        var fileName = f.Name()

        var extension = filepath.Ext(fileName)

		var name = fileName[0:len(fileName)-len(extension)]

		jsonFile, err := os.Open(bb.Path+"/"+fileName)

		byteValue, _ := ioutil.ReadAll(jsonFile)

		if err != nil {

		    panic(err)

		    return

		}

		defer jsonFile.Close()

		var body = make(map[string]interface{})

		json.Unmarshal(byteValue, &body)

		bb.Template[name] = body

		fmt.Println("Template "+name+" loaded")

    }
}

func (bb *ByteParser) ParseXml(){

	files, err := ioutil.ReadDir(bb.Path)

    if err != nil {

        panic(err)

        return
    }

    bb.Template = make(map[string]interface{})

    for _, f := range files {

        var fileName = f.Name()

        var extension = filepath.Ext(fileName)

		var name = fileName[0:len(fileName)-len(extension)]

		xmlFile, err := os.Open(bb.Path+"/"+fileName)

		byteValue, _ := ioutil.ReadAll(xmlFile)

		if err != nil {

		    panic(err)

		    return

		}

		defer xmlFile.Close()

		var body = make(map[string]interface{})

		xml.Unmarshal(byteValue, &body)

		bb.Template[name] = body

		fmt.Println("Template "+name+" loaded")

    }
}

func (bb *ByteParser) ParseJsonToByte(jsonMap map[string]interface{}, templateID string) []byte{

	var byteBuffer ByteBuffer.Buffer

	byteBuffer = ByteBuffer.Buffer{
		Endian:"big",
	}

	if bb.EncodingType == "littleendian"{
		byteBuffer = ByteBuffer.Buffer{
			Endian:"little",
		}
	}

	if bb.Template[templateID] == nil{
		panic("No template id found")
		return nil
	}

	var template = bb.Template[templateID].(map[string]interface{})

	if template == nil{
		panic("Invalid json, no template key found")
		return nil
	}

	if template["Template"] == nil{
		panic("Invalid json, no template key found")
		return nil
	}

	var templateArray = template["Template"].([]interface {})

	for index := range templateArray{

		var hashMap = templateArray[index].(map[string]interface{})

		var key = hashMap["Key"].(string)

		var dataType = hashMap["DataType"].(string)

		var size = int(hashMap["Size"].(float64))

		if jsonMap[key] == nil{
			panic("Invalid key, no key found in template file")
			break
		}

		if dataType == "byte"{

			var byteData byte

			if reflect.TypeOf(jsonMap[key]).String() == "int"{

				byteData = byte(jsonMap[key].(int))

			}else if reflect.TypeOf(jsonMap[key]).String() == "float64"{

				byteData = byte(jsonMap[key].(float64))

			}else if reflect.TypeOf(jsonMap[key]).String() == "string"{

				var byteTemp = []byte(jsonMap[key].(string))

				byteData = byteTemp[0]

			}

			byteBuffer.PutByte(byteData)

		}else if dataType == "string"{

			var stringVal = jsonMap[key].(string)

			if len(stringVal) < size{

				var diff = size - len(stringVal)

				for i := 0; i < diff; i++ {

					stringVal += " "

				}

			}

			if len(stringVal) != size{
				panic("Invalid size, does not match with the size defined in the template key: "+key)
				break
			}

			byteBuffer.Put([]byte(stringVal))

		}else if dataType == "short"{

			byteBuffer.PutShort(jsonMap[key].(int))

		}else if dataType == "int"{

			byteBuffer.PutInt(jsonMap[key].(int))

		}else if dataType == "long"{

			byteBuffer.PutLong(jsonMap[key].(int))

		}else if dataType == "float"{
			
			byteBuffer.PutFloat(jsonMap[key].(float32))

		}else if dataType == "double"{

			byteBuffer.PutDouble(jsonMap[key].(float64))

		}else{

			panic(dataType)

		}

	}

	return byteBuffer.Array()

}

func (bb *ByteParser) ParseToObject(byteArr []byte, templateID string) map[string]interface{}{

	if bb.Template[templateID] == nil{
		panic("No template id found")
		return nil
	}

	var template = bb.Template[templateID].(map[string]interface{})

	if template == nil{
		panic("Invalid json, no template key found")
		return nil
	}

	var size = int(template["size"].(float64))

	if size != len(byteArr){
		panic("Length, does not match with the template invalid packet")
		return nil
	}

	if template["Template"] == nil{
		panic("Invalid json, no template key found")
		return nil
	}

	var templateArray = template["Template"].([]interface {})

	var byteBuffer ByteBuffer.Buffer

	byteBuffer = ByteBuffer.Buffer{
		Endian:"big",
	}

	if bb.EncodingType == "littleendian"{
		byteBuffer = ByteBuffer.Buffer{
			Endian:"little",
		}
	}

	byteBuffer.Wrap(byteArr)

	var returnHashMap = make(map[string]interface{})

	for index := range templateArray{

		var hashMap = templateArray[index].(map[string]interface{})

		var key = hashMap["Key"].(string)

		var dataType = hashMap["DataType"].(string)

		var size = int(hashMap["Size"].(float64))

		if dataType == "byte"{

			var byteArr = byteBuffer.GetByte()

			returnHashMap[key] = int(byteArr[0])

		}else if dataType == "string"{

			returnHashMap[key] = strings.TrimSpace(string(byteBuffer.Get(size)))

		}else if dataType == "short"{

			returnHashMap[key] = byteBuffer.Bytes2Short(byteBuffer.GetShort())

		}else if dataType == "int"{

			returnHashMap[key] = byteBuffer.Bytes2Int(byteBuffer.GetInt())

		}else if dataType == "long"{

			returnHashMap[key] = byteBuffer.Bytes2Long(byteBuffer.GetLong())

		}else if dataType == "float"{

			returnHashMap[key] = byteBuffer.Bytes2Float(byteBuffer.GetFloat())

		}else if dataType == "double"{

			returnHashMap[key] = byteBuffer.Bytes2Double(byteBuffer.GetDouble())

		}
	}

	return returnHashMap
}
