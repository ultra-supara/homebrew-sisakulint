require_relative "strategy" # typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Sisakulint < Formula
  desc "Support tools for GitHub Actions workflow files"
  homepage "https://github.com/ultra-supara/sisakulint"
  version "0.0.21"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/ultra-supara/homebrew-sisakulint/releases/download/v0.0.21/sisakulint_0.0.21_darwin_arm64.tar.gz", using: GitHubPrivateRepositoryReleaseDownloadStrategy
      sha256 "d6d386b048b0e36148063498846f829b782f4083849f630ed920fa04d9468e16"

      def install
        bin.install "sisakulint"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/ultra-supara/homebrew-sisakulint/releases/download/v0.0.21/sisakulint_0.0.21_darwin_amd64.tar.gz", using: GitHubPrivateRepositoryReleaseDownloadStrategy
      sha256 "0c14decf7d8fb56aa0dca86c92469fb8dc4bed5c9e14990b4f6c5463f8f36e94"

      def install
        bin.install "sisakulint"
      end
    end
  end

  on_linux do
    if Hardware::CPU.arm? && !Hardware::CPU.is_64_bit?
      url "https://github.com/ultra-supara/homebrew-sisakulint/releases/download/v0.0.21/sisakulint_0.0.21_linux_armv6.tar.gz", using: GitHubPrivateRepositoryReleaseDownloadStrategy
      sha256 "b0de1f1accb90512849528d063c029daca2201a8d02c0e021b8d31b833da17dc"

      def install
        bin.install "sisakulint"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/ultra-supara/homebrew-sisakulint/releases/download/v0.0.21/sisakulint_0.0.21_linux_amd64.tar.gz", using: GitHubPrivateRepositoryReleaseDownloadStrategy
      sha256 "79ee09dcc6cb46314d47c88cf202af0789cca05907629d7fea2a873274a4e969"

      def install
        bin.install "sisakulint"
      end
    end
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/ultra-supara/homebrew-sisakulint/releases/download/v0.0.21/sisakulint_0.0.21_linux_arm64.tar.gz", using: GitHubPrivateRepositoryReleaseDownloadStrategy
      sha256 "57a4d55431d12c11990076d0bf34166b9427d74d3bc64611b3ab9b3350d14471"

      def install
        bin.install "sisakulint"
      end
    end
  end

  test do
    system "#{bin}/sisakulint -version"
  end
end
