class Bsdoc < Formula
  desc "Generate Backstage catalog documentation using GitHub Copilot"
  homepage "https://github.com/grahambrooks/bsdoc"
  version "2026.11.22"

  on_macos do
    on_arm do
      url "https://github.com/grahambrooks/bsdoc/releases/download/v2026.11.22/bsdoc_2026.11.22_darwin_arm64.tar.gz"
      sha256 "0d4f4a3999476d3162a0a8c3c938637a8c938598fb328fdc99e0c3bd6f786be0"
    end
    on_intel do
      url "https://github.com/grahambrooks/bsdoc/releases/download/v2026.11.22/bsdoc_2026.11.22_darwin_amd64.tar.gz"
      sha256 "a9b139ba86fafe14ab40e49ff88a8d4e5d44bd64959fb52e9f0dacf9b092deed"
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/grahambrooks/bsdoc/releases/download/v2026.11.22/bsdoc_2026.11.22_linux_arm64.tar.gz"
      sha256 "3c894ed52566434394a97d962977a6640605b18da9448f6e60cf87ffcdc813e1"
    end
    on_intel do
      url "https://github.com/grahambrooks/bsdoc/releases/download/v2026.11.22/bsdoc_2026.11.22_linux_amd64.tar.gz"
      sha256 "0440517ee83dcf9f0f508a13991240505901a29edc8c72c3a0adf393c8a88178"
    end
  end

  def install
    bin.install "bsdoc"
  end

  test do
    assert_match version.to_s, shell_output("#{bin}/bsdoc version")
  end
end
