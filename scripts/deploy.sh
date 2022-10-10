gcloud run deploy \
  trophy-comp
  --region us-central1
  --source .
  --set-env-vars BUILD_MODE=production,CORS_ALLOW_ORIGIN=$CORS_ALLOW_ORIGIN,STEAM_API_KEY=$STEAM_API_KEY
