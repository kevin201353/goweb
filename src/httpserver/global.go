package main

import (
    "io/ioutil"
    "log"
    "encoding/json"
)


//constant cmd
const GSE_LOGIN_REQUEST              string = "76"
const GSE_GET_FLIGHT                 string = "78"
const GSE_GET_ORIGIN                 string = "9A"
const GSE_GET_FLISITE                string = "9B"
const GSE_GET_FLIGHT9C               string = "9C"
const GSE_GET_ATC                    string = "111"
const GSE_UPLOAD_PHOTO               string = "C2"




/*
//登录响应
*/
type LoginResp struct {
	Errorcode  int   `json:"errorcode"`
	Operatortype int  `json:"operatortype"`
	Operatorname  string  `json:"operatorname"`
	Airport  string   `json:"airport"`
	Onduty   int  `json:"onduty"`
}

func NewLoginResp() *LoginResp {
    return &LoginResp{0,0,"吴磊", "CGO", 1}
}


/*
//航班列表
*/
type FlightInfoResp struct {
    Errorcode int `json:"errorcode"`
    Flight []FlightList `json:"flightinfolist"`
}

type FlightList struct {
    AirLine         string
    ArriveTime      string
    BridgeName      string
    CmbID           int
    CraftNo         string 
    CraftType       string
    FlightSite      string
    FlightSiteType  int
    InFlightID      string
    InFlightNo      string
    InFlightState   string
    LeaveTime       string
    OutFlightID     string
    OutFlightNo     string
    OutFlightState  string
}


type JsonStruct struct {
    
}

func (*JsonStruct)Load(filename string, v interface{}){
    data, err := ioutil.ReadFile(filename)
    if err != nil {
       log.Fatal(err)
    }
    err = json.Unmarshal(data, v)
    if err != nil {
       log.Fatal(err)
    }
}

func NewJsonStruct() *JsonStruct {
    return &JsonStruct{}
}

