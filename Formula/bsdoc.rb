class Bsdoc < Formula
  desc "Generate Backstage catalog documentation using GitHub Copilot"
  homepage "https://github.com/grahambrooks/bsdoc"
  version "0.0.0"

  on_macos do
    on_arm do
      url "https://github.com/grahambrooks/bsdoc/releases/download/v0.0.0/bsdoc_0.0.0_darwin_arm64.tar.gz"
      sha256 "0000000000000000000000000000000000000000000000000000000000000000"
    end
    on_intel do
      url "https://github.com/grahambrooks/bsdoc/releases/download/v0.0.0/bsdoc_0.0.0_darwin_amd64.tar.gz"
      sha256 "0000000000000000000000000000000000000000000000000000000000000000"
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/grahambrooks/bsdoc/releases/download/v0.0.0/bsdoc_0.0.0_linux_arm64.tar.gz"
      sha256 "0000000000000000000000000000000000000000000000000000000000000000"
    end
    on_intel do
      url "https://github.com/grahambrooks/bsdoc/releases/download/v0.0.0/bsdoc_0.0.0_linux_amd64.tar.gz"
      sha256 "0000000000000000000000000000000000000000000000000000000000000000"
    end
  end

  def install
    bin.install "bsdoc"
  end

  test do
    assert_match version.to_s, shell_output("#{bin}/bsdoc version")
  end
end
