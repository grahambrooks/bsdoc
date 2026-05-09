class Bsdoc < Formula
  desc "Generate Backstage catalog documentation using GitHub Copilot"
  homepage "https://github.com/grahambrooks/bsdoc"
  version "2026.5.9"

  on_macos do
    on_arm do
      url "https://github.com/grahambrooks/bsdoc/releases/download/v2026.5.9/bsdoc_2026.5.9_darwin_arm64.tar.gz"
      sha256 "e9aaec1bbd51d8804c288a986aef06f1dc33acae6fc7ddc381b18ae44d51488c"
    end
    on_intel do
      url "https://github.com/grahambrooks/bsdoc/releases/download/v2026.5.9/bsdoc_2026.5.9_darwin_amd64.tar.gz"
      sha256 "f2332ded6c63758d43431eb749219f0f370d9bfcfe3c24f63ca57733ed224137"
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/grahambrooks/bsdoc/releases/download/v2026.5.9/bsdoc_2026.5.9_linux_arm64.tar.gz"
      sha256 "05befce1f69abcde6e877c76ea3971fa9f67c61b1ccd97cfac325f3da9a21a7c"
    end
    on_intel do
      url "https://github.com/grahambrooks/bsdoc/releases/download/v2026.5.9/bsdoc_2026.5.9_linux_amd64.tar.gz"
      sha256 "864f512f7d4568ef9b8bb6ede7a9d8739adb68cb77f66c7b8fca18466e4cd9fd"
    end
  end

  def install
    bin.install "bsdoc"
  end

  test do
    assert_match version.to_s, shell_output("#{bin}/bsdoc version")
  end
end
