# csvgen

Generate Go struct definitions from CSV headers, with deduplication across structs.

For example, these three headers from Amazon order history reports:

```csv
Order Date,Order ID,Title,Category,ASIN/ISBN,UNSPSC Code,Website,Release Date,Condition,Seller,Seller Credentials,List Price Per Unit,Purchase Price Per Unit,Quantity,Payment Instrument Type,Purchase Order Number,PO Line Number,Ordering Customer Email,Shipment Date,Shipping Address Name,Shipping Address Street 1,Shipping Address Street 2,Shipping Address City,Shipping Address State,Shipping Address Zip,Order Status,Carrier Name & Tracking Number,Item Subtotal,Item Subtotal Tax,Item Total,Tax Exemption Applied,Tax Exemption Type,Exemption Opt-Out,Buyer Name,Currency,Group Name
Order Date,Order ID,Payment Instrument Type,Website,Purchase Order Number,Ordering Customer Email,Shipment Date,Shipping Address Name,Shipping Address Street 1,Shipping Address Street 2,Shipping Address City,Shipping Address State,Shipping Address Zip,Order Status,Carrier Name & Tracking Number,Subtotal,Shipping Charge,Tax Before Promotions,Total Promotions,Tax Charged,Total Charged,Buyer Name,Group Name
Order ID,Order Date,Title,Category,ASIN/ISBN,Website,Purchase Order Number,Seller,Seller Credentials,Ordering Customer Email,Shipment Date,Shipping Address Name,Shipping Address Street 1,Shipping Address Street 2,Shipping Address City,Shipping Address State,Shipping Address Zip,Return Date,Return Reason,Quantity,Buyer Name,Group Name
```

Results in three component structs for deduplication (`ABC, AB, AC`) and three base structs (`A, B, C`):

```go
package main

type ABC struct {
	BuyerName              string `csv:"Buyer Name"`
	GroupName              string `csv:"Group Name"`
	OrderDate              string `csv:"Order Date"`
	OrderId                string `csv:"Order ID"`
	OrderingCustomerEmail  string `csv:"Ordering Customer Email"`
	PurchaseOrderNumber    string `csv:"Purchase Order Number"`
	ShipmentDate           string `csv:"Shipment Date"`
	ShippingAddressCity    string `csv:"Shipping Address City"`
	ShippingAddressName    string `csv:"Shipping Address Name"`
	ShippingAddressState   string `csv:"Shipping Address State"`
	ShippingAddressStreet1 string `csv:"Shipping Address Street 1"`
	ShippingAddressStreet2 string `csv:"Shipping Address Street 2"`
	ShippingAddressZip     string `csv:"Shipping Address Zip"`
	Website                string `csv:"Website"`
}

type AB struct {
	CarrierNameTrackingNumber string `csv:"Carrier Name & Tracking Number"`
	OrderStatus               string `csv:"Order Status"`
	PaymentInstrumentType     string `csv:"Payment Instrument Type"`
}

type AC struct {
	AsinIsbn          string `csv:"ASIN/ISBN"`
	Category          string `csv:"Category"`
	Quantity          string `csv:"Quantity"`
	Seller            string `csv:"Seller"`
	SellerCredentials string `csv:"Seller Credentials"`
	Title             string `csv:"Title"`
}

type A struct {
	ABC
	AB
	AC

	Condition            string `csv:"Condition"`
	Currency             string `csv:"Currency"`
	ExemptionOptOut      string `csv:"Exemption Opt-Out"`
	ItemSubtotal         string `csv:"Item Subtotal"`
	ItemSubtotalTax      string `csv:"Item Subtotal Tax"`
	ItemTotal            string `csv:"Item Total"`
	ListPricePerUnit     string `csv:"List Price Per Unit"`
	PoLineNumber         string `csv:"PO Line Number"`
	PurchasePricePerUnit string `csv:"Purchase Price Per Unit"`
	ReleaseDate          string `csv:"Release Date"`
	TaxExemptionApplied  string `csv:"Tax Exemption Applied"`
	TaxExemptionType     string `csv:"Tax Exemption Type"`
	UnspscCode           string `csv:"UNSPSC Code"`
}

type B struct {
	ABC
	AB

	ShippingCharge      string `csv:"Shipping Charge"`
	Subtotal            string `csv:"Subtotal"`
	TaxBeforePromotions string `csv:"Tax Before Promotions"`
	TaxCharged          string `csv:"Tax Charged"`
	TotalCharged        string `csv:"Total Charged"`
	TotalPromotions     string `csv:"Total Promotions"`
}

type C struct {
	ABC
	AC

	ReturnDate   string `csv:"Return Date"`
	ReturnReason string `csv:"Return Reason"`
}
```
