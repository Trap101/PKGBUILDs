# Maintainer: Orhun Parmaksız <orhun@archlinux.org>

pkgname=arduino-language-server
pkgver=0.7.4
pkgrel=2
pkgdesc="An Arduino Language Server based on Clangd to Arduino code autocompletion"
arch=('x86_64')
url="https://github.com/arduino/arduino-language-server"
license=('AGPL3')
depends=('glibc' 'arduino-cli' 'clang')
makedepends=('go')
source=("$pkgname-$pkgver.tar.gz::$url/archive/$pkgver.tar.gz")
sha512sums=('c1f4a4ff1f525a812495418b95742357954426018366f6f3f211736291508a254e675fadcd0fdb6c21826f5aa60519c94c169b1e338ab5afd4d130ea1f794937')

build() {
  cd "$pkgname-$pkgver"
  export CGO_CPPFLAGS="${CPPFLAGS}"
  export CGO_CFLAGS="${CFLAGS}"
  export CGO_CXXFLAGS="${CXXFLAGS}"
  export CGO_LDFLAGS="${LDFLAGS}"
  export GOFLAGS="-buildmode=pie -trimpath -ldflags=-linkmode=external -mod=readonly -modcacherw"
  go build -o "$pkgname" .
}

package() {
  cd "$pkgname-$pkgver"
  install -Dm 755 "$pkgname" -t "$pkgdir/usr/bin"
  install -Dm 644 README.md -t "$pkgdir/usr/share/doc/$pkgname"
}

# vim: ts=2 sw=2 et:
