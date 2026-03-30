const { execSync } = require('child_process');
const fs = require('fs');
const path = require('path');
const os = require('os');

const BIN_DIR = path.join(__dirname, '..', 'bin');
const BINARY_NAME = 'whros';

function getPlatform() {
  const platform = os.platform();
  const arch = os.arch();
  const platformMap = {
    'darwin': { name: 'darwin', ext: '' },
    'linux': { name: 'linux', ext: '' },
    'win32': { name: 'windows', ext: '.exe' }
  };
  const archMap = {
    'x64': 'amd64',
    'arm64': 'arm64'
  };
  const p = platformMap[platform];
  const a = archMap[arch] || arch;
  return { platform: p.name, arch: a, ext: p.ext };
}

function getDownloadUrl(version) {
  const { platform, arch, ext } = getPlatform();
  const name = `whros-${platform}-${arch}${ext}`;
  return `https://github.com/你的用户名/whros-cli/releases/download/${version}/${name}`;
}

function downloadBinary(version) {
  const url = getDownloadUrl(version);
  const { ext } = getPlatform();
  const outputPath = path.join(BIN_DIR, BINARY_NAME + ext);

  console.log(`Downloading whros ${version} for ${os.platform()} ${os.arch()}...`);
  console.log(`URL: ${url}`);

  try {
    execSync(`curl -L -o "${outputPath}" "${url}"`, { stdio: 'inherit' });

    if (os.platform() !== 'win32') {
      fs.chmodSync(outputPath, 0o755);
    }

    console.log(`Installed to: ${outputPath}`);
  } catch (error) {
    console.error('Download failed. Please manually download from:', url);
    process.exit(1);
  }
}

function main() {
  const version = process.env.npm_package_version || '1.0.0';

  if (!fs.existsSync(BIN_DIR)) {
    fs.mkdirSync(BIN_DIR, { recursive: true });
  }

  downloadBinary(version);
}

main();
