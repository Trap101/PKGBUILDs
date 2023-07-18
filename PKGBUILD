# Maintainer: Orhun Parmaksız <orhun@archlinux.org>
# Maintainer: Caleb Maclennan <caleb@alerque.com>

pkgname=himalaya
pkgver=0.8.4
pkgrel=1
pkgdesc="A CLI email client"
arch=('x86_64')
url="https://github.com/soywod/himalaya"
license=('BSD')
depends=('gcc-libs' 'notmuch' 'openssl')
makedepends=('cargo')
source=("$url/archive/v$pkgver/$pkgname-$pkgver.tar.gz")
sha512sums=('7b67bee9ebfd5fdfefd6dc2caa4337fb9174d81d2fed13e124ca8e8de538e87f8a39fd93f7c81bf560f5bcaddc1d7408af4be58ee551cd29bb259ca1cfa12645')

prepare() {
  cd "$pkgname-$pkgver"
  cargo fetch --locked --target "$CARCH-unknown-linux-gnu"
  mkdir -p {completions,man}
}

build() {
  cd "$pkgname-$pkgver"
  CFLAGS+=" -ffat-lto-objects"
  cargo build --frozen --release --features notmuch-backend
  local _completion="target/release/$pkgname completion"
  $_completion bash > "completions/$pkgname"
  $_completion fish > "completions/$pkgname.fish"
  $_completion zsh  > "completions/_$pkgname"
  local _mangen="target/release/$pkgname man"
  $_mangen man
}

check() {
  cd "$pkgname-$pkgver"
  cargo test --frozen --features notmuch-backend -- --skip "test_imap_backend"
}

package_himalaya() {
  cd "$pkgname-$pkgver"
  install -Dm 755 "target/release/$pkgname" -t "$pkgdir/usr/bin"
  install -Dm 664 "completions/$pkgname" -t "$pkgdir/usr/share/bash-completion/completions/"
  install -Dm 664 "completions/$pkgname.fish" -t "$pkgdir/usr/share/fish/vendor_completions.d/"
  install -Dm 664 "completions/_$pkgname" -t "$pkgdir/usr/share/zsh/site-functions/"
  install -Dm 644 README.md -t "$pkgdir/usr/share/doc/$pkgname"
  install -Dm 644 LICENSE -t "$pkgdir/usr/share/licenses/$pkgname"
  install -Dm 644 "assets/$pkgname.desktop" -t "$pkgdir/usr/share/applications"
  find man/ -type f -exec install -Dm 644 -t "$pkgdir/usr/share/man/man1" {} \;
}

# vim:set ts=2 sw=2 et:
