# API endpoints

* /api/
    * /v1/
        * /product/
            * /?product={int}           (GET) - Get the product from the AH API based on a barcode                      (READY)
        * /recipe/
            * /?recipe={int}            (GET) - Get a recipe saved in the database                                      (READY)
            * /search                   (GET) - Get the recipe ID by sending ingredients # USES GET REQUEST BODY        (READY)