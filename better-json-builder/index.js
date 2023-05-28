const http = require('https');
const fs = require('fs');
const { execSync } = require('child_process');

const productInfoJsonUrl =
  'https://releases-cdn.liferay.com/tools/workspace/.product_info.json';

const productJsonFileName = '.product_info.json';
const productJsonFile = fs.createWriteStream(productJsonFileName);
http.get(productInfoJsonUrl, function (response) {
  response.pipe(productJsonFile);
  productJsonFile.on('finish', () => {
    productJsonFile.close();
    console.log(`Downloaded ${productJsonFileName}`);
    fs.readFile(productJsonFileName, 'utf8', (err, data) => {
      if (err) {
        console.error(err);
        return;
      }
      const betterJson = buildBetterJson(JSON.parse(data));
      writePromotedProductsJson(betterJson);
      writeTypedProductsJson(betterJson, 'dxp');
      writeTypedProductsJson(betterJson, 'portal');
      writeTypedProductsJson(betterJson, 'commerce');
    });
  });
});

const buildBetterJson = (productJson) => {
  const betterJson = [];
  Object.keys(productJson).forEach((key) => {
    productJson[key]['name'] = key;
    productJson[key]['bundleUrl'] = decodeBundleUrl(
      productJson[key].bundleUrl,
      productJson[key].releaseDate
    );
    productJson[key]['bundleChecksumMD5Url'] = decodeBundleUrl(
      productJson[key].bundleChecksumMD5Url,
      productJson[key].releaseDate
    );
    betterJson.push(productJson[key]);
  });

  fs.writeFile(
    '../better_product_info.json',
    JSON.stringify(betterJson, null, '\t'),
    function (err) {
      if (err) throw err;
      console.log('Write better_product_info.json');
    }
  );

  return betterJson;
};

const writePromotedProductsJson = (betterJson) => {
  const promotedProducts = [];
  betterJson.forEach((product) => {
    if (product.promoted === 'true') {
      promotedProducts.push(product);
    }
  });

  fs.writeFile(
    '../promoted_product_info.json',
    JSON.stringify(promotedProducts, null, '\t'),
    function (err) {
      if (err) throw err;
      console.log('Write promoted_product_info.json');
    }
  );
};

const writeTypedProductsJson = (betterJson, productType) => {
  const specificProducts = [];
  betterJson.forEach((product) => {
    if (product.name.startsWith(productType)) {
      specificProducts.push(product);
    }
  });

  fs.writeFile(
    `../${productType}_product_info.json`,
    JSON.stringify(specificProducts, null, '\t'),
    function (err) {
      if (err) throw err;
      console.log(`Write ${productType}_product_info.json`);
    }
  );
};

const decodeBundleUrl = (encodedBundleUrl, releaseDate) => {
  const bundleUrlDecodeResultBytes = execSync(
    `java -cp com.liferay.workspace.bundle.url.codec-1.0.0.jar Main.java ${encodedBundleUrl} ${releaseDate}`
  );
  return bundleUrlDecodeResultBytes.toString().trim();
};
