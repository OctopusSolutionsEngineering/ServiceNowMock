# ServiceNow Mock

This web app mocks the ServiceNow endpoints used by Octopus for ITSM integration. It accepts any credentials. Change requests will automatically move into the implementation state after 5 requests.

This app has been deployed to https://mockservicenow.gentlemushroom-1ab8bf8e.westus2.azurecontainerapps.io/.

To use the app in Octopus, create a new ServiceNow connection and use the URL above. You can use any username and password, and any Oauth client ID and secret.