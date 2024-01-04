# sisakulint
sisakulint is a private repository tool, but it can be installed from brew.
![logo](https://github.com/ultra-supara/homebrew-sisakulint/assets/67861004/e9801cbb-fbe1-4822-a5cd-d1daac33e90f)

## JSON schema for normal system confirmation for GitHub Actions syntax
paste yours `settings.json`
```
 "yaml.schemas": {
     "https://ultra-supara/homebrew-sisakulint/settings.json": "/.github/workflows/*.{yml,yaml}"
 }
```


## install your machine directly for macOS user

```bash
$ brew tap ultra-supara/homebrew-sisakulint
$ brew install sisakulint
```

## install from release page for Linux user

```bash
# visit release page of this repository and download for yours.
$ cd < sisakulintがあるところ >
$ mv ./sisakulint /usr/local/bin/sisakulint
```

## Run sisakulint on linux brew using docker
I am so sorry, this method is currently unavailable.
Run the following commands:

```bash
# Usage
$ git clone https://github.com/ultra-supara/homebrew-sisakulint.git
$ cd homebrew-sisakulint
$ docker compose up -d --build
$ docker attach <createされたcontainer nameに各自で変更してください>
# ここで rootに入ると思います。
$ ls -la
$ brew install sisakulint
```

## Usage test
Create a file called test.yaml in the `.github/workflows` directory or go to your repository where your workflows file is located.
```yaml
name: Upload Release Archive

on:
  push:
    tags:
      - "v[0-9]+\\.[0-9]+\\.[0-9]+"

jobs:
  build:
    name: Upload Release Asset
    runs-on: macos-latest
    steps:
      - name: Set version
        id: version
        run: |
          REPOSITORY=$(echo ${{ github.repository }} | sed -e "s#.*/##")
          VERSION=$(echo ${{ github.ref }} | sed -e "s#refs/tags/##g")
          echo ::set-output name=version::$VERSION
          echo ::set-output name=filename::$REPOSITORY-$VERSION
          echo "Version $VERSION"
      - name: Checkout code
        uses: actions/checkout@v2
        with:
          token: ${{ secrets.GITHUB_TOKEN }}
          submodules: true
      - name: Archive
        run: |
          zip -r ${{ steps.version.outputs.filename }}.zip ./ -x "*.git*"
      - name: Create Release
        id: create_release
        uses: actions/create-release@v1
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
        with:
          tag_name: ${{ github.ref }}
          release_name: Release ${{ github.ref }}
          draft: false
          prerelease: false
      - name: Upload Release Asset
        id: upload-release-asset
        uses: actions/upload-release-asset@v1
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
        with:
          upload_url: ${{ steps.create_release.outputs.upload_url }}
          asset_path: ./${{ steps.version.outputs.filename }}.zip
          asset_name: ${{ steps.version.outputs.filename }}.zip
          asset_content_type: application/zip
```
execute following commands
```bash
$ sisakulint -h
$ sisakulint -debug
```
likely, you will receive the following result...
```bash
[sisaku:🤔] linting repository... .
[sisaku:🤔] Detected project: /Users/para/go/src/github.com/ultra-supara/go_rego
[sisaku:🤔] the number of corrected yaml file 1 yaml files
[sisaku:🤔] validating workflow... .github/workflows/test.yaml
[sisaku:🤔] Detected project: /Users/para/go/src/github.com/ultra-supara/go_rego
[linter mode] no configuration file
[sisaku:🤔] parsed workflow in 2 0 ms .github/workflows/test.yaml
[SyntaxTreeVisitor] VisitStep was tooking line:13,col:9 steps, at step "2024-01-01 02:41:55.137615 +0900 JST m=+0.004782396" took 0 ms
[SyntaxTreeVisitor] VisitStep was tooking line:21,col:9 steps, at step "2024-01-01 02:41:55.137858 +0900 JST m=+0.005025666" took 0 ms
[SyntaxTreeVisitor] VisitStep was tooking line:26,col:9 steps, at step "2024-01-01 02:41:55.137874 +0900 JST m=+0.005041658" took 0 ms
[SyntaxTreeVisitor] VisitStep was tooking line:29,col:9 steps, at step "2024-01-01 02:41:55.137882 +0900 JST m=+0.005050131" took 0 ms
[SyntaxTreeVisitor] VisitStep was tooking line:39,col:9 steps, at step "2024-01-01 02:41:55.137901 +0900 JST m=+0.005069150" took 0 ms
[SyntaxTreeVisitor] VisitJobPost was tooking 5 jobs, at job "build" took 0 ms
[SyntaxTreeVisitor] VisitStep was tooking 5 steps took 0 ms
[SyntaxTreeVisitor] VisitJobPre took 0 ms
[SyntaxTreeVisitor] VisitWorkflowPost took 0 ms
[SyntaxTreeVisitor] VisitJob was tooking 1 jobs took 0 ms
[SyntaxTreeVisitor] VisitWorkflowPre took 0 ms
[linter mode] env-var found 0 errors
[linter mode] id found 0 errors
[linter mode] permissions found 0 errors
[linter mode] workflow-call found 0 errors
[linter mode] expression found 1 errors
[linter mode] deprecated-commands found 2 errors
[linter mode] cond found 0 errors
[sisaku:🤔] Found total 5 errors found in 0 found in ms .github/workflows/test.yaml
.github/workflows/test.yaml:11:14: one ${{ }} expression should be included in "runner label at \"runs-on\" section" value but got 0 expressions [expression]
       11 👈|    runs-on: macos-latest

.github/workflows/test.yaml:15:14: workflow command "set-output" was deprecated. You should use `echo "{name}={value}" >> $GITHUB_OUTPUT` reference: https://docs.github.com/en/actions/using-workflows/workflow-commands-for-github-actions [deprecated-commands]
       15 👈|        run: |

.github/workflows/test.yaml:15:14: workflow command "set-output" was deprecated. You should use `echo "{name}={value}" >> $GITHUB_OUTPUT` reference: https://docs.github.com/en/actions/using-workflows/workflow-commands-for-github-actions [deprecated-commands]
       15 👈|        run: |

.github/workflows/test.yaml:32:9: unexpected key "env" for "element of steps sequence" section. expected one of  [syntax]
       32 👈|        env:

.github/workflows/test.yaml:42:9: unexpected key "env" for "element of steps sequence" section. expected one of  [syntax]
       42 👈|        env:
```

## Try the query processing used for static analysis of sisakulint from opa cli
I only partially cut out the opa cli part of the query processing that is combined with the static analysis processing. You can use it by installing opa.
```
# Usage
$ brew install opa
$ git clone https://github.com/ultra-supara/homebrew-sisakulint.git
$ cd homebrew-sisakulint/script
# Example : commitsha.rego
$ opa eval --format pretty --data commitsha.rego --input commitsha.yaml "data.core.missing_action_ref_sha_warnings"
```
you can get such results...
```
[
  "Warning: The action ref in 'uses' should be a full length commit SHA for immutability and security. see documents : https://docs.github.com/ja/actions/security-guides/security-hardening-for-github-actions#using-third-party-actions"
]
```
This rule can proactively prompt you to use the commit shan feature when using third-party actions. This implementation has proven to be extremely successful in processing queries.

If you want to use it more widely...
```
$ opa eval --format pretty --data commitsha.rego --input ./github/workflows/*.yaml "data.core.missing_action_ref_sha_warnings"
```
this part `"data.core.missing_action_ref_sha_warnings"` changes depending on the query you want to run. Please rewrite each rego file you want to try.

## Links

- [developer document](https://www.notion.so/ultra-supara/sisakulint-user-document-d3f28d427cf9456dbe3c0f063a7d3baf?pvs=4)
- [user document](https://www.notion.so/ultra-supara/sisakulint-c18505b443254ee5a3e5e3751b810a33?pvs=4)

- slides
- poster
- video
