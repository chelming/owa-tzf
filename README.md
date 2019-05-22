# owa-tzf
OWA Timezone Fixer
Fixes the timezone for a OWA ics feed that's reporting UTC instead of the actual timezone.

Example Docker run:
```
docker run -d --name "owa-tzf" -p 8080:8080 \
  -e URL='https://url_to_feed.ics' \
  -e TZID='Eastern Standard Time' \
  cwhits/owa-tzf:latest
```
