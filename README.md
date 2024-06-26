![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/lgdd/liferay-product-info/builder.yml?label=auto-update&style=flat)
![GitHub last commit](https://img.shields.io/github/last-commit/lgdd/liferay-product-info?color=informational&label=latest%20update)

# Liferay Product Info

Liferay produces and uses a JSON file called [.product_info.json](https://releases-cdn.liferay.com/tools/workspace/.product_info.json) for Blade CLI. The JSON format is annoying to parse and bundle URLs are encoded.

This repository offers a better format for this JSON (i.e. root is an array and not an object) in which URLs are decoded. It also offers different JSON files for promoted bundles, dxp bundles (ee), portal bundles (ce) and commerce bundles (prior to 7.4).

> [!IMPORTANT]
> A new JSON file for releases (and finally easy to parse) is being built by Liferay to take into account the new quarterly releases. This file is also synced into this repository.
> 
> Each release in the [original releases.json](https://releases.liferay.com/releases.json) contains a URL where you can find a list of artifacts, including a `release.properties` file. This repository also mirror them under the [releases](releases) directory. This repository also create a set of new files where each release json entry contains the information from the `release.properties` file.


> [!NOTE]
> **All the files are automatically updated on a daily basis.**

## Releases

- [releases.json](https://raw.githubusercontent.com/lgdd/liferay-product-info/main/releases.json)
- [dxp_releases.json](https://raw.githubusercontent.com/lgdd/liferay-product-info/main/releases/dxp_releases.json)
  - [dxp_74_releases.json](https://raw.githubusercontent.com/lgdd/liferay-product-info/main/releases/dxp_74_releases.json)
  - [dxp_73_releases.json](https://raw.githubusercontent.com/lgdd/liferay-product-info/main/releases/dxp_73_releases.json)
  - [dxp_72_releases.json](https://raw.githubusercontent.com/lgdd/liferay-product-info/main/releases/dxp_72_releases.json)
  - [dxp_71_releases.json](https://raw.githubusercontent.com/lgdd/liferay-product-info/main/releases/dxp_71_releases.json)
  - [dxp_70_releases.json](https://raw.githubusercontent.com/lgdd/liferay-product-info/main/releases/dxp_70_releases.json)
- [portal_releases.json](https://raw.githubusercontent.com/lgdd/liferay-product-info/main/releases/portal_releases.json)
  - [portal_74_releases.json](https://raw.githubusercontent.com/lgdd/liferay-product-info/main/releases/portal_74_releases.json)
  - [portal_73_releases.json](https://raw.githubusercontent.com/lgdd/liferay-product-info/main/releases/portal_73_releases.json)
  - [portal_72_releases.json](https://raw.githubusercontent.com/lgdd/liferay-product-info/main/releases/portal_72_releases.json)
  - [portal_71_releases.json](https://raw.githubusercontent.com/lgdd/liferay-product-info/main/releases/portal_71_releases.json)
  - [portal_70_releases.json](https://raw.githubusercontent.com/lgdd/liferay-product-info/main/releases/portal_70_releases.json)

## Product Info

- [.product_info.json](https://raw.githubusercontent.com/lgdd/liferay-product-info/main/.product_info.json)
- [better_product_info.json](https://raw.githubusercontent.com/lgdd/liferay-product-info/main/better_product_info.json)
- [promoted_product_info.json](https://raw.githubusercontent.com/lgdd/liferay-product-info/main/promoted_product_info.json)
- [dxp_product_info.json](https://raw.githubusercontent.com/lgdd/liferay-product-info/main/dxp_product_info.json)
  - [dxp_74_product_info.json](https://raw.githubusercontent.com/lgdd/liferay-product-info/main/dxp_74_product_info.json)
  - [dxp_73_product_info.json](https://raw.githubusercontent.com/lgdd/liferay-product-info/main/dxp_73_product_info.json)
  - [dxp_72_product_info.json](https://raw.githubusercontent.com/lgdd/liferay-product-info/main/dxp_72_product_info.json)
  - [dxp_71_product_info.json](https://raw.githubusercontent.com/lgdd/liferay-product-info/main/dxp_71_product_info.json)
  - [dxp_70_product_info.json](https://raw.githubusercontent.com/lgdd/liferay-product-info/main/dxp_70_product_info.json)
- [portal_product_info.json](https://raw.githubusercontent.com/lgdd/liferay-product-info/main/portal_product_info.json)
  - [portal_74_product_info.json](https://raw.githubusercontent.com/lgdd/liferay-product-info/main/portal_74_product_info.json)
  - [portal_73_product_info.json](https://raw.githubusercontent.com/lgdd/liferay-product-info/main/portal_73_product_info.json)
  - [portal_72_product_info.json](https://raw.githubusercontent.com/lgdd/liferay-product-info/main/portal_72_product_info.json)
  - [portal_71_product_info.json](https://raw.githubusercontent.com/lgdd/liferay-product-info/main/portal_71_product_info.json)
  - [portal_70_product_info.json](https://raw.githubusercontent.com/lgdd/liferay-product-info/main/portal_70_product_info.json)
- [commerce_product_info.json](https://raw.githubusercontent.com/lgdd/liferay-product-info/main/commerce_product_info.json)
- [liferaycloud_latest_docker_images.json](https://raw.githubusercontent.com/lgdd/liferay-product-info/main/liferaycloud_latest_docker_images.json)
