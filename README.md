# Wasp
A desktop notification reminder for Google Calender tasks. Sends alerts
for specified tasks if they are late and incomplete.

## Limitations
Due to limitations with Google's Task API, it is not possible to
know the exact time a task is due as Google only provides the date.
As a result, one would need to pass in the time manually into wasp.

## Flags
The following flags are available as part of wasp:
- **-token-file** The path to the token file. If no such file exists, one
  will be created the first time wasp is run. In that case, a prompt
  will be displayed with a link to authorize wasp with Google. It is also
  possible to set the environment variable **WASP_TOKEN_FILE** instead of
  passing in this flag.
- **-creds-file** The path to the credentials file. This is the file that is       
  downloaded after creating [OAuth 2.0 pair](#installation). It is also
  possible to set the environment variable **WASP_CREDS_FILE** instead of
  passing in this flag.
- **-time** The time the task is due. It should follow the format `XXhXXm`,
  where `XXh` is the hour the task is due (in 24h format), and `XXm` is the
  minute of the hour the task is due.
- **-name** The name of the task. This should be the same as the `Title` of the
  task as it appears on Google Calender.

## Installation
1. Create a new [Google Cloud project](https://console.cloud.google.com)
   and enable the Tasks API.
2. Create a new OAuth 2.0 pair under  **OAuth 2.0 Client IDs**.
3. Add the gmail account that will have its tasks read under 
   **Test users**. This can be found in the **OAuth consent screen** tab.
3. Clone the repository and install dependencies with `go get`.
4. Launch wasp with `go run main.go` and include the appropriate flags.
5. As wasp is structured as a command that only checks if a task is
   past due when run, it is best to schedule it with cron.
   This way, wasp will automatically run periodically and check if the
   task is past due. 