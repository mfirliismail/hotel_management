package utils

import (
    "github.com/midtrans/midtrans-go"
    "github.com/midtrans/midtrans-go/snap"
    "log"
    "os"
)

var SnapGateway *snap.Client

func InitMidtrans() {
    serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
    clientKey := os.Getenv("MIDTRANS_CLIENT_KEY")

    if serverKey == "" || clientKey == "" {
        log.Fatal("MIDTRANS_SERVER_KEY or MIDTRANS_CLIENT_KEY not set in environment variables")
    }

    SnapGateway = &snap.Client{}
    SnapGateway.New(serverKey, midtrans.Sandbox)
}
