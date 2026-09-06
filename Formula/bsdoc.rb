class Bsdoc < Formula
  desc "Generate Backstage catalog documentation using GitHub Copilot"
  homepage "https://github.com/grahambrooks/bsdoc"
  version "2026.5.24"

  on_macos do
    on_arm do
      url "https://github.com/grahambrooks/bsdoc/archive/refs/tags/v2026.9.5.tar.gz"
      sha256 "f91d54ee1dc21b046a5eb5a03d71ac7d0be9f99ed6387234f381f89449f9fb09"
    end
    on_intel do
      url "https://github.com/grahambrooks/bsdoc/releases/download/v2026.5.24/bsdoc_2026.5.24_darwin_amd64.tar.gz"
      sha256 "b06d47bbaaea69935116ac6dbdabdea5719c0ad66f2803400d4c3b2eff95cf39"
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/grahambrooks/bsdoc/releases/download/v2026.5.24/bsdoc_2026.5.24_linux_arm64.tar.gz"
      sha256 "8c1a9848461f83baa67c31bf517ab6f92bed271b73c4492e5c967cd5c11b5f3d"
    end
    on_intel do
      url "https://github.com/grahambrooks/bsdoc/releases/download/v2026.5.24/bsdoc_2026.5.24_linux_amd64.tar.gz"
      sha256 "2007ab2cdc07d3927c27fba0dc0adaf932c7e1170d8a73eb018095676c3a3f53"
    end
  end

  def install
    bin.install "bsdoc"
  end

  test do
    assert_match version.to_s, shell_output("#{bin}/bsdoc version")
  end
end
