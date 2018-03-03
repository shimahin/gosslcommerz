package core

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strings"
)

/*
Sample Request

curl -X POST localhost:8080/ipn -d 'tran_id=5a16c68b23783&val_id=1711231900331kHP17lnrr9T8Gt&amount=100&card_type=VISA-Dutch Bangla&store_amount=97&card_no=425272XXXXXX3456&bank_tran_id=1711231900331S0R8atkhAZksmM&status=VALID&tran_date=2017-11-23 18:59:55&currency=BDT&card_issuer=Standard Chartered Bank&card_brand=VISA&card_issuer_country=Bangladesh&card_issuer_country_code=BD&store_id=testbox&verify_sign=8070c0cefed9e629b01100d8a92afda2&verify_key=amount,bank_tran_id,base_fair,card_brand,card_issuer,card_issuer_country,card_issuer_country_code,card_no,card_type,currency,currency_amount,currency_rate,currency_type,risk_level,risk_title,status,store_amount,store_id,tran_date,tran_id,val_id,value_a,value_b,value_c,value_d&cus_fax=01711111111&currency_type=BDT&currency_amount=100.00&currency_rate=1.0000&base_fair=0.00&value_a=ref001_A&value_b=ref002_B&value_c=ref003_C&value_d=ref004_D&risk_level=0&risk_title=Safe'

*/
/*
	Sample IPN POST Request Response
	2/12/2018 1:06:13 PM INI IPN
	2/12/2018 1:06:13 PM verify_key [amount,bank_tran_id,base_fair,card_brand,card_issuer,card_issuer_country,card_issuer_country_code,card_no,card_type,currency,currency_amount,currency_rate,currency_type,risk_level,risk_title,status,store_amount,store_id,tran_date,tran_id,val_id,value_a,value_b,value_c,value_d]
	2/12/2018 1:06:13 PM card_issuer [BRAC BANK, LTD.]
	2/12/2018 1:06:13 PM tran_date [2018-02-12 13:00:49]
	2/12/2018 1:06:13 PM value_a [ref001_A]
	2/12/2018 1:06:13 PM risk_title [Safe]
	2/12/2018 1:06:13 PM status [VALID]
	2/12/2018 1:06:13 PM card_issuer_country_code [BD]
	2/12/2018 1:06:13 PM currency_amount [150.00]
	2/12/2018 1:06:13 PM risk_level [0]
	2/12/2018 1:06:13 PM store_amount [145.5]
	2/12/2018 1:06:13 PM val_id [180212130454LSoAaVvP591wES1]
	2/12/2018 1:06:13 PM value_d [ref004_D]
	2/12/2018 1:06:13 PM amount [150]
	2/12/2018 1:06:13 PM card_brand [MASTERCARD]
	2/12/2018 1:06:13 PM verify_sign_sha2 [933e92c724f9bd85d2423898cea9de1e6b549f273eebcfb95da769f864ad2634]
	2/12/2018 1:06:13 PM currency_rate [1.0000]
	2/12/2018 1:06:13 PM verify_sign [36f3f9b7231964216fc5fd7a6518b5ba]
	2/12/2018 1:06:13 PM store_id [test_shophobe]
	2/12/2018 1:06:13 PM base_fair [0.00]
	2/12/2018 1:06:13 PM currency [BDT]
	2/12/2018 1:06:13 PM card_issuer_country [Bangladesh]
	2/12/2018 1:06:13 PM value_c [ref003_C]
	2/12/2018 1:06:13 PM card_no [545610XXXXXX4362]
	2/12/2018 1:06:13 PM tran_id [CY123]
	2/12/2018 1:06:13 PM currency_type [BDT]
	2/12/2018 1:06:13 PM value_b [ref002_B]
	2/12/2018 1:06:13 PM bank_tran_id [180212130454TulpU0ADylAw4ci]
*/

func GetSslCommerzIPNListener(storeId string, storePass string) *SslCommerz {
	return &SslCommerz{
		StoreId:   storeId,
		StorePass: storePass,
	}
}

func (s *SslCommerz) IPNListener(request *http.Request) (map[string]interface{}, error) {

	// IPN sends the data to IPN URL as Query Request URI so ,

	ipnResponse := make(map[string]interface{})
	if err := request.ParseForm(); err != nil {
		return nil, err
	}
	for key, values := range request.PostForm {
		fmt.Println(key, values)
		ipnResponse[key] = values
	}

	if _, ok := ipnResponse["verify_key"]; ok {

		hasher := md5.New()

		// Hash Verification
		/*
			Method to validate the hash
			Catch two POST parameters verify_sign and verify_key.
			Explode the parameter verify_key by comma (,). Here, this parameter contents all the parameters which are returned from SSLCommerz and combination of these value generates the verify_sign value.
			Serialize all these parameters by their name
			Example: value of verify_key (before the serial): key1,key3,key2,key5,key4
			Add the store_passwd and your store password value with verify_key
			Example: Now the verify_key (after adding store_passwd): key1,key3,key2,key5,key4,store_passwd
			Example: value of verify_key (after the serialized): key1,key2,key3,key4,key5,store_passwd
			Make a string by combining the parameters' key and value. Example: key1=value1&key2=value2&key3=value3&key4=value4&key5=value5&store_passwd=Your Store Password
			Generate md5 hash of the value and match with verify_sign
		*/
		keys := strings.Split(ipnResponse["verify_key"].([]string)[0], ",")

		var queryStrings []string
		for _, key := range keys {
			if found, ok := ipnResponse[key]; ok {
				queryStrings = append(queryStrings, fmt.Sprintf("%s=%s", key, found.([]string)[0]))
			}
		}

		hasher.Write([]byte(s.StorePass))
		passmd5Hash := hex.EncodeToString(hasher.Sum(nil))
		// add store pass as per doc said
		queryStrings = append(queryStrings, fmt.Sprintf("store_passwd=%s", passmd5Hash))

		sort.Strings(queryStrings) // sort

		queryString := strings.Join(queryStrings, "&")

		hasher.Write([]byte(queryString))
		md5Hash := hex.EncodeToString(hasher.Sum(nil))

		if ipnResponse["verify_sign"] == md5Hash {
			fmt.Println("verified")
		}

		fmt.Println(queryString)
		fmt.Println(md5Hash)

		return ipnResponse, nil
	}

	return nil, errors.New("Bad Request `verify_key` not found")
}
