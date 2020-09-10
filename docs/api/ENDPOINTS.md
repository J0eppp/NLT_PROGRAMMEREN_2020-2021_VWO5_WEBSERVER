# API endpoints

* /api/
    * /v1/
        * /product/
            * /?product={int}           (GET) - Get the product from the AH API based on a barcode                                          (READY)
            * /?product={string}        (GET) - Get the product from the AH API based on a name                                             (NOT READY)
            * /search?query={string}    (GET) - Search in the local database for a product (by title, barcode, category and subcategory)    (NOT READY)
            