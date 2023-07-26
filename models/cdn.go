package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CDN struct {
	Id              primitive.ObjectID `json:"id,omitempty"`
	CDN             string             `json:"cdn,omitempty" validate:"required"`
	Create_system   string             `json:"create_system,omitempty" validate:"required"`
	Ip_address      string             `json:"ip_address,omitempty" validate:"required"`
	Datetime_create primitive.DateTime `json:"datetime_create,omitempty" validate:"required"`
	Detail          Detail             `json:"detail"`
}

type Detail struct {
	Name         string `json:"name"`
	Doc_type     string `json:"doc_type"`
	Case_no      string `json:"case_no"`
	Red_case_no  string `json:"red_case_no"`
	Court_code   string `json:"court_code"`
	Req_name     string `json:"req_name"`
	Indss_ref_id string `json:"indss_ref_id"`
	Pathsv       string `json:"pathsv"`
	Approve_date string `json:"approve_date"`
	Expire_date  string `json:"expire_date"`
	Typestorage  string `json:"typestorage"`
}

// {
//     "_id": ObjectId("63ff1c892a37e02398b92bd6"),
//     "cdn": "CDN-E2303018575880",
//     "create_system": "Efiling",
//     "ip_address": "10.1.2.89",
//     "datetime_create": "2023-03-01 16:36:09",
//     "detail": {
//         "name": "หนังสือรับรองคดีถึงที่สุด (สำหรับทำสิ่งพิมพ์ออก)",
//         "doc_type": "23",
//         "case_no": "พE276/2565",
//         "red_case_no": "พE281/2565",
//         "court_code": "lpgc",
//         "req_name": "สุจิตรา บุญธรรม อุปนันท์",
//         "indss_ref_id": "",
//         "pathsv": "/home/100937/20230301/1320221114685916/100937_3720230228000280_20230301163608000_FJ007_CDN_E2303018575880.pdf",
//         "approve_date": "2023-03-01",
//         "expire_date": "2023-05-30",
//         "typestorage": "SFTP"
//     }
// }
