# `backend/` for blami.net
This directory contains [Google App Engine][1] backend Go app for
https://blami.net. Purpose of this app is so far:

- Serve presence status for [Twitch][2] (whether channel is on-line or not) so
  it can be shown on my homepage without visitor needing to authenticate with
  Twitch first to obtain access token.

- Serve presence status for [Steam][3] (whether user is in game or not) so it
  can be shown on my homepage - again without visitor needing to authenticate
  with Steam.

## Twitch Presence
  Twitch requires OAuth2 bearer token on every API request. Twitch also
  provides PubSub where one can

## Steam Presence
  
