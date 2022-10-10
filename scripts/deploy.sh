echo "gcloud run deploy trophy-comp --region us-central1 --source . --image us-central1-docker.pkg.dev/trophy-comp/cloud-run-source-deploy/trophy-comp --set-env-vars BUILD_MODE=production,CORS_ALLOW_ORIGIN=$CORS_ALLOW_ORIGIN,STEAM_API_KEY=$STEAM_API_KEY
"

gcloud run deploy trophy-comp --region us-central1 --image us-central1-docker.pkg.dev/trophy-comp/cloud-run-source-deploy/trophy-comp --set-env-vars BUILD_MODE=production,CORS_ALLOW_ORIGIN=$CORS_ALLOW_ORIGIN,STEAM_API_KEY=$STEAM_API_KEY
