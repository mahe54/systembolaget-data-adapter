package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/mahe54/systembolaget-data-adapter"
)

func main() {

	// webUrl := "https://www.systembolaget.se/api/gateway/productsearch/search/"
	// categories := []string{"Öl", "Sprit", "Cider & blanddrycker", "Alkoholfritt", "Presentartiklar", "Vin"}
	// vinSubCategories := []string{"Rött vin", "Vitt vin", "Mousserande vin", "Rosévin", "Vinlåda", "Starkvin"}

	// for _, category := range categories {
	// 	if category == "Vin" {
	// 		for _, subCategory := range vinSubCategories {
	// 			productPages = append(productPages, getAllProductPages(webUrl, 1, category, subCategory)...)
	// 		}
	// 	} else {
	// 		productPages = append(productPages, getAllProductPages(webUrl, 1, category, "")...)
	// 	}
	// 	saveProductPagesToFile(productPages, category)
	// }

	client, err := systembolaget.NewClient(&systembolaget.Client{})
	if err != nil {
		fmt.Println(err)
	}

	var productPages []systembolaget.ProductPage
	productPages = append(productPages, client.GetProductPages(systembolaget.Ol, systembolaget.Mousserande_vin)...)

}

func saveProductPagesToFile(productPages []systembolaget.ProductPage, category string) {
	productPagesJson, _ := json.Marshal(productPages)
	filename := "productPages_" + category + ".json"
	err := ioutil.WriteFile(filename, productPagesJson, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

// type MockCaller struct {
// 	BaseURL string
// }

// func (cc *MockCaller) DoCall(url string, verb cmdb.HttpVerb, headers []cmdb.KeyVal, payload []byte) ([]byte, error) {

// 	url = strings.Replace(url, cc.BaseURL, "", -1)

// 	switch url {
// 	case "jwt/login":
// 		return []byte(`eyJhbGciO///FAKE TOKEN///iJIUzI1NiJ9.eyJzdWIiOiIwRWZmQm1GeDdRMHBpckExVXlFT05iRFN2b1Y4TGtkNVRuaFNpa3FvQmg5akIrSTJxZ1o5K0RqRlBXNVZHNWRUSFByQVg5VFJ6R2t1eEVsQXpkT2FUOHBZTGM1N0hGZjRxYTVPZlhzSDF6cXJSeFdaQzNoaWpBPT0iLCJuYmYiOjE2MzkxMzIzNzYsImlzcyI6ImFyMDIwY21kYnRhczIuZGRjLnRlbGlhc29uZXJhLm5ldCIsImV4cCI6MTYzOTEzNjA5NiwiX2NhY2hlSWQiOjUwNjMwLCJpYXQiOjE2MzkxMzI0OTYsImp0aSI6IklER0FBNVYwRkREV1pBUkRXMFFJUkNYSDdTVUdTNCJ9.VmmxT7mlh///FAKE TOKEN///QAPfvqkGpEPM_wLA2SrAdnMuVG6xNdbHqo`), nil

// 	case "arsys/v1/entry/AST:TS_H2/?fields=values(Asset%20ID%2B,Instance%20Id,Name,TS_H2_ReleaseNumber,TS_H2_Alias,TS_H2_SystemFullName,TS_H2_Object,AssetLifecycleStatus,TS_H2_Status,TS_H2_Description)&q=%27Data%20Set%20Id%27+%3D+%22BMC.ASSET%22AND%27Asset%20ID%2B%27+%3D+%22Hid100006532%22":
// 		return []byte(`{"entries":[{"values":{"Asset ID+":"Hid100006532","Instance Id":"ASGAA5V0ABRPEAPNINPXPM155WIJ2P","Name":"public-cloud-common-services-prod","TS_H2_ReleaseNumber":null,"TS_H2_Alias":null,"TS_H2_SystemFullName":null,"TS_H2_Object":"System","AssetLifecycleStatus":"Deployed","TS_H2_Status":"In production","TS_H2_Description":"Public cloud common services such as CI/CD, Artifactory, DB etc"},"_links":{"self":[{"href":"https://ar020cmdbtas2.ddc.teliasonera.net:8443/api/arsys/v1/entry/AST:TS_H2/000000000015817%7C000000002287474%7C000000012925700%7C000000004148777"}]}}],"_links":{"self":[{"href":"https://ar020cmdbtas2.ddc.teliasonera.net:8443/api/arsys/v1/entry/AST:TS_H2/?fields=values(Asset%20ID%2B,Instance%20Id,Name,TS_H2_ReleaseNumber,TS_H2_Alias,TS_H2_SystemFullName,TS_H2_Object,AssetLifecycleStatus,TS_H2_Status,TS_H2_Description)&q=%27Data%20Set%20Id%27+%3D+%22BMC.ASSET%22AND%27Asset%20ID%2B%27+%3D+%22Hid100006532%22"}]}}`), nil

// 	case "arsys/v1/entry/TS:H2:PersonRoleSystem?fields=values(System%20HID,System%20Name,Login%20Name,Full%20Name,Role,Email%20Address,Status)&q=%27System%20HID%27%3D%22Hid100006532%22":
// 		return []byte(`{"entries":[{"values":{"System HID":"Hid100006532","System Name":"public-cloud-common-services-prod","Login Name":"ojg469","Full Name":"John Johnsson","Role":"IT Solution Manager","Email Address":"John.Johnsson@teliacompany.com","Status":"Active"},"_links":{"self":[{"href":"https://ar020cmdbtas2.ddc.teliasonera.net:8443/api/arsys/v1/entry/TS:H2:PersonRoleSystem/000000000067370"}]}},{"values":{"System HID":"Hid100006532","System Name":"public-cloud-common-services-prod","Login Name":"ojg469","Full Name":"John Johnsson","Role":"Application Manager","Email Address":"John.Johnsson@teliacompany.com","Status":"Active"},"_links":{"self":[{"href":"https://ar020cmdbtas2.ddc.teliasonera.net:8443/api/arsys/v1/entry/TS:H2:PersonRoleSystem/000000000067371"}]}},{"values":{"System HID":"Hid100006532","System Name":"public-cloud-common-services-prod","Login Name":"ftn1234","Full Name":"Eve Svensson","Role":"Application Manager","Email Address":"Eve.Svensson@telia.no","Status":"Active"},"_links":{"self":[{"href":"https://ar020cmdbtas2.ddc.teliasonera.net:8443/api/arsys/v1/entry/TS:H2:PersonRoleSystem/000000000067684"}]}},{"values":{"System HID":"Hid100006532","System Name":"public-cloud-common-services-prod","Login Name":"ftn1234","Full Name":"Eve Svensson","Role":"IT Solution Manager","Email Address":"Eve.Svensson@telia.no","Status":"Active"},"_links":{"self":[{"href":"https://ar020cmdbtas2.ddc.teliasonera.net:8443/api/arsys/v1/entry/TS:H2:PersonRoleSystem/000000000070805"}]}},{"values":{"System HID":"Hid100006532","System Name":"public-cloud-common-services-prod","Login Name":"ncr849","Full Name":"Jonathan Rosenberg","Role":"IT Solution Owner","Email Address":"Jonathan.Rosenberg@teliacompany.com","Status":"Active"},"_links":{"self":[{"href":"https://ar020cmdbtas2.ddc.teliasonera.net:8443/api/arsys/v1/entry/TS:H2:PersonRoleSystem/000000000078912"}]}}],"_links":{"self":[{"href":"https://ar020cmdbtas2.ddc.teliasonera.net:8443/api/arsys/v1/entry/TS:H2:PersonRoleSystem?fields=values(System%20HID,System%20Name,Login%20Name,Full%20Name,Role,Email%20Address,Status)&q=%27System%20HID%27%3D%22Hid100006532%22"}]}}`), nil

// 	case "arsys/v1/entry/BMC.CORE:BMC_BaseRelationship/?fields=values(Source.InstanceId)&q=%27DatasetId%27+%3D+%22BMC.ASSET%22AND%27Source.ClassId%27%3D%22BMC_Application%22AND%27Destination.InstanceId%27%3D%20%22ASGAA5V0ABRPEAPNINPXPM155WIJ2P%22":
// 		return []byte(`{"entries":[{"values":{"Source.InstanceId":"OI-649D9194A671411EAAC51C5900F96F15"},"_links":{"self":[{"href":"https://ar020cmdbtas2.ddc.teliasonera.net:8443/api/arsys/v1/entry/BMC.CORE:BMC_BaseRelationship/000000021586429"}]}}],"_links":{"self":[{"href":"https://ar020cmdbtas2.ddc.teliasonera.net:8443/api/arsys/v1/entry/BMC.CORE:BMC_BaseRelationship/?fields=values(Source.InstanceId)&q=%27DatasetId%27+%3D+%22BMC.ASSET%22AND%27Source.ClassId%27%3D%22BMC_Application%22AND%27Destination.InstanceId%27%3D%20%22ASGAA5V0ABRPEAPNINPXPM155WIJ2P%22"}]}}`), nil

// 	case "arsys/v1/entry/BMC.CORE:BMC_Application?fields=values(InstanceId,Name,Category,Type,Item,TS_CloudServiceProvider,TS_CloudAccountID,TS_RelatedToHID,TS_RelatedToH2Name)&q=%27DatasetId%27+%3D+%22BMC.ASSET%22AND%27InstanceId%27%3D%22OI-649D9194A671411EAAC51C5900F96F15%22":
// 		return []byte(`{"entries":[{"values":{"InstanceId":"OI-649D9194A671411EAAC51C5900F96F15","Name":"public-cloud-aws-dc-dev-01","Category":"Software","Type":"Application","Item":"Application Platform","TS_CloudServiceProvider":"AWS","TS_CloudAccountID":"0123456789ABCDEF","TS_RelatedToHID":"Hid100006532","TS_RelatedToH2Name":"public-cloud-common-services-prod"},"_links":{"self":[{"href":"https://ar020cmdbtas2.ddc.teliasonera.net:8443/api/arsys/v1/entry/BMC.CORE:BMC_Application/000000000076678%7C000000002288362%7C000000012934530"}]}}],"_links":{"self":[{"href":"https://ar020cmdbtas2.ddc.teliasonera.net:8443/api/arsys/v1/entry/BMC.CORE:BMC_Application?fields=values(InstanceId,Name,Category,Type,Item,TS_CloudServiceProvider,TS_CloudAccountID,TS_RelatedToHID,TS_RelatedToH2Name)&q=%27DatasetId%27+%3D+%22BMC.ASSET%22AND%27InstanceId%27%3D%22OI-649D9194A671411EAAC51C5900F96F15%22"}]}}`), nil

// 	case "arsys/v1/entry/BMC.CORE:BMC_Application_?fields=values(InstanceId,RequestId,TS_CloudServiceProvider,TS_CloudAccountID)&q=%27DatasetId%27+%3D+%22BMC.ASSET%22AND%27InstanceId%27%3D%22OI-649D9194A671411EAAC51C5900F96F15%22":
// 		return []byte(`{"entries":[{"values":{"InstanceId":"OI-649D9194A671411EAAC51C5900F96F15","RequestId":"000000000076678","TS_CloudServiceProvider":"AWS","TS_CloudAccountID":"0123456789ABCDEF"},"_links":{"self":[{"href":"https://ar020cmdbtas2.ddc.teliasonera.net:8443/api/arsys/v1/entry/BMC.CORE:BMC_Application_/000000000076678"}]}}],"_links":{"self":[{"href":"https://ar020cmdbtas2.ddc.teliasonera.net:8443/api/arsys/v1/entry/BMC.CORE:BMC_Application_?fields=values(InstanceId,RequestId,TS_CloudServiceProvider,TS_CloudAccountID)&q=%27DatasetId%27+%3D+%22BMC.ASSET%22AND%27InstanceId%27%3D%22OI-649D9194A671411EAAC51C5900F96F15%22"}]}}`), nil

// 	case "arsys/v1/entry/BMC.CORE:BMC_Application_/000000000076678":
// 		return []byte("{}"), nil

// 	case "jwt/logout":
// 		return nil, nil
// 	}

// 	return nil, errors.New("case not found - ?")
// }
