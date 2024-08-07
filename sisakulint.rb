require_relative "strategy" # typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Sisakulint < Formula
  desc "Support tools for GitHub Actions workflow files"
  homepage "https://github.com/ultra-supara/sisakulint"
  version "0.0.7"

  on_macos do
    on_intel do
      url "https://github.com/ultra-supara/sisakulint/releases/download/v0.0.7/sisakulint_0.0.7_darwin_amd64.tar.gz", using: GitHubPrivateRepositoryReleaseDownloadStrategy
      sha256 "b0a6a2061fb61b64dc3c0b951cffcf01aea8b51465566b0184b52242984ba48d"

      def install
        bin.install "sisakulint"
      end
    end
    on_arm do
      url "https://github.com/ultra-supara/sisakulint/releases/download/v0.0.7/sisakulint_0.0.7_darwin_arm64.tar.gz", using: GitHubPrivateRepositoryReleaseDownloadStrategy
      sha256 "f20be9809f34c11d0c47692de02beff27e8b33a546f09cf1909232c0ffd72ed9"

      def install
        bin.install "sisakulint"
      end
    end
  end

  on_linux do
    on_intel do
      if Hardware::CPU.is_64_bit?
        url "https://github.com/ultra-supara/sisakulint/releases/download/v0.0.7/sisakulint_0.0.7_linux_amd64.tar.gz", using: GitHubPrivateRepositoryReleaseDownloadStrategy
        sha256 "077f20ae27aed6a4182865f779c8799cb747f0ac1e45593e8bff3a7edc2bf8b1"

        def install
          bin.install "sisakulint"
        end
      end
    end
    on_arm do
      if !Hardware::CPU.is_64_bit?
        url "https://github.com/ultra-supara/sisakulint/releases/download/v0.0.7/sisakulint_0.0.7_linux_armv6.tar.gz", using: GitHubPrivateRepositoryReleaseDownloadStrategy
        sha256 "62803f3d7f12beb6f1cf96d9afd2bd9e8d324014e4d3f5752f687f6a2de8f63c"

        def install
          bin.install "sisakulint"
        end
      end
    end
    on_arm do
      if Hardware::CPU.is_64_bit?
        url "https://github.com/ultra-supara/sisakulint/releases/download/v0.0.7/sisakulint_0.0.7_linux_arm64.tar.gz", using: GitHubPrivateRepositoryReleaseDownloadStrategy
        sha256 "0c734c8c8d3e34b432c031d3705651d935d7923786c13a0f73a5ff20c5c8075b"

        def install
          bin.install "sisakulint"
        end
      end
    end
  end

  test do
    system "#{bin}/sisakulint -version"
  end
end
