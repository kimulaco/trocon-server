echo "gcloud run deploy $DEPLOY_PROJECT_NAME --region $DEPLOY_REGION --source . --set-env-vars BUILD_MODE=production,CORS_ALLOW_ORIGIN=$CORS_ALLOW_ORIGIN,STEAM_API_KEY=$STEAM_API_KEY,STEAM_API_BASE_URL=$STEAM_API_BASE_URL,SENTRY_DSN=$SENTRY_DSN
"

gcloud run deploy $DEPLOY_PROJECT_NAME --region $DEPLOY_REGION --source . --set-env-vars BUILD_MODE=production,CORS_ALLOW_ORIGIN=$CORS_ALLOW_ORIGIN,STEAM_API_KEY=$STEAM_API_KEY,STEAM_API_BASE_URL=$STEAM_API_BASE_URL,SENTRY_DSN=$SENTRY_DSN
