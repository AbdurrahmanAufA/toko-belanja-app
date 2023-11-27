package handler

import (
	// "token-belanja/docs"
	"fmt"
	"toko-belanja/infra/config"
	"toko-belanja/infra/database"
	"toko-belanja/repository/category_repository/category_pg"
	"toko-belanja/repository/product_repository/product_pg"
	"toko-belanja/repository/transactionHistory_repository/transactionHistory_pg"
	"toko-belanja/repository/user_repository/user_pg"
	"toko-belanja/service"

	"github.com/gin-gonic/gin"
	// swaggerfiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	// swagger embed files
)

func SeedAdmin() {
	config.LoadAppConfig()

	database.InitiliazeDatabase()

	db := database.GetDatabaseInstance()

	userRepo := user_pg.NewUserPG(db)

	userService := service.NewUserService(userRepo)

	_, err := userService.SeedAdminUser()
	if err != nil {
		fmt.Printf("Error seeding admin user: %v\n", err)
		return
	}

}

func StartApp() {

	config.LoadAppConfig()

	database.InitiliazeDatabase()

	var port = config.GetAppConfig().Port

	db := database.GetDatabaseInstance()

	userRepo := user_pg.NewUserPG(db)
	userService := service.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	categoryRepo := category_pg.NewCategoryPG(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := NewCategoryHandler(categoryService)

	productRepo := product_pg.NewProductPG(db)
	productService := service.NewProductService(productRepo)
	productHandler := NewProductHandler(productService)

	transactionHistoryRepo := transactionHistory_pg.NewTransactionHistoryPG(db)
	transactionHistoryService := service.NewTransactionHistoryService(transactionHistoryRepo, productRepo, userRepo)
	transactionHistoryHandler := NewTransactionHistoryHandler(transactionHistoryService)

	authService := service.NewAuthService(userRepo, categoryRepo, productRepo, transactionHistoryRepo)

	route := gin.Default()

	// docs.SwaggerInfo.Title = "Belajar DDD"
	// docs.SwaggerInfo.Description = "Ini adalah API dengan pattern DDD"
	// docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = "h8-movies-production.up.railway.app"
	// docs.SwaggerInfo.Schemes = []string{"https", "http"}

	// route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	transactionHistoryRoute := route.Group("/transactionHistory")
	{
		transactionHistoryRoute.Use(authService.Authentication())
		transactionHistoryRoute.POST("/", transactionHistoryHandler.CreateTransaction)
		transactionHistoryRoute.GET("/my-transactions", transactionHistoryHandler.GetTransactionWithProducts)
		transactionHistoryRoute.GET("/user-transactions", authService.AuthorizationAdmin(), transactionHistoryHandler.GetTransactionWithProductsAndUser)
	}

	userRoute := route.Group("/users")
	{
		userRoute.POST("/register", userHandler.Register)
		userRoute.POST("/login", userHandler.Login)
		userRoute.Use(authService.Authentication())
		userRoute.PATCH("/topup", userHandler.TopUpBalance)
	}

	categoryRoute := route.Group("/categories")
	{
		categoryRoute.Use(authService.Authentication())
		categoryRoute.POST("/", authService.AuthorizationAdmin(), categoryHandler.CreateNewCategory)
		categoryRoute.GET("/", categoryHandler.GetCategory)
		categoryRoute.PATCH("/:categoryId", authService.AuthorizationAdmin(), categoryHandler.PatchCategory)
		categoryRoute.DELETE("/:categoryId", authService.AuthorizationAdmin(), categoryHandler.DeleteCategory)
	}

	productRoute := route.Group("/product")
	{
		productRoute.Use(authService.Authentication())
		productRoute.POST("/", authService.AuthorizationAdmin(), productHandler.CreateProduct)
		productRoute.GET("/", productHandler.GetProducts)
		productRoute.PATCH("/:productId", authService.AuthorizationAdmin(), productHandler.UpdateProductById)
		productRoute.DELETE(":productId", authService.AuthorizationAdmin(), productHandler.DeleteProduct)

	}

	route.Run(":" + port)
}
