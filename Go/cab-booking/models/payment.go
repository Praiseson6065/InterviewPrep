package models

type PaymentStatus string

const (
    PaymentPending   PaymentStatus = "PENDING"
    PaymentSuccess   PaymentStatus = "SUCCESS"
    PaymentFailed    PaymentStatus = "FAILED"
    PaymentRefunded  PaymentStatus = "REFUNDED"
)

type Payment struct {
    ID        string
    RideID    string
    Amount    float64
    Method    string
    Status    PaymentStatus
}