# Maintainer: Orhun Parmaksız <orhun@archlinux.org>
# Contributor: Arne Beer <privat@arne.beer>

pkgname=pueue
pkgver=3.3.0
pkgrel=1
pkgdesc="A CLI tool for managing long running shell commands"
arch=('x86_64')
url="https://github.com/nukesor/pueue"
license=('MIT')
depends=('gcc-libs')
makedepends=('cargo')
source=("$pkgname-$pkgver.tar.gz::$url/archive/v$pkgver.tar.gz")
sha256sums=('95f9c2744c6b9c43cf9e78864fcc05478aad65527786cb9ab5c58c2b998547de')

prepare() {
  cd "$pkgname-$pkgver"
  cargo fetch --locked --target "$CARCH-unknown-linux-gnu"
  mkdir -p utils/completions/
}

build() {
  cd "$pkgname-$pkgver"
  CFLAGS+=" -ffat-lto-objects"
  cargo build --release --frozen
  ./target/release/pueue completions bash utils/completions/
  ./target/release/pueue completions fish utils/completions/
  ./target/release/pueue completions zsh utils/completions/
}

check() {
  cd "$pkgname-$pkgver"
  cargo test --frozen
}

package() {
  cd "$pkgname-$pkgver"

  # Install binaries
  install -Dm755 "target/release/pueue" "$pkgdir/usr/bin/pueue"
  install -Dm755 "target/release/pueued" "$pkgdir/usr/bin/pueued"

  # Place systemd user service
  install -Dm644 "utils/pueued.service" "$pkgdir/usr/lib/systemd/user/pueued.service"

  # Install shell completions file
  install -Dm644 "utils/completions/_pueue" "$pkgdir/usr/share/zsh/site-functions/_pueue"
  install -Dm644 "utils/completions/pueue.bash" "$pkgdir/usr/share/bash-completion/completions/pueue.bash"
  install -Dm644 "utils/completions/pueue.fish" "$pkgdir/usr/share/fish/completions/pueue.fish"

  # Install License
  install -Dm644 "LICENSE" "$pkgdir/usr/share/licenses/pueue/LICENSE"
}

# vim: ts=2 sw=2 et:
