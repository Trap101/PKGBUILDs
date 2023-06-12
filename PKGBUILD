# Maintainer: Orhun Parmaksız <orhun@archlinux.org>
# Contributor: Mahor Foruzesh <mahor1221 at gmail dot com>

pkgname=rye
_commit=ce17cf574e76d220306a8cffea237c7b0d6d5686
pkgver=0.7.0
pkgrel=1
pkgdesc="An experimental alternative to poetry, pip, pipenv, venv, virtualenv, pdm, hatch"
arch=('x86_64')
url="https://github.com/mitsuhiko/rye"
license=('MIT')
depends=(
	zlib
	bzip2
	openssl
	gcc-libs
	libxcrypt-compat
)
makedepends=('cargo' 'git')
source=("$pkgname::git+$url.git#commit=$_commit")
sha512sums=('SKIP')
options=('!lto')

prepare() {
	cd "$pkgname"
	cargo fetch --locked --target "$CARCH-unknown-linux-gnu"
}

build() {
	cd "$pkgname"
	cargo build --frozen --release
}

check() {
	cd "$pkgname"
	cargo test --frozen
}

package() {
	cd "$pkgname"
	install -Dm 755 "target/release/$pkgname" -t "$pkgdir/usr/bin"
	install -Dm 644 README.md -t "$pkgdir/usr/share/doc/$pkgname"
	install -Dm 644 LICENSE -t "$pkgdir/usr/share/licenses/$pkgname"
}

# vim: ts=2 sw=2 et:
