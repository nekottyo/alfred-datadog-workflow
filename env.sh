#!/usr/bin/env zsh
root="$( git rev-parse --show-toplevel )"
testdir="${root}/testenv"

# Absolute bare-minimum for AwGo to function...
export alfred_workflow_bundleid="net.deanishe.awgo"
export alfred_workflow_data="${testdir}/data"
export alfred_workflow_cache="${testdir}/cache"
export alfred_workflow_version="$(cat ${root}/VERSION)"
export alfred_workflow_name="alfred-datadog-workflow"
