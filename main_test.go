// main_test.go

package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	sw "./go"
)

func TestMain(m *testing.M) {
	sw.Initialize("postgres", "1", "orderlunch")

	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := sw.AppObj.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	sw.AppObj.DB.Exec("DELETE FROM orders")
	sw.AppObj.DB.Exec("ALTER SEQUENCE orders_id_seq RESTART WITH 1")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS public.orders
(
    id integer NOT NULL DEFAULT nextval('orders_id_seq'::regclass),
    userid integer,
    dishid integer,
    orderdate timestamp with time zone,
    quantity integer,
    status text COLLATE pg_catalog."default",
    CONSTRAINT orders_pkey PRIMARY KEY (id)
)`

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	sw.AppObj.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestEmptyTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/order/getList", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func TestCreateOrder(t *testing.T) {
	clearTable()

	var jsonStr = []byte(`{"UserID": 1, "DishID": 1, "Quantity": 10, "Status": "ordered"}`)
	req, _ := http.NewRequest("POST", "/order", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	response := executeRequest(req)
	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["id"] != 1.0 {
		t.Errorf("Expected product ID to be '1'. Got '%v'", m["id"])
	}
}
