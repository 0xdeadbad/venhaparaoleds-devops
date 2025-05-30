# ---- locals ----
locals {
  POSTGRES_ENV = var.POSTGRES_ENV
}
# ---- end ----

# ---- images ----
resource "docker_image" "nginx" {
  name = "nginx:1.27.5-alpine"
}

resource "docker_image" "postgres" {
  name = "postgres:17.5-alpine"
}

resource "docker_image" "ledsproj" {
  name = "ledsproj"
  build {
    context    = "../../"
    tag        = ["ledsproj:latest"]
    build_args = []
    label      = {}
  }
}
# ---- end ----

# ---- containers ----
resource "docker_container" "postgres" {
  name  = "postgres-ledsproj"
  image = docker_image.postgres.image_id
  networks = [
    docker_network.internal_network
  ]
  env = POSTGRES_ENV
}

resource "docker_container" "ledsproj" {
  name  = "ledsproj"
  image = docker_image.ledsproj.image_id
  networks = [
    docker_network.internal_network
  ]
  env = POSTGRES_ENV
}

resource "docker_container" "nginx" {
  name  = "nginx-ledsproj"
  image = docker_image.nginx.image_id
  networks = [
    docker_network.default_network,
    docker_network.internal_network
  ]

  ports {
    internal = 8080
    external = 8080
  }
}
# ---- end ----

# ---- networks ----
resource "docker_network" "internal_network" {
  name     = "internal_network"
  driver   = "bridge"
  internal = true
}

resource "docker_network" "default_network" {
  name     = "default"
  driver   = "bridge"
  internal = false
}
# ---- end ----
