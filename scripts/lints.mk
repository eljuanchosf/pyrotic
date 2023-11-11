#####################
##@ Lints   
#####################

lint: ## Lint tools
	go vet ./...
	golangci-lint run ./...

scan: ## Security scanning
	gosec ./...
	govulncheck ./...

trivy: ## Trivy secrets scanning
	@trivy fs .