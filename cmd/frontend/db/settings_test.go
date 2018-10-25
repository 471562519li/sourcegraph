package db

import (
	"testing"

	dbtesting "github.com/sourcegraph/sourcegraph/cmd/frontend/db/testing"
	"github.com/sourcegraph/sourcegraph/pkg/api"
)

func TestSettings_ListAll(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	ctx := dbtesting.TestContext(t)

	user, err := Users.Create(ctx, NewUser{
		Email:                 "a@a.com",
		Username:              "u",
		Password:              "p",
		EmailVerificationCode: "c",
	})
	if err != nil {
		t.Fatal(err)
	}

	if _, err := Settings.CreateIfUpToDate(ctx, api.ConfigurationSubject{User: &user.ID}, nil, &user.ID, ""); err != nil {
		t.Fatal(err)
	}

	settings, err := Settings.ListAll(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if want := 1; len(settings) != want {
		t.Errorf("got %d settings, want %d", len(settings), want)
	}
}
