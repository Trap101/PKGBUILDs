# Maintainer: Orhun Parmaksız <orhun@archlinux.org>

pkgname=arch-repro-status
pkgver=1.3.9
pkgrel=1
pkgdesc="Check the reproducibility status of your Arch Linux packages"
arch=('x86_64')
url="https://gitlab.archlinux.org/archlinux/arch-repro-status"
license=('MIT')
depends=('pacman')
makedepends=('cargo')
groups=('archlinux-tools')
source=("$pkgname-$pkgver.tar.gz::$url/-/archive/v$pkgver/$pkgname-v$pkgver.tar.gz")
sha512sums=('d6df4e10a6f9083d50020009ffa5b1908afc56e87f0b4f9e5084ead02fa97e5b04bb6b56cb8a8cded6180b60d5b08488cb74dcbede607883e3c89b6cd9b21de6')

prepare() {
  cd "$pkgname-v$pkgver"
  mkdir completions/
  cargo fetch --locked --target "$CARCH-unknown-linux-gnu"
}

build() {
  cd "$pkgname-v$pkgver"
  cargo build --frozen --release
  OUT_DIR=completions/ ./target/release/$pkgname-completions
}

check() {
  cd "$pkgname-v$pkgver"
  cargo test --frozen
}

package() {
  cd "$pkgname-v$pkgver"
  install -Dm 755 "target/release/$pkgname" -t "$pkgdir/usr/bin"
  install -Dm 644 README.md -t "$pkgdir/usr/share/doc/$pkgname"
  install -Dm 644 LICENSE -t "$pkgdir/usr/share/licenses/$pkgname"
  install -Dm 644 "man/$pkgname.1" -t "$pkgdir/usr/share/man/man1"
  install -Dm 644 "completions/$pkgname.bash" "$pkgdir/usr/share/bash-completion/completions/$pkgname"
  install -Dm 644 "completions/$pkgname.fish" -t "$pkgdir/usr/share/fish/vendor_completions.d"
  install -Dm 644 "completions/_$pkgname" -t "$pkgdir/usr/share/zsh/site-functions"
}
