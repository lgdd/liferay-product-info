const http = require('https');
const axios = require('axios');
const fs = require('fs');
const { execSync } = require('child_process');
const cliProgress = require('cli-progress');

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
      writeTypedProductsJson(betterJson, 'dxp', '7.4');
      writeTypedProductsJson(betterJson, 'dxp', '7.3');
      writeTypedProductsJson(betterJson, 'dxp', '7.2');
      writeTypedProductsJson(betterJson, 'dxp', '7.1');
      writeTypedProductsJson(betterJson, 'dxp', '7.0');

      writeTypedProductsJson(betterJson, 'portal');
      writeTypedProductsJson(betterJson, 'portal', '7.4');
      writeTypedProductsJson(betterJson, 'portal', '7.3');
      writeTypedProductsJson(betterJson, 'portal', '7.2');
      writeTypedProductsJson(betterJson, 'portal', '7.1');
      writeTypedProductsJson(betterJson, 'portal', '7.0');

      writeTypedProductsJson(betterJson, 'commerce');
    });
  });
});

const buildBetterJson = (productJson) => {
  const betterJson = [];
  Object.keys(productJson).forEach((key) => {
    productJson[key]['name'] = key;
    betterJson.push(productJson[key]);
  });

  if (process.env.DECODE_BUNDLE_URLS) {
    const progressbar = new cliProgress.SingleBar(
      {
        format:
          'Decoding Bundle URLs [{bar}] {percentage}% | ETA: {eta}s | {value}/{total}',
      },
      cliProgress.Presets.shades_classic
    );
    progressbar.start(betterJson.length, 0);

    betterJson.forEach((product) => {
      product.bundleUrl = decodeBundleUrl(
        product.bundleUrl,
        product.releaseDate
      );
      product.bundleChecksumMD5Url = decodeBundleUrl(
        product.bundleChecksumMD5Url,
        product.releaseDate
      );
      progressbar.increment();
    });

    progressbar.stop();
  }

  fs.writeFile(
    '../../better_product_info.json',
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
    '../../promoted_product_info.json',
    JSON.stringify(promotedProducts, null, '\t'),
    function (err) {
      if (err) throw err;
      console.log('Write promoted_product_info.json');
    }
  );
};

const writeTypedProductsJson = (betterJson, productType, version) => {
  version = version === undefined ? '' : version;
  const prefix =
    version === ''
      ? productType
      : `${productType}_${version.replaceAll('.', '')}`;
  let specificProducts = [];
  betterJson.forEach((product) => {
    if (product.name.startsWith(`${productType}-${version}`)) {
      specificProducts.push(product);
    }
  });

  if(productType === "portal" && version === "7.2") {
    specificProducts = specificProducts.reverse()
  }

  if(productType === "portal" && version === "7.1") {
    specificProducts = specificProducts.reverse()
  }

  fs.writeFile(
    `../../${prefix}_product_info.json`,
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

// Liferay Cloud Images

const latestStableVersionRegex =
  /^(\d+\.\d+\.\d+(-jdk\d+)?|^\d+\.\d+(-jdk\d+)?)(-\d+\.\d+\.\d+)?$/;
const dockerHubRepoApiBaseUrl =
  'https://registry.hub.docker.com/v2/repositories';

const liferayCloudImages = [
  'liferaycloud/backup',
  'liferaycloud/jenkins',
  'liferaycloud/database',
  'liferaycloud/liferay-dxp',
  'liferaycloud/elasticsearch',
  'liferaycloud/nginx',
];

const fetchLatestTag = async (liferayCloudImage, page) => {
  try {
    const response = await axios.get(
      `${dockerHubRepoApiBaseUrl}/${liferayCloudImage}/tags?page_size=1024&page=${page}`
    );
    let results = response.data.results;
    latestTag = findLatestTag(liferayCloudImage, results);
    if (latestTag === '') {
      return await fetchLatestTag(liferayCloudImage, page + 1);
    } else {
      return latestTag;
    }
  } catch (error) {
    console.error(error);
  }
};

const findLatestTag = (liferayCloudImage, results) => {
  for (let i = 0; i < results.length; i++) {
    if (latestStableVersionRegex.test(results[i].name)) {
      results[i][
        'latest_docker_image'
      ] = `${liferayCloudImage}:${results[i].name}`;
      return results[i];
    }
  }
  return '';
};

const fetchLatestTagPromises = liferayCloudImages.map(
  async (liferayCloudImage) => {
    return await fetchLatestTag(liferayCloudImage, 1);
  }
);

const writeLatestLiferayCloudImages = async () => {
  const latestDockerImages = await Promise.all(fetchLatestTagPromises);
  fs.writeFile(
    '../../liferaycloud_latest_docker_images.json',
    JSON.stringify(latestDockerImages, null, '\t'),
    function (err) {
      if (err) throw err;
      console.log('Write liferaycloud_latest_docker_images.json');
    }
  );
};

writeLatestLiferayCloudImages();
