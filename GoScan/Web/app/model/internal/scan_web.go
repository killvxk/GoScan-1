// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// ScanWeb is the golang structure for table scan_web.
type ScanWeb struct {
    Id             int         `orm:"id"              json:"id"`              //   
    CusName        string      `orm:"cus_name"        json:"cus_name"`        //   
    Url            string      `orm:"url"             json:"url"`             //   
    Code           int         `orm:"code"            json:"code"`            //   
    Title          string      `orm:"title"           json:"title"`           //   
    ContentLength  int         `orm:"content_length"  json:"content_length"`  //   
    Fingerprint    string      `orm:"fingerprint"     json:"fingerprint"`     //   
    Image          string      `orm:"image"           json:"image"`           //   
    ScreenshotFlag bool        `orm:"screenshot_flag" json:"screenshot_flag"` //   
    Js             string      `orm:"js"              json:"js"`              //   
    Urls           string      `orm:"urls"            json:"urls"`            //   
    Forms          string      `orm:"forms"           json:"forms"`           //   
    Secret         string      `orm:"secret"          json:"secret"`          //   
    Flag           bool        `orm:"flag"            json:"flag"`            //   
    NsqFlag        bool        `orm:"nsq_flag"        json:"nsq_flag"`        //   
    ScanFlag       bool        `orm:"scan_flag"       json:"scan_flag"`       //   
    ScanNsqFlag    bool        `orm:"scan_nsq_flag"   json:"scan_nsq_flag"`   //   
    CreateAt       *gtime.Time `orm:"create_at"       json:"create_at"`       //   
}