# Directory structure
# API endpoints
## Login
* POST /auth/login
## User
* GET /user/:id
* POST /user/:id
* DELETE /user/:id
## Product
* GET /products
* GET /product/:id
* POST /product
* DELETE /product/:id
## Cart
* POST /cart: Create a new cart
* POST /cart/:id: Put a product in the cart
* GET /cart/:id: Get a cart's contents
* POST /cart/:id/checkout: Take an order
## Order
* GET /orders
* GET /order/:id
* POST /order/:id/pay