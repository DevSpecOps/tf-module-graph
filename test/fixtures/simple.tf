module "vpc" {
  source = "./modules/vpc"
}

module "app" {
  source = "./modules/app"
}