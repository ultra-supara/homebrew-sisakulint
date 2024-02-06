# sisakulint
Before moving on, please consider giving us a GitHub star ⭐️. Thank you!

![logo](https://github.com/ultra-supara/homebrew-sisakulint/assets/67861004/e9801cbb-fbe1-4822-a5cd-d1daac33e90f)

## install macOS user

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
    env:
          SIIISA=AAKUUU: foo
    steps:
      - name: Set version
        id: version
        run: |
          REPOSITORY=$(echo ${{ github.repository }} | sed -e "s#.*/##")
          echo ::set-output name=filename::$REPOSITORY-$VERSION
      - name: Checkout code
        uses: actions/checkout@v2
        with:
          token: ${{ secrets.GITHUB_TOKEN }}
          submodules: true
      - name: Archive
        run: |
          zip -r ${{ steps.version.outputs.filename }}.zip ./ -x "*.git*"
      - run: echo 'Commit is pushed'
        # ERROR: It is always evaluated to true
        if: |
          ${{ github.event_name == 'push' }}
      - name: Create Release
        id: create_release
        uses: actions/create-release@v1
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
          FOO=BAR: foo
          FOO BAR: foo
        with:
          tag_name: ${{ github.ref }}
          release_name: Release ${{ github.ref }}
          draft: false
          prerelease: false
      - name: Upload Release Asset
        id: upload-release-asset
        uses: actions/upload-release-asset@v1
        with:
          upload_url: ${{ steps.create_release.outputs.upload_url }}
          asset_path: ./${{ steps.version.outputs.filename }}.zip
          asset_name: ${{ steps.version.outputs.filename }}.zip
          asset_content_type: application/zip

  test:
    runs-on: ubuntu-latest
    permissions:
      # ERROR: "checks" is correct scope name
      check: write
      # ERROR: Available values are "read", "write" or "none"
      issues: readable
    steps:
      - run: echo '${{ "hello" }}'
      - run: echo "${{ toJson(hashFiles('**/lock', '**/cache/') }}"
      - run: echo '${{ github.event. }}'

  run shell:
    steps:
      - run: echo 'hello'
```
execute following commands
```bash
$ sisakulint -h
$ sisakulint -debug
```
you will likely receive the following result...
```bash
[sisaku:🤔] linting repository... .
[sisaku:🤔] Detected project: /Users/para/go/src/github.com/ultra-supara/go_rego
[sisaku:🤔] the number of corrected yaml file 1 yaml files
[sisaku:🤔] validating workflow... .github/workflows/a.yaml
[sisaku:🤔] Detected project: /Users/para/go/src/github.com/ultra-supara/go_rego
[linter mode] no configuration file
[sisaku:🤔] parsed workflow in 2 0 ms .github/workflows/a.yaml
[SyntaxTreeVisitor] VisitStep was tooking line:15,col:9 steps, at step "2024-01-25 21:51:17.572627 +0900 JST m=+0.011103420" took 0 ms
[SyntaxTreeVisitor] VisitStep was tooking line:20,col:9 steps, at step "2024-01-25 21:51:17.57293 +0900 JST m=+0.011406635" took 0 ms
[SyntaxTreeVisitor] VisitStep was tooking line:25,col:9 steps, at step "2024-01-25 21:51:17.572954 +0900 JST m=+0.011430674" took 0 ms
[SyntaxTreeVisitor] VisitStep was tooking line:28,col:9 steps, at step "2024-01-25 21:51:17.572973 +0900 JST m=+0.011449393" took 0 ms
[SyntaxTreeVisitor] VisitStep was tooking line:32,col:9 steps, at step "2024-01-25 21:51:17.572998 +0900 JST m=+0.011474986" took 0 ms
[SyntaxTreeVisitor] VisitStep was tooking line:44,col:9 steps, at step "2024-01-25 21:51:17.573019 +0900 JST m=+0.011495668" took 0 ms
[SyntaxTreeVisitor] VisitJobPost was tooking 6 jobs, at job "build" took 0 ms
[SyntaxTreeVisitor] VisitStep was tooking 6 steps took 0 ms
[SyntaxTreeVisitor] VisitJobPre took 0 ms
[SyntaxTreeVisitor] VisitStep was tooking line:61,col:9 steps, at step "2024-01-25 21:51:17.573077 +0900 JST m=+0.011553316" took 0 ms
[SyntaxTreeVisitor] VisitStep was tooking line:62,col:9 steps, at step "2024-01-25 21:51:17.57309 +0900 JST m=+0.011566716" took 0 ms
[SyntaxTreeVisitor] VisitStep was tooking line:63,col:9 steps, at step "2024-01-25 21:51:17.573098 +0900 JST m=+0.011575087" took 0 ms
[SyntaxTreeVisitor] VisitJobPost was tooking 3 jobs, at job "test" took 0 ms
[SyntaxTreeVisitor] VisitStep was tooking 3 steps took 0 ms
[SyntaxTreeVisitor] VisitJobPre took 0 ms
[SyntaxTreeVisitor] VisitStep was tooking line:67,col:9 steps, at step "2024-01-25 21:51:17.573159 +0900 JST m=+0.011636302" took 0 ms
[SyntaxTreeVisitor] VisitJobPost was tooking 1 jobs, at job "run shell" took 0 ms
[SyntaxTreeVisitor] VisitStep was tooking 1 steps took 0 ms
[SyntaxTreeVisitor] VisitJobPre took 0 ms
[SyntaxTreeVisitor] VisitWorkflowPost took 0 ms
[SyntaxTreeVisitor] VisitJob was tooking 3 jobs took 0 ms
[SyntaxTreeVisitor] VisitWorkflowPre took 0 ms
[linter mode] env-var found 1 errors
[linter mode] id found 1 errors
[linter mode] permissions found 2 errors
[linter mode] workflow-call found 0 errors
[linter mode] expression found 5 errors
[linter mode] deprecated-commands found 1 errors
[linter mode] cond found 1 errors
[sisaku:🤔] Found total 13 errors found in 1 found in ms .github/workflows/a.yaml
.github/workflows/a.yaml:11:14: one ${{ }} expression should be included in "runner label at \"runs-on\" section" value but got 0 expressions [expression]
       11 👈|    runs-on: macos-latest
                    
.github/workflows/a.yaml:13:11: Environment variable name '"SIIISA=AAKUUU"' is not formatted correctly. Please ensure that it does not include characters such as '&', '=', or spaces, as these are not allowed in variable names. [env-var]
       13 👈|          SIIISA=AAKUUU: foo
                 
.github/workflows/a.yaml:17:14: workflow command "set-output" was deprecated. You should use `echo "{name}={value}" >> $GITHUB_OUTPUT` reference: https://docs.github.com/en/actions/using-workflows/workflow-commands-for-github-actions [deprecated-commands]
       17 👈|        run: |
                    
.github/workflows/a.yaml:30:13: The condition '${{ github.event_name == 'push' }}
' will always evaluate to true. If you intended to use a literal value, please use ${{ true }}. Ensure there are no extra characters within the ${{ }} brackets in conditions. [cond]
       30 👈|        if: |
                   
.github/workflows/a.yaml:35:9: unexpected key "env" for "element of \"steps\" sequence" section. expected one of  [syntax]
       35 👈|        env:
               
.github/workflows/a.yaml:54:14: one ${{ }} expression should be included in "runner label at \"runs-on\" section" value but got 0 expressions [expression]
       54 👈|    runs-on: ubuntu-latest
                    
.github/workflows/a.yaml:57:7: unknown permission scope "check". all available permission scopes are "actions", "checks", "contents", "deployments", "discussions", "id-token", "issues", "packages", "pages", "pull-requests", "repository-projects", "security-events", "statuses" [permissions]
       57 👈|      check: write
             
.github/workflows/a.yaml:59:15: The value "readable" is not a valid permission for the scope "issues". Only 'read', 'write', or 'none' are acceptable values. [permissions]
       59 👈|      issues: readable
                     
.github/workflows/a.yaml:61:24: got unexpected char '"' while lexing expression, expecting 'a'..'z', 'A'..'Z', '_', '0'..'9', '', '}', '(', ')', '[', ']', '.', '!', '<', '>', '=', '&', '|', '*', ',', ' '. do you mean string literals? only single quotes are available for string delimiter [expression]
       61 👈|      - run: echo '${{ "hello" }}'
                              
.github/workflows/a.yaml:62:65: unexpected end of expression, while parsing arguments of function call, expected ",", ")" [expression]
       62 👈|      - run: echo "${{ toJson(hashFiles('**/lock', '**/cache/') }}"
                                                                       
.github/workflows/a.yaml:63:38: unexpected end of expression, while parsing expected an object property dereference (like 'a.b') or an array element dereference (like 'a.*'), expected "IDENT", "*" [expression]
       63 👈|      - run: echo '${{ github.event. }}'
                                            
.github/workflows/a.yaml:65:3: "runs-on" section is missing in job "run shell" [syntax]
       65 👈|  run shell:
         
.github/workflows/a.yaml:65:3: Invalid job ID "run shell". job IDs must start with a letter or '_', and may contain only alphanumeric characters, '-', or '_'. [id]
       65 👈|  run shell:
```

## Try rego query from CI integration
I partially cut out the opa cli part of the query processing that is combined with the static analysis processing.
set your `.github/workflows` dir any name (example : query.yaml)
```yaml
name: policy check for oss repo from GitHub Actions
on: [push]

jobs:
  opa-check:
    name: opa cli
    runs-on: ubuntu-latest
    steps:
    - name: Checkout sisakulint Repository
      uses: actions/checkout@v2
      with:
        repository: ultra-supara/homebrew-sisakulint
        path: sisakulint

    - name: Install OPA
      run: |
        curl -L -o opa https://openpolicyagent.org/downloads/latest/opa_linux_amd64
        chmod +x opa
        sudo mv opa /usr/local/bin/

    - name: Run OPA Policy Check
      run: |
        git clone https://github.com/ultra-supara/homebrew-sisakulint.git
        cd homebrew-sisakulint/script
        echo "-------start---------"
        opa eval --format pretty --data commitsha.rego --input commitsha.yaml "data.core.missing_action_ref_sha_warnings" --explain=full
        echo "---------------------"
        opa eval --format pretty --data credential.rego --input credential.yaml "data.core.check_credentials" --explain=full
        echo "---------------------"
        opa eval --format pretty --data imagetag.rego --input imagetag.yaml "data.core.missing_image_tag_warnings" --explain=full
        echo "---------------------"
        opa eval --format pretty --data jobsecrets.rego --input jobsecrets.yaml "data.core.missing_secrets_warnings" --explain=full
        echo "---------------------"
        opa eval --format pretty --data permission.rego --input permission.yaml "data.core.missing_permissions_warnings" --explain=full
        echo "---------------------"
        opa eval --format pretty --data issueinjection.rego --input issueinjection.yaml "data.core.deny" --explain=full
        echo "---------------------"
        opa eval --format pretty --data pull_req_title.rego --input pull_req_title.yaml "data.core.deny" --explain=full
        echo "---------------------"
        opa eval --format pretty --data supply_chain_protection.rego --input supply_chain_protection.yaml "data.core.generate_error_messages" --explain=full
        echo "---------------------"
        opa eval --format pretty --data timeout-minutes.rego --input timeout-minutes.yaml "data.core.missing_timeout_warnings" --explain=full
        echo "---------------------"
        opa eval --format pretty --data untrusted_image.rego --input untrusted_image.yaml "data.core.deny" --explain=full
        echo "--------end----------"
```

## JSON schema for GitHub Actions syntax
paste yours `settings.json`

```
 "yaml.schemas": {
     "https://ultra-supara/homebrew-sisakulint/settings.json": "/.github/workflows/*.{yml,yaml}"
 }
```

## Links

- slides
- poster
- video
