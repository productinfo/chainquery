// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package model

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

func testTransactions(t *testing.T) {
	t.Parallel()

	query := Transactions(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testTransactionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	transaction := &Transaction{}
	if err = randomize.Struct(seed, transaction, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = transaction.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = transaction.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Transactions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTransactionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	transaction := &Transaction{}
	if err = randomize.Struct(seed, transaction, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = transaction.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Transactions(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Transactions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTransactionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	transaction := &Transaction{}
	if err = randomize.Struct(seed, transaction, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = transaction.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := TransactionSlice{transaction}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Transactions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testTransactionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	transaction := &Transaction{}
	if err = randomize.Struct(seed, transaction, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = transaction.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := TransactionExists(tx, transaction.ID)
	if err != nil {
		t.Errorf("Unable to check if Transaction exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TransactionExistsG to return true, but got false.")
	}
}
func testTransactionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	transaction := &Transaction{}
	if err = randomize.Struct(seed, transaction, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = transaction.Insert(tx); err != nil {
		t.Error(err)
	}

	transactionFound, err := FindTransaction(tx, transaction.ID)
	if err != nil {
		t.Error(err)
	}

	if transactionFound == nil {
		t.Error("want a record, got nil")
	}
}
func testTransactionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	transaction := &Transaction{}
	if err = randomize.Struct(seed, transaction, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = transaction.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Transactions(tx).Bind(transaction); err != nil {
		t.Error(err)
	}
}

func testTransactionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	transaction := &Transaction{}
	if err = randomize.Struct(seed, transaction, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = transaction.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Transactions(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTransactionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	transactionOne := &Transaction{}
	transactionTwo := &Transaction{}
	if err = randomize.Struct(seed, transactionOne, transactionDBTypes, false, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}
	if err = randomize.Struct(seed, transactionTwo, transactionDBTypes, false, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = transactionOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = transactionTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Transactions(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTransactionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	transactionOne := &Transaction{}
	transactionTwo := &Transaction{}
	if err = randomize.Struct(seed, transactionOne, transactionDBTypes, false, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}
	if err = randomize.Struct(seed, transactionTwo, transactionDBTypes, false, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = transactionOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = transactionTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Transactions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testTransactionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	transaction := &Transaction{}
	if err = randomize.Struct(seed, transaction, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = transaction.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Transactions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTransactionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	transaction := &Transaction{}
	if err = randomize.Struct(seed, transaction, transactionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = transaction.Insert(tx, transactionColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := Transactions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTransactionToManyTransactionByHashClaims(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Transaction
	var b, c Claim

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, claimDBTypes, false, claimColumnsWithDefault...)
	randomize.Struct(seed, &c, claimDBTypes, false, claimColumnsWithDefault...)

	b.TransactionByHashID.Valid = true
	c.TransactionByHashID.Valid = true
	b.TransactionByHashID.String = a.Hash
	c.TransactionByHashID.String = a.Hash
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	claim, err := a.TransactionByHashClaims(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range claim {
		if v.TransactionByHashID.String == b.TransactionByHashID.String {
			bFound = true
		}
		if v.TransactionByHashID.String == c.TransactionByHashID.String {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := TransactionSlice{&a}
	if err = a.L.LoadTransactionByHashClaims(tx, false, (*[]*Transaction)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.TransactionByHashClaims); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.TransactionByHashClaims = nil
	if err = a.L.LoadTransactionByHashClaims(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.TransactionByHashClaims); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", claim)
	}
}

func testTransactionToManyInputs(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Transaction
	var b, c Input

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, inputDBTypes, false, inputColumnsWithDefault...)
	randomize.Struct(seed, &c, inputDBTypes, false, inputColumnsWithDefault...)

	b.TransactionID = a.ID
	c.TransactionID = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	input, err := a.Inputs(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range input {
		if v.TransactionID == b.TransactionID {
			bFound = true
		}
		if v.TransactionID == c.TransactionID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := TransactionSlice{&a}
	if err = a.L.LoadInputs(tx, false, (*[]*Transaction)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Inputs); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Inputs = nil
	if err = a.L.LoadInputs(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Inputs); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", input)
	}
}

func testTransactionToManyOutputs(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Transaction
	var b, c Output

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, outputDBTypes, false, outputColumnsWithDefault...)
	randomize.Struct(seed, &c, outputDBTypes, false, outputColumnsWithDefault...)

	b.TransactionID = a.ID
	c.TransactionID = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	output, err := a.Outputs(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range output {
		if v.TransactionID == b.TransactionID {
			bFound = true
		}
		if v.TransactionID == c.TransactionID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := TransactionSlice{&a}
	if err = a.L.LoadOutputs(tx, false, (*[]*Transaction)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Outputs); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Outputs = nil
	if err = a.L.LoadOutputs(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Outputs); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", output)
	}
}

func testTransactionToManyTransactionAddresses(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Transaction
	var b, c TransactionAddress

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	randomize.Struct(seed, &b, transactionAddressDBTypes, false, transactionAddressColumnsWithDefault...)
	randomize.Struct(seed, &c, transactionAddressDBTypes, false, transactionAddressColumnsWithDefault...)

	b.TransactionID = a.ID
	c.TransactionID = a.ID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	transactionAddress, err := a.TransactionAddresses(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range transactionAddress {
		if v.TransactionID == b.TransactionID {
			bFound = true
		}
		if v.TransactionID == c.TransactionID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := TransactionSlice{&a}
	if err = a.L.LoadTransactionAddresses(tx, false, (*[]*Transaction)(&slice)); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.TransactionAddresses); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.TransactionAddresses = nil
	if err = a.L.LoadTransactionAddresses(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.TransactionAddresses); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", transactionAddress)
	}
}

func testTransactionToManyAddOpTransactionByHashClaims(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Transaction
	var b, c, d, e Claim

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, transactionDBTypes, false, strmangle.SetComplement(transactionPrimaryKeyColumns, transactionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Claim{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, claimDBTypes, false, strmangle.SetComplement(claimPrimaryKeyColumns, claimColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Claim{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddTransactionByHashClaims(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.Hash != first.TransactionByHashID.String {
			t.Error("foreign key was wrong value", a.Hash, first.TransactionByHashID.String)
		}
		if a.Hash != second.TransactionByHashID.String {
			t.Error("foreign key was wrong value", a.Hash, second.TransactionByHashID.String)
		}

		if first.R.TransactionByHash != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.TransactionByHash != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.TransactionByHashClaims[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.TransactionByHashClaims[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.TransactionByHashClaims(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testTransactionToManySetOpTransactionByHashClaims(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Transaction
	var b, c, d, e Claim

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, transactionDBTypes, false, strmangle.SetComplement(transactionPrimaryKeyColumns, transactionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Claim{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, claimDBTypes, false, strmangle.SetComplement(claimPrimaryKeyColumns, claimColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.SetTransactionByHashClaims(tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.TransactionByHashClaims(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetTransactionByHashClaims(tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.TransactionByHashClaims(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.TransactionByHashID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.TransactionByHashID.Valid {
		t.Error("want c's foreign key value to be nil")
	}
	if a.Hash != d.TransactionByHashID.String {
		t.Error("foreign key was wrong value", a.Hash, d.TransactionByHashID.String)
	}
	if a.Hash != e.TransactionByHashID.String {
		t.Error("foreign key was wrong value", a.Hash, e.TransactionByHashID.String)
	}

	if b.R.TransactionByHash != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.TransactionByHash != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.TransactionByHash != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.TransactionByHash != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.TransactionByHashClaims[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.TransactionByHashClaims[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testTransactionToManyRemoveOpTransactionByHashClaims(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Transaction
	var b, c, d, e Claim

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, transactionDBTypes, false, strmangle.SetComplement(transactionPrimaryKeyColumns, transactionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Claim{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, claimDBTypes, false, strmangle.SetComplement(claimPrimaryKeyColumns, claimColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	err = a.AddTransactionByHashClaims(tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.TransactionByHashClaims(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveTransactionByHashClaims(tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.TransactionByHashClaims(tx).Count()
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if b.TransactionByHashID.Valid {
		t.Error("want b's foreign key value to be nil")
	}
	if c.TransactionByHashID.Valid {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.TransactionByHash != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.TransactionByHash != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.TransactionByHash != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.TransactionByHash != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.TransactionByHashClaims) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.TransactionByHashClaims[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.TransactionByHashClaims[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testTransactionToManyAddOpInputs(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Transaction
	var b, c, d, e Input

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, transactionDBTypes, false, strmangle.SetComplement(transactionPrimaryKeyColumns, transactionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Input{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, inputDBTypes, false, strmangle.SetComplement(inputPrimaryKeyColumns, inputColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Input{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddInputs(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.TransactionID {
			t.Error("foreign key was wrong value", a.ID, first.TransactionID)
		}
		if a.ID != second.TransactionID {
			t.Error("foreign key was wrong value", a.ID, second.TransactionID)
		}

		if first.R.Transaction != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Transaction != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Inputs[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Inputs[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Inputs(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testTransactionToManyAddOpOutputs(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Transaction
	var b, c, d, e Output

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, transactionDBTypes, false, strmangle.SetComplement(transactionPrimaryKeyColumns, transactionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Output{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, outputDBTypes, false, strmangle.SetComplement(outputPrimaryKeyColumns, outputColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Output{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddOutputs(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.TransactionID {
			t.Error("foreign key was wrong value", a.ID, first.TransactionID)
		}
		if a.ID != second.TransactionID {
			t.Error("foreign key was wrong value", a.ID, second.TransactionID)
		}

		if first.R.Transaction != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Transaction != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Outputs[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Outputs[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Outputs(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testTransactionToManyAddOpTransactionAddresses(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Transaction
	var b, c, d, e TransactionAddress

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, transactionDBTypes, false, strmangle.SetComplement(transactionPrimaryKeyColumns, transactionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*TransactionAddress{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, transactionAddressDBTypes, false, strmangle.SetComplement(transactionAddressPrimaryKeyColumns, transactionAddressColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*TransactionAddress{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddTransactionAddresses(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.TransactionID {
			t.Error("foreign key was wrong value", a.ID, first.TransactionID)
		}
		if a.ID != second.TransactionID {
			t.Error("foreign key was wrong value", a.ID, second.TransactionID)
		}

		if first.R.Transaction != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Transaction != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.TransactionAddresses[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.TransactionAddresses[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.TransactionAddresses(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testTransactionToOneBlockUsingBlockByHash(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Transaction
	var foreign Block

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, blockDBTypes, false, blockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Block struct: %s", err)
	}

	local.BlockByHashID.Valid = true

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.BlockByHashID.String = foreign.Hash
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.BlockByHash(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.Hash != foreign.Hash {
		t.Errorf("want: %v, got %v", foreign.Hash, check.Hash)
	}

	slice := TransactionSlice{&local}
	if err = local.L.LoadBlockByHash(tx, false, (*[]*Transaction)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.BlockByHash == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.BlockByHash = nil
	if err = local.L.LoadBlockByHash(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.BlockByHash == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testTransactionToOneSetOpBlockUsingBlockByHash(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Transaction
	var b, c Block

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, transactionDBTypes, false, strmangle.SetComplement(transactionPrimaryKeyColumns, transactionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, blockDBTypes, false, strmangle.SetComplement(blockPrimaryKeyColumns, blockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, blockDBTypes, false, strmangle.SetComplement(blockPrimaryKeyColumns, blockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Block{&b, &c} {
		err = a.SetBlockByHash(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.BlockByHash != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.BlockByHashTransactions[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.BlockByHashID.String != x.Hash {
			t.Error("foreign key was wrong value", a.BlockByHashID.String)
		}

		zero := reflect.Zero(reflect.TypeOf(a.BlockByHashID.String))
		reflect.Indirect(reflect.ValueOf(&a.BlockByHashID.String)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.BlockByHashID.String != x.Hash {
			t.Error("foreign key was wrong value", a.BlockByHashID.String, x.Hash)
		}
	}
}

func testTransactionToOneRemoveOpBlockUsingBlockByHash(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Transaction
	var b Block

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, transactionDBTypes, false, strmangle.SetComplement(transactionPrimaryKeyColumns, transactionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, blockDBTypes, false, strmangle.SetComplement(blockPrimaryKeyColumns, blockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(tx); err != nil {
		t.Fatal(err)
	}

	if err = a.SetBlockByHash(tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveBlockByHash(tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.BlockByHash(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.BlockByHash != nil {
		t.Error("R struct entry should be nil")
	}

	if a.BlockByHashID.Valid {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.BlockByHashTransactions) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testTransactionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	transaction := &Transaction{}
	if err = randomize.Struct(seed, transaction, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = transaction.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = transaction.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testTransactionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	transaction := &Transaction{}
	if err = randomize.Struct(seed, transaction, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = transaction.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := TransactionSlice{transaction}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testTransactionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	transaction := &Transaction{}
	if err = randomize.Struct(seed, transaction, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = transaction.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Transactions(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	transactionDBTypes = map[string]string{`BlockByHashID`: `varchar`, `Created`: `datetime`, `CreatedTime`: `datetime`, `Fee`: `decimal`, `Hash`: `varchar`, `ID`: `bigint`, `InputCount`: `int`, `LockTime`: `int`, `Modified`: `datetime`, `OutputCount`: `int`, `Raw`: `text`, `TransactionSize`: `bigint`, `TransactionTime`: `bigint`, `Value`: `decimal`, `Version`: `int`}
	_                  = bytes.MinRead
)

func testTransactionsUpdate(t *testing.T) {
	t.Parallel()

	if len(transactionColumns) == len(transactionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	transaction := &Transaction{}
	if err = randomize.Struct(seed, transaction, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = transaction.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Transactions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, transaction, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	if err = transaction.Update(tx); err != nil {
		t.Error(err)
	}
}

func testTransactionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(transactionColumns) == len(transactionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	transaction := &Transaction{}
	if err = randomize.Struct(seed, transaction, transactionDBTypes, true, transactionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = transaction.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Transactions(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, transaction, transactionDBTypes, true, transactionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(transactionColumns, transactionPrimaryKeyColumns) {
		fields = transactionColumns
	} else {
		fields = strmangle.SetComplement(
			transactionColumns,
			transactionPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(transaction))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := TransactionSlice{transaction}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testTransactionsUpsert(t *testing.T) {
	t.Parallel()

	if len(transactionColumns) == len(transactionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	transaction := Transaction{}
	if err = randomize.Struct(seed, &transaction, transactionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = transaction.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Transaction: %s", err)
	}

	count, err := Transactions(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &transaction, transactionDBTypes, false, transactionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Transaction struct: %s", err)
	}

	if err = transaction.Upsert(tx, nil); err != nil {
		t.Errorf("Unable to upsert Transaction: %s", err)
	}

	count, err = Transactions(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}