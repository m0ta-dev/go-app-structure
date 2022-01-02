package main

import (
	"context"
	"fmt"
	"log"
	"os"
	
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	
	
	"gitlab.com/m0ta/benefy/app/config"
	"gitlab.com/m0ta/benefy/app/controller"
	"gitlab.com/m0ta/benefy/app/logger"
	"gitlab.com/m0ta/benefy/app/router"
	"gitlab.com/m0ta/benefy/app/service"
	"gitlab.com/m0ta/benefy/app/store"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()

	// config
	cfg := config.Get()

	// logger
	l := logger.Get()
	ll := log.New(os.Stdout, "INIT\t", log.Ldate|log.Ltime|log.LUTC)

	// Init repository store (with postgresql inside)
	store, err := store.New(ctx)
	if err != nil {
		return errors.Wrap(err, "store.New failed")
	}

	if cfg.InitDB == "true" {
		var a string
		fmt.Print("== IMPORTANT == Start database initialization?(y)")
		fmt.Fscanln(os.Stdin, &a)
		if a == "y" {
			err := initDB()
			if err != nil {
				return errors.Wrap(err, "InitDB failed")
			}
		} else {
			return errors.New("Database initialization canceled")
		}
		ll.Println("Database initialization completed")
		return nil
	}

	// Init service manager
	serviceManager, err := service.NewManager(ctx, store)
	if err != nil {
		return errors.Wrap(err, "manager.New failed")
	}

	// Init controllers
	cUser 		:= controller.NewUsers(ctx, serviceManager, l)
	cEntity 	:= controller.NewEntities(ctx, store)
	cPortfolio 	:= controller.NewPortfolios(ctx, serviceManager)
		//contBalance 	:= controller.NewBalances(ctx, serviceManager, l)
	//contCurrency 	:= controller.NewCurrencies(ctx, store)
	//contLang 		:= controller.NewLangs(ctx, store)
	//cRate 		:= controller.NewRates(ctx, serviceManager, l)
	cUtil 		:= controller.NewUtils(ctx, store, serviceManager)
	cMarketData	:= controller.NewMarketData(ctx, serviceManager, l)
	cLedgerData := controller.NewLedgerData(ctx, serviceManager, store)
		//contAccountType := controller.NewAccountTypes(ctx, serviceManager, l)
		//contDivisionType := controller.NewDivisionTypes(ctx, serviceManager, l)
		//contAccount		:= controller.NewAccounts(ctx, serviceManager, l)
		//contDivision 	:= controller.NewDivisions(ctx, serviceManager, l)
		//contWalletData 	:= controller.NewWalletData1(ctx, serviceManager, l)
	//fileController := controller.NewFiles(ctx, serviceManager, l)
	
	// Initialize Fiber instance

	app := fiber.New()
	app.Use(cors.New())

	api := router.SetupRoutes(app, cUser)

	router.SetupRoutesForLedgerData(api, cLedgerData)
	router.SetupRoutesForMarketData(api, cMarketData)
	//router.SetupRoutesForRate(api, cRate)
	router.SetupRoutesForEntity(api, cEntity)
	
	router.SetupRoutesForUser(api, cUser)
	//router.SetupRoutesForLanguage(api, cUtil)
	router.SetupRoutesForUtils(api, cUtil)
	router.SetupRoutesForPortfolio(api, cPortfolio)
	
	
	startUpdateServices(ctx, serviceManager)
	
	// список api routers
	
	ll = log.New(os.Stdout, "INIT\t", log.Ldate|log.Ltime|log.LUTC)
	s := app.Stack()
	for _, v := range s {
		for _, w := range v {
			//ll.Println(s[i][j])
			if (w.Method == "GET") || (w.Method == "POST") || (w.Method == "DELETE") || (w.Method == "PATCH") {
				ll.Println(w.Method, w.Path)
			}
		}
	}
	
	// запускаем api сервер
	log.Fatal(app.Listen(cfg.HTTPAddr))

	return nil
}