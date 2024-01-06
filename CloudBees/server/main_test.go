package main

import (
	"context"
	"testing"

	"github.com/priyanshu/train-app/proto"
	"google.golang.org/grpc"
)

const serverAddress = "localhost:50051"

// TestServer tests the server functionality
func TestServer(t *testing.T) {
	// Start the server in a goroutine
	go main()

	// Connect to the server
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	// Create a client for the TrainService
	client := proto.NewTrainServiceClient(conn)

	// Define a helper function to run tests
	runTest := func(testName string, testFunc func(client proto.TrainServiceClient, t *testing.T)) {
		t.Run(testName, func(t *testing.T) {
			testFunc(client, t)
		})
	}

	// Run tests for different functionalities
	runTest("AddUserTest", createUserTest)
	runTest("PurchaseTest", purchaseTest)
	runTest("ViewUsersBySectionTest", viewUsersBySectionTest)
	runTest("RemoveUserTest", removeUserTest)

	// Add more test cases as needed
}

// createUserTest tests the AddUser functionality
func createUserTest(client proto.TrainServiceClient, t *testing.T) {
	// Create a user
	user := &proto.User{
		FirstName: "priyanshu",
		LastName:  "Doe",
		Email:     "priyanshu@example.com",
		UserId:    "priyanshu123",
	}

	// Add the user to the system
	response, err := client.AddUser(context.Background(), user)
	if err != nil {
		t.Fatalf("Error adding user: %v", err)
	}

	// Check if the user was added successfully
	if !response.Success {
		t.Error("Failed to add user.")
	} else {
		t.Logf("User Added Successfully: %+v", user)
	}
}

// purchaseTest tests the Purchase functionality
func purchaseTest(client proto.TrainServiceClient, t *testing.T) {
	// Assuming you have a user already added, you can reuse the user ID for the purchase
	userID := "priyanshu123"

	// Create a ticket
	ticket := &proto.Ticket{
		User: &proto.User{UserId: userID},
		From: "StationA",
		To:   "StationB",
	}

	// Purchase the ticket
	request := &proto.PurchaseRequest{Ticket: ticket}
	response, err := client.Purchase(context.Background(), request)
	if err != nil {
		t.Fatalf("Error purchasing ticket: %v", err)
	}
	t.Log("Purchase Function Tested Successfully")

	// Assertion 1: Ensure the response is not nil
	if response == nil {
		t.Error("Purchase response is nil.")
	}

	// Assertion 2: Check if the response ticket is not nil
	if response.Ticket == nil {
		t.Error("Purchase response ticket is nil.")
	}

	// Assertion 3: Validate the seat information (customize based on your logic)
	if response.Ticket.Seat == "" {
		t.Error("Purchase response ticket has no seat information.")
	}

	// Add more assertions based on your specific logic and expected behavior
}

// viewUsersBySectionTest tests the ViewUsersBySection functionality
func viewUsersBySectionTest(client proto.TrainServiceClient, t *testing.T) {
	// ... previous setup code ...

	// Now, test the ViewUsersBySection function
	section := "A" // Provide the section you want to test
	viewRequest := &proto.ViewUsersBySectionRequest{Section: section}
	viewResponse, err := client.ViewUsersBySection(context.Background(), viewRequest)
	if err != nil {
		t.Fatalf("Error viewing users by section: %v", err)
	}

	// Assertion 1: Ensure the response is not nil
	if viewResponse == nil {
		t.Error("ViewUsersBySection response is nil.")
	}

	// Assertion 2: Check if users are returned in the response
	if len(viewResponse.User) > 0 {
		t.Log("ViewUsersBySection Function Returned Users Successfully")
	} else {
		t.Error("ViewUsersBySection response has no users.")
	}

	// Assertion 3: Validate the seat map for users in the section
	for userID, seat := range viewResponse.SeatMap {
		if seat != section {
			t.Errorf("Incorrect seat %s for user %s in section %s", seat, userID, section)
		}
	}
	// Add more assertions based on your specific logic and expected behavior
}

// removeUserTest tests the RemoveUser functionality
func removeUserTest(client proto.TrainServiceClient, t *testing.T) {
	// Assuming you have a user already added, you can reuse the user ID for removal
	userID := "priyanshu123"

	// Remove the user
	request := &proto.RemoveUserRequest{UserId: userID}
	response, err := client.RemoveUser(context.Background(), request)
	if err != nil {
		t.Fatalf("Error removing user: %v", err)
	}

	// Log success message
	t.Log("RemoveUser Function Tested Successfully")

	// Assertions
	// Assertion 1: Ensure the response is not nil
	if response == nil {
		t.Error("RemoveUser response is nil.")
	}

	// Assertion 2: Ensure the success flag is true
	if !response.Success {
		t.Error("RemoveUser operation did not succeed.")
	}

	// You can add more assertions based on your specific logic for this operation
}

// modifySeatTest tests the ModifySeat functionality
func modifySeatTest(client proto.TrainServiceClient, t *testing.T) {
	// Assuming you have a user already added, you can reuse the user ID for modification
	userID := "priyanshu123"

	// Modify the seat to "NewSeat"
	newSeat := "NewSeat"
	request := &proto.ModifySeatRequest{UserId: userID, NewSeat: newSeat}
	response, err := client.ModifySeat(context.Background(), request)
	if err != nil {
		t.Fatalf("Error modifying seat: %v", err)
	}

	// Log success message
	t.Log("ModifySeat Function Tested Successfully")

	// Assertions
	// Assertion 1: Ensure the response is not nil
	if response == nil {
		t.Error("ModifySeat response is nil.")
	}

	// Assertion 2: Ensure the success flag is true
	if !response.Success {
		t.Error("ModifySeat operation did not succeed.")
	}

	// You can add more assertions based on your specific logic for this operation
}

// getReceiptForUserTest tests the GetReceiptForUser functionality
func getReceiptForUserTest(client proto.TrainServiceClient, t *testing.T) {
	// Assuming you have a user with the ID "priyanshu123" who has purchased a ticket
	userID := "priyanshu123"
	user := &proto.User{UserId: userID}

	// Get the receipt for the user
	response, err := client.GetReceiptForUser(context.Background(), user)
	if err != nil {
		t.Fatalf("Error getting receipt: %v", err)
	}

	// Log success message
	t.Log("GetReceiptForUser Function Tested Successfully")

	// Assertions
	// Assertion 1: Ensure the response is not nil
	if response == nil {
		t.Error("GetReceiptForUser response is nil.")
	}

	// Assertion 2: Ensure the ticket in the response is not nil
	if response.Ticket == nil {
		t.Error("Ticket in GetReceiptForUser response is nil.")
	}

	// You can add more assertions based on your specific logic for this operation
}
