package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/xin2025/stock-screener/handlers"
	"github.com/xin2025/stock-screener/models"
)

var stocks []models.Stock

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	allowedOrigin := os.Getenv("ALLOWED_ORIGIN")
	if allowedOrigin == "" {
		log.Fatal("ALLOWED_ORIGIN environment variable is not set")
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: allowedOrigin,
		AllowMethods: "GET,POST,HEAD,OPTIONS",
		AllowHeaders: "*",
	}))

	loadCSVData("data/stocks.csv")

	app.Post("/api/screen", func(c *fiber.Ctx) error {
		return handlers.ScreenHandler(c, stocks)
	})

	log.Fatal(app.Listen(":3000"))
}

func loadCSVData(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open CSV file: %v", err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	_, err = csvReader.Read()
	if err != nil {
		log.Fatalf("Failed to read header row: %v", err)
	}

	for {
		row, err := csvReader.Read()
		if err != nil {

			break
		}

		marketCap, _ := strconv.ParseFloat(row[1], 64)
		peRatio, _ := strconv.ParseFloat(row[2], 64)
		roe, _ := strconv.ParseFloat(row[3], 64)
		debtToEquity, _ := strconv.ParseFloat(row[4], 64)
		divYield, _ := strconv.ParseFloat(row[5], 64)
		revGrowth, _ := strconv.ParseFloat(row[6], 64)
		epsGrowth, _ := strconv.ParseFloat(row[7], 64)
		currentRatio, _ := strconv.ParseFloat(row[8], 64)
		grossMargin, _ := strconv.ParseFloat(row[9], 64)

		s := models.Stock{
			Ticker:        row[0],
			MarketCap:     marketCap,
			PE:            peRatio,
			ROE:           roe,
			DebtToEquity:  debtToEquity,
			DividendYield: divYield,
			RevenueGrowth: revGrowth,
			EPSGrowth:     epsGrowth,
			CurrentRatio:  currentRatio,
			GrossMargin:   grossMargin,
		}
		stocks = append(stocks, s)
	}

	fmt.Printf("Loaded %d rows from %s\n", len(stocks), path)
}
