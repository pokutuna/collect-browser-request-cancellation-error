PROJECT := pokutuna-playground
GCLOUD := gcloud beta --project $(PROJECT)

.PHONY: deploy
deploy:
	$(GCLOUD) app deploy app/app.yaml
