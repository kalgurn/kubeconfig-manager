package azureAKSClient_test

import (
	"errors"
	"os"
	"testing"

	"github.com/kalgurn/kubeconfig-manager/internal/azureAKSClient"
	"github.com/stretchr/testify/assert"
)

func TestIfAdminYes(t *testing.T) {
	assert := assert.New(t)
	admin := "listClusterAdminCredential"
	ifAdmin := azureAKSClient.IfAdmin(true)
	assert.Equal(admin, ifAdmin, "should be equal")
}
func TestIfAdminNo(t *testing.T) {
	assert := assert.New(t)
	admin := "listClusterUserCredential"
	ifAdmin := azureAKSClient.IfAdmin(false)
	assert.Equal(admin, ifAdmin, "should be equal")
}

// I have no idea how to mock this calls at the moment :(
//func TestGetConfig(t *testing.T) {}

// func TestGetToken(t *testing.T) {
// 	os.Setenv("TENANT_ID", "tenant")
// 	os.Setenv("SUBSCRIPTION_ID", "subscription")
// 	os.Setenv("CLIENT_ID", "client")
// 	os.Setenv("CLIENT_SECRET", "secret")

// 	azureTokenResp := `{
// 		"token_type": "Bearer",
// 		"expires_in": "11",
// 		"ext_expires_in": "11",
// 		"expires_on": "11",
// 		"not_before": "11",
// 		"resource": "https://management.azure.com",
// 		"access_token": "Token"
// 	}`

// 	azureValidToken := azureAKSClient.Token{
// 		Type:  "Bearer",
// 		Token: "Token",
// 	}

// 	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte(azureTokenResp))
// 	}))
// 	defer ts.Close()

// 	tokenID := azureAKSClient.GetToken()
// 	if tokenID != azureValidToken {
// 		t.Errorf("Get Token failed - expected %v, got %v", azureValidToken, tokenID)
// 	}
// }
func TestGetOSVar(t *testing.T) {
	assert := assert.New(t)
	os.Setenv("TESTVAR", "TESTVALUE")

	testvalue := "TESTVALUE"
	testvar := azureAKSClient.GetOSVar("TESTVAR")
	assert.Equal(testvalue, testvar, "should be equal")

}
func TestRespErr(t *testing.T) {
	err := errors.New("test")
	azureAKSClient.RespError(err)
}
