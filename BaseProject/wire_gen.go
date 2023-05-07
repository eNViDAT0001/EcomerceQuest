// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/address"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/app/app_accession"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/app/app_file"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/cart/cart"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/cart/cart_items"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/chat"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/comment"
	chat2 "github.com/eNViDAT0001/Thesis/Backend/delivery/http/notification"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/order/order"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/order/order_items"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/product"
	banner2 "github.com/eNViDAT0001/Thesis/Backend/delivery/http/store/banner"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/store/category"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/store/favorite"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/store/provider"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/user"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/verification/jwt"
	"github.com/eNViDAT0001/Thesis/Backend/delivery/http/verification/smtp"
	storage3 "github.com/eNViDAT0001/Thesis/Backend/internal/address/domain/address/storage"
	usecase3 "github.com/eNViDAT0001/Thesis/Backend/internal/address/domain/address/usecase"
	usecase5 "github.com/eNViDAT0001/Thesis/Backend/internal/app/domain/app_accession/usecase"
	storage11 "github.com/eNViDAT0001/Thesis/Backend/internal/cart/domain/cart/storage"
	usecase12 "github.com/eNViDAT0001/Thesis/Backend/internal/cart/domain/cart/usecase"
	storage12 "github.com/eNViDAT0001/Thesis/Backend/internal/cart/domain/cart_item/storage"
	usecase13 "github.com/eNViDAT0001/Thesis/Backend/internal/cart/domain/cart_item/usecase"
	storage17 "github.com/eNViDAT0001/Thesis/Backend/internal/chat/domain/chat/storage"
	usecase18 "github.com/eNViDAT0001/Thesis/Backend/internal/chat/domain/chat/usecase"
	storage7 "github.com/eNViDAT0001/Thesis/Backend/internal/file_storage/domain/media/storage"
	usecase9 "github.com/eNViDAT0001/Thesis/Backend/internal/file_storage/domain/media/usecase"
	storage15 "github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/notification/storage"
	usecase15 "github.com/eNViDAT0001/Thesis/Backend/internal/notify/domain/notification/usecase"
	storage13 "github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/order/storage"
	usecase16 "github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/order/usecase"
	storage16 "github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/order_item/storage"
	usecase17 "github.com/eNViDAT0001/Thesis/Backend/internal/order/domain/order_item/usecase"
	storage10 "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/comment/storage"
	usecase10 "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/comment/usecase"
	storage6 "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/storage"
	usecase7 "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/product/usecase"
	storage8 "github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/banner/storage"
	usecase11 "github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/banner/usecase"
	storage4 "github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/category/storage"
	usecase4 "github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/category/usecase"
	storage9 "github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/favorite/storage"
	usecase8 "github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/favorite/usecase"
	storage5 "github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/provider/storage"
	usecase6 "github.com/eNViDAT0001/Thesis/Backend/internal/store/domain/provider/usecase"
	"github.com/eNViDAT0001/Thesis/Backend/internal/user/domain/user/storage"
	"github.com/eNViDAT0001/Thesis/Backend/internal/user/domain/user/usecase"
	storage2 "github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/jwt/storage"
	usecase2 "github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/jwt/usecase"
	storage14 "github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/smtp/storage"
	usecase14 "github.com/eNViDAT0001/Thesis/Backend/internal/verification/domain/smtp/usecase"
)

// Injectors from wire.go:

