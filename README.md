## Oauth2 Google Account & JWT

### Preparing
- Create OAuth consent screen [Google Cloud APIs](https://console.cloud.google.com/apis)
- Create app information (only the required content on step 1)
- Select scope (select userinfo.email and userinfo.profile on step 2)
- Save and continue on
- Next Click `Credentials` on left button
- Create Credentials , select application type `web application`
- Set 2 URL on `Authorized JavaScript origins` : set `http://localhost` & `http://localhost:<your-port>`
- Next, set `Authorized redirect URIs` your callback URL : example `http://localhost:8080/callback`
- Save
