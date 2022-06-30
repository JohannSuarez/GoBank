package api

import (
    "fmt"
    "bytes"
	"encoding/json"
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/stretchr/testify/require"

    "github.com/golang/mock/gomock"
    mockdb "github.com/JohannSuarez/GoBackend/db/mock"
    db "github.com/JohannSuarez/GoBackend/db/sqlc"
    "github.com/JohannSuarez/GoBackend/util"
)

func TestGetAccountAPI(t *testing.T) {
    account := randomAccount()

    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    store := mockdb.NewMockStore(ctrl)

    // build stubs
    store.EXPECT().
        GetAccount(gomock.Any(), gomock.Eq(account.ID)).
        Times(1).
        Return(account, nil)

    // Start test server and send request
    server := NewServer(store)
    recorder := httptest.NewRecorder()

    url := fmt.Sprintf("/accounts/%d", account.ID)
    request, err := http.NewRequest(http.MethodGet, url, nil)
    require.NoError(t, err)

    server.router.ServeHTTP(recorder, request)

    // check response
    require.Equal(t, http.StatusOK, recorder.Code)
    requireBodyMatchAccount(t, recorder.Body, account)
}

func randomAccount() db.Account {

    return db.Account{
        ID: util.RandomInt(1, 1000),
        Owner: util.RandomOwner(),
        Balance: util.RandomMoney(),
    }
}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Account) {
    data, err := ioutil.ReadAll(body)
    require.NoError(t, err)

    var gotAccount db.Account
    err = json.Unmarshal(data, &gotAccount)

    require.NoError(t, err)
    require.Equal(t, account, gotAccount)
}
