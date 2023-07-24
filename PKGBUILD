# Maintainer: Orhun Parmaksız <orhun@archlinux.org>
# Contributor: Evan McCarthy <evanmccarthy@outlook.com>
# Contributor: Oleksandr Natalenko <oleksandr@natalenko.name>

pkgname=tere
pkgver=1.5.0
pkgrel=1
pkgdesc="A terminal file explorer"
arch=('x86_64')
url="https://github.com/mgunyho/tere"
license=("custom:EUPL")
makedepends=('cargo')
source=("$pkgname-$pkgver.tar.gz::$url/archive/v$pkgver.tar.gz")
sha256sums=('84195f45b738fb7c805d7b348185658d0dc58aa26e7f92fcad9578bc7bd694bf')

prepare() {
  cd "$srcdir/$pkgname-$pkgver"

	cargo fetch --locked --target "$CARCH-unknown-linux-gnu"
}

build() {
	cd "$srcdir/$pkgname-$pkgver"

	cargo build --release --frozen
}

check() {
	cd "$srcdir/$pkgname-$pkgver"

	cargo test --frozen
}

package() {
	cd ${pkgname}-${pkgver}

	install -Dt "${pkgdir}"/usr/bin -m0755 target/release/${pkgname}
	install -Dt "${pkgdir}"/usr/share/licenses/${pkgname} -m0644 LICENSE
	install -Dt "${pkgdir}"/usr/share/doc/${pkgname} -m0644 README.md
	install -Dt "${pkgdir}"/usr/share/doc/${pkgname}/demo -m0644 demo/*
}

# vim:set ts=2 sw=2 et:
