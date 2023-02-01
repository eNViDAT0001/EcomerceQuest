# About Project
#### This project is a simple Golang Back-end build for Ecommerce system follow Clean Architech, domains include:
```text
- Cart
  - Carts
  - CartItems
- Order
  - Orders
  - OrderItems
- Product
  - Properties
    - Options
  - Media
  - Descriptions (using Markdown)
- Store
  - Provider
  - Brand
  - Category
  - Favorite
- Verification
  - JWT
- File Storage
  - Cloudinary
- User
  - User
  - Address
```
Moreover, I'm also write some package in external folder:
```text
- fakeRedis: A simple cache package for cache product's quantity to avoid race condition
- paging: Get Params from URL and do some Paging stuff
- wrap_stuff: Get singleton instance and do stuffs
- request: Reduce some code in Gin
```
