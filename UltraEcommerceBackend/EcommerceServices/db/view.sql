CREATE VIEW FirstProductMediaList AS
(
WITH cte AS (
   SELECT Product.id, ProductMedia.media_path
        , row_number() OVER (PARTITION BY Product.id ORDER BY ProductMedia.media_path DESC) AS rn
   FROM Product JOIN ProductMedia on Product.id = ProductMedia.product_id
   WHERE Product.deleted_at IS NULL AND ProductMedia.media_type = 'IMAGE'
)
 SELECT * FROM cte WHERE rn = 1
)

CREATE VIEW CartDetailView AS
SELECT
	Cart.id,
	Cart.provider_id,
    Cart.user_id,
	Provider.name,
    Provider.image_path,
        IF(COUNT(CartItem.id) = 0, NULL,
       JSON_ARRAYAGG(
           JSON_OBJECT(
               'id', CartItem.id,
               'product_id', Product.id,
               'option_id', ProductOption.id,
               'option_name', ProductOption.name,
               'name', Product.name,
               'price', IF(ProductOption.id = 0, Product.price, Product.price + ProductOption.price),
               'quantity', CartItem.quantity,
               'discount', Product.discount,
               'media_path', FirstProductMediaList.media_path))) AS items
FROM Cart
JOIN Provider ON Provider.id = Cart.provider_id
JOIN CartItem ON CartItem.cart_id = Cart.id
JOIN Product ON Product.id = CartItem.product_id
LEFT JOIN ProductOption ON CartItem.product_option_id = ProductOption.id
LEFT JOIN FirstProductMediaList ON FirstProductMediaList.id = Product.id
WHERE Cart.deleted_at IS NULL AND CartItem.deleted_at IS NULL
GROUP BY Cart.id