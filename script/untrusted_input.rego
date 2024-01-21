package core

# Detect direct usage of ${{ ... }} in run scripts and suggest using env
deny[reason] {
    some i, j
    step := input.jobs[i].steps[j]
    script := step.run
    regex.match(`\$\{\{.*\}\}`, script)
    not step.env
    reason := sprintf("Direct usage of ${{ ... }} in run script detected at jobs.%v.steps.%v. Use env to set variables.", [i, j])
}

# Detect unsafe usage of ${{ ... }} in GitHub Actions 'uses' field
deny[reason] {
    some i, j
    step := input.jobs[i].steps[j]
    uses_field := step.uses
    regex.match(`\$\{\{.*\}\}`, uses_field)
    reason := sprintf("Direct usage of ${{ ... }} in 'uses' field detected at jobs.%v.steps.%v.", [i, j])
}

# Detect unsafe usage of ${{ ... }} in GitHub Actions 'with' field
deny[reason] {
    some i, j
    step := input.jobs[i].steps[j]
    with_field := step["with"]
    regex.match(`\$\{\{.*\}\}`, with_field[_])
    reason := sprintf("Direct usage of ${{ ... }} in 'with' field detected at jobs.%v.steps.%v.", [i, j])
}

# Detect unsafe usage of ${{ ... }} in toJSON function
deny[reason] {
    some i, j
    step := input.jobs[i].steps[j]
    script := step.run
    regex.match(`\$\{\{\s*toJSON\(github\.event\.\*\..*\)\s*\}\}`, script)
    reason := sprintf("Unsafe usage of ${{ toJSON(...) }} detected at jobs.%v.steps.%v.", [i, j])
}
