package core

# pull_request イベントでプルリクエストのタイトルを使用するワークフローを検出
deny[reason] {
    input.on.pull_request
    job := input.jobs[_]
    step := job.steps[_]
    regex.match(`\$\{\{ github\.event\.pull_request\.title \}\}`, step.run)
    reason := "Workflow uses pull_request event with PR title in steps, which may lead to security risks."
}

# pull_request_target イベントと actions/checkout アクションを使用するワークフローを検出
deny[reason] {
    input.on.pull_request_target
    job := input.jobs[_]
    step := job.steps[_]
    step.uses == "actions/checkout@v2"
    reason := "Workflow uses pull_request_target event with actions/checkout, which may lead to security risks."
}
