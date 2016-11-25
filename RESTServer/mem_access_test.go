package RESTServer

import "testing"

func TestCreateAccount(t *testing.T) {
	var account []Account
	var m = NewMemoryDataAccess()
	account = append(account, Account{id: "choi12", pw: "abc123", name: "SungYong"})
	account = append(account, Account{id: "Soon7", pw: "abc777", name: "siri"})

	for _, data := range account {
		id, err := m.CreateAccount(data)
		if err != nil {
			t.Error("CreateAccount Error!")
			return
		}
		if data.id != id {
			t.Error("Fail to create account.")
			return
		}
	}

}

func TestGetAccount(t *testing.T) {
	var account []Account
	var m = NewMemoryDataAccess()
	account = append(account, Account{id: "choi12", pw: "abc123", name: "SungYong"})
	account = append(account, Account{id: "Soon7", pw: "abc777", name: "siri"})

	for _, data := range account {
		_, err := m.CreateAccount(data)
		if err != nil {
			t.Error("CreateAccount Error!")
			return
		}
	}

	for _, data := range account {
		account, err := m.GetAccount(data.id)
		if err != nil {
			t.Error("GetAccount Error!")
			return
		}

		if data.id != account.id || data.pw != account.pw || data.name != account.name {
			t.Error("Fail to get account.")
			return
		}
	}
}

func TestModifyAccount(t *testing.T) {
	var account []Account
	var m = NewMemoryDataAccess()
	account = append(account, Account{id: "choi12", pw: "abc123", name: "SungYong"})

	for _, data := range account {
		_, err := m.CreateAccount(data)
		if err != nil {
			t.Error("CreateAccount Error!")
			return
		}
	}

	var sChangeName = "JungYong"
	account[0].name = sChangeName

	err := m.ModifyAccount(account[0].id, account[0])
	if err != nil {
		t.Error("GetAccount Error!")
		return
	}

	for _, data := range account {
		account, err := m.GetAccount(data.id)
		if err != nil {
			t.Error("GetAccount Error!")
			return
		}

		if data.id != account.id && data.pw != account.pw && data.name != account.name {
			t.Error("Fail to modify account. %s - %s\n", data.name, account.name)
			return
		}
	}
}

func TestRemoveAccount(t *testing.T) {
	var account []Account
	var m = NewMemoryDataAccess()

	account = append(account, Account{id: "choi12", pw: "abc123", name: "SungYong"})

	for _, data := range account {
		_, err := m.CreateAccount(data)
		if err != nil {
			t.Error("CreateAccount Error!")
			return
		}
	}

	var sRemoveID = "choi12"

	if err := m.RemoveAccount(sRemoveID); err != nil {
		t.Error("RemoveAccount Error!")
		return
	}

	if _, err := m.GetAccount(sRemoveID); err == nil {
		t.Error("Fail to remove account. %s\n", sRemoveID)
		return
	}

}
