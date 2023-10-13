# AvoxiCodingChallenge


To Run:

make run or go run ./main.go

Endpoint:
/ip-check GET

Request Body example:
{
    "ipAddress": "5.170.0.1",
    "countries": [
        "US", "UA", "JP"
    ]
}


Response Example:
{
    "ipAddress": "5.170.0.1",
    "countries": [
        {
            "country": "UA",
            "ipInCountry": false
        },
        {
            "country": "JP",
            "ipInCountry": false
        },
        {
            "country": "US",
            "ipInCountry": false
        }
    ]
}
