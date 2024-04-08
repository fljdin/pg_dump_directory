#! /usr/bin/env bats

COMMAND="go run ."

@test "help" {
    $COMMAND -h
    $COMMAND --help
}

@test "failed to open file" {
    run $COMMAND -f null
    [[ "$status" -eq 1 ]]
    [[ "$output" == *"open null: no such file or directory"* ]]
}