func initHandlerCollection() *HandlerCollection {
	userStorage := storage.NewUserStorage()
	useCase := usecase.NewUserUseCase(userStorage)
	jwtStorage := storage2.NewJwtStorage()
	jwtUseCase := usecase2.NewJwtUseCase(userStorage, jwtStorage)
	httpHandler := user.NewUserHandler(useCase, jwtUseCase)
	storage18 := storage3.NewAddressStorage()
	userUseCase := usecase3.NewAddressUseCase(storage18)
	userHttpHandler := address.NewAddressHandler(userUseCase)
	categoryStorage := storage4.NewCategoryStorage()
	categoryUseCase := usecase4.NewCategoryUseCase(categoryStorage)
	categoryHttpHandler := category.NewCategoryHandler(categoryUseCase)
	app_accessionUseCase := usecase5.NewAppAccessionUseCase(userStorage, jwtStorage)
	app_accessionHttpHandler := app_accession.NewAppAccessionHandler(jwtUseCase, useCase, app_accessionUseCase)
	providerStorage := storage5.NewProviderStorage()
	providerUseCase := usecase6.NewProviderUseCase(providerStorage)
	productStorage := storage6.NewProductStorage()
	mediaStorage := storage7.NewMediaStorage()
	bannerStorage := storage8.NewBannerStorage()
	productUseCase := usecase7.NewProductUseCase(productStorage, mediaStorage, bannerStorage)
	jwtHttpHandler := jwt.NewJwtHandler(jwtUseCase, useCase, providerUseCase, productUseCase)
	providerHttpHandler := provider.NewProviderHandler(providerUseCase)
	favoriteStorage := storage9.NewFavoriteStorage()
	favoriteUseCase := usecase8.NewFavoriteUseCase(favoriteStorage)
	favoriteHttpHandler := banner.NewFavoriteHandler(favoriteUseCase)
	mediaUseCase := usecase9.NewMediaUseCase(mediaStorage)
	productHttpHandler := product.NewProductHandler(productUseCase, mediaUseCase, categoryUseCase)
	commentStorage := storage10.NewCommentStorage()
	commentUseCase := usecase10.NewCommentUseCase(commentStorage, mediaStorage)
	commentHttpHandler := comment.NewCommentHandler(commentUseCase)
	mediaHttpHandler := app_file.NewMediaHandler(mediaUseCase)
	bannerUseCase := usecase11.NewBannerUseCase(bannerStorage, productStorage)
	bannerHttpHandler := banner2.NewBannerHandler(bannerUseCase)
	cartStorage := storage11.NewCartStorage()
	cartUseCase := usecase12.NewCartUseCase(cartStorage)
	cartHttpHandler := cart.NewCartHandler(cartUseCase)
	cart_itemStorage := storage12.NewCartItemStorage()
	cart_itemUseCase := usecase13.NewCartItemUseCase(cart_itemStorage)
	cart_itemHttpHandler := cart_items.NewCartItemHandler(cart_itemUseCase)
	orderStorage := storage13.NewOrderStorage()
	smtpStorage := storage14.NewSmtpStorage()
	smtpUseCase := usecase14.NewSmtpUseCase(smtpStorage)
	notificationStorage := storage15.NewNotificationStorage()
	notificationUseCase := usecase15.NewNotificationUseCase(notificationStorage)
	orderUseCase := usecase16.NewOrderUseCase(orderStorage, userStorage, smtpUseCase, notificationUseCase)
	orderHttpHandler := order.NewOrderHandler(orderUseCase, smtpUseCase, useCase, notificationUseCase)
	order_itemStorage := storage16.NewOrderItemStorage()
	order_itemUseCase := usecase17.NewOrderItemUseCase(order_itemStorage)
	order_itemHttpHandler := order_items.NewOrderItemHandler(order_itemUseCase)
	smtpHttpHandler := smtp.NewSmtpHandler(jwtUseCase, useCase, smtpUseCase)
	chatStorage := storage17.NewChatStorage()
	chatUseCase := usecase18.NewChatUseCase(chatStorage)
	chatHttpHandler := chat.NewChatHandler(chatUseCase)
	notificationHttpHandler := chat2.NewNotificationHandler(notificationUseCase)
	handlerCollection := NewHandlerCollection(httpHandler, userHttpHandler, categoryHttpHandler, app_accessionHttpHandler, jwtHttpHandler, providerHttpHandler, favoriteHttpHandler, productHttpHandler, commentHttpHandler, mediaHttpHandler, bannerHttpHandler, cartHttpHandler, cart_itemHttpHandler, orderHttpHandler, order_itemHttpHandler, smtpHttpHandler, chatHttpHandler, notificationHttpHandler)
	return handlerCollection
}
