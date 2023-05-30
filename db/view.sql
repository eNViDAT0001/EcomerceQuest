CREATE VIEW UserMessage AS
SELECT User.*
FROM User
         JOIN Message ON Message.to_user_id = User.id
WHERE User.deleted_at IS NULL
UNION
SELECT User.*
FROM User
         JOIN Message ON Message.from_user_id = User.id
WHERE User.deleted_at IS NULL;
//
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
//
CREATE VIEW CartDetailView AS
SELECT
    Cart.id,
    Cart.provider_id,
    Cart.user_id,
    Provider.name,
    Provider.image_path,
    Provider.province,
    Provider.district,
    Provider.ward,
    Provider.province_id,
    Provider.district_id,
    Provider.ward_code,
    Provider.street,
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
                       'weight', Product.weight,
                       'height', Product.height,
                       'length', Product.length,
                       'width', Product.width,
                       'media_path', FirstProductMediaList.media_path))) AS items
FROM Cart
         JOIN Provider ON Provider.id = Cart.provider_id AND Provider.deleted_at IS NULL
         JOIN CartItem ON CartItem.cart_id = Cart.id
         JOIN Product ON Product.id = CartItem.product_id
         LEFT JOIN ProductOption ON CartItem.product_option_id = ProductOption.id
         LEFT JOIN FirstProductMediaList ON FirstProductMediaList.id = Product.id
WHERE Cart.deleted_at IS NULL AND CartItem.deleted_at IS NULL
GROUP BY Cart.id
//////////////////////////
CREATE VIEW ProductPreview AS
SELECT Product.id,
       Product.provider_id,
       Product.category_id,
       Product.user_id,
       Product.name,
       Product.price,
       Product.discount,
       Product.short_descriptions,
       Product.height,
       Product.weight,
       Product.width,
       Product.length,
       IF(COUNT(ProductMedia.id) = 0, NULL,
          JSON_ARRAYAGG(JSON_OBJECT(
                  'publicID', ProductMedia.public_id,
                  'mediaPath', ProductMedia.media_path,
                  'type', ProductMedia.media_type))) AS media,
       ROUND(AVG(Comment.rating),0) AS rating
FROM `Product`
         LEFT JOIN ProductMedia ON ProductMedia.product_id = Product.id
         LEFT JOIN `Comment` ON `Comment`.product_id = Product.id
         JOIN Provider ON Product.provider_id = Provider.id AND Provider.deleted_at IS NULL
WHERE Product.deleted_at IS NULL GROUP BY `Product`.`id`