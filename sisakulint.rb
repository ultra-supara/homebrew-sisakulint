# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Sisakulint < Formula
  desc "Support tools for GitHub Actions workflow files"
  homepage "https://github.com/ultra-supara/sisakulint"
  version "0.0.4"

  on_macos do
    if Hardware::CPU.intel?
      url "https://github.com/ultra-supara/homebrew-sisakulint/releases/download/v0.0.4/sisakulint_0.0.4_darwin_amd64.tar.gz"
      sha256 "98cdf16893202612a100304add304a0f7f740ba9a96bb24e3796901ffc684a91"

      def install
        bin.install "sisakulint"
      end
    end
    if Hardware::CPU.arm?
      url "https://github.com/ultra-supara/homebrew-sisakulint/releases/download/v0.0.4/sisakulint_0.0.4_darwin_arm64.tar.gz"
      sha256 "74a1764a2d1c183fa8cde931e8ed29b13a38b91edb37bc223993b6da54ed792c"

      def install
        bin.install "sisakulint"
      end
    end
  end

  on_linux do
    if Hardware::CPU.intel?
      url "https://github.com/ultra-supara/homebrew-sisakulint/releases/download/v0.0.4/sisakulint_0.0.4_linux_amd64.tar.gz"
      sha256 "37e65d36b1faaadf52a7ba7f0daab15ede71c7ec3e5ea0c62aa54a869e448aae"

      def install
        bin.install "sisakulint"
      end
    end
    if Hardware::CPU.arm? && !Hardware::CPU.is_64_bit?
      url "https://github.com/ultra-supara/homebrew-sisakulint/releases/download/v0.0.4/sisakulint_0.0.4_linux_armv6.tar.gz"
      sha256 "cd04273192616a5072f24f4d2f8a52d37176f6763be4b1c95eaf3908534edbd0"

      def install
        bin.install "sisakulint"
      end
    end
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/ultra-supara/homebrew-sisakulint/releases/download/v0.0.4/sisakulint_0.0.4_linux_arm64.tar.gz"
      sha256 "f02199c668fb1158f3324fef824e206455fede154a3b7f3ae3f5bf2d794c0bf7"

      def install
        bin.install "sisakulint"
      end
    end
  end

  test do
    system "#{bin}/sisakulint -version"
  end
end
