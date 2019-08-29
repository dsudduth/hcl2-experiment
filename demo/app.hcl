project "Demo" {
  description = "Demo is pretty awesome."
  version     = "1.0.0"
}

automation_set "build_acceptance" {
  description = "This set executes the acceptance tests for the Demo application."
  machines = [
    "ubuntu-server-64",
    "centos-7-64"
  ]
}
