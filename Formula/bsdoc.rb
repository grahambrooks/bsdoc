class Bsdoc < Formula
  desc "Generate Backstage catalog documentation using GitHub Copilot"
  homepage "https://github.com/grahambrooks/bsdoc"
  version "2026.11.7"

  on_macos do
    on_arm do
      url "https://github.com/grahambrooks/bsdoc/releases/download/v2026.11.7/bsdoc_2026.11.7_darwin_arm64.tar.gz"
      sha256 "80b257ed51fc9863d56498ae24f2535f9d723ab9bc3eaa5388f171afd528e506"
    end
    on_intel do
      url "https://github.com/grahambrooks/bsdoc/releases/download/v2026.11.7/bsdoc_2026.11.7_darwin_amd64.tar.gz"
      sha256 "d1c8f49519dfd4d848f27d1884dd69f0cf7541d3601e593d441a00bba38487d8"
    end
  end

  on_linux do
    on_arm do
      url "https://github.com/grahambrooks/bsdoc/releases/download/v2026.11.7/bsdoc_2026.11.7_linux_arm64.tar.gz"
      sha256 "3e49018f736d94077943a3db9c8876a7d45ce3538d6e882465ae38093e395b2f"
    end
    on_intel do
      url "https://github.com/grahambrooks/bsdoc/releases/download/v2026.11.7/bsdoc_2026.11.7_linux_amd64.tar.gz"
      sha256 "a2542c17a64c07509cad897830cfd0c25354c2c94709ed46a866f723698cf993"
    end
  end

  def install
    bin.install "bsdoc"
  end

  test do
    assert_match version.to_s, shell_output("#{bin}/bsdoc version")
  end
end
