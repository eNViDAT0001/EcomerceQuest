package main

import (
	"github.com/eNViDAT0001/Thesis/Backend/socket"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/roylee0704/gron"
	"log"
	"net/http"
	"time"
)

func router(r *gin.Engine) {
	r.Use(requestid.New(requestid.WithCustomHeaderStrKey("da-Best-thesis")))
	allHandler := initHandlerCollection()
	socketManager := socket.GetManager()
	// Validate Token: allHandler.jwtHandler.VerifyToken()
	// Validate User's Token: allHandler.jwtHandler.VerifyUserToken()
	// Validate Admin's Token: allHandler.jwtHandler.VerifyAdminToken()
	v1 := r.Group("/api/v1")
	{
		chatGroup := v1.Group("/chats")
		{
			chatGroup.GET("/channels/users/:user_id", allHandler.chatHandler.ListChannel())
			chatGroup.GET("/channels/user_a/:user_id_a/user_b/:user_id_b", allHandler.chatHandler.ListMessages())
			chatGroup.PATCH("/:message_id/user/:user_id/to/:to_id", allHandler.chatHandler.SeenMessages())
			chatGroup.POST("", allHandler.chatHandler.CreateMessage())
		}
		notifyGroup := v1.Group("/notifications")
		{
			notifyGroup.GET("/users/:user_id", allHandler.notificationHandler.ListNotifications())
			notifyGroup.PATCH("/:notify_id/user/:user_id", allHandler.notificationHandler.SeenNotification())
			notifyGroup.PATCH("/user/:user_id", allHandler.notificationHandler.SeenAllNotification())
		}
		socketGroup := v1.Group("/ws")
		{
			socketGroup.GET("/user/:user_id/token/:token", socketManager.ConnectChatWS())
		}
		appGroup := v1.Group("/app")
		{
			loginGroup := appGroup.Group("/login")
			{
				loginGroup.POST("", allHandler.appAccessHandler.Login())
				loginGroup.GET("/google", allHandler.appAccessHandler.GoogleSSO())
				loginGroup.GET("/google/callback", allHandler.appAccessHandler.CallbackGoogleSSO())
			}
			appGroup.POST("/register", allHandler.appAccessHandler.Register())
		}

		tokenGroup := v1.Group("/token")
		{
			tokenGroup.POST("/refresh", allHandler.jwtHandler.RefreshToken())
		}

		mailGroup := v1.Group("/mail")
		{
			authAminGroup := mailGroup.Group("")

			mailGroup.POST("/verify", allHandler.smtpHandler.VerifyCode())
			mailGroup.POST("/reset", allHandler.smtpHandler.CreateResetPassCode())
			mailGroup.POST("/feedback", allHandler.smtpHandler.CreateEmailFeedback())

			authAminGroup.GET("/feedback", allHandler.smtpHandler.List())
			authAminGroup.POST("", allHandler.smtpHandler.SendEmail())
		}
		fileGroup := v1.Group("/files")
		{
			fileGroup.POST("", allHandler.mediaHandler.UploadMedia())
			fileGroup.DELETE("/:public_id", allHandler.mediaHandler.DeleteMedia())
		}

		providerGroup := v1.Group("/providers")
		{
			// Phân quyền theo Admin
			authAminGroup := providerGroup.Group("")
			//authAminGroup.Use(allHandler.jwtHandler.VerifyAdminToken())
			// Phân quyền theo User ID
			authGroup := providerGroup.Group("")
			//authGroup.Use(allHandler.jwtHandler.VerifyProviderWithUserToken())
			// Phân quyền theo User ID
			authUserGroup := providerGroup.Group("")
			//authUserGroup.Use(allHandler.jwtHandler.VerifyUserToken())

			authAminGroup.GET("", allHandler.providerHandler.ListProvider())
			authUserGroup.GET("/user/:user_id", allHandler.providerHandler.ListProviderByUserID())
			providerGroup.GET("/:provider_id", allHandler.providerHandler.GetProviderByID())
			authGroup.GET("/:provider_id/user/:user_id", allHandler.providerHandler.GetProviderFullDetailByID())
			authUserGroup.POST("/user/:user_id", allHandler.providerHandler.CreateProvider())
			authGroup.PATCH("/:provider_id/user/:user_id", allHandler.providerHandler.UpdateProvider())
			authGroup.DELETE("/:provider_id/user/:user_id", allHandler.providerHandler.DeleteProviderByID())
		}

		favoriteGroup := v1.Group("/favorites")
		{
			// Phân quyền theo User ID
			authGroup := favoriteGroup.Group("")
			authGroup.Use(allHandler.jwtHandler.VerifyUserToken())

			authGroup.GET("/user/:user_id", allHandler.favoriteHandler.ListProvidersByUserID())
			authGroup.POST("", allHandler.favoriteHandler.AddFavorite())
			authGroup.DELETE("/user/:user_id/provider/:provider_id", allHandler.favoriteHandler.DeleteFavorite())
		}

		categoryGroup := v1.Group("/categories")
		{
			// Phân quyền theo Admin
			authAminGroup := categoryGroup.Group("")
			authAminGroup.Use(allHandler.jwtHandler.VerifyAdminToken())

			categoryGroup.GET("/roof", allHandler.categoryHandler.GetCategoryRoofList())
			categoryGroup.GET("/children/:category_id", allHandler.categoryHandler.GetCategoryChildrenTreeWithCategoryID())
			categoryGroup.GET("/parents/:category_id", allHandler.categoryHandler.GetCategoryParentsTreeWithCategoryID())
			categoryGroup.GET("", allHandler.categoryHandler.GetCategoryTree())
			categoryGroup.GET("/list", allHandler.categoryHandler.ListCategories())

			authAminGroup.POST("", allHandler.categoryHandler.CreateCategory())
			authAminGroup.PATCH("/:category_id", allHandler.categoryHandler.UpdateCategory())

		}
		commentGroup := v1.Group("/comments")
		{
			// Phân quyền theo User ID
			authGroup := commentGroup.Group("")
			authGroup.Use(allHandler.jwtHandler.VerifyUserToken())

			commentGroup.GET("product/:product_id", allHandler.commentHandler.ListCommentByProductID())
			commentGroup.GET("/:comment_id", allHandler.commentHandler.GetCommentDetailByID())
			authGroup.POST("/product/:product_id/user/:user_id", allHandler.commentHandler.CreateComment())
		}
		userGroup := v1.Group("/users")
		{
			// Chỉ cần có token là dùng được
			userTokenGroup := userGroup.Group("")
			userTokenGroup.Use(allHandler.jwtHandler.VerifyUserToken())
			// Phân quyền theo User ID
			authGroup := userGroup.Group("")
			authGroup.Use(allHandler.jwtHandler.VerifyUserToken())

			// Phân quyền theo Admin
			authAminGroup := userGroup.Group("")
			authAminGroup.Use(allHandler.jwtHandler.VerifyAdminToken())

			authGroup.GET("/:user_id", allHandler.userHandler.GetUserDetailByID())
			authGroup.PATCH("/:user_id", allHandler.userHandler.AdminUpdateUser())
			authGroup.PATCH("/:user_id/info", allHandler.userHandler.UpdateUserInfo())
			authGroup.PATCH("/:user_id/identity", allHandler.userHandler.UpdateUserIdentity())
			authGroup.PUT("/:user_id", allHandler.userHandler.SetPassword())
			userGroup.PUT("/reset_pass/:email", allHandler.userHandler.ResetPassword())
			authAminGroup.DELETE("/:user_id", allHandler.userHandler.DeleteUserByID())
			authAminGroup.DELETE("", allHandler.userHandler.DeleteUserByIDs())
			userTokenGroup.GET("", allHandler.userHandler.GetUserList())
			authAminGroup.POST("", allHandler.userHandler.CreateUser())
		}
		productGroup := v1.Group("/products")
		{
			// Phân quyền theo User ID
			authGroup := productGroup.Group("")
			authGroup.Use(allHandler.jwtHandler.VerifyProductWithUserToken())

			// Phân quyền theo User ID và ProviderID
			authProviderGroup := productGroup.Group("")
			authProviderGroup.Use(allHandler.jwtHandler.VerifyProviderWithUserToken())

			productGroup.GET("", allHandler.productHandler.ListProduct())
			productGroup.GET("/:product_id/description", allHandler.productHandler.GetProductDescriptionsByProductID())
			productGroup.GET("/:product_id", allHandler.productHandler.GetProductInfoByID())
			productGroup.GET("/:product_id/media", allHandler.productHandler.GetProductMediaByProductID())
			productGroup.GET("/:product_id/specification", allHandler.productHandler.GetSpecificationTreeByProductID())
			productGroup.GET("/preview", allHandler.productHandler.ListProductsPreview())
			productGroup.GET("/category/:category_id/preview", allHandler.productHandler.ListProductPreviewWithCategoryID())
			productGroup.GET("/category/:category_id", allHandler.productHandler.ListProductWithCategoryID())
			productGroup.GET("/banner/:banner_id/preview", allHandler.productHandler.ListProductPreviewWithBannerID())
			productGroup.GET("/banner/:banner_id/preview/no", allHandler.productHandler.ListProductPreviewNotInBannerID())

			authProviderGroup.POST("/provider/:provider_id/user/:user_id", allHandler.productHandler.CreateProduct())
			authGroup.POST("/:product_id/description", allHandler.productHandler.CreateDescriptions())
			authGroup.POST("/:product_id/option", allHandler.productHandler.CreateProductOptions())
			authGroup.POST("/:product_id/media", allHandler.productHandler.CreateProductMedia())
			authGroup.POST("/:product_id/specification", allHandler.productHandler.CreateSpecification())

			authGroup.DELETE("/:product_id", allHandler.productHandler.DeleteProductByID())
			authProviderGroup.DELETE("/provider/:provider_id", allHandler.productHandler.DeleteProductByIDs())

			authGroup.PATCH("/:product_id/option", allHandler.productHandler.UpdateProductOptions())
			authGroup.PATCH("/:product_id/", allHandler.productHandler.UpdateFullProduct())
			authGroup.PATCH("/:product_id/specification/:specification_id", allHandler.productHandler.UpdateProductSpecification())
			authGroup.PATCH("/:product_id/descriptions/:description_id", allHandler.productHandler.UpdateDescriptions())

		}
		bannerGroup := v1.Group("/banners")
		{
			authGroup := bannerGroup.Group("")
			authGroup.Use(allHandler.jwtHandler.VerifyUserToken())

			// Phân quyền theo Admin
			authAminGroup := bannerGroup.Group("")
			authAminGroup.Use(allHandler.jwtHandler.VerifyAdminToken())

			bannerGroup.GET("", allHandler.bannerHandler.ListBanner())
			bannerGroup.GET("/:banner_id", allHandler.bannerHandler.GetBannerByID())
			bannerGroup.GET("/:banner_id/detail", allHandler.bannerHandler.GetBannerDetailByID())
			bannerGroup.GET("/:banner_id/products", allHandler.bannerHandler.ListProductByBannerID())
			bannerGroup.GET("/:banner_id/products/preview", allHandler.bannerHandler.ListProductPreviewByBannerID())
			bannerGroup.GET("/:banner_id/products/no", allHandler.bannerHandler.ListProductNotInBannerID())
			bannerGroup.GET("/:banner_id/products/preview/no", allHandler.bannerHandler.ListProductPreviewNotInBannerID())

			authAminGroup.DELETE("", allHandler.bannerHandler.DeleteBannerByIDs())
			authAminGroup.POST("", allHandler.bannerHandler.CreateBanner())
			authAminGroup.PATCH("/:banner_id/user/:user_id", allHandler.bannerHandler.UpdateBanner())
		}
		cartGroup := v1.Group("/carts")
		{
			cartGroup.GET("/user/:user_id", allHandler.cartHandler.ListCartByUserID())
			cartGroup.DELETE("/:cart_id/user/:user_id", allHandler.cartHandler.DeleteCart())

			cartGroup.POST("/product/:product_id/provider/:provider_id/user/:user_id", allHandler.cartItemHandler.UpsertCartItem())
			cartGroup.PATCH("/:cart_id/items/:cart_item_id", allHandler.cartItemHandler.UpdateCartItem())
			cartGroup.DELETE("/:cart_id/items/:cart_item_id", allHandler.cartItemHandler.DeleteCartItem())
		}
		paymentGroup := v1.Group("/payments")
		{
			paymentGroup.POST("", allHandler.paymentHandler.CreatePayment())
			paymentGroup.GET("/:payment_id", allHandler.paymentHandler.GetPaymentByID())
		}

		orderGroup := v1.Group("/orders")
		{
			orderAdminGroup := orderGroup.Group("")
			orderAdminGroup.Use(allHandler.jwtHandler.VerifyAdminToken())

			orderGroup.POST("", allHandler.orderHandler.CreateOrder())
			orderGroup.PATCH("/:order_id", allHandler.orderHandler.UpdateOrderStatus())
			orderGroup.PATCH("/:order_id/update", allHandler.orderHandler.UpdateOrder())
			orderGroup.PATCH("/payment", allHandler.orderHandler.UpdateOrderPayment())
			orderGroup.PATCH("/:order_id/user/:user_id/cancel", allHandler.orderHandler.CancelOrder())
			orderGroup.PATCH("/:order_id/user/:user_id/verify", allHandler.orderHandler.VerifyDeliveredStatus())
			orderGroup.DELETE("/:order_id", allHandler.orderHandler.DeleteOrder())

			orderGroup.GET("/:order_id/items", allHandler.orderItemHandler.ListByOrderID())
			orderGroup.GET("/user/:user_id", allHandler.orderHandler.ListByUserID())
			orderGroup.GET("/user/:user_id/preview", allHandler.orderHandler.ListPreviewByUserID())
			orderAdminGroup.GET("/admin/:user_id", allHandler.orderHandler.List())
			orderAdminGroup.GET("/admin/:user_id/preview", allHandler.orderHandler.ListPreview())
			orderGroup.GET("/provider/:provider_id", allHandler.orderHandler.ListByProviderID())
			orderGroup.GET("/provider/:provider_id/preview", allHandler.orderHandler.ListPreviewByProviderID())
		}
		addressGroup := v1.Group("/addresses")
		{
			addressGroup.Use(allHandler.jwtHandler.VerifyToken())

			authGroup := addressGroup.Group("")
			authGroup.Use(allHandler.jwtHandler.VerifyUserToken())

			authGroup.GET("/:address_id/user/:user_id", allHandler.addressHandler.GetAddressDetailByID())
			authGroup.GET("/user/:user_id", allHandler.addressHandler.GetAddressesWithUserID())
			authGroup.POST("/user/:user_id", allHandler.addressHandler.CreateAddress())
			authGroup.DELETE("/user/:user_id", allHandler.addressHandler.DeleteAddressByIDs())
			authGroup.PATCH("/:address_id/user/:user_id/", allHandler.addressHandler.UpdateAddress())
		}

	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//Dummy Schedule job
	c := gron.New()
	c.AddFunc(gron.Every(1*time.Hour), func() {
		err := allHandler.orderHandler.RemoveInvalidOrder()
		if err != nil {
			log.Fatal(err)
		}
	})
	c.Start()

}
