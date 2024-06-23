# PriceHistoryAlert

![1_0sq2_OlM7HYqhRZiyqlX5Q](https://github.com/GoliathReaper/PriceHistoryAlert/assets/77969919/aded3816-2686-4ca5-a91a-0c68edf3611a)

This Go program checks the price of a product on Amazon and Flipkart using the website pricehistory.com. It compares the current price with a user-defined maximum price. If the current price is less than or equal to the maximum price, a notification is sent.

## Features

- Fetches the product page from a given URL.
- Extracts the price of the product from the page.
- Compares the extracted price with a predefined maximum price.
- Sends a telegram notification if the price is below or equal to the maximum price.
- Users can compile the Go code into an executable and configure a cron job to run periodically.

## Requirements

- Go 1.16 or later
- `github.com/PuerkitoBio/goquery`
- `github.com/go-telegram-bot-api/telegram-bot-api/v5`

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/GoliathReaper/PriceHistoryAlert.git
    cd PriceHistoryAlert
    ```

2. Install dependencies:

    ```sh
    go get -u github.com/PuerkitoBio/goquery
    go get -u github.com/go-telegram-bot-api/telegram-bot-api/v5
    ```

## Usage

1. Update the URL and maximum price in the `price_url.csv`


2. Run the program:

    ```sh
    go run .
    ```
