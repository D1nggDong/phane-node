package main

import (
	"context"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	waLog "go.mau.fi/whatsmeow/util/log"
	"google.golang.org/protobuf/proto"
	waProto "go.mau.fi/whatsmeow/binary/proto"
)

func main() {
	ctx := context.Background()
	dbLog := waLog.Stdout("Database", "ERROR", true)
	
	container, _ := sqlstore.New(ctx, "sqlite3", "file:phane_wa.db?_foreign_keys=on", dbLog)
	deviceStore, _ := container.GetFirstDevice(ctx)
	client := whatsmeow.NewClient(deviceStore, waLog.Stdout("Client", "ERROR", true))
	
	err := client.Connect()
	if err != nil { log.Fatalf("Failed to connect: %v", err) }

	time.Sleep(2 * time.Second)
	
	// FIXED: Added ctx here
	client.SendPresence(ctx, types.PresenceAvailable)

	// TARGET: Your specific number
	target := types.NewJID("12488263584", types.DefaultUserServer)
	
	fmt.Printf("🚀 Pinging %s via PHANE Hub...\n", target.User)

	resp, err := client.SendMessage(ctx, target, &waProto.Message{
		Conversation: proto.String("🛸 PHANE SIGNAL: Connection established. Pi is live."),
	})

	if err != nil {
		log.Fatalf("❌ Error: %v", err)
	}
	
	fmt.Printf("✅ Success! Server timestamp: %s\n", resp.Timestamp)
	
	time.Sleep(2 * time.Second)
	client.Disconnect()
}
