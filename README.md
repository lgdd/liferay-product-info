# Liferay Product Info

Liferay produces and uses a JSON file called [.product_info.json](https://releases-cdn.liferay.com/tools/workspace/.product_info.json) for Blade CLI. The JSON format is annoying to parse and bundle URLs are encoded.

This repository offers a better format for this JSON (i.e. root is an array and not an object) in which URLs are decoded. It also offers different JSON files for promoted bundles, dxp bundles (ee), portal bundles (ce) and commerce bundles (prior to 7.4).

These files are automatically updated on a daily basis:

- [better_product_info.json](https://raw.githubusercontent.com/lgdd/liferay-product-info/main/better_product_info.json)
- [promoted_product_info.json](https://raw.githubusercontent.com/lgdd/liferay-product-info/main/promoted_product_info.json)
- [dxp_product_info.json](https://raw.githubusercontent.com/lgdd/liferay-product-info/main/dxp_product_info.json)
- [portal_product_info.json](https://raw.githubusercontent.com/lgdd/liferay-product-info/main/portal_product_info.json)
- [commerce_product_info.json](https://raw.githubusercontent.com/lgdd/liferay-product-info/main/commerce_product_info.json)
