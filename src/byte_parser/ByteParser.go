package byte_parser

import(
	"log"
    "io/ioutil"
    "path/filepath"
    "os"
    "encoding/json"
    "ByteBuffer"
    "reflect"
    "strings"
    "encoding/xml"
    "errors"
)

type ByteParser struct{
	ParserType string
	Path string
	EncodingType string
	Template map[string]interface{}
}

func (bb *ByteParser) ParseJson(){

	// reading path

	files, err := ioutil.ReadDir(bb.Path)

    if err != nil {

        log.Println(err)

        return
    }

    // template

    bb.Template = make(map[string]interface{})

    // iterating each files and mapping byte array to hashmap

    for _, f := range files {

        fileName := f.Name()

        extension := filepath.Ext(fileName)

		name := fileName[0:len(fileName)-len(extension)]

		jsonFile, err := os.Open(bb.Path+"/"+fileName)

		byteValue, _ := ioutil.ReadAll(jsonFile)

		if err != nil {

		    log.Println(err)

		    return

		}

		defer jsonFile.Close()

		body := make(map[string]interface{})

		json.Unmarshal(byteValue, &body)

		bb.Template[name] = body

		log.Println("Template "+name+" loaded")

    }
}

func (bb *ByteParser) ParseXml(){

	files, err := ioutil.ReadDir(bb.Path)

    if err != nil {

        log.Println(err)

        return
    }

    bb.Template = make(map[string]interface{})

    for _, f := range files {

        fileName := f.Name()

        extension := filepath.Ext(fileName)

		name := fileName[0:len(fileName)-len(extension)]

		xmlFile, err := os.Open(bb.Path+"/"+fileName)

		byteValue, _ := ioutil.ReadAll(xmlFile)

		if err != nil {

		    log.Println(err)

		    return

		}

		defer xmlFile.Close()

		body := make(map[string]interface{})

		xml.Unmarshal(byteValue, &body)

		bb.Template[name] = body

		log.Println("Template "+name+" loaded")

    }
}

func (bb *ByteParser) ParseJsonToByte(jsonMap map[string]interface{}, templateID string) ([]byte, error){

	byteBuffer := ByteBuffer.Buffer{
		Endian:"big",
	}

	if bb.EncodingType == "littleendian"{
		byteBuffer = ByteBuffer.Buffer{
			Endian:"little",
		}
	}

	if bb.Template[templateID] == nil{
		return nil, errors.New("No template id found")
	}

	template := bb.Template[templateID].(map[string]interface{})

	if template == nil{
		return nil, errors.New("Invalid json, no template key found")
	}

	if template["Template"] == nil{
		return nil, errors.New("Invalid json, no template key found")
	}

	templateArray := template["Template"].([]interface {})

	for index := range templateArray{

		hashMap := templateArray[index].(map[string]interface{})

		key := hashMap["Key"].(string)

		dataType := hashMap["DataType"].(string)

		size := int(hashMap["Size"].(float64))

		if jsonMap[key] == nil{
			return nil, errors.New("Invalid key, no key found in template file, "+key)
		}

		if dataType == "byte"{

			var byteData byte

			if reflect.TypeOf(jsonMap[key]).String() == "int"{

				byteData = byte(jsonMap[key].(int))

			}else if reflect.TypeOf(jsonMap[key]).String() == "float64"{

				byteData = byte(jsonMap[key].(float64))

			}else if reflect.TypeOf(jsonMap[key]).String() == "string"{

				byteTemp := []byte(jsonMap[key].(string))

				byteData = byteTemp[0]

			}

			byteBuffer.PutByte(byteData)

		}else if dataType == "string"{

			stringVal := jsonMap[key].(string)

			if len(stringVal) < size{

				diff := size - len(stringVal)

				for i := 0; i < diff; i++ {

					stringVal += " "

				}

			}

			if len(stringVal) != size{
				return nil, errors.New("Invalid size, does not match with the size defined in the template key: "+key)
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

		}

	}

	return byteBuffer.Array(), nil

}

func (bb *ByteParser) ParseToObject(byteArr []byte, templateID string) (map[string]interface{}, error){

	if bb.Template[templateID] == nil{
		return nil, errors.New("No template id found")
	}

	template := bb.Template[templateID].(map[string]interface{})

	if template == nil{
		return nil, errors.New("Invalid json, no template key found")
	}

	size := int(template["size"].(float64))

	if size != len(byteArr){
		return nil, errors.New("Length, does not match with the template invalid packet")
	}

	if template["Template"] == nil{
		return nil, errors.New("Invalid json, no template key found")
	}

	templateArray := template["Template"].([]interface {})

	var byteBuffer ByteBuffer.Buffer

	if bb.EncodingType == "littleendian"{

		byteBuffer = ByteBuffer.Buffer{
			Endian:"little",
		}

	}else{

		byteBuffer = ByteBuffer.Buffer{
			Endian:"big",
		}
	}

	byteBuffer.Wrap(byteArr)

	returnHashMap := make(map[string]interface{})

	for index := range templateArray{

		hashMap := templateArray[index].(map[string]interface{})

		key := hashMap["Key"].(string)

		dataType := hashMap["DataType"].(string)

		size := int(hashMap["Size"].(float64))

		if dataType == "byte"{

			byteArr := byteBuffer.GetByte()

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

	return returnHashMap, nil
}
