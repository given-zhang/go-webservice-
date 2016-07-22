package main 
import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

//POST到webService
func PostWebService(url string, method string, value string) string {
	res, err := http.Post(url, "text/xml; charset=utf-8", bytes.NewBuffer([]byte(value)))
	if err != nil {
		fmt.Println("post error", err)
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("read error", err)
	}
	res.Body.Close()
	fmt.Printf("result----%s", data)
	return "" //ByteToString(data)
}
func CreateSOAPXml(nameSpace string, methodName string, valueStr string) string {
	soapBody := "<?xml version=\"1.0\" encoding=\"utf-8\"?>"
	soapBody += "<soap12:Envelope xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xsd=\"http://www.w3.org/2001/XMLSchema\" xmlns:soap12=\"http://www.w3.org/2003/05/soap-envelope\">"
	soapBody += "<soap12:Body>"
	soapBody += "<" + methodName + " xmlns=\"" + nameSpace + "\">"
	//以下是具体参数
	soapBody += valueStr
	soapBody += "</" + methodName + "></soap12:Body></soap12:Envelope>"
	return soapBody
}

func main() {
	param := "<datad>df</data>"
	//编码
	param = template.HTMLEscapeString(param)
	postStr := CreateSOAPXml("http://tempuri.org/", "ReceiveYiCheMessage", "<xmlMessage>"+param+"</xmlMessage>")
	PostWebService("http://localhost:5906/NewsYiCheSyncService.asmx", "ReceiveYiCheMessage", postStr)
}
