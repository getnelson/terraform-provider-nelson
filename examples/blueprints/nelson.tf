terraform {
  required_version = "~> 0.11.7"
}

provider nelson {
  address = "https://nelson.local"
  path = "/home/nelson/.nelson/config.yml"
}

resource "nelson_blueprint" "deployment" {
  name = "nelson"
  description = "deployment blueprint"
  file = "${file("${path.module}/deployment.bp")}"
}
