package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/verma29897/bulksms/db"
)

type Account struct {
	BusinessID    string `gorm:"index"`
	WABAID        string `gorm:"index"`
	PhoneNumberID string
	AccessToken   string
	UserID        *int64 `json:"user_id,omitempty"`
}

func FetchWABAData(token string) (*Account, error) {
	url := fmt.Sprintf("https://graph.facebook.com/v20.0/me?fields=id,name,whatsapp_business_account&access_token=%s", token)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var me map[string]interface{}
	body, _ := io.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &me)

	bizID := me["id"].(string)
	wabaID := me["whatsapp_business_account"].(map[string]interface{})["id"].(string)

	phoneURL := fmt.Sprintf("https://graph.facebook.com/v20.0/%s/phone_numbers?access_token=%s", wabaID, token)
	resp2, err := http.Get(phoneURL)
	if err != nil {
		return nil, err
	}
	defer resp2.Body.Close()

	var phoneData map[string]interface{}
	pb, _ := io.ReadAll(resp2.Body)
	_ = json.Unmarshal(pb, &phoneData)

	phones := phoneData["data"].([]interface{})
	phoneID := phones[0].(map[string]interface{})["id"].(string)

	return &Account{
		BusinessID:    bizID,
		WABAID:        wabaID,
		PhoneNumberID: phoneID,
		AccessToken:   token,
	}, nil
}

func StoreAccount(a *Account) error {
	conn := db.GetDB()
	// Try with user_id first; if column missing, fall back without it
	_, err := conn.Exec(`
	INSERT INTO whatsapp_accounts (business_id, waba_id, phone_number_id, access_token, user_id)
	VALUES ($1, $2, $3, $4, $5)`,
		a.BusinessID, a.WABAID, a.PhoneNumberID, a.AccessToken, a.UserID)
	if err != nil {
		if strings.Contains(err.Error(), "column \"user_id\" does not exist") {
			_, e2 := conn.Exec(`
			INSERT INTO whatsapp_accounts (business_id, waba_id, phone_number_id, access_token)
			VALUES ($1, $2, $3, $4)`,
				a.BusinessID, a.WABAID, a.PhoneNumberID, a.AccessToken)
			return e2
		}
		return err
	}
	return nil
}

func GetLatestAccountByUserID(userID int64) (*Account, error) {
	conn := db.GetDB()
	row := conn.QueryRow(`
		SELECT business_id, waba_id, phone_number_id, access_token, user_id
		FROM whatsapp_accounts
		WHERE user_id = $1
		ORDER BY id DESC
		LIMIT 1`, userID)
	var a Account
	if err := row.Scan(&a.BusinessID, &a.WABAID, &a.PhoneNumberID, &a.AccessToken, &a.UserID); err != nil {
		// Fallback: return the latest account globally if per-user not available
		row2 := conn.QueryRow(`
			SELECT business_id, waba_id, phone_number_id, access_token
			FROM whatsapp_accounts
			ORDER BY id DESC
			LIMIT 1`)
		var b Account
		if e2 := row2.Scan(&b.BusinessID, &b.WABAID, &b.PhoneNumberID, &b.AccessToken); e2 != nil {
			return nil, e2
		}
		return &b, nil
	}
	return &a, nil
}
