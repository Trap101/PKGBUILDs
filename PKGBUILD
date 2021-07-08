# Maintainer: Kaizhao Zhang <zhangkaizhao@gmail.com>

libyaml_version=0.2.5

pkgname=zigmod
pkgver=49
pkgrel=1
pkgdesc="A package manager for the Zig programming language"
url="https://github.com/nektro/zigmod"
license=('MIT')
makedepends=('git' 'zig')
arch=('x86_64')
source=(
  "${pkgname}-${pkgver}.tar.gz::${url}/archive/refs/tags/v${pkgver}.tar.gz"
  "https://github.com/yaml/libyaml/releases/download/${libyaml_version}/yaml-${libyaml_version}.tar.gz"
)
sha256sums=(
  '7690ac0ce06dc43a84223c06cda9b16e88e7019ff4542e371a1e05253f1f2291'
  'c642ae9b75fee120b2d96c712538bd2cf283228d2337df2cf2988e3c02678ef4'
)

prepare() {
  cd "${srcdir}/${pkgname}-${pkgver}"
  rm -r libs/yaml
  ln -s "${srcdir}/yaml-${libyaml_version}" libs/yaml
  zig build -Dbootstrap
  ./zig-out/bin/zigmod fetch
}

build() {
  cd "${srcdir}/${pkgname}-${pkgver}"
  zig build -Dtag=v${pkgver}
}

package() {
  cd "${srcdir}/${pkgname}-${pkgver}"
  install -D -m755 zig-out/bin/zigmod "${pkgdir}/usr/bin/zigmod"
  install -D -m644 README.md "${pkgdir}/usr/share/doc/${pkgname}/README.md"
  cp -r docs "${pkgdir}/usr/share/doc/${pkgname}/docs"
  install -D -m644 LICENSE "${pkgdir}/usr/share/licenses/${pkgname}/LICENSE"
}
