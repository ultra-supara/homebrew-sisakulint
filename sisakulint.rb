# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Sisakulint < Formula
  desc "Support tools for GitHub Actions workflow files"
  homepage "https://github.com/ultra-supara/sisakulint"
  version "0.0.15"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/ultra-supara/homebrew-sisakulint/releases/download/v0.0.15/sisakulint_0.0.15_darwin_arm64.tar.gz", using: GitHubPrivateRepositoryReleaseDownloadStrategy
      sha256 "397c292c2a318e17ebfd9c67af03f21010f83bccf6f1ee0da36d9e478eae29a8"

      def install
        bin.install "sisakulint"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/ultra-supara/homebrew-sisakulint/releases/download/v0.0.15/sisakulint_0.0.15_darwin_amd64.tar.gz", using: GitHubPrivateRepositoryReleaseDownloadStrategy
      sha256 "1eb3a550883739e4ad72fbe6b20600f61131cb128833186a819ae58b13e3acbf"

      def install
        bin.install "sisakulint"
      end
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/ultra-supara/homebrew-sisakulint/releases/download/v0.0.15/sisakulint_0.0.15_linux_arm64.tar.gz", using: GitHubPrivateRepositoryReleaseDownloadStrategy
      sha256 "e77036ba1471af71b88224245115365a3789a4423316055c27471c890aae8a96"

      def install
        bin.install "sisakulint"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/ultra-supara/homebrew-sisakulint/releases/download/v0.0.15/sisakulint_0.0.15_linux_amd64.tar.gz", using: GitHubPrivateRepositoryReleaseDownloadStrategy
      sha256 "392df7476bc41c3282a2ba7c7ba9e23c8c269396ee3a646f1681a59a1244bea4"

      def install
        bin.install "sisakulint"
      end
    end
    if Hardware::CPU.arm? && !Hardware::CPU.is_64_bit?
      url "https://github.com/ultra-supara/homebrew-sisakulint/releases/download/v0.0.15/sisakulint_0.0.15_linux_armv6.tar.gz", using: GitHubPrivateRepositoryReleaseDownloadStrategy
      sha256 "b97471c89b16c954c03164562d2fdf762788382ab44b0ec7bc29f18517c2b312"

      def install
        bin.install "sisakulint"
      end
    end
  end

  test do
    system "#{bin}/sisakulint -version"
  end
end
