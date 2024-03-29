package test_tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/NovanHsiu/goutil"
	"github.com/NovanHsiu/goutil/cipher"
	"github.com/NovanHsiu/goutil/network"
)

func TestUtil(t *testing.T) {
	sqlnowtime := goutil.SQLTimeFormatToString(time.Now())
	if sqlnowtime == "" {
		t.Errorf("SQLTimeFormatToString failed! SQL Now Time: %s", sqlnowtime)
	}
}

func TestErrorHandler(t *testing.T) {
	fmt.Println(goutil.CreateResponse("200.1"))
}

func TestHttpRequest(t *testing.T) {
	fmt.Println(network.RequestTimeOutSecond)
}

func TestCipher(t *testing.T) {
	cp := cipher.DefaultCipher()
	passwd := cp.EncodePassword("abcdefg")
	if passwd == "" {
		t.Errorf("TestCipher failed! password: %s", passwd)
	}
}

func TestPostFormDataWithFilesRequest(t *testing.T) {
	/*url := "http://localhost:3000/api"
	scode, resBody, err := network.PostFormDataWithFilesRequest(url+"/files/common", map[string]string{"dir": "test"}, "upload_files", []string{"run_test.sh"})
	if err != nil {
		t.Errorf("TestPostFormDataWithFilesRequest failed! %v", err)
	} else {
		fmt.Println(scode, resBody)
	}*/
	fmt.Println("pass")
}

func TestGetSystemLanguage(t *testing.T) {
	langCode := goutil.GetSystemLanguage()
	fmt.Println(langCode)
	//t.Errorf("TestGetSystemLanguage failed!")
}

func TestSetModuelLanguage(t *testing.T) {
	goutil.SetModuleLanguage("en")
	fmt.Println(goutil.ErrorCodeTable)
}
