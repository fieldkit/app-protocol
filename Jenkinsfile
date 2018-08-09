@Library('conservify') _

properties([
    [$class: 'BuildDiscarderProperty', strategy: [$class: 'LogRotator', numToKeepStr: '5']],
    pipelineTriggers([[$class: 'GitHubPushTrigger']]),
])

timestamps {
    node () {
        conservifyBuild(name: 'app-protocol', archive: false)
    }
}
