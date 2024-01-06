package main

import (
	"bufio"
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/priyanshu/train-app/proto"
	"google.golang.org/grpc"
)

const (
	addr = "localhost:50051"
)

// TestClient is a test function to run various client-side tests
func TestClient(t *testing.T) {
	// Establish a connection to the server
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a gRPC client
	client := proto.NewTrainServiceClient(conn)

	// Run AddUserTest
	t.Run("AddUserTest", func(t *testing.T) {
		createUserTest(client, t)
	})

	// Run PurchaseTicketTest
	t.Run("PurchaseTicketTest", func(t *testing.T) {
		// Simulate user input for the purchaseTicket test
		scanner := bufio.NewScanner(strings.NewReader("priyanshu123\nStationA\nStationB\n"))
		purchaseTicket(client, scanner)
	})

	// Add more test cases as needed
}

// createUserTest tests the AddUser functionality
func createUserTest(client proto.TrainServiceClient, t *testing.T) {
	fmt.Println("Running AddUserTest...")

	// Create a sample user
	user := &proto.User{
		FirstName: "test",
		LastName:  "user",
		Email:     "testuser@example.com",
		UserId:    "test123",
	}

	// Call the AddUser gRPC method
	response, err := client.AddUser(context.Background(), user)
	if err != nil {
		t.Fatalf("Error adding user: %v", err)
	}

	// Check the response and log the result
	if !response.Success {
		t.Error("Failed to add user.")
	} else {
		t.Log("User Added Successfully:", user)
	}
}
