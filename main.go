package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

//creating a meeting strcture
type Meeting struct {
	title_             string
	starttime_         string
	endtime_           string
	creationtimestamp_ string
	participants_      []Participant
}

//creating participant structure
type Participant struct {
	name_  string
	email_ string
	RSPV_  string
}

func meetingID(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("Meeting with id")

	} else {
		//message = "Method not allowed"
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func newMeeting(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		w.Header().Set("content-type", "application/json")
		var meeting Meeting

		json.NewDecoder(r.Body).Decode(&meeting)

		collection := client.Database("test").Collection("meeting")
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		result, _ := collection.InsertOne(ctx, meeting)

		json.NewEncoder(w).Encode(result)
	} else {
		//message = "method not allowed"
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func meetingBtnTimes(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("meeting with tym")
	} else {
		//message = "Method not allowed"
		http.Redirect(w, r, "/", http.StatusFound)
	}
}
func meetingOfParticipant(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("participant's mmeeting")
	} else {
		//message = "method not allowed"
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func main() {

	//Setting Up the connection to the database of mongodb

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://ashish:ashish12345@cluster0.7zjfc.mongodb.net/test?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// creating a new database

	//createNewDatabase := client.Database("meetingsApi")

	//creating two collections (meetings and participants)
	//meetingsCollections := createNewDatabase.Collection("meetings")
	//participantsCollections := createNewDatabase.Collection("participants")

	//demo data for meeting

	/*meetingResult, err := meetingsCollections.InsertOne(ctx, bson.D{

		{"title", "meeting 1"},
		{"participants", bson.A{"participant1", "participant2", "participant3"}},
		{"startTime", "9:00am"},
		{"endTime", "10:00am"},
		//{"creationTimestamp", getTimestamp()},
	})
	*/

	//printing demodata id(unique) if their's no error

	/*if err != nil {
		log.Fatal(err)
	}
	fmt.Println(meetingResult)
	*/

	//demo data for meeting

	/*	participantResult, err := participantsCollections.InsertMany(ctx, []interface{}{

			bson.D{
				{"name", "participant 1"},
				{"Email", "participant1@gmail.com"},
				{"RSVP", "Yes"},
			},
			bson.D{
				{"name", "participant 2"},
				{"Email", "participant2@gmail.com"},
				{"RSVP", "No"},
			},
			bson.D{
				{"name", "participant 3"},
				{"Email", "participant3@gmail.com"},
				{"RSVP", "Maybe"},
			},
		})
	*/
	//Printing demo data of participants if errorfree

	/*if err != nil {
		log.Fatal(err)
	}
	fmt.Println(participantResult.InsertedIDs)
	*/

	//print the participant post request

	/*cursor, err := meetingsCollections.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}*/

	/*defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var meeting bson.M
		if err = cursor.Decode(&meeting); err != nil {
			log.Fatal(err)
		}
		fmt.Println(meeting)
	}
	*/

	/*var meeting bson.M
	if err = meetingsCollections.FindOne(ctx, bson.M{}).Decode(&meeting); err != nil {
		log.Fatal(err)
	}
	fmt.Println(meeting)
	*/

	/*
		filterCursor, err := participantsCollections.Find(ctx, bson.M{"_id": ObjectId("5f8adbe3216be4379ffc5b4b")})
		if err != nil {
			log.Fatal(err)
		}

		//
		var participantsFiltered []bson.M
		if err = filterCursor.Decode(&participantsFiltered); err != nil {
			log.Fatal(err)
		}
		fmt.Println(participantsFiltered)
	*/

	http.HandleFunc("/meeting", newMeeting)
	http.HandleFunc("/meeting/{id}", meetingID)
	http.HandleFunc("/particitant/{id}", meetingOfParticipant)
	http.HandleFunc("/twotimeframe", meetingBtnTimes)
	http.ListenAndServe(":3000", nil) // set listen port

}
