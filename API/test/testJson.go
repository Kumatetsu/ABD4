package test

import (
	"ABD4/API/model"
	"bytes"
	"encoding/json"
	"fmt"
)

func TestJson() {
	fmt.Println(string("\n\n\n==============================================================\n" +
		"====================      TEST  JSON    ====================\n" +
		"==============================================================\n\n\n"))
	TestJsonAcheteur()
	TestJsonGame()
	TestJsonReservation()
	TestJsonTransaction()
	fmt.Println(string("\n\n\n==============================================================" +
		"\n==============================================================" +
		"\n==============================================================\n\n\n"))
}


func TestJsonAcheteur() {
	fmt.Println("=====UnseriaLize bytes into acheteur=====")
	byt := []byte(`{"Civilite":"Monsieur","Nom":"Carmine","Prenom":"Art","Age":64,"Email":"carmine.art@gogole.com"}`)
	acheteur := &model.Acheteur{}

	if err := json.Unmarshal(byt, acheteur ); err != nil {
		panic(err)
	}
	fmt.Println(acheteur.ToString())
	marshallAcheteur, _ := acheteur.Marshal()

	if bytes.Equal(byt,marshallAcheteur) {
		fmt.Printf("marshal Acheteur done\n")
	}

	acheteur2 := &model.Acheteur{}

	if err := json.Unmarshal(marshallAcheteur, acheteur2 ); err != nil {
		panic(err)
	}

	if acheteur2.ToString() == acheteur.ToString() {
		fmt.Printf("Acheteur done\n")
	}
	fmt.Println("=====TEST TestJsonAcheteur Done=====\n\n\n")

}

func TestJsonGame() {
	fmt.Println("=====UnseriaLize bytes into game=====")
	byt2 := []byte(`{"Nom":"Interminable attente chez le medecin","Jour":"2018-09-07","Horaire":"05:30","VR":"Non"}`)
	game := &model.Game{}
	if err := json.Unmarshal(byt2, game ); err != nil {
		panic(err)
	}

	fmt.Println(game.ToString())
	marshallGame, _ := game.Marshal()

	if bytes.Equal(byt2,marshallGame) {
		fmt.Printf("marshal Game done\n")
	}

	game2 := &model.Game{}
	if err := json.Unmarshal(marshallGame, game2 ); err != nil {
		panic(err)
	}

	if game2.ToString() == game.ToString() {
		fmt.Printf("game done\n")
	}
	fmt.Println("=====TEST TestJsonGame Done=====\n\n\n")

}

func TestJsonReservation() {
	fmt.Println("=====UnseriaLize bytes into reservation=====")
	byt3 := []byte(`{"Spectateur":{"Civilite":"Madame","Nom":"Nya","Prenom":"Kayla","Age":22},"Tarif":"Plein tarif"}`)
	reservation := &model.Reservation{}
	if err := json.Unmarshal(byt3, reservation); err != nil {
		panic(err)
	}
	fmt.Println(reservation.ToString())

	marshallReservation, _ := reservation.Marshal()
	if bytes.Equal(byt3,marshallReservation) {
		fmt.Printf("marshal Reservation done\n")
	}

	reservation2 := &model.Reservation{}
	if err := json.Unmarshal(marshallReservation, reservation2); err != nil {
		panic(err)
	}

	if reservation2.ToString() == reservation.ToString() {
		fmt.Printf("Reservation done\n")
	}
	fmt.Println("=====TEST TestJsonReservation Done=====\n\n\n")

}

func TestJsonTransaction() {
	fmt.Println("=====UnseriaLize bytes into transaction=====")
	byt4 := []byte(`{"Acheteur":{"Civilite":"Monsieur","Nom":"Carmine","Prenom":"Art","Age":64,"Email":"carmine.art@gogole.com"},"Game":{"Nom":"Interminable attente chez le medecin","Jour":"2018-09-07","Horaire":"05:30","VR":"Non"},"Reservation":[{"Spectateur":{"Civilite":"Monsieur","Nom":"Carmine","Prenom":"Art","Age":64},"Tarif":"Senior"},{"Spectateur":{"Civilite":"Madame","Nom":"Nya","Prenom":"Kayla","Age":22},"Tarif":"Plein tarif"}]}`)
	transaction := &model.Transaction{}
	if err := json.Unmarshal(byt4, &transaction); err != nil {
		panic(err)
	}
	fmt.Println(transaction.ToString())

	marshallTransaction, _ := transaction.Marshal()
	if bytes.Equal(byt4, marshallTransaction) {
		fmt.Printf("marshal Transaction done\n")
	}

	transaction2 := &model.Transaction{}
	if err := json.Unmarshal(marshallTransaction, transaction2); err != nil {
		panic(err)
	}

	if transaction2.ToString() == transaction.ToString() {
		fmt.Printf("transaction done\n")
	}
	fmt.Println("=====Test TestJsonTransaction  Done=====\n\n\n")
}