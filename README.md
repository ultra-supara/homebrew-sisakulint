# sisakulint
sisakulint is a private repository tool, but it can be installed from brew.

## install your machine directly for macOS user

```
$ brew tap ultra-supara/homebrew-sisakulint
$ brew install sisakulint
```

## install from release page for Linux user

```
# visit release page of this repository and download for yours.
$ cd < sisakulintがあるところ >
$ mv ./sisakulint /usr/local/bin/sisakulint
```

## Run sisakulint on linux brew using docker
I am so sorry, this method is currently unavailable.
Run the following commands:
```
# Usage
$ git clone https://github.com/ultra-supara/homebrew-sisakulint.git
$ cd homebrew-sisakulint
$ docker compose up -d --build
$ docker attach <createされたcontainer nameに各自で変更してください>
# ここで rootに入ると思います。
$ ls -la
$ brew install sisakulint
```

Brief description of the project. Explain what it does and the problem it solves.

## Links

- [user document](https://www.notion.so/ultra-supara/sisakulint-user-document-d3f28d427cf9456dbe3c0f063a7d3baf?pvs=4)
- [developer document](https://www.notion.so/ultra-supara/sisakulint-c18505b443254ee5a3e5e3751b810a33?pvs=4)

- slides
- poster
- video
